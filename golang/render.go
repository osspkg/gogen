/*
 *  Copyright (c) 2025-2026 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package golang

import (
	"bufio"
	"bytes"
	"fmt"
	"go/format"
	"io"
	"regexp"
	"strconv"
	"sync/atomic"

	"go.osspkg.com/gogen/types"
)

var rawMode = atomic.Bool{}

func SetRawMode() {
	rawMode.Store(true)
}

func SetDefaultMode() {
	rawMode.Store(false)
}

func Render(w io.Writer, arg types.Token) error {
	buf := bytes.NewBuffer(nil)

	if err := arg.Render(buf); err != nil {
		return err
	}

	var (
		b   []byte
		err error
	)

	if rawMode.Load() {
		b = buf.Bytes()
	} else {
		if b, err = format.Source(buf.Bytes()); err != nil {
			return fmt.Errorf(
				"gofmt error: %w%s",
				err, getBadCode(buf.Bytes(), err),
			)
		}
	}

	if _, err = w.Write(b); err != nil {
		return err
	}

	return nil
}

var rex = regexp.MustCompile(`(?iU)^(\d+)\:(\d+)\:`)

func getBadCode(b []byte, err error) string {
	if err == nil {
		return ""
	}

	found := rex.FindStringSubmatch(err.Error())
	if len(found) != 3 {
		return ""
	}

	line, err := strconv.ParseInt(found[1], 10, 64)
	if err != nil {
		return ""
	}

	scanner := bufio.NewScanner(bytes.NewReader(b))
	var (
		currentLine int64
		code        string
	)

	for scanner.Scan() {
		if currentLine == line-1 {
			code += scanner.Text()
			break
		}
		currentLine++
	}

	return ", code:`" + code + "`"
}
