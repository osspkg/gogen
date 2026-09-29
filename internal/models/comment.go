/*
 *  Copyright (c) 2025-2026 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package models

import (
	"io"
	"strings"

	"go.osspkg.com/gogen/internal/config"
	"go.osspkg.com/gogen/internal/gen"
)

type Comment[C config.Config] struct {
	c C
	D string
}

func (v *Comment[C]) Render(w io.Writer) error {
	oc := v.c.CommentSingle()
	if strings.Contains(v.D, "\n") {
		oc = v.c.CommentMulti()
	}

	if oc.Open == "//" && len(v.D) > 0 && v.D[0] != ' ' && v.D[0] != '\t' {
		return gen.Render(w, oc.Open, " ", v.D, oc.Close)
	}
	return gen.Render(w, oc.Open, v.D, oc.Close)
}

func (v *Comment[C]) RenderLayout() gen.Layout {
	style := gen.Style{Kind: gen.KindComment, Text: v.D}
	return gen.Layout{First: style, Last: style}
}
