package toolrun

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// MCPExecutor executes MCP tool calls using MCP Go SDK types.
// Implementations typically wrap an MCP client connection.
type MCPExecutor interface {
	// CallTool executes a tool call and returns the result.
	CallTool(ctx context.Context, serverName string, params *mcp.CallToolParams) (*mcp.CallToolResult, error)

	// CallToolStream executes a tool call with streaming.
	// Implementations may return ErrStreamNotSupported when streaming is unavailable.
	// Contract: if err is nil, the returned channel MUST be non-nil.
	CallToolStream(ctx context.Context, serverName string, params *mcp.CallToolParams) (<-chan StreamEvent, error)
}

// ProviderExecutor executes provider-bound tools.
// It is intentionally generic but uses canonical tool IDs and args.
type ProviderExecutor interface {
	// CallTool executes a provider tool and returns the result.
	CallTool(ctx context.Context, providerID, toolID string, args map[string]any) (any, error)

	// CallToolStream executes a provider tool with streaming.
	// Implementations may return ErrStreamNotSupported when streaming is unavailable.
	// Contract: if err is nil, the returned channel MUST be non-nil.
	CallToolStream(ctx context.Context, providerID, toolID string, args map[string]any) (<-chan StreamEvent, error)
}

// LocalHandler is the function signature for local tool execution.
// It receives a context and arguments, and returns a result or error.
type LocalHandler func(ctx context.Context, args map[string]any) (any, error)

// LocalRegistry resolves local handlers by name.
// Implementations provide a mapping from handler names to LocalHandler functions.
type LocalRegistry interface {
	// Get returns the handler for the given name, or false if not found.
	Get(name string) (LocalHandler, bool)
}
