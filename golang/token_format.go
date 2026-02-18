/*
 *  Copyright (c) 2025-2026 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package golang

import (
	"go.osspkg.com/gogen/internal/models"
	"go.osspkg.com/gogen/types"
)

func (v *Tokens) Comment(arg string) *Tokens {
	*v = append(*v, &models.Comment[config]{D: arg})
	return v
}

func Comment(arg string) *Tokens {
	return create().Comment(arg)
}

//------------------------------------------------------

func (v *Tokens) Line() *Tokens {
	*v = append(*v, &models.Letter{D: "\n"})
	return v
}

func Line() *Tokens {
	return create().Line()
}

//------------------------------------------------------

func (v *Tokens) Block(args ...types.Token) *Tokens {
	*v = append(*v, &models.Block{D: args})
	return v
}

func Block(args ...types.Token) *Tokens {
	return create().Block(args...)
}
