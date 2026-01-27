# toolrun

`toolrun` is the execution layer for MCP-style tools defined in
`github.com/jonwraymond/toolmodel` and resolved via
`github.com/jonwraymond/toolindex`.

It:
- Resolves a canonical tool ID to a tool definition and backends
- Validates inputs and outputs against JSON Schema (on by default)
- Dispatches across `mcp`, `provider`, and `local` backends
- Normalizes results into a consistent `RunResult`
- Supports sequential chains with explicit data passing

## Install

```bash
go get github.com/jonwraymond/toolrun
```

## Quick start

```go
runner := toolrun.NewRunner(
  toolrun.WithIndex(idx),
  toolrun.WithMCPExecutor(mcpExec),
)

result, err := runner.Run(ctx, "github:get_repo", map[string]any{
  "owner": "jonwraymond",
  "repo":  "toolmodel",
})
if err != nil {
  // errors include tool ID, backend kind, and operation
  log.Fatal(err)
}

fmt.Println(result.Structured)
```

## Chain semantics

`RunChain` executes steps in order and stops on the first error.

When `usePrevious` is true, the prior step's structured result is injected at
`args["previous"]` (even when the prior result is nil).

## Streaming contract

`RunStream` is executor-defined. If streaming is unsupported, it returns
`ErrStreamNotSupported`.

Executor contract:
- if `err == nil`, the returned channel must be non-nil
- events missing `ToolID` are stamped by the runner

## Alignment

- MCP protocol target: `2025-11-25` (via `toolmodel.MCPVersion`)
- Default backend selection: `toolindex.DefaultBackendSelector`
