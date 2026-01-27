# toolrun

`toolrun` is the execution layer for tools defined in `toolmodel` and resolved
via `toolindex`.

It:
- resolves a canonical tool ID to a tool definition + backends,
- validates inputs and outputs against JSON Schema (on by default),
- dispatches across `mcp`, `provider`, and `local` backends,
- normalizes results into a consistent `RunResult`, and
- supports sequential chains with explicit data passing.

## Install

```bash
go get github.com/jonwraymond/toolrun
```

## Quick start

Define a tool, register it in `toolindex`, and execute it locally:

```go
import (
  "context"
  "fmt"
  "log"

  "github.com/jonwraymond/toolindex"
  "github.com/jonwraymond/toolmodel"
  "github.com/jonwraymond/toolrun"
  "github.com/modelcontextprotocol/go-sdk/mcp"
)

// Minimal LocalRegistry implementation.
type localMap map[string]toolrun.LocalHandler
func (m localMap) Get(name string) (toolrun.LocalHandler, bool) {
  h, ok := m[name]
  return h, ok
}

idx := toolindex.NewInMemoryIndex()

t := toolmodel.Tool{
  Namespace: "math",
  Tool: mcp.Tool{
    Name:        "add",
    Description: "Add two integers",
    InputSchema: map[string]any{
      "type": "object",
      "properties": map[string]any{
        "a": {"type": "integer"},
        "b": {"type": "integer"},
      },
      "required": []string{"a", "b"},
    },
    OutputSchema: map[string]any{
      "type": "object",
      "properties": map[string]any{
        "sum": {"type": "integer"},
      },
      "required": []string{"sum"},
    },
  },
}

backend := toolmodel.ToolBackend{
  Kind:  toolmodel.BackendKindLocal,
  Local: &toolmodel.LocalBackend{Name: "math.add"},
}
_ = idx.RegisterTool(t, backend)

locals := localMap{
  "math.add": func(ctx context.Context, args map[string]any) (any, error) {
    a := args["a"].(int)
    b := args["b"].(int)
    return map[string]any{"sum": a + b}, nil
  },
}

runner := toolrun.NewRunner(
  toolrun.WithIndex(idx),
  toolrun.WithLocalRegistry(locals),
)

result, err := runner.Run(ctx, "math:add", map[string]any{
  "a": 2,
  "b": 3,
})
if err != nil {
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

## Integration notes

- Backend selection defaults to `toolindex.DefaultBackendSelector`
  (`local > provider > mcp`).
- Resolution can be injected via `WithToolResolver` and `WithBackendsResolver`
  when you do not want a hard dependency on `toolindex`.

## Version compatibility (current tags)

- `toolmodel`: `v0.1.0`
- `toolindex`: `v0.1.2`
- `tooldocs`: `v0.1.2`
- `toolrun`: `v0.1.1`
- `toolcode`: `v0.1.1`
- `toolruntime`: `v0.1.1`
- `toolsearch`: `v0.1.1`
- `metatools-mcp`: `v0.1.4`

MCP protocol target: `2025-11-25` (via `toolmodel.MCPVersion`).
