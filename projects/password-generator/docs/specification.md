# Specifications

## **1. Context and Objectives**

**Target audience**: Developers, security-conscious users, and anyone needing robust password generation.
**Main objective**: Build a **TUI** (_Text User Interface_) application that generates random, customizable, and secure passwords with an intuitive interface and copy/export options. This project is also educational and aims to improve skills in [Go](https://go.dev/), [BubbleTea](https://github.com/charmbracelet/bubbletea), and optionally [LipGloss](https://github.com/charmbracelet/lipgloss).

### **Product Vision**
The product should provide a terminal password generator that is:
- **Simple**: usable within seconds with minimal learning curve.
- **Reliable**: based on technical decisions aligned with security usage.
- **Educational**: designed to make code, architecture, and BubbleTea concepts easy to understand.

### **MVP Objectives**
- Provide a simple and fast terminal interface.
- Generate robust passwords from configurable criteria.
- Provide immediate feedback on password quality.
- Establish a modular, testable codebase for future extensions.
- Serve as a practical learning project for Go and BubbleTea architecture.

### **User Problem**
Existing password generators are often either too basic or too feature-heavy for simple use cases. This project aims to provide a local, fast, terminal-oriented tool that lets users:
- generate a secure password without opening a browser;
- immediately understand active settings;
- quickly assess password robustness.

### **Target Personas**
- **Terminal developer**: wants to generate strong passwords quickly without leaving the terminal.
- **Security-aware user**: wants full control over generation rules and to avoid weak passwords.
- **Go/BubbleTea learner**: wants to understand how to structure a clean, testable, maintainable TUI app.

### **Out of MVP Scope**
- Remote account sync or cloud services.
- Long-term history persistence.
- Advanced encryption or full password vault management.
- Native integration with third-party password managers.

---

## **2. Core Features**

### **Base Features (MVP)**
| ID  | Feature                                | Description                                                                                     | Priority |
|-----|----------------------------------------|-------------------------------------------------------------------------------------------------|----------|
| F01 | Random password generation             | Generates a random password from criteria (length, character types).                           | High     |
| F02 | Criteria customization                 | Allows users to configure length and include uppercase, lowercase, digits, symbols.           | High     |
| F03 | Password display                       | Displays the generated password prominently in the interface.                                  | High     |
| F04 | Clipboard copy                         | Allows users to copy generated password to clipboard (if supported).                           | Medium   |
| F05 | Password strength evaluation           | Shows an estimated strength level (weak, medium, strong).                                     | Medium   |
| F06 | Batch generation                       | Generates several passwords in one action (for example: 5).                                   | Low      |
| F07 | File save                              | Saves generated passwords to a file (plain text or encrypted).                                | Low      |

### **Advanced Features (Future iterations)**
| ID  | Feature                                | Description                                                                                     |
|-----|----------------------------------------|-------------------------------------------------------------------------------------------------|
| F08 | Character exclusion mode               | Excludes specific characters (for example: `l`, `1`, `O`, `0`).                               |
| F09 | Passphrase generation                  | Generates passphrases (for example: `CorrectHorseBatteryStaple`).                              |
| F10 | Password history                       | Stores and displays generated passwords during the current session.                             |
| F11 | Visual themes                          | Allows changing interface color/themes.                                                         |
| F12 | Password manager integration           | Direct export to a password manager (for example: Bitwarden, KeePass).                         |

### **Scope for the First Version**
The first version should prioritize **F01**, **F02**, **F03**, and **F05**.

**F04** can be added if a reliable cross-platform library is selected without adding too much complexity. Features **F06** to **F12** are considered later evolutions.

### **MVP Product Value**
The MVP should allow a user to launch the app, configure criteria, generate a strong password, and immediately understand whether the result is acceptable. If this flow is not simple and smooth, the MVP is not considered successful.

### **MVP User Stories**
- As a user, I want to choose password length so I can match service constraints.
- As a user, I want to enable/disable character categories so I can control complexity.
- As a user, I want to generate a password in one action so I can get a usable result quickly.
- As a user, I want to see a strength indicator so I can assess quality immediately.
- As a user, I want a clear validation message for invalid settings so I can fix inputs without leaving the app.

### **Business Rules**
- Minimum password length must be defined and validated. Proposed: **8**.
- Maximum length should remain reasonable for TUI display. Proposed: **128**.
- At least one character category must be selected before generation.
- If multiple categories are selected, generated passwords must contain at least one character from each active category.
- Generation must rely exclusively on cryptographically secure randomness.
- Strength evaluation must be consistent with selected criteria and reproducible for identical inputs.

---

## **3. Technical Specifications**

### **Languages and Libraries**
- **Language**: Go **1.26.2**.
- **TUI**: [BubbleTea](https://github.com/charmbracelet/bubbletea) (interface) and [LipGloss](https://github.com/charmbracelet/lipgloss).
- **Random generation**: `crypto/rand` (secure generation).
- **Parsing/Validation**: standard library (`strings`, `unicode`).
- **Files**: standard library `encoding/json` or `os` (saving).
- **Clipboard**: library to be selected when feature is implemented.

### **Implementation Principles**
- Prefer explicit and readable code over compact but opaque implementations.
- Avoid premature complexity: a clean, simple first version is preferred.
- Isolate business logic from UI for testability and clarity.
- Document non-obvious design decisions in code or project docs.

### **Suggested Logical Architecture**
- **UI/TUI**: display, navigation, keyboard shortcuts.
- **Domain**: config structures, generation, strength evaluation.
- **Infrastructure**: clipboard, file saving, optional serialization.

This separation enables domain testing without terminal UI dependencies.

### **Recommended Project Structure**

```text
password-generator/
├── cmd/
│   └── password-generator/
│       └── main.go
├── internal/
│   ├── app/
│   │   ├── model.go
│   │   ├── view.go
│   │   └── keymap.go
│   ├── domain/
│   │   ├── password/
│   │   │   ├── generator.go
│   │   │   ├── generator_test.go
│   │   │   ├── strength.go
│   │   │   └── strength_test.go
│   │   └── rules/
│   │       ├── validator.go
│   │       └── validator_test.go
│   ├── config/
│   │   └── defaults.go
│   └── infra/
│       └── clipboard/
│           ├── clipboard.go
│           └── clipboard_stub.go
├── docs/
│   ├── decisions.md
│   ├── mvp-tracking.md
│   ├── specifications_fr.md
│   └── specification.md
├── go.mod
├── go.sum
├── README.md
└── README_fr.md
```

### **TDD Development Approach**
- Follow **TDD** whenever possible for business logic.
- For each important rule: write a failing test first, then implement the minimum to pass.
- Focus deep testing on domain behavior: config validation, generation rules, category presence, strength logic.
- Keep TUI as an orchestration layer over tested domain code.

### **Educational Constraints**
- The project explicitly targets skill growth in Go and BubbleTea.
- Implementation choices must remain easy to review and learn from.
- Preferred contribution modes:
  - either the project owner writes the code directly;
  - or external contributions remain explicit and easy to explain.
- Avoid unnecessary abstractions, implicit shortcuts, and excessive dependencies.
- Each significant step should answer: what was done, why it was done, and how it works.

### **Non-Functional Requirements**
- **Portability**: Must work on Windows, Linux, macOS.
- **Security**: Use `crypto/rand` for generation (no `math/rand`).
- **Extensibility**: Modular code that supports future features.
- **Tests**: Unit tests for generation and strength logic.
- **Performance**: Single-password generation should feel instantaneous to users.
- **UX**: Main actions must be keyboard-accessible.
- **Understandability**: Code must stay readable for learners.

### **Technical Validation Criteria**
- Project builds successfully on at least one target environment.
- Unit tests cover nominal and validation-error scenarios.
- No `math/rand` usage in password generation path.
- User-facing error messages are clear and non-blocking.
- Core domain logic can be explained clearly in writing or verbally.

---

## **4. Mockups and User Flow**

### **Main Flow**
1. User launches the app.
2. Interface displays password criteria options.
3. User validates settings.
4. Password is generated and displayed.
5. User can:
   - copy password;
   - save password;
   - generate another password;
   - quit app.

### **Target User Journey**
- User understands the main screen without external documentation.
- User quickly identifies editable parameters and primary action.
- User gets a result in a small number of actions.
- User clearly understands whether generated password is weak/medium/strong.
- User can restart or quit without confusion.

### **Expected Interface Behavior**
- Keyboard navigation between fields.
- Clear generate shortcut (for example: `Enter` or `g`).
- Clear exit shortcut (`q`, optionally `Ctrl+C`).
- Invalid settings display readable errors without exiting the app.
- Password and strength refresh immediately after generation.

### **Product UX Principles**
- Emphasize the primary action: generate.
- Reduce cognitive load by showing only useful information.
- Keep vocabulary simple and consistent across labels.
- Make current configuration state explicit.
- Avoid ambiguity between automatic and manual actions.

### **Example TUI Mockup**
```text
+-------------------------------------+
|         PASSWORD GENERATOR          |
+-------------------------------------+
| Length      : [12]                  |
| Uppercase   : [x]                   |
| Lowercase   : [x]                   |
| Digits      : [x]                   |
| Symbols     : [x]                   |
|                                     |
| [Generate]   [Quit]                 |
+-------------------------------------+
| Password    : 7H#k9Lm$2pQ!          |
| Strength    : Strong                |
|                                     |
| [Copy]       [Save]                 |
+-------------------------------------+
```

---

## **5. MVP Acceptance Criteria**

### **Functional**
- User can define password length within valid range.
- User can enable/disable character categories.
- A valid password is generated on request and displayed.
- Strength is recalculated and displayed after each generation.
- Invalid configuration blocks generation with a clear message.

### **Technical**
- Generation logic is covered by automated tests.
- Strength logic is covered by automated tests.
- UI code remains decoupled from domain logic.

### **Ergonomics**
- Application is fully keyboard-usable.
- Main labels are understandable without external docs.
- Display remains readable in standard terminal size.
- Main generation flow is understandable in under one minute.

### **Pedagogy**
- Main components are easy to explain: config, generation, strength, TUI.
- A new contributor can understand project structure without deep BubbleTea expertise.
- Tests also act as executable behavior documentation.

---

## **6. Open Questions**
- Which clipboard library should be selected for consistent cross-platform behavior?
- Should clipboard support be included in MVP or deferred if cross-platform cost is too high?
- Which strength method should be used: simple heuristics or entropy-style estimation?
- Should save format be plain text, JSON, or deferred outside MVP?

---

## **7. Success Indicators**
- MVP generates valid passwords in only a few interactions.
- Project remains simple enough to preserve educational value.
- Tests actively guide implementation instead of being added only at the end.
- Code structure makes adding copy/save/history features straightforward.