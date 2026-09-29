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

// Tokens is a chainable sequence of Go source tokens.
type Tokens []types.Token

func create() *Tokens {
	return &Tokens{}
}

// Render writes the token sequence to w without applying go/format.
func (v *Tokens) Render(w io.Writer) error {
	return gen.Render(w, []types.Token(*v))
}

// Unwrap returns the underlying token slice without copying it.
func (v *Tokens) Unwrap() []types.Token {
	return *v
}

// Join appends each supplied token to the sequence and returns it.
func (v *Tokens) Join(args ...types.Token) *Tokens {
	*v = append(*v, args...)
	return v
}
