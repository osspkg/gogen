/*
 *  Copyright (c) 2025-2026 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package python

import "go.osspkg.com/gogen/internal/gen"

// NoneValue appends Python's None literal.
func (v *Tokens) NoneValue() *Tokens { return v.Join(keyword("None")) }

// NoneValue creates Python's None literal.
func NoneValue() *Tokens { return create().NoneValue() }

// True appends Python's True literal.
func (v *Tokens) True() *Tokens { return v.Join(keyword("True")) }

// True creates Python's True literal.
func True() *Tokens { return create().True() }

// False appends Python's False literal.
func (v *Tokens) False() *Tokens { return v.Join(keyword("False")) }

// False creates Python's False literal.
func False() *Tokens { return create().False() }

// Ellipsis appends Python's ellipsis literal.
func (v *Tokens) Ellipsis() *Tokens {
	return v.Join(gen.Symbol("...", gen.Style{Kind: gen.KindLiteral, Text: "..."}))
}

// Ellipsis creates Python's ellipsis literal.
func Ellipsis() *Tokens { return create().Ellipsis() }
