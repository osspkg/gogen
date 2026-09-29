/*
 *  Copyright (c) 2025-2026 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package models

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"go.osspkg.com/gogen/internal/gen"
	"go.osspkg.com/gogen/types"
)

type StructField struct {
	Name string
	Type types.Token
	Tags []string
}

func (v *StructField) Render(w io.Writer) error {
	if len(v.Tags)%2 != 0 {
		return fmt.Errorf("invalid struct tags: expected key/value pairs")
	}
	return gen.Render(w, v.tokens())
}

func (v *StructField) RenderLayout() gen.Layout {
	return gen.LayoutOf(v.tokens())
}

func (v *StructField) tokens() []types.Token {
	out := []types.Token{&Keyword{D: v.Name, Verify: true}}
	out = append(out, gen.Params(v.Type)...)
	if len(v.Tags) == 0 || len(v.Tags)%2 != 0 {
		return out
	}

	var tag strings.Builder
	for i := 0; i < len(v.Tags); i += 2 {
		if i > 0 {
			tag.WriteByte(' ')
		}
		tag.WriteString(v.Tags[i])
		tag.WriteByte(':')
		tag.WriteString(strconv.Quote(v.Tags[i+1]))
	}

	tagValue := tag.String()
	if strings.ContainsAny(tagValue, "`\r\n") {
		tagValue = strconv.Quote(tagValue)
	} else {
		tagValue = "`" + tagValue + "`"
	}
	out = append(out, &Keyword{D: tagValue, Raw: true})
	return out
}
