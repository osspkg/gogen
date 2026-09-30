/*
 *  Copyright (c) 2025-2026 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package python

import (
	"fmt"
	"io"

	"go.osspkg.com/gogen/internal/gen"
	"go.osspkg.com/gogen/internal/models"
	"go.osspkg.com/gogen/types"
)

// ListLiteral creates a Python list literal.
func (v *Tokens) ListLiteral(values ...types.Token) *Tokens {
	return v.Join(&models.SquareBracket{D: values})
}

// ListLiteral creates a Python list literal.
func ListLiteral(values ...types.Token) *Tokens { return create().ListLiteral(values...) }

// DictLiteral creates a Python dictionary literal.
func (v *Tokens) DictLiteral(entries ...types.Token) *Tokens {
	return v.Join(&dictLiteral{entries: entries})
}

// DictLiteral creates a Python dictionary literal.
func DictLiteral(entries ...types.Token) *Tokens { return create().DictLiteral(entries...) }

// SetLiteral creates a Python set literal. An empty set is emitted as set().
func (v *Tokens) SetLiteral(values ...types.Token) *Tokens {
	return v.Join(&setLiteral{values: values})
}

// SetLiteral creates a Python set literal. An empty set is emitted as set().
func SetLiteral(values ...types.Token) *Tokens { return create().SetLiteral(values...) }

// ListComp creates a list comprehension from an expression and clauses.
func (v *Tokens) ListComp(expression types.Token, clauses ...types.Token) *Tokens {
	return v.Join(&comprehension{open: "[", close: "]", expression: expression, clauses: clauses})
}

// ListComp creates a list comprehension from an expression and clauses.
func ListComp(expression types.Token, clauses ...types.Token) *Tokens {
	return create().ListComp(expression, clauses...)
}

// SetComp creates a set comprehension from an expression and clauses.
func (v *Tokens) SetComp(expression types.Token, clauses ...types.Token) *Tokens {
	return v.Join(&comprehension{open: "{", close: "}", expression: expression, clauses: clauses})
}

// SetComp creates a set comprehension from an expression and clauses.
func SetComp(expression types.Token, clauses ...types.Token) *Tokens {
	return create().SetComp(expression, clauses...)
}

// DictComp creates a dictionary comprehension from key/value expressions and clauses.
func (v *Tokens) DictComp(key, value types.Token, clauses ...types.Token) *Tokens {
	return v.Join(&dictComprehension{key: key, value: value, clauses: clauses})
}

// DictComp creates a dictionary comprehension from key/value expressions and clauses.
func DictComp(key, value types.Token, clauses ...types.Token) *Tokens {
	return create().DictComp(key, value, clauses...)
}

// GeneratorExpr creates a generator expression from an expression and clauses.
func (v *Tokens) GeneratorExpr(expression types.Token, clauses ...types.Token) *Tokens {
	return v.Join(&comprehension{open: "(", close: ")", expression: expression, clauses: clauses})
}

// GeneratorExpr creates a generator expression from an expression and clauses.
func GeneratorExpr(expression types.Token, clauses ...types.Token) *Tokens {
	return create().GeneratorExpr(expression, clauses...)
}

// ForClause creates the for-in clause used by comprehensions.
func ForClause(target, iterable types.Token) types.Token {
	return &forClause{target: target, iterable: iterable}
}

// AsyncForClause creates an asynchronous for-in clause used by comprehensions.
func AsyncForClause(target, iterable types.Token) types.Token {
	return &forClause{target: target, iterable: iterable, async: true}
}

// IfClause creates a filter clause used by comprehensions.
func IfClause(condition types.Token) types.Token { return &ifClause{condition: condition} }

type dictLiteral struct{ entries []types.Token }

func (v *dictLiteral) Render(w io.Writer) error {
	if _, err := io.WriteString(w, "{"); err != nil {
		return err
	}
	if err := renderCommaSeparated(w, v.entries); err != nil {
		return err
	}
	_, err := io.WriteString(w, "}")
	return err
}

func (v *dictLiteral) RenderLayout() gen.Layout {
	return containerLayout(gen.KindBlockOpen, gen.KindBlockClose)
}

type setLiteral struct{ values []types.Token }

func (v *setLiteral) Render(w io.Writer) error {
	if len(v.values) == 0 {
		_, err := io.WriteString(w, "set()")
		return err
	}
	if _, err := io.WriteString(w, "{"); err != nil {
		return err
	}
	if err := renderCommaSeparated(w, v.values); err != nil {
		return err
	}
	_, err := io.WriteString(w, "}")
	return err
}

func (v *setLiteral) RenderLayout() gen.Layout {
	if len(v.values) == 0 {
		style := gen.Style{Kind: gen.KindWord, Text: "set", CanEndExpression: true}
		return gen.Layout{First: style, Last: gen.Style{Kind: gen.KindCloseParen}}
	}
	return containerLayout(gen.KindBlockOpen, gen.KindBlockClose)
}

type comprehension struct {
	open       string
	close      string
	expression types.Token
	clauses    []types.Token
}

func (v *comprehension) Render(w io.Writer) error {
	if v.expression == nil || len(v.clauses) == 0 {
		return fmt.Errorf("comprehension requires an expression and at least one clause")
	}
	return renderContainer(w, v.open, v.close, append([]types.Token{v.expression}, v.clauses...))
}

func (v *comprehension) RenderLayout() gen.Layout {
	return containerLayoutFor(v.open, v.close)
}

type dictComprehension struct {
	key     types.Token
	value   types.Token
	clauses []types.Token
}

func (v *dictComprehension) Render(w io.Writer) error {
	if v.key == nil || v.value == nil || len(v.clauses) == 0 {
		return fmt.Errorf("dictionary comprehension requires key, value, and at least one clause")
	}
	entries := []types.Token{keyed(v.key, v.value)}
	entries = append(entries, v.clauses...)
	return renderContainer(w, "{", "}", entries)
}

func (v *dictComprehension) RenderLayout() gen.Layout {
	return containerLayout(gen.KindBlockOpen, gen.KindBlockClose)
}

type forClause struct {
	target   types.Token
	iterable types.Token
	async    bool
}

func (v *forClause) Render(w io.Writer) error {
	if v.target == nil || v.iterable == nil {
		return fmt.Errorf("comprehension for clause requires target and iterable")
	}
	if v.async {
		return gen.Render(w, []types.Token{keyword("async"), keyword("for"), v.target, keyword("in"), v.iterable})
	}
	return gen.Render(w, []types.Token{keyword("for"), v.target, keyword("in"), v.iterable})
}

func (v *forClause) RenderLayout() gen.Layout {
	return gen.Layout{First: wordStyle("for"), Last: gen.LayoutOf([]types.Token{v.iterable}).Last}
}

type ifClause struct{ condition types.Token }

func (v *ifClause) Render(w io.Writer) error {
	if v.condition == nil {
		return fmt.Errorf("comprehension if clause requires condition")
	}
	return gen.Render(w, []types.Token{keyword("if"), v.condition})
}

func (v *ifClause) RenderLayout() gen.Layout {
	return gen.Layout{First: wordStyle("if"), Last: gen.LayoutOf([]types.Token{v.condition}).Last}
}

func renderCommaSeparated(w io.Writer, values []types.Token) error {
	for index, value := range values {
		if value == nil {
			return fmt.Errorf("container contains a nil value")
		}
		if index > 0 {
			if _, err := io.WriteString(w, ", "); err != nil {
				return err
			}
		}
		if err := value.Render(w); err != nil {
			return err
		}
	}
	return nil
}

func renderContainer(w io.Writer, open, close string, values []types.Token) error {
	if _, err := io.WriteString(w, open); err != nil {
		return err
	}
	if err := gen.Render(w, values); err != nil {
		return err
	}
	_, err := io.WriteString(w, close)
	return err
}

func containerLayout(first, last gen.Kind) gen.Layout {
	return gen.Layout{
		First: gen.Style{Kind: first},
		Last:  gen.Style{Kind: last},
	}
}

func containerLayoutFor(open, close string) gen.Layout {
	first, last := gen.KindUnknown, gen.KindUnknown
	switch open {
	case "[":
		first = gen.KindOpenSquare
	case "(":
		first = gen.KindOpenParen
	case "{":
		first = gen.KindBlockOpen
	}
	switch close {
	case "]":
		last = gen.KindCloseBracket
	case ")":
		last = gen.KindCloseParen
	case "}":
		last = gen.KindBlockClose
	}
	return containerLayout(first, last)
}
