/*
 *  Copyright (c) 2025-2026 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package gen

import (
	"io"

	"go.osspkg.com/gogen/types"
)

type symbol struct {
	text  string
	style Style
}

func Symbol(text string, style Style) types.Token {
	return symbol{text: text, style: style}
}

func (v symbol) Render(w io.Writer) error {
	_, err := io.WriteString(w, v.text)
	return err
}

func (v symbol) RenderLayout() Layout {
	return Layout{First: v.style, Last: v.style}
}
