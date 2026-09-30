# Python 3.10+ code builder plan

## Goal

Provide `go.osspkg.com/gogen/python`, a composable Python source builder alongside the Go and TypeScript adapters. The adapter writes readable source directly and keeps Python-specific syntax policy outside shared internals.

## Work breakdown

- [x] Add generic indentation-prefix support to `internal/gen` while preserving tab-based `Indented` behavior.
- [x] Add Python token configuration for operators, identifiers, keywords, comments, strings, and expression boundaries.
- [x] Add `Tokens`, package-level and fluent constructors, `Render`, `Join`, and `Unwrap`.
- [x] Add import, alias, function, async function, class, decorator, annotation, and suite builders.
- [x] Add control flow, clauses, statements, operators, calls, indexing, common literals, comprehensions, and f-strings.
- [x] Add golden tests for readable output, nested and empty suites, clause alignment, escaping, validation, and writer errors.
- [x] Add package documentation, runnable examples, and a Python usage guide.
- [x] Link the Python guide and installation path from the root README.
- [x] Run `go test ./...`, `go vet ./...`, `make lint`, and `git diff --check`; fix any findings.

## Design constraints

- Python syntax and lexical details stay in `python/`; shared `internal` changes remain language-neutral.
- Suites use four spaces and empty suites emit `pass`.
- Rendering does not invoke a formatter or interpreter.
- `Raw` remains literal source input; callers own its validity.
- Target syntax is Python 3.10+, including structural pattern matching and `|` type unions.
