/*
 *  Copyright (c) 2025-2026 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package golang

import (
	"go.osspkg.com/gogen/internal/models"
)

// Op appends a Go operator token to the token sequence. Render returns an error if the operator is not supported.
func (v *Tokens) Op(arg string) *Tokens {
	*v = append(*v, &models.Operation[config]{D: arg})
	return v
}

// Op creates a token sequence containing a Go operator token. Render returns an error if the operator is not supported.
func Op(arg string) *Tokens {
	return create().Op(arg)
}

// Raw appends the supplied source text without quoting or formatting to the token sequence. Use it when the text must be inserted verbatim.
func (v *Tokens) Raw(arg string) *Tokens {
	*v = append(*v, &models.Keyword{D: arg, Raw: true})
	return v
}

// Raw creates a token sequence containing the supplied source text without quoting or formatting. Use it when the text must be inserted verbatim.
func Raw(arg string) *Tokens {
	return create().Raw(arg)
}

// Text appends a quoted Go string literal to the token sequence. Use it for string values that must be escaped as Go source.
func (v *Tokens) Text(arg string) *Tokens {
	*v = append(*v, &models.Text{D: arg})
	return v
}

// Text creates a token sequence containing a quoted Go string literal. Use it for string values that must be escaped as Go source.
func Text(arg string) *Tokens {
	return create().Text(arg)
}
