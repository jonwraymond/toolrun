# Changelog

## Unreleased

### Added
- Context cancellation checks for `Run`, `RunChain`, and tool resolution.
- Tests covering cancellation behavior between chain steps and before execution.
- Progress callbacks via `ProgressRunner` with start/end and per-step updates.
