/*
 *  Copyright (c) 2025-2026 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package golang

import (
	"io"

	"go.osspkg.com/gogen/internal/gen"
	"go.osspkg.com/gogen/types"
)

type structField struct {
	name string
	typ  types.Token
	tags []string
}

func (v *structField) Render(w io.Writer) error {
	parts, err := v.tokens()
	if err != nil {
		return err
	}
	return gen.Render(w, parts)
}

func (v *structField) RenderLayout() gen.Layout {
	parts, err := v.tokens()
	if err != nil {
		return gen.Layout{}
	}
	return gen.LayoutOf(parts)
}

func (v *structField) tokens() ([]types.Token, error) {
	tag, err := (config{}).structTag(v.tags)
	if err != nil {
		return nil, err
	}

	out := []types.Token{identifier(v.name)}
	out = append(out, gen.Params(v.typ)...)
	if tag != "" {
		out = append(out, rawKeyword(tag))
	}
	return out, nil
}

// Field appends a named struct field with the supplied type. Tags are alternating
// keys and values; rendering returns an error if their count is odd.
func (v *Tokens) Field(name string, fieldType types.Token, tags ...string) *Tokens {
	*v = append(*v, &structField{name: name, typ: fieldType, tags: tags})
	return v
}

// Field starts a token sequence with a named struct field and optional tags.
// Tags are alternating keys and values; rendering returns an error if their count is odd.
func Field(name string, fieldType types.Token, tags ...string) *Tokens {
	return create().Field(name, fieldType, tags...)
}
