/*
 *  Copyright (c) 2025-2026 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package typescript_test

import (
	"bytes"
	"io"
	"testing"

	"go.osspkg.com/gogen/types"
	ts "go.osspkg.com/gogen/typescript"
)

func render(t *testing.T, token interface{ Render(io.Writer) error }) string {
	t.Helper()
	var out bytes.Buffer
	if err := token.Render(&out); err != nil {
		t.Fatal(err)
	}
	return out.String()
}

func TestTokensRenderExpressions(t *testing.T) {
	tests := []struct {
		name  string
		token interface{ Render(io.Writer) error }
		want  string
	}{
		{name: "call and property selector", token: ts.ID("client").Selector("send").Call(ts.Text("hello")), want: `client.send("hello")`},
		{name: "optional chain and nullish coalescing", token: ts.ID("user").OptionalChain("name").Op("??").Text("anonymous"), want: `user?.name ?? "anonymous"`},
		{name: "generic call with array index", token: ts.ID("lookup").TypeArgs(ts.String()).Call(ts.ID("items").Index(ts.Raw("0"))), want: `lookup<string>(items[0])`},
		{name: "array literal", token: ts.ArrayLiteral(ts.Text("a"), ts.Text("b")), want: `["a", "b"]`},
		{name: "array type", token: ts.ArrayType(ts.String()), want: `string[]`},
		{name: "object literal", token: ts.ObjectLiteral(ts.KeyValue(ts.ID("name"), ts.Text("Ada"))), want: "{\n\tname: \"Ada\"\n}"},
		{name: "optional property", token: ts.ID("name").Optional().Colon().String(), want: "name?: string"},
		{name: "non-null assertion", token: ts.ID("value").NonNull(), want: "value!"},
		{name: "template interpolation", token: ts.Template(ts.TemplateText("hello "), ts.TemplateExpr(ts.ID("name"))), want: "`hello ${name}`"},
		{name: "raw is preserved", token: ts.Raw("first\n  second"), want: "first\n  second"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := render(t, test.token); got != test.want {
				t.Fatalf("render() = %q, want %q", got, test.want)
			}
		})
	}

	got := render(t, ts.Template(ts.TemplateText("` ${value} \\")))
	want := string([]byte{'`', '\\', '`', ' ', '\\', '$', '{', 'v', 'a', 'l', 'u', 'e', '}', ' ', '\\', '\\', '`'})
	if got != want {
		t.Fatalf("escaped template = %q, want %q", got, want)
	}
}

func TestTokensRenderDeclarationsAndControlFlow(t *testing.T) {
	function := ts.Export().Function().ID("greet").
		Bracket(ts.ID("name").Colon().String()).Colon().String().
		Block(
			ts.If().Bracket(ts.ID("name").Op("===").Text("")).Block(
				ts.Return().Text("anonymous").Semicolon(),
			).Else().Block(
				ts.Return().Text("hello ").Op("+").ID("name").Semicolon(),
			),
		)
	file := ts.Import("react", ts.ID("React")).Line().Join(function)

	want := "import React from \"react\";\nexport function greet(name: string): string {\n\tif (name === \"\") {\n\t\treturn \"anonymous\";\n\t} else {\n\t\treturn \"hello \" + name;\n\t}\n}"
	if got := render(t, file); got != want {
		t.Fatalf("render() =\n%s\nwant:\n%s", got, want)
	}
}

func TestTypesAndClassDeclarations(t *testing.T) {
	typeAlias := ts.Export().Type().ID("Identifier").Op("=").Union(ts.String(), ts.Number()).Semicolon()
	interfaceDecl := ts.Interface().ID("Repository").TypeArgs(ts.ID("T")).Block(
		ts.ID("get").Bracket(ts.ID("id").Colon().String()).Colon().ID("T").Semicolon(),
	)
	classDecl := ts.Export().Class().ID("UserRepository").Implements().ID("Repository").TypeArgs(ts.ID("User")).Block(
		ts.Private().ID("items").Colon().ArrayType(ts.ID("User")).Semicolon(),
		ts.Constructor().Bracket(ts.ID("items").Colon().ArrayType(ts.ID("User"))).Block(
			ts.This().Selector("items").Op("=").ID("items").Semicolon(),
		),
	)
	want := "export type Identifier = string | number;\ninterface Repository<T> {\n\tget(id: string): T;\n}\nexport class UserRepository implements Repository<User> {\n\tprivate items: User[];\n\tconstructor(items: User[]) {\n\t\tthis.items = items;\n\t}\n}"
	got := render(t, typeAlias.Line().Join(interfaceDecl, ts.Line(), classDecl))
	if got != want {
		t.Fatalf("render() =\n%s\nwant:\n%s", got, want)
	}
}

func TestEmptyAndNestedBlocks(t *testing.T) {
	empty := render(t, ts.Block())
	if empty != "{}" {
		t.Fatalf("empty block = %q, want %q", empty, "{}")
	}

	nested := ts.Function().ID("outer").Bracket().Block(
		ts.If().Bracket(ts.True()).Block(),
		ts.If().Bracket(ts.False()).Block(ts.ID("nested").Call()),
	)
	want := "function outer() {\n\tif (true) {}\n\tif (false) {\n\t\tnested()\n\t}\n}"
	if got := render(t, nested); got != want {
		t.Fatalf("nested blocks =\n%s\nwant:\n%s", got, want)
	}
}

func TestControlFlowAndComments(t *testing.T) {
	token := ts.ForOf("item", ts.ID("items")).Block(
		ts.Comment("skip non-positive values"),
		ts.If().Bracket(ts.ID("item").Op("<=").Raw("0")).Block(
			ts.Continue().Semicolon(),
		),
	)
	want := "for (const item of items) {\n\t// skip non-positive values\n\tif (item <= 0) {\n\t\tcontinue;\n\t}\n}"
	if got := render(t, token); got != want {
		t.Fatalf("render() =\n%s\nwant:\n%s", got, want)
	}
}

func TestIdentifierAndOperatorValidation(t *testing.T) {
	if got := render(t, ts.ID("变量").Op("+").ID("value")); got != "变量 + value" {
		t.Fatalf("Unicode identifier render = %q", got)
	}

	for _, token := range []interface{ Render(io.Writer) error }{
		ts.ID("not-an-identifier"),
		ts.ID("value").Op("???"),
	} {
		var out bytes.Buffer
		if err := token.Render(&out); err == nil {
			t.Errorf("Render(%T) succeeded, want validation error", token)
		}
	}
}

func TestStringLiteralEscaping(t *testing.T) {
	token := ts.Text("quote: \"; bell: \a; line: \u2028")
	want := `"quote: \"; bell: \x07; line: \u2028"`
	if got := render(t, token); got != want {
		t.Fatalf("string literal = %q, want %q", got, want)
	}
}

func TestImportAndExportBuilders(t *testing.T) {
	tests := []struct {
		name  string
		token interface{ Render(io.Writer) error }
		want  string
	}{
		{name: "named import", token: ts.Import("react", ts.ImportNames(ts.ID("useState"), ts.ID("useEffect").As().ID("effect"))), want: `import { useState, useEffect as effect } from "react";`},
		{name: "combined default and named import", token: ts.Import("react", ts.ID("React"), ts.ImportNames(ts.ID("useState"))), want: `import React, { useState } from "react";`},
		{name: "namespace import", token: ts.ImportNamespace("fs", "node:fs"), want: `import * as fs from "node:fs";`},
		{name: "type-only import", token: ts.ImportType("./types", ts.ImportNames(ts.ID("User"))), want: `import type { User } from "./types";`},
		{name: "side-effect import", token: ts.Import("./setup"), want: `import "./setup";`},
		{name: "export names", token: ts.ExportNames(ts.ID("User"), ts.ID("Name").As().ID("UserName")), want: `export { User, Name as UserName };`},
		{name: "re-export names", token: ts.ExportFrom("./model", ts.ID("User")), want: `export { User } from "./model";`},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := render(t, test.token); got != test.want {
				t.Fatalf("render() = %q, want %q", got, test.want)
			}
		})
	}
}

func TestRenderMatchesTokensRender(t *testing.T) {
	token := ts.Const().ID("count").Op("=").Raw("2").Semicolon()
	want := render(t, token)
	var got bytes.Buffer
	if err := ts.Render(&got, token); err != nil {
		t.Fatal(err)
	}
	if got.String() != want {
		t.Fatalf("Render() = %q, Tokens.Render() = %q", got.String(), want)
	}
}

func TestTSXRenderingAndEscaping(t *testing.T) {
	element := ts.JSXElement(
		ts.ID("main"),
		[]types.Token{
			ts.JSXAttribute("title", `say "hi" & <`),
			ts.JSXAttributeExpr("count", ts.ID("count")),
			ts.JSXSpreadAttribute(ts.ID("props")),
			ts.JSXBooleanAttribute("hidden"),
		},
		ts.JSXText("A < B & {value} "),
		ts.JSXExpr(ts.ID("name")),
		ts.JSXElement(ts.ID("img"), []types.Token{ts.JSXAttribute("alt", "avatar")}),
	)
	want := `<main title="say &quot;hi&quot; &amp; &lt;" count={count} {...props} hidden>A &lt; B &amp; &#123;value&#125; {name}<img alt="avatar" /></main>`
	if got := render(t, element); got != want {
		t.Fatalf("JSX render() = %q, want %q", got, want)
	}

	fragment := ts.JSXFragment(ts.JSXElement(ts.ID("span"), nil, ts.JSXText("ok")))
	if got, want := render(t, fragment), "<><span>ok</span></>"; got != want {
		t.Fatalf("fragment render() = %q, want %q", got, want)
	}
}

func TestTSXRejectsInvalidAttributesAndMissingExpressions(t *testing.T) {
	tests := []struct {
		name  string
		token interface{ Render(io.Writer) error }
	}{
		{name: "invalid attribute name", token: ts.JSXAttribute(`title" onclick="x`, "value")},
		{name: "nil expression", token: ts.JSXExpr(nil)},
		{name: "nil spread expression", token: ts.JSXSpreadAttribute(nil)},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var out bytes.Buffer
			if err := test.token.Render(&out); err == nil {
				t.Fatal("Render() succeeded, want error")
			}
		})
	}
}
