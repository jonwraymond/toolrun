# Design Notes

This page documents the key tradeoffs and error semantics behind `toolrun`.

## Design tradeoffs

- **Backend-agnostic execution.** The runner dispatches to MCP, provider, or local backends through narrow executor interfaces. This keeps tool execution pluggable without embedding transport details.
- **Index-first resolution.** When configured, `toolindex` is the primary source for tool definitions and backends. Optional resolvers allow ad-hoc tools or dynamic backends without an index.
- **Validation on by default.** Input and output validation is enabled by default to catch schema errors early. This trades a small amount of latency for correctness and safety.
- **Strict chaining policy.** Chains stop on the first error (v1) and inject previous results at `args["previous"]` when requested. This is deterministic and easy to reason about, but not a branching workflow engine.
- **Structured-first results.** For MCP backends, `StructuredContent` is preferred; otherwise text content is best-effort parsed as JSON. This avoids losing structure while preserving fallback behavior.
- **Context-aware execution.** Runner methods check `context.Context` and pass it to backends; full cancellation depends on backend support.
- **Optional progress callbacks.** `ProgressRunner` provides coarse progress updates for chains and long-running tools without changing the core `Runner` API.

## Error semantics

`toolrun` uses sentinel errors and wrapped context via `ToolError`:

- `ErrToolNotFound` – no tool definition could be resolved.
- `ErrNoBackends` – tool exists but no backend is available.
- `ErrValidation` – input validation failed.
- `ErrExecution` – backend execution failed.
- `ErrOutputValidation` – output validation failed.
- `ErrStreamNotSupported` – streaming is not supported by the selected backend.

`ToolError` wraps these with `ToolID`, `Backend`, and `Op` (e.g., `resolve`, `execute`, `validate_input`).

## Extension points

- **Custom backend selection:** provide `BackendSelector` to prioritize specific backends.
- **Custom validation:** plug in a different `SchemaValidator` implementation.
- **Custom executors:** implement `MCPExecutor`, `ProviderExecutor`, or `LocalRegistry` for your runtime.

## Operational guidance

- Prefer registering tools in `toolindex` for consistent resolution and backends.
- Enable output validation in production to catch schema drift early.
- Use streaming only when your backend supports it; otherwise expect `ErrStreamNotSupported`.
