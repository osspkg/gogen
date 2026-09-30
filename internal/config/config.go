/*
 *  Copyright (c) 2025-2026 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package config

type TokenKind uint8

const (
	TokenDefault TokenKind = iota
	TokenDot
	TokenComma
	TokenColon
	TokenSemicolon
	TokenOpenParen
	TokenCloseParen
	TokenOpenSquare
	TokenCloseSquare
	TokenBlockOpen
	TokenBlockClose
	TokenTypePrefix
	TokenPrefixOperator
	TokenPostfixOperator
)

type Config interface {
	CommentSingle() OpenClose
	CommentMulti() OpenClose
	OperationAvailable(op string) bool
	OperationKind(op string) TokenKind
	RawKind(text string, verified bool) TokenKind
	IsIdentifier(text string) bool
	CanEndExpression(word string) bool
	QuoteString(text string) string
}

type OpenClose struct {
	Open                     string
	Close                    string
	SpaceAfterOpenWhenNeeded bool
}
