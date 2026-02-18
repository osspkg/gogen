/*
 *  Copyright (c) 2025-2026 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package golang_test

import (
	"bytes"
	"fmt"
	"testing"

	. "go.osspkg.com/gogen/golang"
)

func Test_newStmt(t *testing.T) {
	//SetRawMode()

	buf := bytes.NewBuffer(nil)

	stmt := Comment("Autogenerate code").
		Package("main").
		Import("fmt", "fmt").
		Import("fmt2", "fmt").
		Join(
			Func().ID("main").Bracket().Block(
				Comment("d").
					Var().ID("a").Op("=").
					Map(Interface().Block(), Interface().Block(ID("Do").Bracket().Error())).
					Block(
						Text("a").Op(":").Text("").Op(","),
						Array(2).Int8().Op(":").Raw("1").Op(","),
					).
					Comment("d"),
				Var().ID("b").Op("=").New(Int8()),
				Var().ID("b").Op("=").Make(Slice().Int8(), 0, 10),
				Op("*").ID("aaa").Op("=").Append(Op("*").ID("aaa"), Slice().Byte().Op("...")),
				For().List(ID("a"), ID("b")).Op(":=").Range().Pkg("fmt").ID("aa").Block(),
				For().Block(
					Select().Block(),
				),
				Switch().ID("aa").Op(".").Bracket(ID("type")).Block(
					Case().Slice().Byte().Op(":").Block(
						Pkg("fmt").ID("Println").Call(ID("aa")),
					),
					Default().Op(":"),
				),
			),
		)

	if err := Render(buf, stmt); err != nil {
		t.Fatal(err)
	}

	fmt.Println("______________________________")
	fmt.Print(buf.String())
	fmt.Println("______________________________")
}
