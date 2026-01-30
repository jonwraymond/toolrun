# API Reference

## Runner interface

```go
type Runner interface {
  Run(ctx context.Context, toolID string, args map[string]any) (RunResult, error)
  RunStream(ctx context.Context, toolID string, args map[string]any) (<-chan StreamEvent, error)
  RunChain(ctx context.Context, steps []ChainStep) (RunResult, []StepResult, error)
}

// Optional progress interface
type ProgressRunner interface {
  RunWithProgress(ctx context.Context, toolID string, args map[string]any, onProgress ProgressCallback) (RunResult, error)
  RunChainWithProgress(ctx context.Context, steps []ChainStep, onProgress ProgressCallback) (RunResult, []StepResult, error)
}

type ProgressCallback func(ProgressEvent)
```

### Runner contract

- Concurrency: implementations are safe for concurrent use.
- Context: methods honor cancellation/deadlines and return `ctx.Err()` when canceled.
- Errors: use `errors.Is` with `ErrInvalidToolID`, `ErrToolNotFound`, `ErrValidation`,
  `ErrExecution`, `ErrOutputValidation`, and `ErrStreamNotSupported`.
- Ownership: args are read-only; returned results are caller-owned snapshots.
- Determinism: identical inputs over stable backends should yield stable results.

## Config

```go
type Config struct {
  Index           toolindex.Index
  ToolResolver    func(id string) (*toolmodel.Tool, error)
  BackendsResolver func(id string) ([]toolmodel.ToolBackend, error)
  BackendSelector toolindex.BackendSelector
  Validator       toolmodel.SchemaValidator
  ValidateInput   bool
  ValidateOutput  bool
  MCP      MCPExecutor
  Provider ProviderExecutor
  Local    LocalRegistry
}
```

## Executors

```go
type MCPExecutor interface {
  CallTool(ctx context.Context, serverName string, params *mcp.CallToolParams) (*mcp.CallToolResult, error)
  CallToolStream(ctx context.Context, serverName string, params *mcp.CallToolParams) (<-chan StreamEvent, error)
}

type ProviderExecutor interface {
  CallTool(ctx context.Context, providerID, toolID string, args map[string]any) (any, error)
  CallToolStream(ctx context.Context, providerID, toolID string, args map[string]any) (<-chan StreamEvent, error)
}

type LocalRegistry interface {
  Get(name string) (LocalHandler, bool)
}
```

### Executor contracts

- MCPExecutor/ProviderExecutor must return `ErrStreamNotSupported` for unsupported streaming.
- If streaming is supported and error is nil, the returned channel must be non-nil.
- LocalRegistry must return `(nil, false)` for unknown names.

## Results

```go
type ChainStep struct {
  ToolID      string
  Args        map[string]any
  UsePrevious bool
}

type RunResult struct {
  Tool       toolmodel.Tool
  Backend    toolmodel.ToolBackend
  Structured any
  MCPResult  *mcp.CallToolResult
}
```

## Streaming

```go
type StreamEvent struct {
  Kind  StreamEventKind
  ToolID string
  Data  any
  Err   error
}
```

## Errors

- `ErrToolNotFound`
- `ErrInvalidToolID`
- `ErrNoBackends`
- `ErrValidation`
- `ErrOutputValidation`
- `ErrExecution`
- `ErrStreamNotSupported`
