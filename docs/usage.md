# Usage

## Configure a runner

```go
runner := toolrun.NewRunner(
  toolrun.WithIndex(idx),
  toolrun.WithMCPExecutor(mcpExec),
  toolrun.WithProviderExecutor(providerExec),
  toolrun.WithLocalRegistry(localRegistry),
)
```

## Validation control

```go
runner := toolrun.NewRunner(
  toolrun.WithIndex(idx),
  toolrun.WithValidation(true, false), // validate input, skip output
)
```

## Run a tool

```go
res, err := runner.Run(ctx, "github:get_repo", map[string]any{
  "owner": "octo",
  "repo":  "hello",
})
```

## Run a chain

```go
steps := []toolrun.ChainStep{
  {ToolID: "user:get", Args: map[string]any{"user_id": "123"}},
  {ToolID: "orders:list", UsePrevious: true},
}

final, all, err := runner.RunChain(ctx, steps)
```
