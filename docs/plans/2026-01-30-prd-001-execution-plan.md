# PRD-001 Execution Plan — toolrun (TDD)

**Status:** Ready
**Date:** 2026-01-30
**PRD:** `2026-01-30-prd-001-interface-contracts.md`


## TDD Workflow (required)
1. Red — write failing contract tests
2. Red verification — run tests
3. Green — minimal code/doc changes
4. Green verification — run tests
5. Commit — one commit per task


## Tasks
### Task 0 — Inventory + contract outline
- Confirm interface list and method signatures.
- Draft explicit contract bullets for each interface.
- Update docs/plans/README.md with this PRD + plan.
### Task 1 — Contract tests (Red/Green)
- Add `*_contract_test.go` with tests for each interface listed below.
- Use stub implementations where needed.
### Task 2 — GoDoc contracts
- Add/expand GoDoc on each interface with explicit contract clauses (thread-safety, errors, context, ownership).
- Update README/design-notes if user-facing.
### Task 3 — Verification
- Run `go test ./...`
- Run linters if configured (golangci-lint / gosec).


## Test Skeletons (contract_test.go)
### Runner
```go
func TestRunner_Contract(t *testing.T) {
    // Methods:
    // - Run(ctx context.Context, toolID string, args map[string]any) (RunResult, error)
    // - RunStream(ctx context.Context, toolID string, args map[string]any) (<-chan StreamEvent, error)
    // - RunChain(ctx context.Context, steps []ChainStep) (RunResult, []StepResult, error)
    // Contract assertions:
    // - Concurrency guarantees documented and enforced
    // - Error semantics (types/wrapping) validated
    // - Context cancellation respected (if applicable)
    // - Deterministic ordering where required
    // - Nil/zero input handling specified
}
```
### ProgressRunner
```go
func TestProgressRunner_Contract(t *testing.T) {
    // Methods:
    // - RunWithProgress(ctx context.Context, toolID string, args map[string]any, onProgress ProgressCallback) (RunResult, error)
    // - RunChainWithProgress(ctx context.Context, steps []ChainStep, onProgress ProgressCallback) (RunResult, []StepResult, error)
    // Contract assertions:
    // - Concurrency guarantees documented and enforced
    // - Error semantics (types/wrapping) validated
    // - Context cancellation respected (if applicable)
    // - Deterministic ordering where required
    // - Nil/zero input handling specified
}
```
### MCPExecutor
```go
func TestMCPExecutor_Contract(t *testing.T) {
    // Methods:
    // - CallTool(ctx context.Context, serverName string, params *mcp.CallToolParams) (*mcp.CallToolResult, error)
    // - CallToolStream(ctx context.Context, serverName string, params *mcp.CallToolParams) (<-chan StreamEvent, error)
    // Contract assertions:
    // - Concurrency guarantees documented and enforced
    // - Error semantics (types/wrapping) validated
    // - Context cancellation respected (if applicable)
    // - Deterministic ordering where required
    // - Nil/zero input handling specified
}
```
### ProviderExecutor
```go
func TestProviderExecutor_Contract(t *testing.T) {
    // Methods:
    // - CallTool(ctx context.Context, providerID, toolID string, args map[string]any) (any, error)
    // - CallToolStream(ctx context.Context, providerID, toolID string, args map[string]any) (<-chan StreamEvent, error)
    // Contract assertions:
    // - Concurrency guarantees documented and enforced
    // - Error semantics (types/wrapping) validated
    // - Context cancellation respected (if applicable)
    // - Deterministic ordering where required
    // - Nil/zero input handling specified
}
```
### LocalRegistry
```go
func TestLocalRegistry_Contract(t *testing.T) {
    // Methods:
    // - Get(name string) (LocalHandler, bool)
    // Contract assertions:
    // - Concurrency guarantees documented and enforced
    // - Error semantics (types/wrapping) validated
    // - Context cancellation respected (if applicable)
    // - Deterministic ordering where required
    // - Nil/zero input handling specified
}
```
