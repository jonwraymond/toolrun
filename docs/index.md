# toolrun

`toolrun` executes tools and chains. It resolves tools/backends, validates
inputs/outputs against JSON Schema, dispatches to the correct executor, and
normalizes results.

## What this library provides

- `Runner` interface for run, stream, and chain execution
- Default runner with validation hooks
- Backend dispatch (mcp, provider, local)
- Consistent error wrapping

## Quickstart

```go
runner := toolrun.NewRunner(
  toolrun.WithIndex(idx),
  toolrun.WithMCPExecutor(mcpExec),
  toolrun.WithLocalRegistry(localRegistry),
)

res, _ := runner.Run(ctx, "github:get_repo", map[string]any{"owner": "o", "repo": "r"})
```

## Next

- Execution pipeline: `architecture.md`
- Configuration and options: `usage.md`
- Examples and chains: `examples.md`
