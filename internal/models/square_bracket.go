/*
 *  Copyright (c) 2025-2026 Mikhail Knyazhev <markus621@yandex.ru>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package models

import (
	"io"

	"go.osspkg.com/gogen/internal/gen"
	"go.osspkg.com/gogen/types"
)

type SquareBracket struct {
	D []types.Token
}

func (v *SquareBracket) Render(w io.Writer) error {
	out := make([]types.Token, 0, len(v.D)*2+2)
	out = append(out, gen.Symbol("[", gen.Style{Kind: gen.KindOpenSquare}))
	for i, token := range v.D {
		if i > 0 {
			out = append(out, gen.Symbol(",", gen.Style{Kind: gen.KindComma}))
		}
		out = append(out, gen.Params(token)...)
	}
	out = append(out, gen.Symbol("]", gen.Style{Kind: gen.KindCloseBracket}))
	return gen.Render(w, out)
}

func (v *SquareBracket) RenderLayout() gen.Layout {
	return gen.Layout{
		First: gen.Style{Kind: gen.KindOpenSquare},
		Last:  gen.Style{Kind: gen.KindCloseBracket},
	}
}
