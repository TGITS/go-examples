# Architecture Decisions

## Initial Decisions
- Keep domain logic in `internal/domain` and test it first (TDD).
- Keep UI orchestration in `internal/app`.
- Keep clipboard dependency abstracted behind an interface.
- Start with a minimal BubbleTea app and grow incrementally.

## Pending Decisions
- Clipboard concrete implementation and package selection.
- Strength scoring strategy (simple heuristic vs entropy-inspired approach).
- Save/export format for future MVP+ iterations.

## Next session

- Next session starts with: a first mini-sprint TDD-style (30-45 min)