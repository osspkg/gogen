/*
 *  Copyright (c) 2025-2026 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package models

import (
	"fmt"
	"io"

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

type Raw struct {
	D      string
	T      types.Token
	AT     []types.Token
	Verify bool
}

func (v *Raw) Render(w io.Writer) error {
	if v.Verify && !rexLetter.MatchString(v.D) {
		return fmt.Errorf("invalid letter: %s", v.D)
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

func (v *Raw) RenderLayout() gen.Layout {
	var style gen.Style
	switch {
	case v.Verify:
		style = gen.Style{Kind: gen.KindWord, Text: v.D}
	case v.D == "]":
		style = gen.Style{Kind: gen.KindCloseBracket, Text: v.D}
	case len(v.D) > 0 && (v.D[0] == '[' || v.D == "map["):
		style = gen.Style{Kind: gen.KindTypePrefix, Text: v.D}
	case v.D == "" && v.T != nil:
		return gen.LayoutOf([]types.Token{v.T})
	default:
		style = gen.Style{Kind: gen.KindFragment, Text: v.D}
	}
	return gen.Layout{First: style, Last: style}
}
