# TypeScript and TSX builders

The `go.osspkg.com/gogen/typescript` package builds TypeScript and TSX source from Go. It provides package-level constructors and chainable `*Tokens` methods, backed by the shared gogen token renderer.

```sh
go get go.osspkg.com/gogen/typescript
```

## Build and render TypeScript

Compose declarations from builders, then render the token sequence to an `io.Writer`:

```go
package main

import (
	"bytes"
	"fmt"

	ts "go.osspkg.com/gogen/typescript"
)

func main() {
	userType := ts.Export().Type().ID("User").Op("=").Block(
		ts.ID("id").Colon().String().Semicolon(),
		ts.ID("name").Colon().String().Semicolon(),
	).Semicolon()
	file := ts.Import("node:fs/promises", ts.ImportNames(ts.ID("readFile"))).Line().Join(userType)

	var source bytes.Buffer
	if err := ts.Render(&source, file); err != nil {
		panic(err)
	}
	fmt.Print(source.String())
}
```

Output:

```typescript
import { readFile } from "node:fs/promises";
export type User = {
	id: string;
	name: string;
};
```

`Tokens.Render(w)` and `typescript.Render(w, token)` produce readable token output without formatting, compiling, or type-checking it. `Block` creates explicit indented layout; nested token sequences do not imply line breaks. Use `.ts` when the generated file has no JSX.

## Build TSX

TSX nodes are structural tokens. The API does not depend on React or another framework. Use `JSXText` for literal text and `JSXExpr` for expressions; use the attribute builder that matches the value kind:

```go
package main

import (
	"bytes"
	"fmt"

	"go.osspkg.com/gogen/types"
	ts "go.osspkg.com/gogen/typescript"
)

func main() {
	node := ts.JSXElement(
		ts.ID("button"),
		[]types.Token{
			ts.JSXAttribute("type", "button"),
			ts.JSXAttributeExpr("disabled", ts.ID("isDisabled")),
			ts.JSXSpreadAttribute(ts.ID("props")),
			ts.JSXBooleanAttribute("autoFocus"),
		},
		ts.JSXText("Save "),
		ts.JSXExpr(ts.ID("label")),
	)
	var source bytes.Buffer
	if err := ts.Render(&source, node); err != nil {
		panic(err)
	}
	fmt.Print(source.String())
}
```

The rendered node is:

```tsx
<button type="button" disabled={isDisabled} {...props} autoFocus>Save {label}</button>
```

Use `JSXFragment(children...)` for a fragment. An element with no children is rendered as self-closing JSX. Text and quoted attribute values are escaped for JSX; expression tokens are rendered inside braces. Write generated code containing JSX to a `.tsx` file. This package does not run a formatter; invoke TypeScript or Prettier separately when needed.

## Builder reference

Builders are available both as package-level functions and as chainable `*Tokens` methods, except where noted. `TemplateText` and `TemplateExpr` return tokens to pass to `Template`; `From` is fluent-only.

| Area | Builders | Notes |
| --- | --- | --- |
| Modules | `Import`, `ImportType`, `ImportNames`, `ImportNamespace`, `Export`, `ExportDefault`, `ExportNames`, `ExportFrom`, `From` | Supports side-effect, default, named, namespace, type-only, and re-export declarations. Use `ID(name).As().ID(alias)` for aliases. |
| Declarations | `Const`, `Let`, `Var`, `Function`, `Type`, `Interface`, `Class`, `Namespace`, `Async`, `Await`, `Return`, `Throw`, `Extends`, `Implements`, `New`, `This` | Compose declarations and common statement keywords. Use `Bracket` for parameters, `Colon` for annotations, and `Block` for bodies. |
| Class modifiers | `Public`, `Private`, `Protected`, `Static`, `Readonly`, `Abstract`, `Override`, `Constructor` | Compose common class members and constructors. |
| Control flow | `If`, `Else`, `ElseIf`, `For`, `ForOf`, `ForIn`, `While`, `Do`, `Switch`, `Case`, `Default`, `Break`, `Continue`, `Try`, `Catch`, `Finally` | `ForOf` and `ForIn` create loops with a `const` binding. |
| Expressions | `ID`, `Raw`, `Text`, `Op`, `Call`, `Bracket`, `List`, `Index`, `TypeArgs`, `Selector`, `OptionalChain`, `NewCall`, `KeyValue`, `Spread`, `Arrow`, `As`, `Optional`, `NonNull` | `Text` creates an escaped string literal; `Raw` inserts caller-provided source verbatim. |
| Values and types | `ObjectLiteral`, `ArrayLiteral`, `Template`, `TemplateText`, `TemplateExpr`, `Any`, `Unknown`, `Never`, `Void`, `String`, `Number`, `Boolean`, `BigInt`, `Symbol`, `ObjectType`, `Null`, `Undefined`, `True`, `False`, `ArrayType`, `RecordType`, `MapType`, `SetType`, `PromiseType`, `Union`, `Intersection` | Compose object/array values, template strings, primitive and utility types, and unions or intersections. |
| TSX | `JSXElement`, `JSXFragment`, `JSXAttribute`, `JSXAttributeExpr`, `JSXBooleanAttribute`, `JSXSpreadAttribute`, `JSXText`, `JSXExpr` | Build elements, fragments, attributes, text nodes, and expression children. |
| Layout and output | `Block`, `Comment`, `Line`, `Join`, `Render`, `Unwrap` | Set explicit layout, combine token sequences, render source, or access the underlying token slice. |

## Rendering errors and raw source

Rendering returns errors from the destination writer and from invalid builder input. `ID` validates identifier syntax, and `Op` rejects operators outside the adapter's supported set. TSX rendering rejects invalid attribute names, nil tags, nil expressions, and nil children. An element with no children is valid and self-closing.

`Raw(source)` writes the supplied source verbatim. It does not parse or validate it, so keep raw input under the caller's control. `Text(value)` emits an escaped JavaScript string literal; `JSXText(value)` and `JSXAttribute(name, value)` escape literal values for their JSX contexts.

For the repository overview and Go adapter reference, see the [root README](../README.md). The package's exported symbols are also documented on [pkg.go.dev](https://pkg.go.dev/go.osspkg.com/gogen/typescript).
