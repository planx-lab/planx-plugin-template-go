# Planx Plugin Template (Go)

This repository is a template for Planx 4.0 Go plugins.

## How it works

- SDK starts gRPC server
- Engine launches plugin as a process
- Plugin only implements business logic

## Run locally

```bash
go run ./cmd/plugin
```

The plugin will start and print a handshake JSON to STDOUT.

## Specification Authority

The authoritative specification for Planx 4.0 lives in the [planx-spec](github.com/planx-lab/planx-spec) repository.
All behavior, formats, and contracts in this repository MUST conform to it. Local documentation must not redefine system contracts.