/*
 *  Copyright (c) 2025-2026 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package gen

import (
	"bytes"
	"io"
)

type IndentWriter struct {
	w          io.Writer
	indent     int
	indentText string
	lineStart  bool
}

func Indented(w io.Writer) *IndentWriter {
	return IndentedWith(w, "\t")
}

// IndentedWith returns a writer that prefixes each indented line with indentText.
// If w is already an IndentWriter, it reuses that writer and its indentation.
func IndentedWith(w io.Writer, indentText string) *IndentWriter {
	if indented, ok := w.(*IndentWriter); ok {
		return indented
	}
	return &IndentWriter{w: w, indentText: indentText}
}

func (w *IndentWriter) Push() {
	w.indent++
}

func (w *IndentWriter) Pop() {
	if w.indent > 0 {
		w.indent--
	}
}

func (w *IndentWriter) AtLineStart() bool {
	return w.lineStart
}

func (w *IndentWriter) EnsureNewline() error {
	if w.lineStart {
		return nil
	}
	_, err := w.Write([]byte{'\n'})
	return err
}

func (w *IndentWriter) Write(p []byte) (int, error) {
	consumed := 0
	for consumed < len(p) {
		if w.lineStart && p[consumed] != '\n' {
			for i := 0; i < w.indent; i++ {
				n, err := io.WriteString(w.w, w.indentText)
				if err != nil {
					return consumed, err
				}
				if n != len(w.indentText) {
					return consumed, io.ErrShortWrite
				}
			}
			w.lineStart = false
		}

		start := consumed
		nextLine := bytes.IndexByte(p[start:], '\n')
		end := len(p)
		if nextLine >= 0 {
			end = start + nextLine + 1
		}
		written, err := w.w.Write(p[start:end])
		consumed += written
		if written > 0 && p[consumed-1] == '\n' {
			w.lineStart = true
		}
		if err != nil {
			return consumed, err
		}
		if written < end-start {
			return consumed, io.ErrShortWrite
		}
	}
	return consumed, nil
}

// WriteVerbatim writes token text without inserting indentation after its own
// newlines. It still indents the token's first line when it begins a block line.
func WriteVerbatim(w io.Writer, text string) error {
	if indented, ok := w.(*IndentWriter); ok {
		if indented.lineStart && text != "" && text[0] != '\n' {
			for i := 0; i < indented.indent; i++ {
				n, err := io.WriteString(indented.w, indented.indentText)
				if err != nil {
					return err
				}
				if n != len(indented.indentText) {
					return io.ErrShortWrite
				}
			}
			indented.lineStart = false
		}
		n, err := io.WriteString(indented.w, text)
		if n > 0 {
			indented.lineStart = text[n-1] == '\n'
		}
		if err != nil {
			return err
		}
		if n != len(text) {
			return io.ErrShortWrite
		}
		return nil
	}
	_, err := io.WriteString(w, text)
	return err
}
