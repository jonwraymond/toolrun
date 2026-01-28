# Examples

## Local tool execution

```go
registry := toolrun.NewLocalRegistry()
registry.Register("ping", func(ctx context.Context, args map[string]any) (any, error) {
  return map[string]any{"ok": true}, nil
})

runner := toolrun.NewRunner(
  toolrun.WithIndex(idx),
  toolrun.WithLocalRegistry(registry),
)

res, _ := runner.Run(ctx, "local:ping", map[string]any{})
```

## Streaming

```go
ch, err := runner.RunStream(ctx, "logs:stream", map[string]any{"tail": 100})
if err != nil {
  // handle ErrStreamNotSupported
}
for ev := range ch {
  fmt.Println(ev.Data)
}
```

## Backend resolution fallback

```go
runner := toolrun.NewRunner(
  toolrun.WithToolResolver(func(id string) (*toolmodel.Tool, error) {
    // resolve from external store
    return &tool, nil
  }),
  toolrun.WithBackendsResolver(func(id string) ([]toolmodel.ToolBackend, error) {
    return []toolmodel.ToolBackend{backend}, nil
  }),
)
```
