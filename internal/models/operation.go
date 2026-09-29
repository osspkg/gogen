/*
 *  Copyright (c) 2025-2026 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package models

import (
	"fmt"
	"io"

	"go.osspkg.com/gogen/internal/config"
	"go.osspkg.com/gogen/internal/gen"
)

type Operation[C config.Config] struct {
	c C
	D string
}

func (v *Operation[C]) Render(w io.Writer) error {
	if !v.c.OperationAvailable(v.D) {
		return fmt.Errorf("invalid operation: %s", v.D)
	}
	return gen.Render(w, v.D)
}

func (v *Operation[C]) RenderLayout() gen.Layout {
	style := gen.Style{Kind: gen.KindOperator, Text: v.D}
	switch v.D {
	case ".":
		style.Kind = gen.KindDot
	case ",":
		style.Kind = gen.KindComma
	case ":":
		style.Kind = gen.KindColon
	case ";":
		style.Kind = gen.KindSemicolon
	case "(":
		style.Kind = gen.KindOpenParen
	case ")":
		style.Kind = gen.KindCloseParen
	case "[":
		style.Kind = gen.KindOpenSquare
	case "]":
		style.Kind = gen.KindCloseBracket
	case "{":
		style.Kind = gen.KindBlockOpen
	case "}":
		style.Kind = gen.KindBlockClose
	case "++", "--", "...":
		style.Kind = gen.KindPostfixOperator
	}
	return gen.Layout{First: style, Last: style}
}
