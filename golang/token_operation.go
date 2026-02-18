/*
 *  Copyright (c) 2025-2026 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package golang

import (
	"go.osspkg.com/gogen/internal/models"
)

func (v *Tokens) Op(arg string) *Tokens {
	*v = append(*v, &models.Operation[config]{D: arg})
	return v
}

func Op(arg string) *Tokens {
	return create().Op(arg)
}

func (v *Tokens) Raw(arg string) *Tokens {
	*v = append(*v, &models.Keyword{D: arg})
	return v
}

func Raw(arg string) *Tokens {
	return create().Raw(arg)
}

func (v *Tokens) Text(arg string) *Tokens {
	*v = append(*v, &models.Text{D: arg})
	return v
}

func Text(arg string) *Tokens {
	return create().Text(arg)
}
