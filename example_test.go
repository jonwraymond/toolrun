package toolrun_test

import (
	"context"
	"fmt"

	"github.com/jonwraymond/toolmodel"
	"github.com/jonwraymond/toolrun"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// simpleLocalRegistry is a basic LocalRegistry implementation for examples.
type simpleLocalRegistry struct {
	handlers map[string]toolrun.LocalHandler
}

func newSimpleLocalRegistry() *simpleLocalRegistry {
	return &simpleLocalRegistry{handlers: make(map[string]toolrun.LocalHandler)}
}

func (r *simpleLocalRegistry) Get(name string) (toolrun.LocalHandler, bool) {
	h, ok := r.handlers[name]
	return h, ok
}

func (r *simpleLocalRegistry) Register(name string, h toolrun.LocalHandler) {
	r.handlers[name] = h
}

func Example_basicRun() {
	// Create a tool
	tool := toolmodel.Tool{
		Tool: mcp.Tool{
			Name:        "greet",
			InputSchema: map[string]any{"type": "object"},
		},
	}

	// Create a backend
	backend := toolmodel.ToolBackend{
		Kind:  toolmodel.BackendKindLocal,
		Local: &toolmodel.LocalBackend{Name: "greeter"},
	}

	// Create a local registry with a handler
	localReg := newSimpleLocalRegistry()
	localReg.Register("greeter", func(_ context.Context, args map[string]any) (any, error) {
		name, _ := args["name"].(string)
		if name == "" {
			name = "World"
		}
		return map[string]any{"greeting": "Hello, " + name + "!"}, nil
	})

	// Create resolvers (in production, you'd use toolindex.Index)
	toolResolver := func(id string) (*toolmodel.Tool, error) {
		if id == "greet" {
			return &tool, nil
		}
		return nil, fmt.Errorf("tool not found: %s", id)
	}
	backendsResolver := func(id string) ([]toolmodel.ToolBackend, error) {
		if id == "greet" {
			return []toolmodel.ToolBackend{backend}, nil
		}
		return nil, fmt.Errorf("no backends for: %s", id)
	}

	// Create runner
	runner := toolrun.NewRunner(
		toolrun.WithToolResolver(toolResolver),
		toolrun.WithBackendsResolver(backendsResolver),
		toolrun.WithLocalRegistry(localReg),
		toolrun.WithValidation(false, false), // Disable validation for example
	)

	// Execute the tool
	result, err := runner.Run(context.Background(), "greet", map[string]any{"name": "Claude"})
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// Access the structured result
	m := result.Structured.(map[string]any)
	fmt.Println(m["greeting"])

	// Output:
	// Hello, Claude!
}

func Example_chainExecution() {
	// Create tools for a processing pipeline
	tools := map[string]toolmodel.Tool{
		"fetch": {Tool: mcp.Tool{
			Name:        "fetch",
			InputSchema: map[string]any{"type": "object"},
		}},
		"transform": {Tool: mcp.Tool{
			Name:        "transform",
			InputSchema: map[string]any{"type": "object"},
		}},
		"store": {Tool: mcp.Tool{
			Name:        "store",
			InputSchema: map[string]any{"type": "object"},
		}},
	}

	backends := map[string]toolmodel.ToolBackend{}
	for name := range tools {
		backends[name] = toolmodel.ToolBackend{
			Kind:  toolmodel.BackendKindLocal,
			Local: &toolmodel.LocalBackend{Name: name + "-handler"},
		}
	}

	// Create handlers
	localReg := newSimpleLocalRegistry()

	localReg.Register("fetch-handler", func(_ context.Context, args map[string]any) (any, error) {
		return map[string]any{"data": []string{"item1", "item2", "item3"}}, nil
	})

	localReg.Register("transform-handler", func(_ context.Context, args map[string]any) (any, error) {
		prev, _ := args["previous"].(map[string]any)
		data, _ := prev["data"].([]string)
		transformed := make([]string, len(data))
		for i, item := range data {
			transformed[i] = "processed-" + item
		}
		return map[string]any{"data": transformed}, nil
	})

	localReg.Register("store-handler", func(_ context.Context, args map[string]any) (any, error) {
		prev, _ := args["previous"].(map[string]any)
		data, _ := prev["data"].([]string)
		return map[string]any{
			"stored": len(data),
			"status": "success",
		}, nil
	})

	// Create resolvers
	toolResolver := func(id string) (*toolmodel.Tool, error) {
		if t, ok := tools[id]; ok {
			return &t, nil
		}
		return nil, fmt.Errorf("tool not found: %s", id)
	}
	backendsResolver := func(id string) ([]toolmodel.ToolBackend, error) {
		if b, ok := backends[id]; ok {
			return []toolmodel.ToolBackend{b}, nil
		}
		return nil, fmt.Errorf("no backends for: %s", id)
	}

	// Create runner
	runner := toolrun.NewRunner(
		toolrun.WithToolResolver(toolResolver),
		toolrun.WithBackendsResolver(backendsResolver),
		toolrun.WithLocalRegistry(localReg),
		toolrun.WithValidation(false, false),
	)

	// Define the chain
	steps := []toolrun.ChainStep{
		{ToolID: "fetch"},
		{ToolID: "transform", UsePrevious: true},
		{ToolID: "store", UsePrevious: true},
	}

	// Execute the chain
	final, _, err := runner.RunChain(context.Background(), steps)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	m := final.Structured.(map[string]any)
	fmt.Printf("Stored %v items: %s\n", m["stored"], m["status"])

	// Output:
	// Stored 3 items: success
}

func Example_customBackendSelector() {
	// Custom selector that prefers provider backends for specific tools
	customSelector := func(backends []toolmodel.ToolBackend) toolmodel.ToolBackend {
		// First try to find a provider backend
		for _, b := range backends {
			if b.Kind == toolmodel.BackendKindProvider {
				return b
			}
		}
		// Fall back to any available backend
		return backends[0]
	}

	// Create a tool with multiple backends
	tool := toolmodel.Tool{
		Tool: mcp.Tool{
			Name:        "multi-backend-tool",
			InputSchema: map[string]any{"type": "object"},
		},
	}

	localBackend := toolmodel.ToolBackend{
		Kind:  toolmodel.BackendKindLocal,
		Local: &toolmodel.LocalBackend{Name: "local-handler"},
	}
	providerBackend := toolmodel.ToolBackend{
		Kind: toolmodel.BackendKindProvider,
		Provider: &toolmodel.ProviderBackend{
			ProviderID: "my-provider",
			ToolID:     "provider-tool",
		},
	}

	backends := []toolmodel.ToolBackend{localBackend, providerBackend}

	// Create resolvers
	toolResolver := func(_ string) (*toolmodel.Tool, error) {
		return &tool, nil
	}
	backendsResolver := func(_ string) ([]toolmodel.ToolBackend, error) {
		return backends, nil
	}

	// Create runner with custom selector
	runner := toolrun.NewRunner(
		toolrun.WithToolResolver(toolResolver),
		toolrun.WithBackendsResolver(backendsResolver),
		toolrun.WithBackendSelector(customSelector),
		toolrun.WithValidation(false, false),
	)

	// At this point, the runner would use the provider backend
	// even though a local backend is also available
	_ = runner

	fmt.Println("Runner configured with custom backend selector")

	// Output:
	// Runner configured with custom backend selector
}
