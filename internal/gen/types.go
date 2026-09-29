/*
 *  Copyright (c) 2025-2026 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package gen

import "go.osspkg.com/gogen/types"

type Unwrap interface {
	Unwrap() []types.Token
}

type Kind uint8

const (
	KindUnknown Kind = iota
	KindWord
	KindLiteral
	KindFragment
	KindTypePrefix
	KindCloseBracket
	KindOperator
	KindPrefixOperator
	KindPostfixOperator
	KindDot
	KindComma
	KindColon
	KindSemicolon
	KindOpenParen
	KindCloseParen
	KindOpenSquare
	KindBlockOpen
	KindBlockClose
	KindLine
	KindComment
)

type Style struct {
	Kind Kind
	Text string
}

type Layout struct {
	First Style
	Last  Style
}

type Styled interface {
	RenderLayout() Layout
}
