/*
 *  Copyright (c) 2025-2026 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package golang_test

import (
	"bytes"
	"io"
	"testing"

	gogen "go.osspkg.com/gogen/golang"
)

func TestTokensRenderSpacing(t *testing.T) {
	tests := []struct {
		name  string
		token interface{ Render(io.Writer) error }
		want  string
	}{
		{
			name:  "declaration and assignment",
			token: gogen.Var().ID("count").Op("=").Raw("1"),
			want:  "var count = 1",
		},
		{
			name:  "binary operator",
			token: gogen.ID("left").Op("+").ID("right"),
			want:  "left + right",
		},
		{
			name:  "unary operator",
			token: gogen.Op("*").ID("value"),
			want:  "*value",
		},
		{
			name:  "postfix increment",
			token: gogen.ID("index").Op("++"),
			want:  "index++",
		},
		{
			name:  "variadic expansion",
			token: gogen.Slice().Byte().Op("..."),
			want:  "[]byte...",
		},
		{
			name:  "raw content is preserved",
			token: gogen.Raw("alpha + beta\n  gamma"),
			want:  "alpha + beta\n  gamma",
		},
		{
			name:  "selector and call",
			token: gogen.Pkg("fmt").ID("Println").Call(gogen.Text("hello")),
			want:  "fmt.Println(\"hello\")",
		},
		{
			name:  "call argument separators",
			token: gogen.ID("combine").Call(gogen.ID("left"), gogen.ID("right")),
			want:  "combine(left, right)",
		},
		{
			name:  "index expression",
			token: gogen.ID("items").Index(gogen.ID("index")),
			want:  "items[index]",
		},
		{
			name:  "generic type arguments",
			token: gogen.ID("Map").TypeArgs(gogen.ID("string"), gogen.ID("int")),
			want:  "Map[string, int]",
		},
		{
			name:  "keyed element",
			token: gogen.KeyValue(gogen.Text("name"), gogen.ID("value")),
			want:  `"name": value`,
		},
		{
			name:  "empty function block",
			token: gogen.Func().ID("run").Bracket().Block(),
			want:  "func run() {}",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got bytes.Buffer
			if err := tt.token.Render(&got); err != nil {
				t.Fatal(err)
			}
			if got.String() != tt.want {
				t.Errorf("rendered source:\n%s\nwant:\n%s", got.String(), tt.want)
			}
		})
	}
}

func TestRenderFormatsCompositeExpressions(t *testing.T) {
	token := gogen.Package("main").Join(
		gogen.Func().ID("main").Bracket().Block(
			gogen.Var().ID("value").Op("=").
				Map(gogen.String(), gogen.Int()).Op("{").
				KeyValue(gogen.Text("key"), gogen.ID("items").Index(gogen.Raw("0"))).
				Op("}"),
		),
	)

	var got bytes.Buffer
	if err := gogen.Render(&got, token); err != nil {
		t.Fatal(err)
	}

	want := "package main\n\nfunc main() {\n\tvar value = map[string]int{\"key\": items[0]}\n}\n"
	if got.String() != want {
		t.Errorf("formatted source:\n%s\nwant:\n%s", got.String(), want)
	}
}

func TestTokensRenderFileAndNestedBlocks(t *testing.T) {
	file := gogen.Package("main").
		Import("fmt", "fmt").
		Join(
			gogen.Func().ID("main").Bracket().Block(
				gogen.If().ID("ready").Block(
					gogen.Pkg("fmt").ID("Println").Call(gogen.Text("ready")),
				).Else().Block(
					gogen.Pkg("fmt").ID("Println").Call(gogen.Text("waiting")),
				),
			),
		)

	var got bytes.Buffer
	if err := file.Render(&got); err != nil {
		t.Fatal(err)
	}

	want := "package main\nimport fmt \"fmt\"\nfunc main() {\n\tif ready {\n\t\tfmt.Println(\"ready\")\n\t} else {\n\t\tfmt.Println(\"waiting\")\n\t}\n}"
	if got.String() != want {
		t.Errorf("rendered source:\n%s\nwant:\n%s", got.String(), want)
	}
}

func TestTokensRenderDeclarationBlocks(t *testing.T) {
	file := gogen.Package("main").
		ImportBlock(
			gogen.Text("fmt"),
			gogen.ID("json").Text("encoding/json"),
		).
		TypeBlock(
			gogen.ID("Name").String(),
			gogen.ID("Count").Int(),
		)

	var got bytes.Buffer
	if err := file.Render(&got); err != nil {
		t.Fatal(err)
	}

	want := "package main\nimport (\n\t\"fmt\"\n\tjson \"encoding/json\"\n)\ntype (\n\tName string\n\tCount int\n)\n"
	if got.String() != want {
		t.Errorf("rendered source:\n%s\nwant:\n%s", got.String(), want)
	}
}

func TestRenderFormatsDeclarationBlocks(t *testing.T) {
	file := gogen.Package("main").
		ImportBlock(gogen.Text("fmt"), gogen.ID("json").Text("encoding/json")).
		TypeBlock(gogen.ID("Name").String(), gogen.ID("Count").Int())

	var got bytes.Buffer
	if err := gogen.Render(&got, file); err != nil {
		t.Fatal(err)
	}

	want := "package main\n\nimport (\n\tjson \"encoding/json\"\n\t\"fmt\"\n)\n\ntype (\n\tName  string\n\tCount int\n)\n"
	if got.String() != want {
		t.Errorf("formatted source:\n%s\nwant:\n%s", got.String(), want)
	}
}

func TestTokensRenderStructFieldTags(t *testing.T) {
	token := gogen.Struct().Block(
		gogen.Field("ID", gogen.Uint64(), "json", "id,omitempty", "db", "user_id"),
		gogen.Field("Name", gogen.String()),
	)

	var got bytes.Buffer
	if err := token.Render(&got); err != nil {
		t.Fatal(err)
	}

	want := "struct {\n\tID uint64 `json:\"id,omitempty\" db:\"user_id\"`\n\tName string\n}"
	if got.String() != want {
		t.Errorf("rendered source:\n%s\nwant:\n%s", got.String(), want)
	}
}

func TestRenderFormatsStructFieldTags(t *testing.T) {
	token := gogen.Package("main").Join(
		gogen.Type().ID("User").Struct().Block(
			gogen.Field("ID", gogen.Uint64(), "json", "id,omitempty", "db", "user_id"),
			gogen.Field("Name", gogen.String(), "json", "name"),
		),
	)

	var got bytes.Buffer
	if err := gogen.Render(&got, token); err != nil {
		t.Fatal(err)
	}

	want := "package main\n\ntype User struct {\n\tID   uint64 `json:\"id,omitempty\" db:\"user_id\"`\n\tName string `json:\"name\"`\n}\n"
	if got.String() != want {
		t.Errorf("formatted source:\n%s\nwant:\n%s", got.String(), want)
	}
}

func TestTokensRenderStructFieldRejectsUnpairedTag(t *testing.T) {
	var got bytes.Buffer
	if err := gogen.Field("ID", gogen.Uint64(), "json").Render(&got); err == nil {
		t.Fatal("Render() accepted an unpaired struct tag key")
	}
}

func TestTokensRenderPreservesMultilineRawInBlock(t *testing.T) {
	token := gogen.Block(gogen.Raw("first\n  second"))

	var got bytes.Buffer
	if err := token.Render(&got); err != nil {
		t.Fatal(err)
	}

	want := "{\n\tfirst\n  second\n}"
	if got.String() != want {
		t.Errorf("rendered source:\n%s\nwant:\n%s", got.String(), want)
	}
}

func TestTokensRenderCommentInBlock(t *testing.T) {
	token := gogen.Func().ID("run").Bracket().Block(
		gogen.Comment("marker"),
		gogen.Return(),
	)

	var got bytes.Buffer
	if err := token.Render(&got); err != nil {
		t.Fatal(err)
	}

	want := "func run() {\n\t// marker\n\treturn\n}"
	if got.String() != want {
		t.Errorf("rendered source:\n%s\nwant:\n%s", got.String(), want)
	}
}

func TestRenderFormatsTokenOutput(t *testing.T) {
	var got bytes.Buffer
	token := gogen.Package("main").Join(
		gogen.Func().ID("main").Bracket().Block(
			gogen.Var().ID("count").Op("=").Raw("1"),
		),
	)

	if err := gogen.Render(&got, token); err != nil {
		t.Fatal(err)
	}

	want := "package main\n\nfunc main() {\n\tvar count = 1\n}\n"
	if got.String() != want {
		t.Errorf("formatted source:\n%s\nwant:\n%s", got.String(), want)
	}
}

func TestTokensRenderRejectsInvalidOperator(t *testing.T) {
	var got bytes.Buffer
	if err := gogen.ID("value").Op("??").Render(&got); err == nil {
		t.Fatal("Render() accepted an invalid operator")
	}
}
