# toolrun

`toolrun` executes tools and chains. It resolves tools/backends, validates
inputs/outputs against JSON Schema, dispatches to the correct executor, and
normalizes results.

## Key APIs

- `Runner` interface
- `DefaultRunner` (via `NewRunner`)
- `Run`, `RunStream`, `RunChain`
- `ChainStep`, `RunResult`, `StreamEvent`

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
- Examples: `examples.md`
