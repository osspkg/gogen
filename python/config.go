/*
 *  Copyright (c) 2025-2026 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package python

import (
	"fmt"
	"strings"
	"unicode"

	"go.osspkg.com/gogen/internal/config"
)

var _ config.Config = languageConfig{}

type languageConfig struct{}

var reservedWords = map[string]struct{}{
	"False": {}, "None": {}, "True": {}, "and": {}, "as": {}, "assert": {}, "async": {},
	"await": {}, "break": {}, "class": {}, "continue": {}, "def": {}, "del": {}, "elif": {},
	"else": {}, "except": {}, "finally": {}, "for": {}, "from": {}, "global": {}, "if": {},
	"import": {}, "in": {}, "is": {}, "lambda": {}, "nonlocal": {}, "not": {}, "or": {},
	"pass": {}, "raise": {}, "return": {}, "try": {}, "while": {}, "with": {}, "yield": {},
}

func (languageConfig) OperationAvailable(op string) bool {
	switch op {
	case "+", "-", "*", "**", "/", "//", "%", "@", "<<", ">>", "&", "|", "^", "~",
		"+=", "-=", "*=", "**=", "/=", "//=", "%=", "@=", "<<=", ">>=", "&=", "|=", "^=",
		"==", "!=", "<", ">", "<=", ">=", "=", ":=", "and", "or", "not", "in", "is", "is not", "not in",
		"->", ".", ",", ":", "(", ")", "[", "]", "{", "}":
		return true
	default:
		return false
	}
}

func (languageConfig) OperationKind(op string) config.TokenKind {
	switch op {
	case ".":
		return config.TokenDot
	case ",":
		return config.TokenComma
	case ":":
		return config.TokenColon
	case "(":
		return config.TokenOpenParen
	case ")":
		return config.TokenCloseParen
	case "[":
		return config.TokenOpenSquare
	case "]":
		return config.TokenCloseSquare
	case "{":
		return config.TokenBlockOpen
	case "}":
		return config.TokenBlockClose
	case "+", "-", "*", "**", "~":
		return config.TokenPrefixOperator
	default:
		return config.TokenDefault
	}
}

func (languageConfig) RawKind(text string, verified bool) config.TokenKind {
	if verified {
		return config.TokenDefault
	}
	switch text {
	case "]":
		return config.TokenCloseSquare
	case "}":
		return config.TokenBlockClose
	case ".":
		return config.TokenDot
	default:
		return config.TokenDefault
	}
}

func (languageConfig) IsIdentifier(text string) bool {
	if text == "" {
		return false
	}
	if _, reserved := reservedWords[text]; reserved {
		return false
	}
	for index, r := range text {
		if index == 0 {
			if !isIdentifierStart(r) {
				return false
			}
			continue
		}
		if !isIdentifierContinue(r) {
			return false
		}
	}
	return true
}

func isIdentifierStart(r rune) bool {
	return r == '_' || unicode.IsLetter(r) || unicode.In(r, unicode.Nl)
}

func isIdentifierContinue(r rune) bool {
	return isIdentifierStart(r) || unicode.IsDigit(r) || unicode.In(r, unicode.Mn, unicode.Mc, unicode.Pc) || r == '\u200c' || r == '\u200d'
}

func (languageConfig) CanEndExpression(word string) bool {
	switch strings.TrimSpace(word) {
	case "and", "as", "assert", "async", "await", "break", "class", "continue", "def", "del", "elif", "else", "except", "finally", "for", "from", "global", "if", "import", "in", "is", "lambda", "nonlocal", "not", "or", "pass", "raise", "return", "try", "while", "with", "yield":
		return false
	default:
		return true
	}
}

func (languageConfig) QuoteString(value string) string {
	var out strings.Builder
	out.Grow(len(value) + 2)
	out.WriteByte('"')
	for _, r := range value {
		switch r {
		case '\\', '"':
			out.WriteByte('\\')
			out.WriteRune(r)
		case '\n':
			out.WriteString("\\n")
		case '\r':
			out.WriteString("\\r")
		case '\t':
			out.WriteString("\\t")
		case '\b':
			out.WriteString("\\x08")
		case '\f':
			out.WriteString("\\x0c")
		default:
			if r < 0x20 || r == 0x7f {
				out.WriteString(fmt.Sprintf("\\x%02x", r))
			} else {
				out.WriteRune(r)
			}
		}
	}
	out.WriteByte('"')
	return out.String()
}

func (languageConfig) CommentSingle() config.OpenClose {
	return config.OpenClose{Open: "#", Close: "\n", SpaceAfterOpenWhenNeeded: true}
}

func (languageConfig) CommentMulti() config.OpenClose {
	return config.OpenClose{Open: "#", Close: "\n", SpaceAfterOpenWhenNeeded: true}
}
