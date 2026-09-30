/*
 *  Copyright (c) 2025-2026 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package main

import (
	"bytes"
	"fmt"
	"os"

	gogen "go.osspkg.com/gogen/golang"
)

func main() {
	expression := gogen.ID("lookup").TypeArgs(gogen.String()).Call(
		gogen.ID("items").Index(gogen.Raw("0")),
	)

	var fragment bytes.Buffer
	if err := expression.Render(&fragment); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Printf("Tokens.Render: %s\n", fragment.String())

	file := gogen.Package("example").Join(
		gogen.Func().ID("lookup").Bracket(
			gogen.ID("items").Slice().String(),
			gogen.ID("index").Int(),
		).String().Block(
			gogen.Return().ID("items").Index(gogen.ID("index")),
		),
	)

	if err := gogen.Render(os.Stdout, file); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
