# Go code builder

The `go.osspkg.com/gogen/golang` package builds Go source from composable tokens. Package-level constructors start token sequences; fluent `*Tokens` methods append declarations, expressions, and statements.

```sh
go get go.osspkg.com/gogen/golang
```

## Build and render a Go file

The following builds a small source file with grouped imports and type declarations:

```go
package main

import (
	"bytes"
	"fmt"

	gogen "go.osspkg.com/gogen/golang"
)

func main() {
	file := gogen.Package("main").
		ImportBlock(
			gogen.Text("fmt"),
			gogen.ID("json").Text("encoding/json"),
		).
		TypeBlock(
			gogen.ID("Name").String(),
			gogen.ID("Count").Int(),
		).
		Join(
			gogen.Func().ID("main").Bracket().Block(
				gogen.Pkg("fmt").ID("Println").Call(gogen.Text("generated")),
			),
		)

	var source bytes.Buffer
	if err := gogen.Render(&source, file); err != nil {
		panic(err)
	}
	fmt.Print(source.String())
}
```

Output:

```go
package main

import (
	json "encoding/json"
	"fmt"
)

type (
	Name  string
	Count int
)

func main() {
	fmt.Println("generated")
}
```

The `ImportBlock` specs are import paths, optionally preceded by an alias. `TypeBlock` specs contain a type name followed by its underlying type. Both builders add the grouping syntax and layout.

## Struct fields and tags

`Field(name, type, tags...)` adds a named struct field. Tag arguments are alternating key/value strings; the builder quotes and escapes values for a Go struct tag:

```go
user := gogen.Package("main").Join(
	gogen.Type().ID("User").Struct().Block(
		gogen.Field("ID", gogen.Uint64(), "json", "id,omitempty", "db", "user_id"),
		gogen.Field("Name", gogen.String(), "json", "name"),
	),
)
```

This generates:

```go
package main

type User struct {
	ID   uint64 `json:"id,omitempty" db:"user_id"`
	Name string `json:"name"`
}
```

An odd number of tag arguments causes rendering to return an error.

## Expressions and generic calls

Use `Text` for escaped string literals, `ID` for identifiers, `Index` for indexed access, and `TypeArgs` for generic type arguments:

```go
expression := gogen.ID("lookup").TypeArgs(gogen.String()).Call(
	gogen.ID("items").Index(gogen.Raw("0")),
)
```

Output:

```go
lookup[string](items[0])
```

`Call` adds call parentheses, `Bracket` adds parentheses to a sequence, and `List` creates a comma-separated sequence without delimiters. `KeyValue` creates a keyed element in a composite literal.

## Builder reference

Most builders are available both as package-level constructors and as chainable `*Tokens` methods. A few are intentionally available in only one form; those are noted below.

| Area | Builders | Notes |
| --- | --- | --- |
| File and declarations | `Package`, `Import`, `ImportBlock`, `Type`, `TypeBlock`, `Var`, `Const` | Start a package, add imports, and declare values or types. Group declaration specs omit the `import` or `type` keyword. |
| Functions and statements | `Func`, `Return`, `Defer`, `Go`, `If`, `Else`, `ElseIf`, `For`, `Range`, `Switch`, `Select`, `Case`, `Default`, `Break`, `Continue`, `Fallthrough`, `Goto` | Compose functions and control flow. `Defer`, `Go`, `Switch`, and `Select` are package-level only; `Range`, `Else`, `ElseIf`, `Break`, `Continue`, `Fallthrough`, and `Goto` are fluent-only. |
| Calls and lists | `Call`, `Bracket`, `Params`, `List` | Build calls, parenthesized expressions, parameter lists, or comma-separated values. `Params` is package-level only. |
| Expressions | `ID`, `Pkg`, `Op`, `Raw`, `Text`, `Index`, `TypeArgs`, `KeyValue` | Build identifiers, package selectors, operators, literals, indexes, generic arguments, and keyed values. |
| Composite values and allocation | `Struct`, `Field`, `New`, `Make`, `Append` | Build structs, tagged fields, and common `new`, `make`, and `append` expressions. |
| Types | `Any`, `Chan`, `Interface`, `Slice`, `Array`, `Map`, `Bool`, `Byte`, `Rune`, `String`, `Error`, `Int`, `Int8`, `Int16`, `Int32`, `Int64`, `Uint`, `Uint8`, `Uint16`, `Uint32`, `Uint64`, `Uintptr`, `Float32`, `Float64`, `Complex64`, `Complex128`, `Nil` | Add predeclared types and common Go type forms. |
| Layout and output | `Block`, `Comment`, `Line`, `Join`, `Render`, `Unwrap` | Build indented blocks, add comments or line breaks, combine tokens, render, or access the token slice. |
| Package render mode | `Render`, `SetRawMode`, `SetDefaultMode` | `Render` applies `go/format` by default. The mode setters affect subsequent package-level render calls. |

## Rendering modes and errors

`Tokens.Render(w)` writes the readable token layout without formatting. Package-level `golang.Render(w, token)` applies `go/format` by default. `SetRawMode()` disables formatting for subsequent package-level calls; `SetDefaultMode()` enables it again. These mode setters change package-wide behavior, so prefer `Tokens.Render` when a call should explicitly skip formatting without changing that behavior.

Rendering returns errors from tokens or the destination writer. With formatting enabled, `Render` also returns errors from `go/format` when the generated text is not valid Go syntax. Formatting does not type-check the generated program. `ID` validates identifier syntax, and `Op` rejects operators unsupported by this adapter. Rendering a `Field` returns an error if its tag arguments do not form key/value pairs.

`Raw` writes source text verbatim and does not validate it. Use it only for source controlled by the caller. `Array(n)` emits the supplied array length; it does not check whether that length is a legal Go constant expression.

To provide a custom token, implement [`types.Token`](../types/token.go) and write its source in `Render(io.Writer) error`.

For the repository overview and other language guides, see the [root README](../README.md), [TypeScript/TSX guide](../typescript/README.md), and [Python guide](../python/README.md). The Go package reference is available on [pkg.go.dev](https://pkg.go.dev/go.osspkg.com/gogen/golang).
