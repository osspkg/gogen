/*
 *  Copyright (c) 2025-2026 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package typescript

import (
	"io"
	"strings"

	"go.osspkg.com/gogen/internal/gen"
	"go.osspkg.com/gogen/internal/models"
	"go.osspkg.com/gogen/types"
)

// Any appends the any type.
func (v *Tokens) Any() *Tokens { return v.Join(keyword("any")) }

// Any appends the any type.
func Any() *Tokens { return create().Any() }

// Unknown appends the unknown type.
func (v *Tokens) Unknown() *Tokens { return v.Join(keyword("unknown")) }

// Unknown appends the unknown type.
func Unknown() *Tokens { return create().Unknown() }

// Never appends the never type.
func (v *Tokens) Never() *Tokens { return v.Join(keyword("never")) }

// Never appends the never type.
func Never() *Tokens { return create().Never() }

// Void appends the void type.
func (v *Tokens) Void() *Tokens { return v.Join(keyword("void")) }

// Void appends the void type.
func Void() *Tokens { return create().Void() }

// String appends the string type.
func (v *Tokens) String() *Tokens { return v.Join(keyword("string")) }

// String appends the string type.
func String() *Tokens { return create().String() }

// Number appends the number type.
func (v *Tokens) Number() *Tokens { return v.Join(keyword("number")) }

// Number appends the number type.
func Number() *Tokens { return create().Number() }

// Boolean appends the boolean type.
func (v *Tokens) Boolean() *Tokens { return v.Join(keyword("boolean")) }

// Boolean appends the boolean type.
func Boolean() *Tokens { return create().Boolean() }

// BigInt appends the bigint type.
func (v *Tokens) BigInt() *Tokens { return v.Join(keyword("bigint")) }

// BigInt appends the bigint type.
func BigInt() *Tokens { return create().BigInt() }

// Symbol appends the symbol type.
func (v *Tokens) Symbol() *Tokens { return v.Join(keyword("symbol")) }

// Symbol appends the symbol type.
func Symbol() *Tokens { return create().Symbol() }

// ObjectType appends the object type.
func (v *Tokens) ObjectType() *Tokens { return v.Join(keyword("object")) }

// ObjectType appends the object type.
func ObjectType() *Tokens { return create().ObjectType() }

// Null appends the null value.
func (v *Tokens) Null() *Tokens { return v.Join(keyword("null")) }

// Null appends the null value.
func Null() *Tokens { return create().Null() }

// Undefined appends the undefined value.
func (v *Tokens) Undefined() *Tokens { return v.Join(identifier("undefined")) }

// Undefined appends the undefined value.
func Undefined() *Tokens { return create().Undefined() }

// True appends the true literal.
func (v *Tokens) True() *Tokens { return v.Join(keyword("true")) }

// True appends the true literal.
func True() *Tokens { return create().True() }

// False appends the false literal.
func (v *Tokens) False() *Tokens { return v.Join(keyword("false")) }

// False appends the false literal.
func False() *Tokens { return create().False() }

// RecordType creates a Record type from its key and value types.
func (v *Tokens) RecordType(key, value types.Token) *Tokens {
	return v.Join(keyword("Record"), &angleArgs{args: []types.Token{key, value}})
}

// RecordType creates a Record type from its key and value types.
func RecordType(key, value types.Token) *Tokens { return create().RecordType(key, value) }

// MapType creates a Map type from its key and value types.
func (v *Tokens) MapType(key, value types.Token) *Tokens {
	return v.Join(keyword("Map"), &angleArgs{args: []types.Token{key, value}})
}

// MapType creates a Map type from its key and value types.
func MapType(key, value types.Token) *Tokens { return create().MapType(key, value) }

// SetType creates a Set type from its element type.
func (v *Tokens) SetType(value types.Token) *Tokens {
	return v.Join(keyword("Set"), &angleArgs{args: []types.Token{value}})
}

// SetType creates a Set type from its element type.
func SetType(value types.Token) *Tokens { return create().SetType(value) }

// PromiseType creates a Promise type from its resolved value type.
func (v *Tokens) PromiseType(value types.Token) *Tokens {
	return v.Join(keyword("Promise"), &angleArgs{args: []types.Token{value}})
}

// PromiseType creates a Promise type from its resolved value type.
func PromiseType(value types.Token) *Tokens { return create().PromiseType(value) }

// Union creates a union type from its alternatives.
func (v *Tokens) Union(types ...types.Token) *Tokens {
	return appendSeparated(v, types, "|")
}

// Union creates a union type from its alternatives.
func Union(types ...types.Token) *Tokens { return create().Union(types...) }

// Intersection creates an intersection type from its constituent types.
func (v *Tokens) Intersection(types ...types.Token) *Tokens {
	return appendSeparated(v, types, "&")
}

// Intersection creates an intersection type from its constituent types.
func Intersection(types ...types.Token) *Tokens { return create().Intersection(types...) }

func appendSeparated(v *Tokens, values []types.Token, separator string) *Tokens {
	for index, value := range values {
		if index > 0 {
			v = v.Op(separator)
		}
		v = v.Join(value)
	}
	return v
}

type objectLiteral struct{ properties []types.Token }

func (v *objectLiteral) Render(w io.Writer) error { return (&models.Block{D: v.properties}).Render(w) }
func (v *objectLiteral) RenderLayout() gen.Layout {
	return gen.LayoutOf([]types.Token{&models.Block{}})
}

type templateLiteral struct{ parts []types.Token }

func (v *templateLiteral) Render(w io.Writer) error {
	if _, err := io.WriteString(w, "`"); err != nil {
		return err
	}
	for _, part := range v.parts {
		if err := part.Render(w); err != nil {
			return err
		}
	}
	_, err := io.WriteString(w, "`")
	return err
}

func (v *templateLiteral) RenderLayout() gen.Layout {
	style := gen.Style{Kind: gen.KindLiteral}
	return gen.Layout{First: style, Last: style}
}

type templateText struct{ value string }

func (v *templateText) Render(w io.Writer) error {
	var escaped strings.Builder
	for index, r := range v.value {
		if r == '\\' || r == '`' || (r == '$' && index+1 < len(v.value) && v.value[index+1] == '{') {
			escaped.WriteByte('\\')
		}
		escaped.WriteRune(r)
	}
	_, err := io.WriteString(w, escaped.String())
	return err
}
func (v *templateText) RenderLayout() gen.Layout { return gen.Layout{} }

type templateExpression struct{ expression types.Token }

func (v *templateExpression) Render(w io.Writer) error {
	if _, err := io.WriteString(w, "${"); err != nil {
		return err
	}
	if err := v.expression.Render(w); err != nil {
		return err
	}
	_, err := io.WriteString(w, "}")
	return err
}
func (v *templateExpression) RenderLayout() gen.Layout { return gen.Layout{} }

type angleArgs struct{ args []types.Token }

func (v *angleArgs) Render(w io.Writer) error {
	if _, err := io.WriteString(w, "<"); err != nil {
		return err
	}
	for index, arg := range v.args {
		if index > 0 {
			if _, err := io.WriteString(w, ", "); err != nil {
				return err
			}
		}
		if err := arg.Render(w); err != nil {
			return err
		}
	}
	_, err := io.WriteString(w, ">")
	return err
}

func (v *angleArgs) RenderLayout() gen.Layout {
	return gen.Layout{
		First: gen.Style{Kind: gen.KindDot, Text: "<"},
		Last:  gen.Style{Kind: gen.KindFragment, Text: ">"},
	}
}

type lineToken struct{}

func (lineToken) Render(w io.Writer) error {
	_, err := io.WriteString(w, "\n")
	return err
}
func (lineToken) RenderLayout() gen.Layout {
	style := gen.Style{Kind: gen.KindLine, Text: "\n"}
	return gen.Layout{First: style, Last: style}
}
