/*
 *  Copyright (c) 2025-2026 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package golang

import (
	"go.osspkg.com/gogen/internal/models"
	"go.osspkg.com/gogen/types"
)

// Field appends a named struct field with the supplied type. Tags are alternating
// keys and values; rendering returns an error if their count is odd.
func (v *Tokens) Field(name string, fieldType types.Token, tags ...string) *Tokens {
	*v = append(*v, &models.StructField{Name: name, Type: fieldType, Tags: tags})
	return v
}

// Field starts a token sequence with a named struct field and optional tags.
// Tags are alternating keys and values; rendering returns an error if their count is odd.
func Field(name string, fieldType types.Token, tags ...string) *Tokens {
	return create().Field(name, fieldType, tags...)
}
