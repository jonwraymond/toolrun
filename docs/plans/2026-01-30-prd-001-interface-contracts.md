# PRD-001 Interface Contracts — toolrun

**Status:** Draft
**Date:** 2026-01-30


## Overview
Define explicit interface contracts (GoDoc + documented semantics) for all interfaces in this repo. Contracts must state concurrency guarantees, error semantics, ownership of inputs/outputs, and context handling.


## Goals
- Every interface has explicit GoDoc describing behavioral contract.
- Contract behavior is codified in tests (contract tests).
- Docs/README updated where behavior is user-facing.


## Non-Goals
- No API shape changes unless required to satisfy the contract tests.
- No new features beyond contract clarity and tests.


## Interface Inventory
| Interface | File | Methods |
| --- | --- | --- |
| `Runner` | `toolrun/runner.go:8` | Run(ctx context.Context, toolID string, args map[string]any) (RunResult, error)<br/>RunStream(ctx context.Context, toolID string, args map[string]any) (<-chan StreamEvent, error)<br/>RunChain(ctx context.Context, steps []ChainStep) (RunResult, []StepResult, error) |
| `ProgressRunner` | `toolrun/runner.go:34` | RunWithProgress(ctx context.Context, toolID string, args map[string]any, onProgress ProgressCallback) (RunResult, error)<br/>RunChainWithProgress(ctx context.Context, steps []ChainStep, onProgress ProgressCallback) (RunResult, []StepResult, error) |
| `MCPExecutor` | `toolrun/executor.go:11` | CallTool(ctx context.Context, serverName string, params *mcp.CallToolParams) (*mcp.CallToolResult, error)<br/>CallToolStream(ctx context.Context, serverName string, params *mcp.CallToolParams) (<-chan StreamEvent, error) |
| `ProviderExecutor` | `toolrun/executor.go:23` | CallTool(ctx context.Context, providerID, toolID string, args map[string]any) (any, error)<br/>CallToolStream(ctx context.Context, providerID, toolID string, args map[string]any) (<-chan StreamEvent, error) |
| `LocalRegistry` | `toolrun/executor.go:41` | Get(name string) (LocalHandler, bool) |

## Contract Template (apply per interface)
- **Thread-safety:** explicitly state if safe for concurrent use.
- **Context:** cancellation/deadline handling (if context is a parameter).
- **Errors:** classification, retryability, and wrapping expectations.
- **Ownership:** who owns/allocates inputs/outputs; mutation expectations.
- **Determinism/order:** ordering guarantees for returned slices/maps/streams.
- **Nil/zero handling:** behavior for nil inputs or empty values.


## Acceptance Criteria
- All interfaces have GoDoc with explicit behavioral contract.
- Contract tests exist and pass.
- No interface contract contradictions across repos.

