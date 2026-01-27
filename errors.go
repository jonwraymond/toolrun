package toolrun

import (
	"errors"
	"fmt"

	"github.com/jonwraymond/toolindex"
	"github.com/jonwraymond/toolmodel"
)

// Sentinel errors for common failure conditions.
var (
	// ErrToolNotFound is returned when a tool cannot be resolved.
	ErrToolNotFound = errors.New("tool not found")

	// ErrNoBackends is returned when a tool has no available backends.
	ErrNoBackends = errors.New("no backends available")

	// ErrValidation is returned when input validation fails.
	ErrValidation = errors.New("validation error")

	// ErrExecution is returned when tool execution fails.
	ErrExecution = errors.New("execution error")

	// ErrOutputValidation is returned when output validation fails.
	ErrOutputValidation = errors.New("output validation error")

	// ErrStreamNotSupported is returned when streaming is not supported
	// by the executor or backend.
	ErrStreamNotSupported = errors.New("streaming not supported")
)

// ToolError wraps an error with tool execution context.
// It preserves the tool ID, backend, and operation for debugging.
type ToolError struct {
	// ToolID is the canonical tool identifier.
	ToolID string

	// Backend is the backend that was used (may be nil for resolution errors).
	Backend *toolmodel.ToolBackend

	// Op is the operation that failed (e.g., "resolve", "validate_input", "execute").
	Op string

	// Err is the underlying error.
	Err error
}

// Error returns a formatted error message including context.
func (e *ToolError) Error() string {
	if e.Backend != nil {
		return fmt.Sprintf("toolrun: %s %s [%s]: %v", e.Op, e.ToolID, e.Backend.Kind, e.Err)
	}
	return fmt.Sprintf("toolrun: %s %s: %v", e.Op, e.ToolID, e.Err)
}

// Unwrap returns the underlying error for errors.Unwrap.
func (e *ToolError) Unwrap() error {
	return e.Err
}

// Is reports whether this error matches the target.
// It provides special handling to map toolindex.ErrNotFound to ErrToolNotFound.
func (e *ToolError) Is(target error) bool {
	// Map toolindex.ErrNotFound to our ErrToolNotFound
	if errors.Is(target, ErrToolNotFound) && errors.Is(e.Err, toolindex.ErrNotFound) {
		return true
	}

	// Standard error matching through Unwrap chain
	return errors.Is(e.Err, target)
}

// WrapError wraps an error with tool context.
// Returns nil if err is nil.
func WrapError(toolID string, backend *toolmodel.ToolBackend, op string, err error) error {
	if err == nil {
		return nil
	}
	return &ToolError{
		ToolID:  toolID,
		Backend: backend,
		Op:      op,
		Err:     err,
	}
}
