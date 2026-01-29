package toolrun

import "context"

// Runner is the main execution interface for running tools.
// It provides methods for single tool execution, streaming execution,
// and sequential chain execution.
type Runner interface {
	// Run executes a single tool and returns the normalized result.
	// It resolves the tool, validates input, executes via the appropriate backend,
	// normalizes the result, and validates output.
	Run(ctx context.Context, toolID string, args map[string]any) (RunResult, error)

	// RunStream executes a tool with streaming support.
	// Returns a channel that receives streaming events.
	// May return ErrStreamNotSupported if the backend doesn't support streaming.
	RunStream(ctx context.Context, toolID string, args map[string]any) (<-chan StreamEvent, error)

	// RunChain executes a sequence of tool steps.
	// Returns the final result and a slice of step results.
	// Stops on the first error (v1 policy).
	// If UsePrevious is true for a step, the previous step's Structured result
	// is injected at args["previous"], overwriting any existing value,
	// even when the previous result is nil.
	RunChain(ctx context.Context, steps []ChainStep) (RunResult, []StepResult, error)
}

// ProgressCallback receives progress updates during execution.
// Implementations should be fast and non-blocking.
type ProgressCallback func(ProgressEvent)

// ProgressRunner is an optional interface that provides progress callbacks
// for long-running tool executions and chains.
type ProgressRunner interface {
	// RunWithProgress executes a single tool and emits progress updates.
	RunWithProgress(ctx context.Context, toolID string, args map[string]any, onProgress ProgressCallback) (RunResult, error)

	// RunChainWithProgress executes a chain and emits progress updates.
	RunChainWithProgress(ctx context.Context, steps []ChainStep, onProgress ProgressCallback) (RunResult, []StepResult, error)
}
