# toolrun

`toolrun` executes tools and chains. It resolves tools/backends, validates
inputs/outputs against JSON Schema, dispatches to the correct executor, and
normalizes results.

[![Docs](https://img.shields.io/badge/docs-ai--tools--stack-blue)](https://jonwraymond.github.io/ai-tools-stack/)

## Deep dives
- Design Notes: [design-notes.md](design-notes.md)
- User Journey: [user-journey.md](user-journey.md)

## Motivation

- **Consistent execution** across MCP, provider, and local tools
- **Safety** via schema validation and normalized results
- **Usability** with clear error wrapping and chain semantics

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

## Usability notes

- Backend selection is deterministic by default
- Chain steps can inject `previous` results
- Streaming is optional; non-streaming backends return `ErrStreamNotSupported`

## Next

- Execution pipeline: `architecture.md`
- Configuration and options: `usage.md`
- Examples: `examples.md`
- Design Notes: [design-notes.md](design-notes.md)
- User Journey: [user-journey.md](user-journey.md)

