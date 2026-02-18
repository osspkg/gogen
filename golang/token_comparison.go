/*
 *  Copyright (c) 2025-2026 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package golang

import "go.osspkg.com/gogen/internal/models"

func (v *Tokens) If() *Tokens {
	*v = append(*v, &models.Keyword{D: "if"})
	return v
}

func If() *Tokens {
	return create().If()
}

func (v *Tokens) Else() *Tokens {
	*v = append(*v, &models.Keyword{D: "else"})
	return v
}

func (v *Tokens) ElseIf() *Tokens {
	*v = append(*v, &models.Keyword{D: "else if"})
	return v
}

func (v *Tokens) Default() *Tokens {
	*v = append(*v, &models.Keyword{D: "default"})
	return v
}

func Default() *Tokens {
	return create().Default()
}

func (v *Tokens) Fallthrough() *Tokens {
	*v = append(*v, &models.Keyword{D: "fallthrough"})
	return v
}

func (v *Tokens) Break() *Tokens {
	*v = append(*v, &models.Keyword{D: "break"})
	return v
}

func (v *Tokens) Case() *Tokens {
	*v = append(*v, &models.Keyword{D: "case"})
	return v
}

func Case() *Tokens {
	return create().Case()
}

func (v *Tokens) Continue() *Tokens {
	*v = append(*v, &models.Keyword{D: "continue"})
	return v
}

func (v *Tokens) Goto() *Tokens {
	*v = append(*v, &models.Keyword{D: "goto"})
	return v
}

func (v *Tokens) For() *Tokens {
	*v = append(*v, &models.Keyword{D: "for"})
	return v
}

func For() *Tokens {
	return create().For()
}

func (v *Tokens) Range() *Tokens {
	*v = append(*v, &models.Keyword{D: "range"})
	return v
}

func Select() *Tokens {
	return create().Join(&models.Keyword{D: "select"})
}

func Switch() *Tokens {
	return create().Join(&models.Keyword{D: "switch"})
}
