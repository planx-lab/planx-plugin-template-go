# AI RULES – PLANX GO PLUGIN (v4)

You are writing a Planx 4.0 Go plugin.

This repository implements ONLY plugin business logic.

---------------------------------

YOU MUST NOT:
- Implement gRPC servers
- Manage sessions or flow control
- Start goroutines for concurrency
- Import planx-engine
- Import planx-proto directly
- Read from STDIN or write to STDOUT (except logging)

YOU MUST:
- Implement SPI interfaces from planx-sdk-go/sdk
- Keep logic synchronous and deterministic
- Treat Batch as opaque

If a requirement seems to need runtime logic:
STOP. That belongs to SDK.
