/*
 *  Copyright (c) 2025-2026 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package golang

import (
	"strconv"

	"go.osspkg.com/gogen/internal/models"
	"go.osspkg.com/gogen/types"
)

// ID appends identifier text checked against the builder's accepted form to the token sequence. Render returns an error if arg is outside the builder's accepted identifier form.
func (v *Tokens) ID(arg string) *Tokens {
	*v = append(*v, &models.Keyword{D: arg, Verify: true})
	return v
}

// ID creates a token sequence containing identifier text checked against the builder's accepted form. Render returns an error if arg is outside the builder's accepted identifier form.
func ID(arg string) *Tokens {
	return create().ID(arg)
}

// Pkg appends a package qualifier ending with a dot to the token sequence. Append the selected identifier after it.
func (v *Tokens) Pkg(arg string) *Tokens {
	if len(arg) == 0 {
		return v
	}
	*v = append(*v, &models.Raw{D: arg, Verify: true})
	return v.Op(".")
}

// Pkg creates a token sequence containing a package qualifier ending with a dot. Append the selected identifier after it.
func Pkg(arg string) *Tokens {
	return create().Pkg(arg)
}

// Type appends the type declaration keyword to the token sequence.
func (v *Tokens) Type() *Tokens {
	*v = append(*v, &models.Keyword{D: "type"})
	return v
}

// Type creates a token sequence containing the type declaration keyword.
func Type() *Tokens {
	return create().Type()
}

// Var appends the var declaration keyword to the token sequence.
func (v *Tokens) Var() *Tokens {
	*v = append(*v, &models.Keyword{D: "var"})
	return v
}

// Var creates a token sequence containing the var declaration keyword.
func Var() *Tokens {
	return create().Var()
}

// Const appends the const declaration keyword to the token sequence.
func (v *Tokens) Const() *Tokens {
	*v = append(*v, &models.Keyword{D: "const"})
	return v
}

// Const creates a token sequence containing the const declaration keyword.
func Const() *Tokens {
	return create().Const()
}

// List appends a comma-separated list without surrounding delimiters to the token sequence. Use Bracket or Call when parentheses are required.
func (v *Tokens) List(args ...types.Token) *Tokens {
	*v = append(*v, &models.Bracket{D: args, Brace: false})
	return v
}

// List creates a token sequence containing a comma-separated list without surrounding delimiters. Use Bracket or Call when parentheses are required.
func List(args ...types.Token) *Tokens {
	return create().List(args...)
}

// Slice appends a slice type prefix to the token sequence.
func (v *Tokens) Slice() *Tokens {
	*v = append(*v, &models.Raw{D: "[]"})
	return v
}

// Slice creates a token sequence containing a slice type prefix.
func Slice() *Tokens {
	return create().Slice()
}

// Array appends an array type with the requested length to the token sequence. Render does not validate that n is a legal Go array length.
func (v *Tokens) Array(n int) *Tokens {
	*v = append(*v, &models.Raw{D: "[" + strconv.Itoa(n) + "]"})
	return v
}

// Array creates a token sequence containing an array type with the requested length. Render does not validate that n is a legal Go array length.
func Array(n int) *Tokens {
	return create().Array(n)
}

// New appends a new call for the supplied type to the token sequence.
func (v *Tokens) New(arg types.Token) *Tokens {
	return v.
		Join(&models.Letter{D: "new"}).
		Bracket(arg)
}

// New creates a token sequence containing a new call for the supplied type.
func New(arg types.Token) *Tokens {
	return create().New(arg)
}

// Make appends a make call for the requested type and sizes to the token sequence. A capacity argument is emitted only when cap is greater than len.
func (v *Tokens) Make(arg types.Token, len, cap int) *Tokens {
	args := make([]types.Token, 0, 3)
	args = append(args, arg, &models.Raw{D: strconv.Itoa(len)})

	if cap > len {
		args = append(args, &models.Raw{D: strconv.Itoa(cap)})
	}

	return v.Join(&models.Letter{D: "make"}).
		Bracket(args...)
}

// Make creates a token sequence containing a make call for the requested type and sizes. A capacity argument is emitted only when cap is greater than len.
func Make(arg types.Token, len, cap int) *Tokens {
	return create().Make(arg, len, cap)
}

// Append appends an append call with the destination and values to the token sequence.
func (v *Tokens) Append(to types.Token, from ...types.Token) *Tokens {
	args := make([]types.Token, 0, len(from)+1)
	args = append(args, to)
	args = append(args, from...)

	return v.Join(&models.Letter{D: "append"}).
		Bracket(args...)
}

// Append creates a token sequence containing an append call with the destination and values.
func Append(to types.Token, from ...types.Token) *Tokens {
	return create().Append(to, from...)
}
