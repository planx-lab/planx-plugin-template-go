# Planx Connector Template (Go)

This repository is a template for Planx 4.0 Go connectors. A connector is one
self-describing plugin binary that may expose multiple components (Source,
Processor, Sink) — see ADR-008/009.

## How it works

- `cmd/plugin/main.go` declares `sdk.Serve(sdk.Plugin{Components: [...]})`
- The SDK starts the gRPC server; the binary describes itself via `Discover`
- The Engine launches the plugin as a process and calls the lifecycle RPCs
- Each component under `internal/{source,processor,sink}/` only implements
  business logic (`sdk.*SPI`)

## Run locally

```bash
go run ./cmd/plugin
```

The plugin will start and print a handshake JSON to STDOUT.

## Specification Authority

The authoritative specification for Planx 4.0 lives in the [planx-spec](github.com/planx-lab/planx-spec) repository.
All behavior, formats, and contracts in this repository MUST conform to it. Local documentation must not redefine system contracts.