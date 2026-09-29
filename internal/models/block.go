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

type Block struct {
	D []types.Token
}

func (v *Block) Render(w io.Writer) error {
	iw := gen.Indented(w)
	if len(v.D) == 0 {
		return gen.Render(iw, []types.Token{
			gen.Symbol("{", gen.Style{Kind: gen.KindBlockOpen}),
			gen.Symbol("}", gen.Style{Kind: gen.KindBlockClose}),
		})
	}

	if err := gen.Render(iw, gen.Symbol("{", gen.Style{Kind: gen.KindBlockOpen})); err != nil {
		return err
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
	return gen.Render(iw, gen.Symbol("}", gen.Style{Kind: gen.KindBlockClose}))
}

func (v *Block) RenderLayout() gen.Layout {
	return gen.Layout{
		First: gen.Style{Kind: gen.KindBlockOpen},
		Last:  gen.Style{Kind: gen.KindBlockClose},
	}
}

type Bracket struct {
	D     []types.Token
	Brace bool
}

func (v *Bracket) Render(w io.Writer) error {
	out := make([]types.Token, 0, len(v.D)*2+2)
	if v.Brace {
		out = append(out, gen.Symbol("(", gen.Style{Kind: gen.KindOpenParen}))
	}
	for i, token := range v.D {
		if i > 0 {
			out = append(out, gen.Symbol(",", gen.Style{Kind: gen.KindComma}))
		}
		out = append(out, gen.Params(token)...)
	}
	if v.Brace {
		out = append(out, gen.Symbol(")", gen.Style{Kind: gen.KindCloseParen}))
	}
	return gen.Render(w, out)
}

func (v *Bracket) RenderLayout() gen.Layout {
	if v.Brace {
		return gen.Layout{
			First: gen.Style{Kind: gen.KindOpenParen},
			Last:  gen.Style{Kind: gen.KindCloseParen},
		}
	}
	return gen.Layout{}
}
