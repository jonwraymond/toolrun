package toolrun

import (
	"github.com/jonwraymond/toolindex"
	"github.com/jonwraymond/toolmodel"
)

// Config controls resolution, validation, and dispatch behavior.
type Config struct {
	// Resolution

	// Index is the tool registry for lookup.
	Index toolindex.Index

	// ToolResolver is a fallback function to resolve tools when Index is not
	// configured or returns ErrNotFound.
	ToolResolver func(id string) (*toolmodel.Tool, error)

	// BackendsResolver is a fallback function to resolve backends when Index
	// is not configured or returns ErrNotFound.
	BackendsResolver func(id string) ([]toolmodel.ToolBackend, error)

	// BackendSelector chooses which backend to use when multiple are available.
	// Defaults to toolindex.DefaultBackendSelector (local > provider > mcp).
	BackendSelector toolindex.BackendSelector

	// Validation

	// Validator validates tool inputs and outputs against JSON Schema.
	// Defaults to toolmodel.NewDefaultValidator().
	Validator toolmodel.SchemaValidator

	// ValidateInput enables input validation before execution.
	// Defaults to true.
	ValidateInput bool

	// ValidateOutput enables output validation after execution.
	// Defaults to true.
	ValidateOutput bool

	// Executors

	// MCP is the executor for MCP backend tools.
	MCP MCPExecutor

	// Provider is the executor for provider backend tools.
	Provider ProviderExecutor

	// Local is the registry for local handler functions.
	Local LocalRegistry
}

// applyDefaults sets default values for unset Config fields.
func (c *Config) applyDefaults() {
	if c.Validator == nil {
		c.Validator = toolmodel.NewDefaultValidator()
	}
	if c.BackendSelector == nil {
		c.BackendSelector = toolindex.DefaultBackendSelector
	}
}

// ConfigOption is a functional option for configuring a Runner.
type ConfigOption func(*Config)

// WithIndex sets the tool index for resolution.
func WithIndex(idx toolindex.Index) ConfigOption {
	return func(c *Config) {
		c.Index = idx
	}
}

// WithValidator sets a custom schema validator.
func WithValidator(v toolmodel.SchemaValidator) ConfigOption {
	return func(c *Config) {
		c.Validator = v
	}
}

// WithMCPExecutor sets the MCP executor.
func WithMCPExecutor(exec MCPExecutor) ConfigOption {
	return func(c *Config) {
		c.MCP = exec
	}
}

// WithProviderExecutor sets the provider executor.
func WithProviderExecutor(exec ProviderExecutor) ConfigOption {
	return func(c *Config) {
		c.Provider = exec
	}
}

// WithLocalRegistry sets the local handler registry.
func WithLocalRegistry(reg LocalRegistry) ConfigOption {
	return func(c *Config) {
		c.Local = reg
	}
}

// WithValidation sets whether to validate inputs and outputs.
func WithValidation(input, output bool) ConfigOption {
	return func(c *Config) {
		c.ValidateInput = input
		c.ValidateOutput = output
	}
}

// WithBackendSelector sets a custom backend selector function.
func WithBackendSelector(selector toolindex.BackendSelector) ConfigOption {
	return func(c *Config) {
		c.BackendSelector = selector
	}
}

// WithToolResolver sets a fallback tool resolver function.
func WithToolResolver(resolver func(id string) (*toolmodel.Tool, error)) ConfigOption {
	return func(c *Config) {
		c.ToolResolver = resolver
	}
}

// WithBackendsResolver sets a fallback backends resolver function.
func WithBackendsResolver(resolver func(id string) ([]toolmodel.ToolBackend, error)) ConfigOption {
	return func(c *Config) {
		c.BackendsResolver = resolver
	}
}
