/*
 *  Copyright (c) 2025-2026 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package python

import (
	"io"

	"go.osspkg.com/gogen/internal/gen"
	"go.osspkg.com/gogen/types"
)

// ID appends a validated Python identifier.
func (v *Tokens) ID(name string) *Tokens { return v.Join(identifier(name)) }

// ID creates a token sequence containing a validated Python identifier.
func ID(name string) *Tokens { return create().ID(name) }

// Raw appends source text verbatim without validating it.
func (v *Tokens) Raw(source string) *Tokens { return v.Join(raw(source)) }

// Raw creates a token sequence containing source text verbatim.
func Raw(source string) *Tokens { return create().Raw(source) }

// Text appends an escaped Python string literal.
func (v *Tokens) Text(value string) *Tokens { return v.Join(text(value)) }

// Text creates a token sequence containing an escaped Python string literal.
func Text(value string) *Tokens { return create().Text(value) }

// Op appends a supported Python operator.
func (v *Tokens) Op(op string) *Tokens { return v.Join(operation(op)) }

// Op creates a token sequence containing a supported Python operator.
func Op(op string) *Tokens { return create().Op(op) }

// Call appends a parenthesized call argument list.
func (v *Tokens) Call(args ...types.Token) *Tokens { return v.Join(brackets(args...)) }

// Call creates a parenthesized call argument list.
func Call(args ...types.Token) *Tokens { return create().Call(args...) }

// Bracket appends a parenthesized token list.
func (v *Tokens) Bracket(args ...types.Token) *Tokens { return v.Join(brackets(args...)) }

// Bracket creates a parenthesized token list.
func Bracket(args ...types.Token) *Tokens { return create().Bracket(args...) }

// List appends a comma-separated sequence without delimiters.
func (v *Tokens) List(args ...types.Token) *Tokens { return v.Join(list(args...)) }

// List creates a comma-separated sequence without delimiters.
func List(args ...types.Token) *Tokens { return create().List(args...) }

// Index appends a square-bracketed index or type argument.
func (v *Tokens) Index(index types.Token) *Tokens { return v.Join(square(index)) }

// Index creates a square-bracketed index or type argument.
func Index(index types.Token) *Tokens { return create().Index(index) }

// KeyValue creates a dictionary or keyword-argument entry.
func KeyValue(key, value types.Token) *Tokens { return create().Join(keyed(key, value)) }

// Selector appends an attribute selector.
func (v *Tokens) Selector(name string) *Tokens { return v.Join(operation("."), identifier(name)) }

// Selector creates an attribute selector.
func Selector(name string) *Tokens { return create().Selector(name) }

// Pkg appends a dotted module or package name.
func (v *Tokens) Pkg(module string) *Tokens {
	if module == "" {
		return v
	}
	for i, part := range splitDots(module) {
		if i > 0 {
			v = v.Join(operation("."))
		}
		v = v.Join(identifier(part))
	}
	return v
}

// Pkg creates a dotted module or package name.
func Pkg(module string) *Tokens { return create().Pkg(module) }

// TupleLiteral creates a Python tuple literal.
func (v *Tokens) TupleLiteral(values ...types.Token) *Tokens {
	return v.Join(&tupleLiteral{values: values})
}

// TupleLiteral creates a Python tuple literal.
func TupleLiteral(values ...types.Token) *Tokens { return create().TupleLiteral(values...) }

// TypeUnion creates a Python 3.10 union annotation.
func (v *Tokens) TypeUnion(values ...types.Token) *Tokens {
	for index, value := range values {
		if index > 0 {
			v = v.Op("|")
		}
		v = v.Join(value)
	}
	return v
}

// TypeUnion creates a Python 3.10 union annotation.
func TypeUnion(values ...types.Token) *Tokens { return create().TypeUnion(values...) }

func splitDots(value string) []string {
	parts := make([]string, 0, 1)
	start := 0
	for i, r := range value {
		if r == '.' {
			parts = append(parts, value[start:i])
			start = i + 1
		}
	}
	return append(parts, value[start:])
}

type tupleLiteral struct{ values []types.Token }

func (v *tupleLiteral) Render(w io.Writer) error {
	out := make([]types.Token, 0, len(v.values)*2+3)
	out = append(out, gen.Symbol("(", gen.Style{Kind: gen.KindOpenParen}))
	for index, value := range v.values {
		if index > 0 {
			out = append(out, gen.Symbol(",", gen.Style{Kind: gen.KindComma}))
		}
		out = append(out, value)
	}
	if len(v.values) == 1 {
		out = append(out, gen.Symbol(",", gen.Style{Kind: gen.KindComma}))
	}
	out = append(out, gen.Symbol(")", gen.Style{Kind: gen.KindCloseParen}))
	return gen.Render(w, out)
}

func (v *tupleLiteral) RenderLayout() gen.Layout {
	return gen.Layout{
		First: gen.Style{Kind: gen.KindOpenParen},
		Last:  gen.Style{Kind: gen.KindCloseParen},
	}
}
