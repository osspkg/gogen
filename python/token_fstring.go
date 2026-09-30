/*
 *  Copyright (c) 2025-2026 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package python

import (
	"fmt"
	"io"
	"strings"

	"go.osspkg.com/gogen/internal/gen"
	"go.osspkg.com/gogen/types"
)

// FString creates a Python f-string from text and expression parts.
func (v *Tokens) FString(parts ...types.Token) *Tokens { return v.Join(&fString{parts: parts}) }

// FString creates a Python f-string from text and expression parts.
func FString(parts ...types.Token) *Tokens { return create().FString(parts...) }

// FStringText creates an escaped text part for FString.
func FStringText(value string) types.Token { return fStringText(value) }

// FStringExpr creates an expression part for FString.
func FStringExpr(expression types.Token) types.Token { return fStringExpr{expression: expression} }

type fString struct{ parts []types.Token }

func (v *fString) Render(w io.Writer) error {
	if err := writeString(w, "f\""); err != nil {
		return err
	}
	for _, part := range v.parts {
		if part == nil {
			return fmt.Errorf("f-string contains a nil part")
		}
		if err := part.Render(w); err != nil {
			return err
		}
	}
	err := writeString(w, "\"")
	return err
}

func (v *fString) RenderLayout() gen.Layout {
	style := gen.Style{Kind: gen.KindLiteral}
	return gen.Layout{First: style, Last: style}
}

type fStringText string

func (v fStringText) Render(w io.Writer) error {
	var out strings.Builder
	for _, r := range string(v) {
		switch r {
		case '\\':
			out.WriteString("\\\\")
		case '"':
			out.WriteString("\\\"")
		case '{':
			out.WriteString("{{")
		case '}':
			out.WriteString("}}")
		case '\n':
			out.WriteString("\\n")
		case '\r':
			out.WriteString("\\r")
		case '\t':
			out.WriteString("\\t")
		default:
			out.WriteRune(r)
		}
	}
	err := writeString(w, out.String())
	return err
}

type fStringExpr struct{ expression types.Token }

func (v fStringExpr) Render(w io.Writer) error {
	if v.expression == nil {
		return fmt.Errorf("f-string expression is nil")
	}
	if err := writeString(w, "{"); err != nil {
		return err
	}
	if err := v.expression.Render(w); err != nil {
		return err
	}
	err := writeString(w, "}")
	return err
}
