/*
 *  Copyright (c) 2025-2026 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package models

import (
	"io"

	"go.osspkg.com/gogen/internal/gen"
	"go.osspkg.com/gogen/types"
)

type DeclarationBlock struct {
	Keyword string
	D       []types.Token
}

func (v *DeclarationBlock) Render(w io.Writer) error {
	iw := gen.Indented(w)
	if err := gen.Render(iw, v.Keyword+" ("); err != nil {
		return err
	}
	if len(v.D) == 0 {
		return gen.Render(iw, gen.Symbol(")", gen.Style{Kind: gen.KindCloseParen}))
	}

	iw.Push()
	for _, token := range v.D {
		if err := iw.EnsureNewline(); err != nil {
			iw.Pop()
			return err
		}
		if err := gen.Render(iw, token); err != nil {
			iw.Pop()
			return err
		}
	}
	if err := iw.EnsureNewline(); err != nil {
		iw.Pop()
		return err
	}
	iw.Pop()
	return gen.Render(iw, gen.Symbol(")", gen.Style{Kind: gen.KindCloseParen}))
}

func (v *DeclarationBlock) RenderLayout() gen.Layout {
	return gen.Layout{
		First: gen.Style{Kind: gen.KindWord, Text: v.Keyword},
		Last:  gen.Style{Kind: gen.KindCloseParen, Text: ")"},
	}
}
