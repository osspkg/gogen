# Python code builder

The `go.osspkg.com/gogen/python` package builds Python 3.10+ source from composable tokens. Package-level constructors create expressions and statement headers; fluent `*Tokens` methods append more tokens. `Tokens.Render` and package-level `python.Render` write readable source directly without formatting or interpreter validation.

```sh
go get go.osspkg.com/gogen/python
```

## Build a module

```go
package main

import (
    "bytes"
    "fmt"

    py "go.osspkg.com/gogen/python"
)

func main() {
    file := py.FromImport("dataclasses", py.ID("dataclass")).Line().Join(
        py.Decorator(py.ID("dataclass")),
        py.Class("User").Block(
            py.ID("name").Colon().String(),
            py.ID("active").Colon().Bool().Op("=").True(),
        ),
    ).Line().Join(
        py.AsyncDef("load_user").Bracket(py.ID("user_id").Colon().Int()).Arrow().TypeUnion(py.ID("User"), py.NoneValue()).Block(
            py.If(py.ID("user_id").Op("<").Raw("0")).Block(py.Return().NoneValue()),
            py.Return().Await().ID("repository").Selector("load").Call(py.ID("user_id")),
        ),
    )

    var source bytes.Buffer
    if err := py.Render(&source, file); err != nil {
        panic(err)
    }
    fmt.Print(source.String())
}
```

The rendered source uses four spaces for each suite level:

```python
from dataclasses import dataclass
@dataclass
class User:
    name: str
    active: bool = True

async def load_user(user_id: int) -> User | None:
    if user_id < 0:
        return None
    return await repository.load(user_id)
```

`Line()` explicitly separates top-level statements. Nested suite builders add their own line breaks and indentation. Empty suites render as `pass`.

## Declarations and control flow

Use `Import`, `ImportAs`, and `FromImport` for imports. `Def`, `AsyncDef`, and `Class` return headers that can be extended with `Bracket`, `Colon`, `Arrow`, and `Block`. `Decorator` emits a decorator on its own line. Control headers include `If`, `Elif`, `Else`, `ForEach`, `AsyncForEach`, `While`, `With`, `AsyncWith`, `Try`, `Except`, `Finally`, `Match`, and `Case`. Call `Block` on a header to add its suite. Clause builders stay aligned after the preceding suite:

```go
py.If(py.ID("ready")).Block(py.Return().True()).Else().Block(py.Return().False())
```

## Expressions and values

| Builder | Use |
| --- | --- |
| `ID`, `Pkg`, `Selector` | Validate identifiers or compose dotted access. |
| `Raw`, `Text`, `FString`, `FStringText`, `FStringExpr` | Add caller-provided source, escaped strings, and interpolated strings. |
| `Op`, `Call`, `Bracket`, `Index`, `List`, `KeyValue` | Compose operators, calls, grouped arguments, indexing, and key/value entries. |
| `ListLiteral`, `TupleLiteral`, `SetLiteral`, `DictLiteral` | Build common collection literals. |
| `ListComp`, `SetComp`, `DictComp`, `GeneratorExpr` | Build comprehensions with `ForClause`, `AsyncForClause`, and `IfClause`. |
| `TypeUnion` | Compose Python 3.10 union annotations with `|`. |
| `Int`, `Float`, `Complex`, `Bool`, `String`, `Str`, `Bytes`, `ByteArray`, `MemoryView`, `Range`, `ObjectType`, `TypeType`, `Any` | Emit built-in types; `Any` emits the name from `typing`. |
| `ListType`, `DictType`, `SetType`, `FrozenSetType`, `TupleType` | Emit PEP 585 generic container types. |
| `NoneValue`, `True`, `False`, `Ellipsis` | Emit Python's built-in literal values. |

`Text` and f-string text parts escape their respective string contexts. F-string expressions accept tokens; advanced conversion flags and format specifications can be composed with `Raw`.

## Built-in types

Use the type builders in annotations and generic expressions. They emit Python's built-in type names and PEP 585 forms such as `list[str]`, `dict[str, int]`, and `tuple[str, ...]`:

```go
py.Def("load_names").Bracket().Arrow().ListType(py.String()).Block(
    py.Return().ListLiteral(py.Text("Ada")),
)
```

`Int`, `Float`, `Complex`, `Bool`, `String`/`Str`, `Bytes`, `ByteArray`, `MemoryView`, `Range`, `ObjectType`, and `TypeType` emit built-in type names. Use `ListType`, `DictType`, `SetType`, `FrozenSetType`, and `TupleType` for generic containers. `TupleType()` without arguments emits bare `tuple`; `TupleType(py.Int(), py.Ellipsis())` emits `tuple[int, ...]`. `Any()` emits `Any` and expects the generated module to import it from `typing`. For `None`, use `NoneValue()`; unions can be composed with `TypeUnion`.

## Builder reference

Package-level constructors and chainable `*Tokens` methods are available unless the table identifies a component token. Component constructors return `types.Token`; use them as arguments to the parent builder.

| Area | Builders | Notes |
| --- | --- | --- |
| Imports and declarations | `Import`, `ImportAs`, `FromImport`, `Decorator`, `Def`, `AsyncDef`, `Class` | Build imports, decorators, function headers, and class headers. |
| Control flow | `If`, `Elif`, `Else`, `ForEach`, `AsyncForEach`, `While`, `With`, `AsyncWith`, `Try`, `Except`, `Finally`, `Match`, `Case` | Append `Block` to a header to emit its suite. |
| Statements and syntax | `Return`, `Yield`, `Raise`, `Pass`, `Break`, `Continue`, `Assert`, `Del`, `Global`, `Nonlocal`, `Await`, `As`, `Arrow`, `Colon`, `Comma` | Compose statements, async expressions, annotations, aliases, and punctuation. |
| Identifiers and expressions | `ID`, `Pkg`, `Selector`, `Raw`, `Text`, `Op`, `Call`, `Bracket`, `Index`, `List`, `KeyValue`, `TypeUnion` | `Raw` writes source verbatim; `Text` escapes a Python string literal. `List` is a comma-separated sequence without delimiters. |
| Built-in types | `Any`, `Int`, `Float`, `Complex`, `Bool`, `String`, `Str`, `Bytes`, `ByteArray`, `MemoryView`, `Range`, `ObjectType`, `TypeType` | `Any` emits `typing.Any`; generated modules must import it. `String` and `Str` both emit `str`. |
| Generic types | `ListType`, `DictType`, `SetType`, `FrozenSetType`, `TupleType` | Emit Python 3.10-compatible PEP 585 annotations. |
| Collection values | `ListLiteral`, `TupleLiteral`, `SetLiteral`, `DictLiteral` | Build list, tuple, set, and dictionary expressions. |
| Comprehensions | `ListComp`, `SetComp`, `DictComp`, `GeneratorExpr` | Combine with the component constructors `ForClause`, `AsyncForClause`, and `IfClause`. |
| F-strings and values | `FString`, `NoneValue`, `True`, `False`, `Ellipsis` | `FStringText` and `FStringExpr` create f-string parts; they return `types.Token`. |
| Layout and output | `Block`, `Comment`, `Line`, `Join`, `Render`, `Unwrap` | `Block` emits a colon and an indented suite; `Line` adds an explicit top-level line break. |


## Errors and custom tokens

`ID` validates Unicode identifier syntax and rejects Python keywords. `Op` rejects operators outside the adapter's supported set. Rendering propagates token and writer errors, but does not check Python grammar, names, types, or runtime behavior. `Raw` is written verbatim; use it for syntax controlled by the caller. To provide a custom builder, implement [`types.Token`](../types/token.go) and write its source in `Render(io.Writer) error`.

For other language adapters and repository information, see the [root README](../README.md), [Go guide](../golang/README.md), and [TypeScript/TSX guide](../typescript/README.md). The package reference is on [pkg.go.dev](https://pkg.go.dev/go.osspkg.com/gogen/python).
