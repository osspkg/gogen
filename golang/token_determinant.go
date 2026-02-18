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

func (v *Tokens) ID(arg string) *Tokens {
	*v = append(*v, &models.Keyword{D: arg, Verify: true})
	return v
}

func ID(arg string) *Tokens {
	return create().ID(arg)
}

func (v *Tokens) Pkg(arg string) *Tokens {
	*v = append(*v, &models.Keyword{D: arg, Verify: true})
	return v.Op(".")
}

func Pkg(arg string) *Tokens {
	return create().Pkg(arg)
}

func (v *Tokens) Type() *Tokens {
	*v = append(*v, &models.Keyword{D: "type"})
	return v
}

func Type() *Tokens {
	return create().Type()
}

func (v *Tokens) Var() *Tokens {
	*v = append(*v, &models.Keyword{D: "var"})
	return v
}

func Var() *Tokens {
	return create().Var()
}

func (v *Tokens) Const() *Tokens {
	*v = append(*v, &models.Keyword{D: "const"})
	return v
}

func (v *Tokens) List(args ...types.Token) *Tokens {
	*v = append(*v, &models.Bracket{D: args, Brace: false})
	return v
}

func List(args ...types.Token) *Tokens {
	return create().List(args...)
}

func (v *Tokens) Slice() *Tokens {
	*v = append(*v, &models.Raw{D: "[]"})
	return v
}

func Slice() *Tokens {
	return create().Slice()
}

func (v *Tokens) Array(n int) *Tokens {
	*v = append(*v, &models.Raw{D: "[" + strconv.Itoa(n) + "]"})
	return v
}

func Array(n int) *Tokens {
	return create().Array(n)
}

func (v *Tokens) New(arg types.Token) *Tokens {
	return v.
		Join(&models.Letter{D: "new"}).
		Bracket(arg)
}

func New(arg types.Token) *Tokens {
	return create().New(arg)
}

func (v *Tokens) Make(arg types.Token, len, cap int) *Tokens {
	args := make([]types.Token, 0, 3)
	args = append(args, arg, &models.Raw{D: strconv.Itoa(len)})

	if cap > len {
		args = append(args, &models.Raw{D: strconv.Itoa(cap)})
	}

	return v.Join(&models.Letter{D: "make"}).
		Bracket(args...)
}

func Make(arg types.Token, len, cap int) *Tokens {
	return create().Make(arg, len, cap)
}

func (v *Tokens) Append(to types.Token, from ...types.Token) *Tokens {
	args := make([]types.Token, 0, len(from)+1)
	args = append(args, to)
	args = append(args, from...)

	return v.Join(&models.Letter{D: "append"}).
		Bracket(args...)
}

func Append(to types.Token, from ...types.Token) *Tokens {
	return create().Append(to, from...)
}
