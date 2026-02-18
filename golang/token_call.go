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

func Defer() *Tokens {
	return create().__defer()
}

func (v *Tokens) __go() *Tokens {
	*v = append(*v, &models.Keyword{D: "go"})
	return v
}

func Go() *Tokens {
	return create().__go()
}

func (v *Tokens) Bracket(args ...types.Token) *Tokens {
	*v = append(*v, &models.Bracket{D: args, Brace: true})
	return v
}

func Params(args ...types.Token) *Tokens {
	return create().Bracket(args...)
}

func (v *Tokens) Call(args ...types.Token) *Tokens {
	return v.Bracket(args...)
}

func Call(args ...types.Token) *Tokens {
	return create().Call(args...)
}

func (v *Tokens) Func() *Tokens {
	*v = append(*v, &models.Keyword{D: "func"})
	return v
}

func Func() *Tokens {
	return create().Func()
}

func (v *Tokens) Return() *Tokens {
	*v = append(*v, &models.Keyword{D: "return"})
	return v
}

func Return() *Tokens {
	return create().Return()
}
