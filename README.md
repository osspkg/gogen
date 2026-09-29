# gogen

[![Go Version](https://img.shields.io/github/go-mod/go-version/osspkg/gogen)](https://go.dev/) [![License](https://img.shields.io/github/license/osspkg/gogen)](LICENSE) [![CI](https://img.shields.io/github/actions/workflow/status/osspkg/gogen/ci.yml?branch=master&label=CI)](https://github.com/osspkg/gogen/actions/workflows/ci.yml) [![Go Reference](https://pkg.go.dev/badge/go.osspkg.com/gogen/golang.svg)](https://pkg.go.dev/go.osspkg.com/gogen/golang)

`gogen` is a Go library for building Go source from composable tokens. Its fluent API covers declarations, types, expressions, comments, and control flow; generated source is formatted with `go/format` by default.

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

## Features and API

Each package-level constructor starts a `*Tokens` sequence. Methods with the same names append to an existing sequence, so builders can be chained. `Join` combines token sequences.

| Area | Methods and constructors | Use |
| --- | --- | --- |
| File declarations | `Package`, `Import`, `ImportBlock` | Start a file and add one import or a grouped import declaration. |
| Type declarations | `Type`, `TypeBlock`, `Struct`, `Interface`, `Field` | Declare named or grouped types, define struct/interface bodies, and add named struct fields. |
| Variables and constants | `Var`, `Const` | Start variable and constant declarations. |
| Functions and calls | `Func`, `Params`, `Bracket`, `Call`, `Return`, `Defer`, `Go` | Build function declarations, parameter lists, calls, and common function statements. |
| Control flow | `If`, `Else`, `ElseIf`, `For`, `Range`, `Switch`, `Select`, `Case`, `Default`, `Break`, `Continue`, `Fallthrough`, `Goto` | Compose conditional statements, loops, and switch/select clauses. |
| Expressions | `ID`, `Pkg`, `Op`, `Raw`, `Text`, `Index`, `TypeArgs`, `KeyValue`, `List`, `New`, `Make`, `Append` | Build identifiers, package selectors, operators, literals, indexing, generic arguments, and common built-in calls. |
| Types | `Any`, `Nil`, `Chan`, `Map`, `Slice`, `Array`, `Bool`, `Byte`, `Rune`, `String`, `Error`, `Int`, `Int8`, `Int16`, `Int32`, `Int64`, `Uint`, `Uint8`, `Uint16`, `Uint32`, `Uint64`, `Uintptr`, `Float32`, `Float64`, `Complex64`, `Complex128` | Add predeclared Go types and map/slice/array type syntax. |
| Layout and rendering | `Block`, `Comment`, `Line`, `Tokens.Render`, `Tokens.Unwrap`, `Tokens.Join`, `Render`, `SetRawMode`, `SetDefaultMode` | Add blocks and comments, inspect token output, combine tokens, and choose formatted or raw output. |
| Custom tokens | `types.Token` | Implement `Render(io.Writer) error` to add a token renderer. |

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

Tag arguments must be key/value pairs; rendering returns an error when the argument count is odd. `Text` quotes a Go string literal. `Raw` inserts source text verbatim. `ID` validates identifier text and `Op` rejects unsupported operators during rendering.

`Tokens.Render(w)` writes readable token layout before formatting. `golang.Render(w, token)` applies `go/format` by default and returns syntax errors when the generated source is invalid; it does not type-check the program. `SetRawMode` and `SetDefaultMode` switch package-wide behavior for subsequent `Render` calls.

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md) for prerequisites and local checks.

## Contributors

[![Contributors](https://img.shields.io/github/contributors/osspkg/gogen)](https://github.com/osspkg/gogen/graphs/contributors)

## License

This project is licensed under the [BSD 3-Clause License](LICENSE).
