/*
 *  Copyright (c) 2025-2026 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package typescript_test

import (
	"bytes"
	"fmt"

	"go.osspkg.com/gogen/types"
	ts "go.osspkg.com/gogen/typescript"
)

func ExampleRender() {
	file := ts.ExportDefault(
		ts.Function().ID("Greeting").Bracket(
			ts.ID("name").Colon().String(),
		).Block(
			ts.Return().Text("Hello, ").Op("+").ID("name").Semicolon(),
		),
	)

	var source bytes.Buffer
	if err := ts.Render(&source, file); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print(source.String())

	// Output:
	// export default function Greeting(name: string) {
	// 	return "Hello, " + name;
	// }
}

func ExampleJSXElement() {
	component := ts.ExportDefault(
		ts.Function().ID("Greeting").Bracket(
			ts.ID("name").Colon().String(),
		).Block(
			ts.Return().Join(ts.JSXElement(
				ts.ID("div"),
				[]types.Token{ts.JSXAttribute("className", "greeting")},
				ts.JSXText("Hello, "),
				ts.JSXExpr(ts.ID("name")),
			)),
		),
	)

	var source bytes.Buffer
	if err := ts.Render(&source, component); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print(source.String())

	// Output:
	// export default function Greeting(name: string) {
	// 	return <div className="greeting">Hello, {name}</div>
	// }
}

func ExampleImport() {
	userType := ts.Export().Type().ID("User").Op("=").Block(
		ts.ID("id").Colon().String().Semicolon(),
		ts.ID("name").Colon().String().Semicolon(),
	).Semicolon()
	file := ts.Import("node:fs/promises", ts.ImportNames(ts.ID("readFile"))).Line().Join(userType)

	var source bytes.Buffer
	if err := ts.Render(&source, file); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print(source.String())

	// Output:
	// import { readFile } from "node:fs/promises";
	// export type User = {
	//	id: string;
	//	name: string;
	// };
}

func ExampleJSXAttribute() {
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
		fmt.Println(err)
		return
	}
	fmt.Print(source.String())

	// Output:
	// <button type="button" disabled={isDisabled} {...props} autoFocus>Save {label}</button>
}
