# gogen

[![Go Version](https://img.shields.io/github/go-mod/go-version/osspkg/gogen)](https://go.dev/) [![License](https://img.shields.io/github/license/osspkg/gogen)](LICENSE) [![CI](https://img.shields.io/github/actions/workflow/status/osspkg/gogen/ci.yml?branch=master&label=CI)](https://github.com/osspkg/gogen/actions/workflows/ci.yml) [![Go Reference](https://pkg.go.dev/badge/go.osspkg.com/gogen/golang.svg)](https://pkg.go.dev/go.osspkg.com/gogen/golang) [![TypeScript Reference](https://pkg.go.dev/badge/go.osspkg.com/gogen/typescript.svg)](https://pkg.go.dev/go.osspkg.com/gogen/typescript)

`gogen` is a token-based source generation library for Go, TypeScript, TSX, and Python. Language adapters provide composable builders while shared internals handle reusable token rendering and layout.

## Demo

Generate a formatted Go source file:

```go
package main

import (
	"bytes"
	"fmt"

	gogen "go.osspkg.com/gogen/golang"
)

func main() {
	file := gogen.Package("main").Import("fmt", "fmt").Join(
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

## Getting started

Requires Go 1.26 or newer:

```sh
go get go.osspkg.com/gogen/golang go.osspkg.com/gogen/typescript go.osspkg.com/gogen/python
```

## Language guides

| Language | Package | Guide |
| --- | --- | --- |
| Go | [`go.osspkg.com/gogen/golang`](https://pkg.go.dev/go.osspkg.com/gogen/golang) | [Go builder guide](golang/README.md) |
| TypeScript and TSX | [`go.osspkg.com/gogen/typescript`](https://pkg.go.dev/go.osspkg.com/gogen/typescript) | [TypeScript and TSX guide](typescript/README.md) |
| Python 3.10+ | [`go.osspkg.com/gogen/python`](https://pkg.go.dev/go.osspkg.com/gogen/python) | [Python guide](python/README.md) |

The Go adapter applies `go/format` by default. The TypeScript and Python adapters write readable source without invoking a formatter; use the language toolchain or a formatter separately when needed.

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md) for setup and local checks.

## Contributors

[![Contributors](https://img.shields.io/github/contributors/osspkg/gogen)](https://github.com/osspkg/gogen/graphs/contributors)

## License

This project is licensed under the [BSD 3-Clause License](LICENSE).
