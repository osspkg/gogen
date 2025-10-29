/*
 *  Copyright (c) 2025 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
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
	out := make([]any, 0)
	out = append(out, "{\n")
	for _, token := range v.D {
		out = append(out, token)
	}
	out = append(out, "}")
	return gen.Render(w, out...)
}

type Params struct {
	D     []types.Token
	Brace bool
}

func (v *Params) Render(w io.Writer) error {
	count := len(v.D) - 1
	out := make([]any, 0, count*2+2)
	if v.Brace {
		out = append(out, "(")
	}
	for i, token := range v.D {
		out = append(out, gen.Params(token))
		if i < count {
			out = append(out, ", ")
		}
	}
	if v.Brace {
		out = append(out, ")")
	}
	return gen.Render(w, out...)
}
