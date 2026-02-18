/*
 *  Copyright (c) 2025-2026 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package models

import (
	"fmt"
	"io"
	"regexp"

	"go.osspkg.com/gogen/internal/gen"
)

type Keyword struct {
	Verify bool
	D      string
}

var rexLetter = regexp.MustCompile(`(?mUi)^[a-z][0-9a-z\_]{0,}$`)

func (v *Keyword) Render(w io.Writer) error {
	if v.Verify && !rexLetter.MatchString(v.D) {
		return fmt.Errorf("invalid letter: %s", v.D)
	}
	return gen.Render(w, v.D)
}
