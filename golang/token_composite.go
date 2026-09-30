/*
 *  Copyright (c) 2025-2026 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package golang

import (
	"go.osspkg.com/gogen/internal/models"
	"go.osspkg.com/gogen/types"
)

// Index appends an index expression such as [i].
func (v *Tokens) Index(index types.Token) *Tokens {
	*v = append(*v, &models.SquareBracket{D: []types.Token{index}})
	return v
}

// Index creates an index expression such as [i].
func Index(index types.Token) *Tokens {
	return create().Index(index)
}

// TypeArgs appends type arguments such as [T, U].
func (v *Tokens) TypeArgs(args ...types.Token) *Tokens {
	*v = append(*v, &models.SquareBracket{D: args})
	return v
}

// TypeArgs creates a list of type arguments such as [T, U].
func TypeArgs(args ...types.Token) *Tokens {
	return create().TypeArgs(args...)
}

// KeyValue appends a keyed element such as key: value.
func (v *Tokens) KeyValue(key, value types.Token) *Tokens {
	*v = append(*v, &models.KeyValue{Key: key, Value: value})
	return v
}

// KeyValue creates a keyed element such as key: value.
func KeyValue(key, value types.Token) *Tokens {
	return create().KeyValue(key, value)
}
