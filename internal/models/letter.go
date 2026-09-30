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
	"go.osspkg.com/gogen/types"
)

type Letter struct {
	D string
}

func (v *Letter) Render(w io.Writer) error {
	return gen.Render(w, v.D)
}

func (v *Letter) RenderLayout() gen.Layout {
	kind := gen.KindWord
	if v.D == "\n" {
		kind = gen.KindLine
	}
	style := gen.Style{Kind: kind, Text: v.D}
	return gen.Layout{First: style, Last: style}
}

type Raw[C config.Config] struct {
	C      C
	D      string
	T      types.Token
	AT     []types.Token
	Verify bool
}

func (v *Raw[C]) Render(w io.Writer) error {
	if v.Verify && !v.C.IsIdentifier(v.D) {
		return fmt.Errorf("invalid identifier: %s", v.D)
	}
	if err := gen.Render(w, v.D); err != nil {
		return err
	}
	if v.T != nil {
		if err := v.T.Render(w); err != nil {
			return err
		}
	}
	if len(v.AT) > 0 {
		for _, token := range v.AT {
			if err := token.Render(w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (v *Raw[C]) RenderLayout() gen.Layout {
	if v.D == "" && v.T != nil {
		return gen.LayoutOf([]types.Token{v.T})
	}
	kind := layoutKind(v.C.RawKind(v.D, v.Verify))
	if v.Verify {
		kind = gen.KindWord
	}
	style := gen.Style{
		Kind:             kind,
		Text:             v.D,
		CanEndExpression: v.C.CanEndExpression(v.D),
	}
	return gen.Layout{First: style, Last: style}
}
