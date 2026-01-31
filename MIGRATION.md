# Migration Guide: toolrun to toolexec/run

This guide covers migrating from the deprecated `toolrun` package to `toolexec/run`.

## Import Path Changes

Update your imports as follows:

| Old Import | New Import |
|------------|------------|
| `github.com/jonwraymond/toolrun` | `github.com/jonwraymond/toolexec/run` |

## Steps

1. **Update go.mod**

   Remove the old dependency:
   ```bash
   go mod edit -droprequire github.com/jonwraymond/toolrun
   ```

   Add the new dependency:
   ```bash
   go get github.com/jonwraymond/toolexec
   ```

2. **Update imports**

   Replace all occurrences in your codebase:
   ```bash
   find . -name "*.go" -exec sed -i '' 's|github.com/jonwraymond/toolrun|github.com/jonwraymond/toolexec/run|g' {} +
   ```

   Or manually update each file:
   ```go
   // Before
   import "github.com/jonwraymond/toolrun"

   // After
   import "github.com/jonwraymond/toolexec/run"
   ```

3. **Update package references**

   If you referenced the package name directly:
   ```go
   // Before
   runner := toolrun.NewRunner(...)

   // After
   runner := run.NewRunner(...)
   ```

4. **Run tests**

   ```bash
   go mod tidy
   go test ./...
   ```

## API Compatibility

The `toolexec/run` package maintains API compatibility with `toolrun`. The following types and functions are preserved:

- `NewRunner` and `Runner` type
- `WithIndex`, `WithLocalRegistry`, `WithToolResolver`, `WithBackendsResolver` options
- `Run`, `RunChain`, `RunStream` methods
- `RunResult`, `LocalHandler`, `LocalRegistry` types
- `ErrStreamNotSupported` error

## Questions

If you encounter issues during migration, please open an issue in the [toolexec repository](https://github.com/jonwraymond/toolexec/issues).
