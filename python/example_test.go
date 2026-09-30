/*
 *  Copyright (c) 2025-2026 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package python_test

import (
	"bytes"
	"fmt"

	py "go.osspkg.com/gogen/python"
)

func ExampleRender() {
	file := py.FromImport("dataclasses", py.ID("dataclass")).Line().Join(
		py.Decorator(py.ID("dataclass")),
		py.Class("User").Block(
			py.ID("name").Colon().ID("str"),
			py.ID("active").Colon().ID("bool").Op("=").True(),
		),
		py.Def("greet").Bracket(py.ID("user").Colon().ID("User")).Arrow().ID("str").Block(
			py.Return().FString(py.FStringText("Hello, "), py.FStringExpr(py.ID("user").Selector("name"))),
		),
	)

	var source bytes.Buffer
	if err := py.Render(&source, file); err != nil {
		panic(err)
	}
	fmt.Print(source.String())
	// Output:
	// from dataclasses import dataclass
	// @dataclass
	// class User:
	//     name: str
	//     active: bool = True
	// def greet(user: User) -> str:
	//     return f"Hello, {user.name}"
}
