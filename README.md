# gogen

`gogen` is a Go library for building Go source code from composable tokens. Its fluent API covers common declarations, types, expressions, comments, and control-flow blocks. The generated source is formatted with `go/format` by default.

The module requires **Go 1.26 or newer**.

## Installation

From your Go module root, add the public builder package:

```sh
go get go.osspkg.com/gogen/golang
```

## Quick start

```go
package main

import (
	"bytes"
	"fmt"

	gogen "go.osspkg.com/gogen/golang"
)

func main() {
	var source bytes.Buffer

	file := gogen.Package("main").
		Import("fmt", "fmt").
		Join(
			gogen.Func().ID("main").Bracket().Block(
				gogen.Pkg("fmt").ID("Println").Call(gogen.Text("hello, gogen")),
			),
		)

	if err := gogen.Render(&source, file); err != nil {
		panic(err)
	}

	fmt.Print(source.String())
}
```

The program writes formatted Go source to `source`, equivalent to:

```go
package main

import fmt "fmt"

func main() {
	fmt.Println("hello, gogen")
}
```

Imports are explicit: add them with `Import`; package-qualified identifiers can be built with `Pkg`.

## Building tokens

Every constructor returns a `*Tokens` value. Chain methods to extend it, or combine token values with `Join`:

```go
statement := gogen.Var().ID("count").Op("=").Raw("1")
```

Common constructors include:

- **Declarations and types:** `Package`, `Import`, `Func`, `Type`, `Var`, `Const`, `Struct`, `Interface`, `Map`, `Slice`, and `Array`.
- **Expressions:** `ID`, `Pkg`, `Text`, `Raw`, `Op`, `Call`, `Params`, `New`, `Make`, and `Append`.
- **Statements and blocks:** `If`, `For`, `Range`, `Switch`, `Select`, `Case`, `Return`, and `Block`.
- **Formatting:** `Comment` and `Line`.

`Text` quotes its argument as a Go string literal. `ID` checks identifier text, and `Op` accepts operators supported by the Go builder. `Raw` inserts text directly. The public `types.Token` interface (`Render(io.Writer) error`) can be implemented to add custom tokens.

## Rendering

`golang.Render(w, token)` renders a token to an `io.Writer` and formats the result using `go/format`. Formatting also reports Go syntax errors in the generated source. It does not type-check or compile the generated program.

For unformatted output, call `SetRawMode()`. Call `SetDefaultMode()` to restore formatting. These functions switch a package-wide mode that applies to subsequent renders.

## Development

The repository's Makefile wraps its checks with `goppy`. Run commands from the repository root:

```sh
make tests
make lint
make build
```

The workflow used by GitHub Actions is:

```sh
make ci
```

`make ci` includes license, lint, test, and build steps. It also installs the latest `goppy` command and runs `goppy setup-lib`, so it may require network access and affect local setup. The CI workflow currently uses Go 1.26.

## License

This project is licensed under the [BSD 3-Clause License](LICENSE).
