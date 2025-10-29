/*
 *  Copyright (c) 2025 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package gen

import (
	"fmt"
	"io"

	"go.osspkg.com/gogen/types"
)

func Render(w io.Writer, args ...any) error {
	for _, arg := range args {
		if err := drawAny(w, arg); err != nil {
			return err
		}
	}
	return nil
}

func drawAny(w io.Writer, t any) error {
	if v, ok := t.(Unwrap); ok {
		return drawList(w, v.Unwrap(), true)
	} else if v, ok := t.([]types.Token); ok {
		return drawList(w, v, false)
	} else if v, ok := t.(types.Token); ok {
		return v.Render(w)
	} else if v, ok := t.(string); ok {
		_, err := io.WriteString(w, v)
		return err
	} else {
		fmt.Printf("[X] %T\n", t)
	}
	return nil
}

func drawList(w io.Writer, list []types.Token, line bool) error {
	count := len(list) - 1
	for i, token := range list {
		if vt, ok := token.(Unwrap); ok {
			if err := drawList(w, vt.Unwrap(), true); err != nil {
				return err
			}
			continue
		}

		if err := token.Render(w); err != nil {
			return err
		}
		if i >= count {
			continue
		}
		if _, err := io.WriteString(w, " "); err != nil {
			return err
		}
	}
	if line {
		if _, err := io.WriteString(w, "\n"); err != nil {
			return err
		}
	}
	return nil
}

func Params(token types.Token) (out []types.Token) {
	if unw, ok := token.(Unwrap); ok {
		out = append(out, unw.Unwrap()...)
	} else {
		out = append(out, token)
	}
	return
}
