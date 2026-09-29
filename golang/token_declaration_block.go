/*
 *  Copyright (c) 2025-2026 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package golang

import (
	"go.osspkg.com/gogen/internal/models"
	"go.osspkg.com/gogen/types"
)

// ImportBlock appends a grouped import declaration. Each spec is an import
// path, optionally preceded by an alias.
func (v *Tokens) ImportBlock(specs ...types.Token) *Tokens {
	*v = append(*v, &models.DeclarationBlock{Keyword: "import", D: specs})
	return v.Line()
}

// ImportBlock starts a token sequence with a grouped import declaration. Each
// spec is an import path, optionally preceded by an alias.
func ImportBlock(specs ...types.Token) *Tokens {
	return create().ImportBlock(specs...)
}

// TypeBlock appends a grouped type declaration. Each spec contains a type
// name followed by its underlying type.
func (v *Tokens) TypeBlock(specs ...types.Token) *Tokens {
	*v = append(*v, &models.DeclarationBlock{Keyword: "type", D: specs})
	return v.Line()
}

// TypeBlock starts a token sequence with a grouped type declaration. Each spec
// contains a type name followed by its underlying type.
func TypeBlock(specs ...types.Token) *Tokens {
	return create().TypeBlock(specs...)
}
