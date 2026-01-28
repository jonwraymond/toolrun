# User Journey

This journey shows how `toolrun` executes tools and chains in a full end-to-end agent flow.

## End-to-end flow (stack view)

![Diagram](assets/diagrams/user-journey.svg)

## Step-by-step

1. **Resolve** the tool and its backends (via `toolindex` or resolver callbacks).
2. **Select** the backend (default priority: local > provider > MCP).
3. **Validate input** against JSON Schema.
4. **Dispatch** to the selected backend.
5. **Normalize output** (structured content preferred).
6. **Validate output** (optional).

## Example: run a tool

```go
runner := toolrun.NewRunner(
  toolrun.WithIndex(idx),
  toolrun.WithMCPExecutor(mcpExec),
)

res, err := runner.Run(ctx, "github:get_repo", map[string]any{
  "owner": "acme",
  "repo":  "app",
})
```

## Example: chain two tools

```go
steps := []toolrun.ChainStep{
  {ToolID: "github:get_repo", Args: map[string]any{"owner": "acme", "repo": "app"}},
  {ToolID: "github:list_issues", UsePrevious: true},
}

final, stepsOut, err := runner.RunChain(ctx, steps)
```

## Expected outcomes

- Consistent execution across backends.
- Structured results suitable for chaining.
- Explicit error classification with `ToolError`.

## Common failure modes

- `ErrToolNotFound` if the tool is missing.
- `ErrValidation` for schema violations.
- `ErrExecution` for backend errors.
- `ErrStreamNotSupported` when streaming is requested but unsupported.
