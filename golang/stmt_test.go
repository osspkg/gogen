/*
 *  Copyright (c) 2025 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
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
	buf := bytes.NewBuffer(nil)

	stmt := Comment("Autogenerate code").
		Package("main").
		Add(
			Import("fmt", "fmt").
				Import("bytes", "bytes"),
			Comment("nolint:errcheck").
				Func().Id("main").
				Params(Id("id").Int64(), Id("name").String()).
				Params(Id("err").Error()).
				Block(
					List(Id("a"), Id("b")).Op(":=").List(Text("123"), Text("321")),
					Defer().Func().Params(Id("a").Int()).Block().Call(Id("a")),
					Go().Id("A").Call(),
					Return().Nil(),
				),
			Func().Id("A").Params().Block(),
		)

	if err := Render(buf, stmt); err != nil {
		t.Fatal(err)
	}

	fmt.Println("______________________________")
	fmt.Print(buf.String())
	fmt.Println("______________________________")
}
