/*
 *  Copyright (c) 2025-2026 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package models

import (
	"io"

	"go.osspkg.com/gogen/internal/config"
	"go.osspkg.com/gogen/internal/gen"
)

type Text[C config.Config] struct {
	C C
	D string
}

func (v *Text[C]) Render(w io.Writer) error {
	return gen.Render(w, v.C.QuoteString(v.D))
}

func (v *Text[C]) RenderLayout() gen.Layout {
	style := gen.Style{Kind: gen.KindLiteral, Text: v.D}
	return gen.Layout{First: style, Last: style}
}
