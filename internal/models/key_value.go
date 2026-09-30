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

type KeyValue struct {
	Key   types.Token
	Value types.Token
}

func (v *KeyValue) Render(w io.Writer) error {
	return gen.Render(w, []types.Token{
		v.Key,
		gen.Symbol(":", gen.Style{Kind: gen.KindColon}),
		v.Value,
	})
}

func (v *KeyValue) RenderLayout() gen.Layout {
	return gen.LayoutOf([]types.Token{
		v.Key,
		gen.Symbol(":", gen.Style{Kind: gen.KindColon}),
		v.Value,
	})
}
