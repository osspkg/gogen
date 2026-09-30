/*
 *  Copyright (c) 2025-2026 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package python

import (
	"io"

	"go.osspkg.com/gogen/internal/gen"
	"go.osspkg.com/gogen/internal/models"
	"go.osspkg.com/gogen/types"
)

// Tokens is a chainable sequence of Python source tokens.
type Tokens []types.Token

func create() *Tokens { return &Tokens{} }

// Render writes arg without formatting or validating the complete Python program.
func Render(w io.Writer, arg types.Token) error { return arg.Render(w) }

// Render writes the token sequence to w without formatting it.
func (v *Tokens) Render(w io.Writer) error { return gen.Render(w, []types.Token(*v)) }

// Unwrap returns the underlying token slice without copying it.
func (v *Tokens) Unwrap() []types.Token { return *v }

// Join appends tokens to the sequence and returns it.
func (v *Tokens) Join(args ...types.Token) *Tokens {
	*v = append(*v, args...)
	return v
}

func wordStyle(value string) gen.Style {
	return gen.Style{Kind: gen.KindWord, Text: value, CanEndExpression: (languageConfig{}).CanEndExpression(value)}
}

func keyword(value string) types.Token {
	return &models.Keyword[languageConfig]{C: languageConfig{}, D: value}
}

func identifier(value string) types.Token {
	return &models.Keyword[languageConfig]{C: languageConfig{}, D: value, Verify: true}
}

func raw(value string) types.Token {
	return &models.Keyword[languageConfig]{C: languageConfig{}, D: value, Raw: true}
}

func text(value string) types.Token {
	return &models.Text[languageConfig]{C: languageConfig{}, D: value}
}

func operation(value string) types.Token { return &models.Operation[languageConfig]{D: value} }

func block(body ...types.Token) types.Token { return &suite{body: body} }

func brackets(args ...types.Token) types.Token { return &models.Bracket{D: args, Brace: true} }

func list(args ...types.Token) types.Token { return &models.Bracket{D: args} }

func square(args ...types.Token) types.Token { return &models.SquareBracket{D: args} }

func keyed(key, value types.Token) types.Token { return &models.KeyValue{Key: key, Value: value} }

func comment(value string) types.Token { return &pythonComment{text: value} }

func line() types.Token { return gen.Symbol("\n", gen.Style{Kind: gen.KindLine}) }

func writeString(w io.Writer, value string) error {
	n, err := io.WriteString(w, value)
	if err != nil {
		return err
	}
	if n != len(value) {
		return io.ErrShortWrite
	}
	return nil
}
