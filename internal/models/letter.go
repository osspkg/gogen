/*
 *  Copyright (c) 2025-2026 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package models

import (
	"fmt"
	"io"

	"go.osspkg.com/gogen/internal/gen"
	"go.osspkg.com/gogen/types"
)

type Letter struct {
	D string
}

func (v *Letter) Render(w io.Writer) error {
	return gen.Render(w, v.D)
}

type Raw struct {
	D      string
	T      types.Token
	AT     []types.Token
	Verify bool
}

func (v *Raw) Render(w io.Writer) error {
	if v.Verify && !rexLetter.MatchString(v.D) {
		return fmt.Errorf("invalid letter: %s", v.D)
	}
	if err := gen.Render(w, v.D); err != nil {
		return err
	}
	if v.T != nil {
		if err := v.T.Render(w); err != nil {
			return err
		}
	}
	if len(v.AT) > 0 {
		for _, token := range v.AT {
			if err := token.Render(w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (v *Raw) NoSpace() {}
