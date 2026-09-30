/*
 *  Copyright (c) 2025-2026 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package typescript

import (
	"io"

	"go.osspkg.com/gogen/internal/gen"
	"go.osspkg.com/gogen/internal/models"
	"go.osspkg.com/gogen/types"
)

// Tokens is a chainable sequence of TypeScript or TSX tokens.
type Tokens []types.Token

func create() *Tokens { return &Tokens{} }

// Render writes arg to w without running a formatter or compiler.
func Render(w io.Writer, arg types.Token) error {
	return arg.Render(w)
}

// Render writes the token sequence to w without formatting it.
func (v *Tokens) Render(w io.Writer) error {
	return gen.Render(w, []types.Token(*v))
}

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

func rawWord(value string) *Tokens {
	return create().Join(&tokenRaw{value: value})
}

type tokenRaw struct{ value string }

func (v *tokenRaw) Render(w io.Writer) error {
	_, err := io.WriteString(w, v.value)
	return err
}

func (v *tokenRaw) RenderLayout() gen.Layout {
	style := gen.Style{Kind: gen.KindFragment, Text: v.value, CanEndExpression: (languageConfig{}).CanEndExpression(v.value)}
	return gen.Layout{First: style, Last: style}
}

func symbol(value string) types.Token {
	kind := models.LayoutKind((languageConfig{}).OperationKind(value))
	return gen.Symbol(value, gen.Style{Kind: kind, Text: value})
}

func postfix(value string) types.Token {
	return gen.Symbol(value, gen.Style{Kind: gen.KindPostfixOperator, Text: value})
}

func keyword(value string) types.Token {
	return &models.Keyword[languageConfig]{C: languageConfig{}, D: value}
}

func raw(value string) types.Token {
	return &models.Keyword[languageConfig]{C: languageConfig{}, D: value, Raw: true}
}

func identifier(value string) types.Token {
	return &models.Keyword[languageConfig]{C: languageConfig{}, D: value, Verify: true}
}

func text(value string) types.Token {
	return &models.Text[languageConfig]{C: languageConfig{}, D: value}
}

func operation(value string) types.Token {
	return &models.Operation[languageConfig]{D: value}
}

func comment(value string) types.Token {
	return &models.Comment[languageConfig]{D: value}
}

func block(args ...types.Token) types.Token {
	return &models.Block{D: args}
}

func brackets(args ...types.Token) types.Token {
	return &models.Bracket{D: args, Brace: true}
}

func list(args ...types.Token) types.Token {
	return &models.Bracket{D: args}
}

func square(args ...types.Token) types.Token {
	return &models.SquareBracket{D: args}
}

func keyed(key, value types.Token) types.Token {
	return &models.KeyValue{Key: key, Value: value}
}
