/*
 *  Copyright (c) 2025-2026 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package golang

import (
	"io"

	"go.osspkg.com/gogen/internal/gen"
	"go.osspkg.com/gogen/types"
)

type Tokens []types.Token

func create() *Tokens {
	return &Tokens{}
}

func (v *Tokens) Render(w io.Writer) error {
	for _, token := range *v {
		if err := gen.Render(w, token); err != nil {
			return err
		}
	}
	return nil
}

func (v *Tokens) Unwrap() []types.Token {
	return *v
}

func (v *Tokens) Join(args ...types.Token) *Tokens {
	*v = append(*v, args...)
	return v
}
