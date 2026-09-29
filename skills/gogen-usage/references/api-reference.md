# gogen Go API reference

Import the Go adapter as `gogen "go.osspkg.com/gogen/golang"`. Constructors return `*gogen.Tokens`; fluent methods append to a sequence and return it. `*Tokens` also implements `types.Token`.

## Choose a builder

| Need | API | Notes |
| --- | --- | --- |
| File/package header | `Package(name)` | Appends a package declaration and line break. |
| One import | `Import(alias, path)` | Both arguments are strings; the path is quoted as a Go string literal. |
| Grouped imports | `ImportBlock(specs...)` | Each spec is `Text(path)` or `ID(alias).Text(path)`. Do not include the `import` keyword. |
| One type declaration | `Type().ID(name).Join(type)` | Compose the type keyword, name, and underlying type. |
| Grouped type declarations | `TypeBlock(specs...)` | Each spec contains the name followed by its type, e.g. `ID("Name").String()`. Do not include the `type` keyword. |
| Struct/interface body | `Struct().Block(...)`, `Interface().Block(...)` | `Block` controls braces, indentation, and line breaks. |
| Struct field and tags | `Field(name, type, key, value, ...)` | Tag arguments are alternating key/value strings. An odd count is a render error. |
| Function declaration | `Func().ID(name).Bracket(params...).Block(body...)` | `Bracket` adds a comma-separated parenthesized list; append result types before `Block`. |
| Function call | `ID(name).Call(args...)` | For package selectors use `Pkg("fmt").ID("Println").Call(...)`. For a value or field selector, add the dot explicitly: `ID("user").Op(".").ID("Name")`. |
| Index/generic arguments | `ID(name).Index(index)`, `ID(name).TypeArgs(args...)` | Both emit square-bracket syntax. |
| Keyed composite element | `KeyValue(key, value)` | Use inside a composite literal, usually between `Op("{")` and `Op("}")`. |
| String literal | `Text(value)` | Escapes and quotes the value as a Go string literal. |
| Identifier | `ID(value)` | Checks the builder's accepted identifier form when rendered. |
| Source fragment | `Raw(source)` | Inserts text verbatim; use only when token-level builders are unsuitable. |
| Operator | `Op(operator)` | Accepts supported Go operator/punctuation strings and reports an error for unsupported values. |
| Common allocation | `New(type)`, `Make(type, length, capacity)`, `Append(dst, values...)` | `Make` emits capacity only if it is greater than length. |
| Common types | `Map(key, value)`, `Slice().Type()`, `Array(n).Type()` and primitive type constructors | `Array(n)` does not validate whether `n` is a legal Go array length. |
| Block/comment/newline | `Block(...)`, `Comment(text)`, `Line()` | Use explicit structure for layout; nesting a token sequence does not itself add a newline. |
| Custom token | `types.Token` | Implement `Render(io.Writer) error`; write source and propagate writer/render errors. |

## Complete constructor inventory

These package-level functions start token sequences:

- Declarations: `Package`, `Import`, `ImportBlock`, `Type`, `TypeBlock`, `Var`, `Const`.
- Functions and control flow: `Func`, `Return`, `Defer`, `Go`, `If`, `For`, `Switch`, `Select`, `Case`, `Default`.
- Calls and lists: `Call`, `Params`, `List`.
- Expressions: `ID`, `Pkg`, `Op`, `Raw`, `Text`, `Index`, `TypeArgs`, `KeyValue`.
- Composite values: `New`, `Make`, `Append`.
- Types and values: `Any`, `Nil`, `Chan`, `Interface`, `Struct`, `Slice`, `Array`, `Map`, `Bool`, `Byte`, `Rune`, `String`, `Error`, `Int`, `Int8`, `Int16`, `Int32`, `Int64`, `Uint`, `Uint8`, `Uint16`, `Uint32`, `Uint64`, `Uintptr`, `Float32`, `Float64`, `Complex64`, `Complex128`.
- Layout: `Block`, `Comment`, `Line`, `Field`.
- Rendering: `Render`, `SetRawMode`, `SetDefaultMode`.

Most constructors also have a same-named fluent method. Exceptions to remember:

- Fluent-only methods: `Bracket`, `Break`, `Continue`, `Else`, `ElseIf`, `Fallthrough`, `Goto`, `Range`, `Join`, `Unwrap`, and `(*Tokens).Render`.
- Package-level only: `Defer`, `Go`, `Switch`, `Select`, `Params`, `Render`, `SetRawMode`, and `SetDefaultMode`.
- `Tokens` is a slice-backed builder. `Unwrap()` returns the backing slice without copying it.

## Rendering choices

- `token.Render(w)` on `*Tokens` emits readable token layout and does not call `go/format`.
- `gogen.Render(w, token)` applies `go/format` by default and returns render, formatting, or writer errors. It does not compile or type-check the result.
- `SetRawMode()` disables formatting for subsequent package-level `Render` calls; `SetDefaultMode()` restores it. This mode is package-wide state.
- `Raw` always preserves its own text verbatim. Raw mode controls the formatter, not the contents of a Raw token.
