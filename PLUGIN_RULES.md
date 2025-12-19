# Planx Plugin Development Rules (Go)

## What you write

- Implement ONE plugin type:
  - Source OR Sink OR Processor
- Write business logic only

## What you never write

- gRPC / networking
- Session lifecycle
- Flow control
- Threading model

## Entry Point

`main.go` MUST only call one SDK function:

- sdk.ServeSource(...)
- sdk.ServeSink(...)
- sdk.ServeProcessor(...)

## Structure Rules

- All logic under internal/plugin/
- No shared code unless absolutely required
