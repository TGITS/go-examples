# Générateur de mots de passe (TUI)

Générateur de mots de passe en terminal, écrit en Go, conçu comme projet d'apprentissage Go et BubbleTea.

## Objectifs
- Construire une application TUI claire et testable.
- Générer des mots de passe robustes avec options configurables.
- Suivre une démarche orientée TDD sur la logique métier.
- Garder une architecture explicite et pédagogique.

## Stack technique
- Go 1.26.2
- BubbleTea (TUI)
- crypto/rand (aléatoire sécurisé)

## Structure du projet
```text
cmd/password-generator/main.go
internal/app/*
internal/domain/password/*
internal/domain/rules/*
internal/config/defaults.go
internal/infra/clipboard/*
```

## Exécution
```bash
go run ./cmd/password-generator
```

## Tests
```bash
go test ./...
```

## Documentation Map
- Spécifications produit et techniques (EN) : [docs/specification.md](docs/specification.md)
- Spécifications produit et techniques (FR) : [docs/specifications_fr.md](docs/specifications_fr.md)
- Suivi d'implémentation MVP : [docs/mvp-tracking.md](docs/mvp-tracking.md)
- Journal des décisions d'architecture : [docs/decisions.md](docs/decisions.md)

## Notes
- Le support presse-papiers est volontairement en mode stub pour l'instant.
- Les tests de domaine sont prioritaires (validator/generator).