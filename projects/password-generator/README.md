# Password Generator (TUI)

Terminal password generator written in Go, designed as a learning project for Go and BubbleTea.

## Goals
- Build a clean and testable TUI application.
- Generate strong passwords with configurable options.
- Follow a TDD-oriented workflow for domain logic.
- Keep architecture explicit and educational.

## Tech Stack
- Go 1.26.2
- BubbleTea (TUI)
- crypto/rand (secure randomness)

## Project Structure
```text
cmd/password-generator/main.go
internal/app/*
internal/domain/password/*
internal/domain/rules/*
internal/config/defaults.go
internal/infra/clipboard/*
```

## Run
```bash
go run ./cmd/password-generator
```

## Test
```bash
go test ./...
```

## Documentation Map
- Product and technical specification (EN): [docs/specification.md](docs/specification.md)
- Product and technical specification (FR): [docs/specifications_fr.md](docs/specifications_fr.md)
- MVP implementation tracking: [docs/mvp-tracking.md](docs/mvp-tracking.md)
- Architecture decisions log: [docs/decisions.md](docs/decisions.md)

## Notes
- Clipboard support is intentionally stubbed for now.
- Domain tests are the first priority (validator/generator).