/*
 *  Copyright (c) 2025 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package golang

import (
	"bytes"
	"go/format"
	"io"

	"go.osspkg.com/gogen/types"
)

func Render(w io.Writer, arg types.Token) error {
	buf := bytes.NewBuffer(nil)
	if err := arg.Render(buf); err != nil {
		return err
	}
	b, err := format.Source(buf.Bytes())
	if err != nil {
		return err
	}
	if _, err = w.Write(b); err != nil {
		return err
	}
	return nil
}
