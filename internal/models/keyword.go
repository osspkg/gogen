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

type Keyword[C config.Config] struct {
	C      C
	Verify bool
	Raw    bool
	D      string
}

func (v *Keyword[C]) Render(w io.Writer) error {
	if v.Verify && !v.C.IsIdentifier(v.D) {
		return fmt.Errorf("invalid identifier: %s", v.D)
	}
	if v.Raw {
		return gen.WriteVerbatim(w, v.D)
	}
	return gen.Render(w, v.D)
}

func (v *Keyword[C]) RenderLayout() gen.Layout {
	kind := gen.KindWord
	if v.Raw {
		kind = gen.KindFragment
	}
	style := gen.Style{
		Kind:             kind,
		Text:             v.D,
		CanEndExpression: v.C.CanEndExpression(v.D),
	}
	return gen.Layout{First: style, Last: style}
}
