# Agent instructions

## Project map

- `golang/` contains the Go-language fluent API, Go formatter, and Go syntax policy in `golang/config.go`.
- `types/` defines the public `Token` interface.
- `internal/gen/` walks token values and applies language-neutral layout rules to semantic token styles.
- `internal/models/` implements reusable token renderers. `internal/config/` defines the language policy contract; language-specific rules belong in the corresponding language package, never as literals or branches in `internal/`.
- `README.md` documents installation and public API usage. Keep examples consistent with the constructors in `golang/`.

## Working in this repository

- Work from the repository root. The module path is `go.osspkg.com/gogen` and `go.mod` requires Go 1.26.
- Preserve the boundary between language packages and reusable `internal` packages. Keep keyword inventories, identifier rules, literal quoting, operator classification, and comment styles in the language package config.
- Extend `internal/config.Config` only with language-neutral capabilities needed by multiple language adapters. Do not add branches for Go (or another target language) to `internal/gen` or `internal/models`.
- Token implementations satisfy `types.Token` by rendering to an `io.Writer`; propagate writer and render errors.
- `golang.Render` formats output by default. `SetRawMode` and `SetDefaultMode` change package-wide rendering behavior; account for that shared state when changing rendering code.
- `ID` and `Op` validate against rules in the Go builder. Update focused examples or tests when changing those rules or exported constructors.

## Persistent project memory

Use the Chroma collection `chat_gogen_memory` for durable repository context.

- Before querying, adding, updating, or deleting memory, ensure the collection exists. Call `chroma_list_collections`; if it is missing, create exactly `chat_gogen_memory` with the default embedding configuration. Do not ask for permission to create it.
- For non-trivial implementation, debugging, architecture, API, or infrastructure work, query the collection with a concise semantic description of the task before making important decisions.
- If a query fails because the collection is missing, list collections, create it if needed, and retry once. For other memory failures, continue from repository evidence and do not invent retrieved information.
- After non-trivial work, store durable decisions or lessons that are likely to help future work. Query for related entries first; update an existing entry when refining a decision and add a new one only when it is distinct.
- Keep one concise, self-contained fact per document. Do not store transcripts, routine command output, facts directly recoverable from source or docs, speculation, or secrets.
- Treat memory as supplemental. Explicit instructions, current source and tests, and current documentation take precedence over it.

## Validation commands

Run these from the repository root:

- `make tests` runs the repository test task through `goppy`.
- `make lint` runs the configured lint task through `goppy`.
- `make build` runs the configured amd64 build through `goppy`.
- `make ci` is the GitHub Actions workflow and runs the pre-commit targets, including license, lint, tests, and build.

`make ci` also runs `make install`, which installs `goppy@latest` and invokes `goppy setup-lib`. This may require network access and affect local setup; inspect `git status` after running it. CI is configured for Go 1.26 in `.github/workflows/ci.yml`.
