/*
 *  Copyright (c) 2025 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package models

import (
	"fmt"
	"io"

	"go.osspkg.com/gogen/internal/config"
	"go.osspkg.com/gogen/internal/gen"
)

type Operation[C config.Config] struct {
	c C
	D string
}

func (v *Operation[C]) Render(w io.Writer) error {
	if !v.c.OperationAvailable(v.D) {
		return fmt.Errorf("invalid operation: %s", v.D)
	}
	return gen.Render(w, v.D)
}
