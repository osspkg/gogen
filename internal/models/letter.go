/*
 *  Copyright (c) 2025 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package models

import (
	"io"

	"go.osspkg.com/gogen/internal/gen"
)

type Letter struct {
	D string
}

func (v *Letter) Render(w io.Writer) error {
	return gen.Render(w, v.D)
}
