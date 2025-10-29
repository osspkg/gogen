/*
 *  Copyright (c) 2025 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package models

import (
	"io"
	"strconv"

	"go.osspkg.com/gogen/internal/gen"
)

type Text struct {
	D string
}

func (v *Text) Render(w io.Writer) error {
	return gen.Render(w, strconv.Quote(v.D))
}
