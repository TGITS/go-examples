# MVP Implementation Tracking

This document is the project tracking source for implementing the MVP with a TDD-first approach.

## 1) Project Goal

Build a terminal password generator MVP focused on:
- F01 Random password generation
- F02 Criteria customization
- F03 Password display
- F05 Password strength evaluation

Optional in MVP (only if low complexity):
- F04 Clipboard copy

## 2) Working Agreement

- Implementation is done manually by the project owner.
- AI assistance is used for planning, reviews, and checkpoints.
- Every meaningful feature starts with tests (Red -> Green -> Refactor).
- Keep code explicit and educational.

## 3) MVP Roadmap (Execution Order)

### Phase 0 - Baseline and Test Harness
Status: Done

- [x] Project scaffold created
- [x] Domain package for rules and generation exists
- [x] Initial unit tests are running
- [x] BubbleTea app entrypoint compiles

Exit criteria:
- `go test ./...` passes

### Phase 1 - Validation Rules (F02 foundation)
Status: In Progress

Scope:
- Validate password length bounds
- Validate at least one character category enabled
- Keep validation errors clear and user-friendly

Tasks:
- [ ] Add/confirm full test matrix for config validation
- [ ] Refine error messages for UI display
- [ ] Add edge-case tests for min/max boundaries

Exit criteria:
- Validation behavior is fully test-covered
- No ambiguity in invalid input handling

### Phase 2 - Secure Generator (F01)
Status: In Progress

Scope:
- Secure random generation using `crypto/rand`
- Guarantee at least one char from each enabled category
- Preserve exact requested length

Tasks:
- [ ] Add test: only selected categories are present
- [ ] Add test: repeated generations produce varied output
- [ ] Add test: invalid config fails fast
- [ ] Refactor generator internals for readability

Exit criteria:
- Generator tests are deterministic in expectations
- No use of `math/rand`

### Phase 3 - Strength Evaluation (F05)
Status: Not Started

Scope:
- Implement simple, consistent scoring heuristic
- Return clear levels: Weak / Medium / Strong

Tasks:
- [ ] Define explicit scoring rules in tests
- [ ] Add threshold boundary tests
- [ ] Align displayed labels with product wording

Exit criteria:
- Score rules are understandable and documented
- All thresholds are test-covered

### Phase 4 - BubbleTea MVP Flow (F03 + orchestration)
Status: Not Started

Scope:
- Connect UI actions to validated domain logic
- Generate and display password + strength
- Display non-blocking validation errors

Tasks:
- [ ] Add app-model tests for key inputs (`g`, `enter`, `q`)
- [ ] Show current config and generated output in view
- [ ] Ensure invalid config message is visible and recoverable

Exit criteria:
- Main user flow works end-to-end in terminal
- Basic interaction behavior is covered by tests

### Phase 5 - Optional Clipboard (F04)
Status: Not Started

Scope:
- Keep clipboard behind interface
- Activate only if platform support is simple and stable

Tasks:
- [ ] Evaluate package options (cross-platform)
- [ ] Add adapter implementation and fallback behavior
- [ ] Add tests around copy command integration

Exit criteria:
- Clipboard works or is safely feature-flagged/deferred

## 4) TDD Workflow Per Task

For each task, use this sequence:

1. Red
- Write one failing test for one behavior.
- Use behavior-oriented test names.

2. Green
- Implement the minimum code needed for the test to pass.
- Avoid broad refactors in this step.

3. Refactor
- Improve naming and structure without changing behavior.
- Keep tests green at all times.

4. Commit checkpoint
- Small commit with one coherent behavior change.
- Update this tracking document.

## 5) Definition of Done (DoD)

A task is done only if:
- [ ] Behavior is covered by automated tests
- [ ] `go test ./...` passes
- [ ] Naming and structure remain understandable
- [ ] No hidden side effects introduced
- [ ] Tracking document status updated

## 6) Current Backlog (Next 10 Actions)

1. Expand validation tests for boundaries and invalid combinations.
2. Confirm validator error messages to be UI-friendly.
3. Add generator test for selected-category-only output.
4. Add generator test for category inclusion guarantee.
5. Add generator test for repeated-run variability.
6. Finalize and document strength scoring thresholds.
7. Add strength threshold boundary tests.
8. Add BubbleTea model tests for `g` and `enter` generation paths.
9. Add BubbleTea model test for quit path `q`.
10. Add first end-to-end manual smoke checklist in README.

## 7) Weekly Progress Log

### Week of 2026-04-14

Completed:
- Initial architecture scaffold.
- Initial validator and generator tests.
- Initial BubbleTea wiring.

In progress:
- Tighten validation matrix.
- Expand generator behavior tests.

Risks:
- Strength scoring might need adjustment after first UI feedback.
- Clipboard complexity may exceed MVP simplicity target.

## 8) Status Snapshot

- Overall MVP progress: 25%
- Phase 0: Done
- Phase 1: In Progress
- Phase 2: In Progress
- Phase 3: Not Started
- Phase 4: Not Started
- Phase 5: Not Started
