/*
 *  Copyright (c) 2025-2026 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package golang

import (
	"go.osspkg.com/gogen/internal/models"
	"go.osspkg.com/gogen/types"
)

func (v *Tokens) __defer() *Tokens {
	*v = append(*v, &models.Keyword{D: "defer"})
	return v
}

// Defer creates a token sequence containing the defer keyword.
func Defer() *Tokens {
	return create().__defer()
}

func (v *Tokens) __go() *Tokens {
	*v = append(*v, &models.Keyword{D: "go"})
	return v
}

// Go creates a token sequence containing the go keyword for starting a goroutine.
func Go() *Tokens {
	return create().__go()
}

// Bracket appends parentheses around the supplied comma-separated tokens to the token sequence. Use it for parameter lists and call arguments.
func (v *Tokens) Bracket(args ...types.Token) *Tokens {
	*v = append(*v, &models.Bracket{D: args, Brace: true})
	return v
}

// Params creates a token sequence containing a parenthesized comma-separated parameter list. Use it when composing a function signature.
func Params(args ...types.Token) *Tokens {
	return create().Bracket(args...)
}

// Call appends a call expression with the supplied arguments to the token sequence.
func (v *Tokens) Call(args ...types.Token) *Tokens {
	return v.Bracket(args...)
}

// Call creates a token sequence containing a call expression with the supplied arguments.
func Call(args ...types.Token) *Tokens {
	return create().Call(args...)
}

// Func appends the func keyword to the token sequence.
func (v *Tokens) Func() *Tokens {
	*v = append(*v, &models.Keyword{D: "func"})
	return v
}

// Func creates a token sequence containing the func keyword.
func Func() *Tokens {
	return create().Func()
}

// Return appends the return keyword to the token sequence.
func (v *Tokens) Return() *Tokens {
	*v = append(*v, &models.Keyword{D: "return"})
	return v
}

// Return creates a token sequence containing the return keyword.
func Return() *Tokens {
	return create().Return()
}
