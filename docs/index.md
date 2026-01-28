# toolrun

Execution and chaining layer for tools.

## What this repo provides

- Run a tool by ID
- Run tool chains with step results
- Schema validation hooks

## Example

```go
result, _ := runner.Run(ctx, "github:get_repo", args)
```
