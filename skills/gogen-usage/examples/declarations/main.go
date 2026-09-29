/*
 *  Copyright (c) 2025-2026 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package main

import (
	"fmt"
	"os"

	gogen "go.osspkg.com/gogen/golang"
)

func main() {
	file := gogen.Package("example").
		ImportBlock(gogen.Text("fmt")).
		TypeBlock(
			gogen.ID("User").Struct().Block(
				gogen.Field("Name", gogen.String(), "json", "name"),
			),
		).
		Join(
			gogen.Func().ID("main").Bracket().Block(
				gogen.Var().ID("user").Op("=").ID("User").Op("{").
					KeyValue(gogen.ID("Name"), gogen.Text("Ada")).Op("}"),
				gogen.Pkg("fmt").ID("Println").Call(gogen.ID("user").Op(".").ID("Name")),
			),
		)

	if err := gogen.Render(os.Stdout, file); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
