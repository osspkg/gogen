/*
 *  Copyright (c) 2025-2026 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package python

import (
	"errors"
	"io"

	"go.osspkg.com/gogen/internal/gen"
	"go.osspkg.com/gogen/types"
)

type suite struct{ body []types.Token }

func (v *suite) Render(w io.Writer) error {
	iw := gen.IndentedWith(w, "    ")
	if err := writeString(iw, ":"); err != nil {
		return err
	}
	if len(v.body) == 0 {
		if err := writeString(iw, " pass\n"); err != nil {
			return err
		}
		return nil
	}
	if err := iw.EnsureNewline(); err != nil {
		return err
	}
	iw.Push()
	for _, token := range v.body {
		if token == nil {
			iw.Pop()
			return errors.New("python suite contains a nil statement")
		}
		if err := iw.EnsureNewline(); err != nil {
			iw.Pop()
			return err
		}
		if err := gen.Render(iw, token); err != nil {
			iw.Pop()
			return err
		}
	}
	if err := iw.EnsureNewline(); err != nil {
		iw.Pop()
		return err
	}
	iw.Pop()
	return nil
}

func (v *suite) RenderLayout() gen.Layout {
	return gen.Layout{
		First: gen.Style{Kind: gen.KindColon, Text: ":"},
		Last:  gen.Style{Kind: gen.KindLine, Text: "\n"},
	}
}

type pythonComment struct{ text string }

func (v *pythonComment) Render(w io.Writer) error {
	lines := splitLines(v.text)
	for _, text := range lines {
		if err := writeString(w, "#"); err != nil {
			return err
		}
		if text != "" && text[0] != ' ' && text[0] != '\t' {
			if err := writeString(w, " "); err != nil {
				return err
			}
		}
		if err := writeString(w, text+"\n"); err != nil {
			return err
		}
	}
	return nil
}

func (v *pythonComment) RenderLayout() gen.Layout {
	style := gen.Style{Kind: gen.KindComment, Text: v.text}
	return gen.Layout{First: style, Last: gen.Style{Kind: gen.KindLine, Text: "\n"}}
}

func splitLines(value string) []string {
	if value == "" {
		return []string{""}
	}
	var lines []string
	for len(value) > 0 {
		index := -1
		for i, r := range value {
			if r == '\n' {
				index = i
				break
			}
		}
		if index < 0 {
			lines = append(lines, value)
			break
		}
		line := value[:index]
		if len(line) > 0 && line[len(line)-1] == '\r' {
			line = line[:len(line)-1]
		}
		lines = append(lines, line)
		value = value[index+1:]
	}
	return lines
}

// Block creates a Python suite that appends a colon and indents its body.
func (v *Tokens) Block(body ...types.Token) *Tokens { return v.Join(block(body...)) }

// Block creates a Python suite that appends a colon and indents its body.
func Block(body ...types.Token) *Tokens { return create().Block(body...) }

// Comment creates a Python comment, prefixing every line with #.
func (v *Tokens) Comment(text string) *Tokens { return v.Join(comment(text)) }

// Comment creates a Python comment, prefixing every line with #.
func Comment(text string) *Tokens { return create().Comment(text) }

// Line appends an explicit line break.
func (v *Tokens) Line() *Tokens { return v.Join(line()) }

// Line creates a token sequence containing an explicit line break.
func Line() *Tokens { return create().Line() }
