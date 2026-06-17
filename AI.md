# AI RULES — PLANX GO PLUGIN (v4)

## Authority Documents

1. [planx-architecture.md](../planx-architecture.md)
2. [planx-ai-guardrails.md](../planx-ai-guardrails.md)
3. [AI_CONTRACT.md](../AI_CONTRACT.md)
4. [planx-sdk-go/AI.md](../planx-sdk-go/AI.md) — SDK semantics

---

## SCOPE
This repository implements Connector business logic. A Connector MAY implement
multiple plugin types (Source, Processor, Sink) in a single repo.

---

## CONNECTOR HARD RULES

AI MUST NOT:
- Implement gRPC servers
- Manage sessions or flow control
- Start goroutines for concurrency
- Import planx-engine
- Import planx-proto directly
- Read from STDIN or write to STDOUT (except logging)

AI MUST:
- Place each plugin kind under its own directory:
  - `internal/source/` — Source SPI implementation
  - `internal/processor/` — Processor SPI implementation
  - `internal/sink/` — Sink SPI implementation
- Each `cmd/{kind}/main.go` calls exactly one SDK function:
  - `sdk.ServeSource(source.New)`
  - `sdk.ServeProcessor(processor.New)`
  - `sdk.ServeSink(sink.New)`
- Implement SPI interfaces from planx-sdk-go/spi
- Keep logic synchronous and deterministic
- Treat Batch as opaque bytes
- Provide one `manifest.yaml` per plugin kind in `manifests/`

If a requirement seems to need runtime logic:
STOP. That belongs to SDK.
