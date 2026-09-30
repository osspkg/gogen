# gogen examples

These examples exercise real package APIs. Run either standalone program from the repository root:

```sh
go run ./skills/gogen-usage/examples/declarations
go run ./skills/gogen-usage/examples/expressions
```

## Build a file with grouped declarations and struct tags

See [`examples/declarations/main.go`](../examples/declarations/main.go). It demonstrates grouped imports and type declarations, struct fields with tags, a keyed composite literal, a package selector, and formatted output.

## Compose an expression and compare render modes

See [`examples/expressions/main.go`](../examples/expressions/main.go). It demonstrates generic type arguments and indexing with `Tokens.Render`, then renders a complete function with `golang.Render`.

## API usage patterns

- Prefer `Text` for Go string literals, including import paths; it handles quoting.
- Prefer `ID` for identifier text and `Op` for supported operators. These builders validate when rendered.
- Use `KeyValue` within composite literal braces; use `Index` for an expression like `items[i]`, and `TypeArgs` for a generic instantiation like `Set[string]`.
- Use `Block` for nested statement or type bodies. Use `Line` when composing top-level declarations manually.
- Keep arbitrary `Raw` fragments small and syntactically complete. `Raw` bypasses escaping and validation, although `golang.Render` still formats the combined output by default.
