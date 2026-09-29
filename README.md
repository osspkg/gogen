# gogen

[![Go Version](https://img.shields.io/github/go-mod/go-version/osspkg/gogen)](https://go.dev/) [![License](https://img.shields.io/github/license/osspkg/gogen)](LICENSE) [![CI](https://img.shields.io/github/actions/workflow/status/osspkg/gogen/ci.yml?branch=master&label=CI)](https://github.com/osspkg/gogen/actions/workflows/ci.yml) [![Go Reference](https://pkg.go.dev/badge/go.osspkg.com/gogen/golang.svg)](https://pkg.go.dev/go.osspkg.com/gogen/golang)

`gogen` is a token-based source generation library. Language packages provide their own fluent APIs and syntax policies, while shared internal renderers handle reusable token and layout behavior. The current public language adapter is `golang`, which formats generated source with `go/format` by default.

## Demo

This program generates a Go file and prints the formatted source:

```go
package main

import (
	"bytes"
	"fmt"

	gogen "go.osspkg.com/gogen/golang"
)

func main() {
	file := gogen.Package("main").
		Import("fmt", "fmt").
		Join(
			gogen.Func().ID("main").Bracket().Block(
				gogen.Pkg("fmt").ID("Println").Call(gogen.Text("hello, gogen")),
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

import fmt "fmt"

func main() {
	fmt.Println("hello, gogen")
}
```

## Getting started

The module requires Go 1.26 or newer. Add the builder package from your module root:

```sh
go get go.osspkg.com/gogen/golang
```

## API reference

Package-level constructors such as `Func()` and `Text("hello")` create a `*Tokens` value. Fluent methods append to that sequence and return it, so expressions can be chained. `Join(tokens...)` appends existing token sequences. A few operations are intentionally fluent-only or package-only; those are called out below.

| Area | Fluent methods | Package-level constructors | Use |
| --- | --- | --- | --- |
| File and declarations | `Package`, `Import`, `ImportBlock`, `Type`, `TypeBlock`, `Var`, `Const` | `Package`, `Import`, `ImportBlock`, `Type`, `TypeBlock`, `Var`, `Const` | Start a package, add imports, and declare values or types. Group builders take declaration specs as tokens. |
| Functions and statements | `Func`, `Return`, `If`, `Else`, `ElseIf`, `For`, `Range`, `Case`, `Default`, `Break`, `Continue`, `Fallthrough`, `Goto` | `Func`, `Return`, `Defer`, `Go`, `If`, `For`, `Switch`, `Select`, `Case`, `Default` | Compose function declarations and control-flow statements. `Defer`, `Go`, `Switch`, and `Select` have package-level constructors only; `Else`, `ElseIf`, `Range`, `Break`, `Continue`, `Fallthrough`, and `Goto` are fluent-only. |
| Calls and lists | `Bracket`, `Call`, `List` | `Call`, `Params`, `List` | `Call` adds call parentheses; `Bracket` adds parentheses to an existing sequence; `Params` creates a parameter list; `List` creates a comma-separated sequence without delimiters. |
| Expressions | `ID`, `Pkg`, `Op`, `Raw`, `Text`, `Index`, `TypeArgs`, `KeyValue` | `ID`, `Pkg`, `Op`, `Raw`, `Text`, `Index`, `TypeArgs`, `KeyValue` | Add identifiers, package selectors, operators, literals, indexes, generic arguments, and keyed composite-literal elements. |
| Composite values and allocation | `New`, `Make`, `Append` | `New`, `Make`, `Append` | Build common `new`, `make`, and `append` calls. `Make` includes capacity only when it is greater than length. |
| Types | `Any`, `Chan`, `Interface`, `Struct`, `Slice`, `Array`, `Map`, `Bool`, `Byte`, `Rune`, `String`, `Error`, `Int`, `Int8`, `Int16`, `Int32`, `Int64`, `Uint`, `Uint8`, `Uint16`, `Uint32`, `Uint64`, `Uintptr`, `Float32`, `Float64`, `Complex64`, `Complex128`, `Nil`, `Field` | Same names as fluent methods | Add predeclared types and common Go type forms. `Field(name, type, tags...)` adds a named struct field. |
| Layout and rendering | `Block`, `Comment`, `Line`, `Join`, `Render`, `Unwrap` | `Block`, `Comment`, `Line` | Create indented blocks and comments, add line breaks, combine tokens, render readable token output, or access the underlying token slice. |
| Package rendering | — | `Render`, `SetRawMode`, `SetDefaultMode` | Format a token with `go/format`, or switch the package-wide render mode. |

`Tokens` is a slice-backed builder. `Unwrap()` returns its underlying slice without copying; changes to that slice are visible to the builder.

### Grouped declarations

Pass import specs as `Text(path)` or `ID(alias).Text(path)`. Pass each type declaration as a token sequence containing its name and type:

```go
file := gogen.Package("main").
	ImportBlock(
		gogen.Text("fmt"),
		gogen.ID("json").Text("encoding/json"),
	).
	TypeBlock(
		gogen.ID("Name").String(),
		gogen.ID("Count").Int(),
	)
```

`ImportBlock` and `TypeBlock` own the parentheses, line breaks, and indentation. Their specs do not include the `import` or `type` keyword.

### Struct tags

`Field` takes struct-tag key/value pairs and quotes the values for Go source:

```go
user := gogen.Struct().Block(
	gogen.Field("ID", gogen.Uint64(), "json", "id,omitempty", "db", "user_id"),
	gogen.Field("Name", gogen.String(), "json", "name"),
)
```

The example renders as:

```go
struct {
	ID uint64 `json:"id,omitempty" db:"user_id"`
	Name string `json:"name"`
}
```

Tag arguments must be key/value pairs; rendering returns an error when the argument count is odd.

### Rendering and errors

`Tokens.Render(w)` writes readable token layout without calling `go/format`. `golang.Render(w, token)` applies `go/format` by default and returns rendering, formatting, or writer errors. Formatting checks syntax only; it does not type-check the generated program. `SetRawMode` and `SetDefaultMode` switch package-wide behavior for subsequent `Render` calls. Raw mode skips formatting, while `Raw` content is always inserted verbatim.

`ID` validates identifier text against the builder's accepted form, and `Op` rejects unsupported Go operators during rendering. `Array(n)` emits the requested length without checking whether it is a legal Go array length.

To add a custom token, implement the public [`types.Token`](types/token.go) interface. Its `Render(io.Writer) error` method should write source text and return writer or rendering errors.

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md) for prerequisites and local checks.

## Contributors

[![Contributors](https://img.shields.io/github/contributors/osspkg/gogen)](https://github.com/osspkg/gogen/graphs/contributors)

## License

This project is licensed under the [BSD 3-Clause License](LICENSE).
