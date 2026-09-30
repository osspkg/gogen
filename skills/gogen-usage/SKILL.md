---
name: gogen-usage
description: Use the go.osspkg.com/gogen/golang token builder to generate, compose, render, or troubleshoot Go source in projects that depend on gogen.
---

# Using gogen

Use this skill when writing Go code that builds Go source with `go.osspkg.com/gogen/golang`. It covers the public token API and its rendering behavior; it does not describe changes to gogen's internal implementation.

## Working rules

- Start with package-level constructors such as `Package`, `Func`, `ID`, `Text`, and `Block`. They return `*Tokens`; fluent methods append tokens and return the sequence for chaining. Use `Join` to append existing token sequences.
- Confirm a method or constructor exists before using it. The package-level functions and `*Tokens` methods overlap but are not identical. See [API reference](references/api-reference.md) for the complete lists and method-only/package-only cases.
- Use typed builders for identifiers, string literals, operators, calls, indexes, keyed elements, and common Go types. Reserve `Raw` for source fragments that must be inserted verbatim; it does not validate or escape its content.
- Build multi-line syntax with `Block`, `Line`, `ImportBlock`, and `TypeBlock`. Nested token sequences are flattened; nesting alone does not create line breaks.
- Use `Tokens.Render(w)` when readable pre-format output or a source fragment is needed. Use package-level `golang.Render(w, token)` for complete Go source; it applies `go/format` by default. Formatting checks syntax, not types.
- `SetRawMode` and `SetDefaultMode` change package-wide render behavior. Avoid toggling these modes around independent or concurrent rendering flows. Raw token contents remain verbatim in either mode.
- Propagate render errors. `ID` and `Op` validate their inputs during rendering; `Field` returns an error for an odd number of tag strings.

## References

- Read [API reference](references/api-reference.md) to choose the correct constructor, fluent method, or rendering path.
- Read [examples](references/examples.md) when composing declarations, struct tags, or expressions. Runnable programs live under `examples/`.
- Keep the API reference and examples aligned with the current package docs; confirm signatures with `go doc ./golang` if the library changes.
