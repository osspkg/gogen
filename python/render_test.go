/*
 *  Copyright (c) 2025-2026 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package python_test

import (
	"bytes"
	"errors"
	"io"
	"testing"

	py "go.osspkg.com/gogen/python"
	"go.osspkg.com/gogen/types"
)

func render(t *testing.T, token interface{ Render(io.Writer) error }) string {
	t.Helper()
	var output bytes.Buffer
	if err := token.Render(&output); err != nil {
		t.Fatal(err)
	}
	return output.String()
}

func TestRenderDeclarationsAndSuites(t *testing.T) {
	file := py.FromImport("typing", py.ID("Sequence").As().ID("Seq")).Line().Join(
		py.Decorator(py.ID("trace")),
		py.Def("greet").Bracket(py.ID("name").Colon().ID("str")).Arrow().TypeUnion(py.ID("str"), py.NoneValue()).Block(
			py.If(py.ID("name").Op("==").Text("")).Block(
				py.Return().Text("anonymous"),
			).Elif(py.ID("name").Op("==").Text("admin")).Block(
				py.Return().Text("welcome"),
			).Else().Block(
				py.Return().Op("not").ID("name"),
			),
		),
	)
	want := "from typing import Sequence as Seq\n@trace\ndef greet(name: str) -> str | None:\n    if name == \"\":\n        return \"anonymous\"\n    elif name == \"admin\":\n        return \"welcome\"\n    else:\n        return not name\n"
	if got := render(t, file); got != want {
		t.Fatalf("render() =\n%s\nwant:\n%s", got, want)
	}
}

func TestRenderImportsClassesAndAsync(t *testing.T) {
	file := py.ImportAs("collections.abc", "abc").Line().Join(
		py.Class("Repository", py.ID("abc").Selector("Sequence")).Block(
			py.AsyncDef("load").Bracket(py.ID("self"), py.ID("key").Colon().ID("str")).Arrow().ID("bytes").Block(
				py.With(py.ID("open").Call(py.ID("key"), py.Text("rb")).As().ID("stream")).Block(
					py.Return().Await().ID("stream").Selector("read").Call(),
				),
			),
		),
	)
	want := "import collections.abc as abc\nclass Repository(abc.Sequence):\n    async def load(self, key: str) -> bytes:\n        with open(key, \"rb\") as stream:\n            return await stream.read()\n"
	if got := render(t, file); got != want {
		t.Fatalf("render() =\n%s\nwant:\n%s", got, want)
	}
}

func TestRenderTryMatchAndEmptySuite(t *testing.T) {
	token := py.Try().Block(
		py.Pass(),
	).Except(py.ID("ValueError")).Block(
		py.Raise(),
	).Finally().Block(
		py.Pass(),
	).Join(
		py.Match(py.ID("value")).Block(
			py.Case(py.Raw("0")).Block(py.Pass()),
			py.Case(py.Raw("_")).Block(py.Pass()),
		),
		py.Def("empty").Bracket().Block(),
	)
	want := "try:\n    pass\nexcept ValueError:\n    raise\nfinally:\n    pass\nmatch value:\n    case 0:\n        pass\n    case _:\n        pass\ndef empty(): pass\n"
	if got := render(t, token); got != want {
		t.Fatalf("render() =\n%s\nwant:\n%s", got, want)
	}
}

func TestRenderContainersComprehensionsAndFString(t *testing.T) {
	tests := []struct {
		name  string
		token types.Token
		want  string
	}{
		{name: "list", token: py.ListLiteral(py.Raw("1"), py.Raw("2")), want: "[1, 2]"},
		{name: "empty set", token: py.SetLiteral(), want: "set()"},
		{name: "set", token: py.SetLiteral(py.Raw("1"), py.Raw("2")), want: "{1, 2}"},
		{name: "dict", token: py.DictLiteral(py.KeyValue(py.Text("x"), py.Raw("1"))), want: `{"x": 1}`},
		{name: "tuple singleton", token: py.TupleLiteral(py.ID("item")), want: "(item,)"},
		{name: "list comprehension", token: py.ListComp(py.ID("x").Op("*").Raw("2"), py.ForClause(py.ID("x"), py.ID("items")), py.IfClause(py.ID("x").Op(">").Raw("0"))), want: "[x * 2 for x in items if x > 0]"},
		{name: "async generator", token: py.GeneratorExpr(py.ID("x"), py.AsyncForClause(py.ID("x"), py.ID("items"))), want: "(x async for x in items)"},
		{name: "dict comprehension", token: py.DictComp(py.ID("x"), py.ID("x").Op("**").Raw("2"), py.ForClause(py.ID("x"), py.ID("values"))), want: "{x: x ** 2 for x in values}"},
		{name: "f-string", token: py.FString(py.FStringText("hello {"), py.FStringExpr(py.ID("name")), py.FStringText("}")), want: `f"hello {{{name}}}"`},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := render(t, test.token); got != test.want {
				t.Fatalf("render() = %q, want %q", got, test.want)
			}
		})
	}
}

func TestRenderEscapingOperatorsAndRaw(t *testing.T) {
	if got, want := render(t, py.Text("quote: \"; slash: \\; line:\n")), "\"quote: \\\"; slash: \\\\; line:\\n\""; got != want {
		t.Fatalf("string = %q, want %q", got, want)
	}
	if got, want := render(t, py.ID("变量").Op("and").Op("not").ID("ready")), "变量 and not ready"; got != want {
		t.Fatalf("word operators = %q, want %q", got, want)
	}
	if got, want := render(t, py.Raw("first\n  second")), "first\n  second"; got != want {
		t.Fatalf("Raw = %q, want %q", got, want)
	}
	if got, want := render(t, py.Comment("first\nsecond")), "# first\n# second\n"; got != want {
		t.Fatalf("comment = %q, want %q", got, want)
	}
}

func TestValidationAndRenderErrors(t *testing.T) {
	for _, token := range []interface{ Render(io.Writer) error }{
		py.ID("class"),
		py.ID("not-valid"),
		py.Op("???"),
		py.ListComp(py.ID("item")),
		py.FString(py.FStringExpr(nil)),
		py.Def("run").Block(nil),
	} {
		var output bytes.Buffer
		if err := token.Render(&output); err == nil {
			t.Errorf("Render(%T) succeeded, want error", token)
		}
	}

	var output bytes.Buffer
	if err := py.Render(&output, py.ID("value")); err != nil {
		t.Fatal(err)
	}
	if got := output.String(); got != "value" {
		t.Fatalf("package Render() = %q, want %q", got, "value")
	}

	wantErr := errors.New("write failed")
	if err := py.Render(errorWriter{err: wantErr}, py.ID("value")); !errors.Is(err, wantErr) {
		t.Fatalf("Render() error = %v, want %v", err, wantErr)
	}
	if err := py.Render(shortWriter{}, py.Comment("comment")); !errors.Is(err, io.ErrShortWrite) {
		t.Fatalf("Render() short-write error = %v, want %v", err, io.ErrShortWrite)
	}
}

type errorWriter struct{ err error }

func (v errorWriter) Write([]byte) (int, error) { return 0, v.err }

type shortWriter struct{}

func (shortWriter) Write([]byte) (int, error) { return 0, nil }
