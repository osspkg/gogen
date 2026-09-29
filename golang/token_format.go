/*
 *  Copyright (c) 2025-2026 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package golang

import (
	"go.osspkg.com/gogen/internal/models"
	"go.osspkg.com/gogen/types"
)

// Comment appends a Go line comment to the token sequence.
func (v *Tokens) Comment(arg string) *Tokens {
	*v = append(*v, &models.Comment[config]{D: arg})
	return v
}

// Comment creates a token sequence containing a Go line comment.
func Comment(arg string) *Tokens {
	return create().Comment(arg)
}

//------------------------------------------------------

// Line appends a line break to the token sequence.
func (v *Tokens) Line() *Tokens {
	*v = append(*v, &models.Letter{D: "\n"})
	return v
}

// Line creates a token sequence containing a line break.
func Line() *Tokens {
	return create().Line()
}

//------------------------------------------------------

// Block appends a brace-delimited block with indented lines to the token sequence. Nested tokens are rendered on separate indented lines.
func (v *Tokens) Block(args ...types.Token) *Tokens {
	*v = append(*v, &models.Block{D: args})
	return v
}

// Block creates a token sequence containing a brace-delimited block with indented lines. Nested tokens are rendered on separate indented lines.
func Block(args ...types.Token) *Tokens {
	return create().Block(args...)
}
