/*
 *  Copyright (c) 2025-2026 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package typescript

import (
	"fmt"
	"strings"
	"unicode"

	"go.osspkg.com/gogen/internal/config"
)

var _ config.Config = languageConfig{}

type languageConfig struct{}

func (languageConfig) OperationAvailable(op string) bool {
	switch op {
	case "+", "-", "*", "/", "%", "**", "++", "--", "==", "===", "!=", "!==", "<", ">", "<=", ">=", "=", "+=", "-=", "*=", "/=", "%=", "**=", "&&", "||", "??", "&&=", "||=", "??=", "!", "~", "&", "|", "^", "<<", ">>", ">>>", "&=", "|=", "^=", "<<=", ">>=", ">>>=", "=>", "?", ":", ";", ",", ".", "?.", "...", "(", ")", "[", "]", "{", "}":
		return true
	default:
		return false
	}
}

func (languageConfig) OperationKind(op string) config.TokenKind {
	switch op {
	case ".", "?.":
		return config.TokenDot
	case ",":
		return config.TokenComma
	case ":":
		return config.TokenColon
	case ";":
		return config.TokenSemicolon
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
	case "++", "--":
		return config.TokenPostfixOperator
	case "!", "~", "+", "-", "*", "&", "...":
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
	case "?.":
		return config.TokenDot
	default:
		return config.TokenDefault
	}
}

func (languageConfig) IsIdentifier(text string) bool {
	if text == "" {
		return false
	}
	for index, r := range text {
		if isIdentifierStart(r) {
			continue
		}
		if index == 0 || !isIdentifierContinue(r) {
			return false
		}
	}
	return true
}

func isIdentifierStart(r rune) bool {
	return r == '$' || r == '_' || unicode.IsLetter(r) || unicode.In(r, unicode.Nl)
}

func isIdentifierContinue(r rune) bool {
	return isIdentifierStart(r) || unicode.IsDigit(r) || unicode.In(r, unicode.Mn, unicode.Mc, unicode.Pc) || r == '\u200c' || r == '\u200d'
}

func (languageConfig) CanEndExpression(word string) bool {
	switch strings.TrimSpace(word) {
	case "as", "asserts", "abstract", "break", "case", "catch", "class", "const", "continue", "debugger", "default", "delete", "do", "else", "enum", "export", "extends", "finally", "for", "function", "if", "implements", "import", "in", "infer", "instanceof", "interface", "keyof", "let", "namespace", "new", "private", "protected", "public", "readonly", "return", "override", "constructor", "satisfies", "static", "switch", "throw", "try", "type", "typeof", "var", "void", "while", "with":
		return false
	default:
		return true
	}
}

func (languageConfig) QuoteString(value string) string {
	var quoted strings.Builder
	quoted.Grow(len(value) + 2)
	quoted.WriteByte('"')
	for _, r := range value {
		switch r {
		case '\\', '"':
			quoted.WriteByte('\\')
			quoted.WriteRune(r)
		case '\b':
			quoted.WriteString("\\b")
		case '\f':
			quoted.WriteString("\\f")
		case '\n':
			quoted.WriteString("\\n")
		case '\r':
			quoted.WriteString("\\r")
		case '\t':
			quoted.WriteString("\\t")
		case '\u2028':
			quoted.WriteString("\\u2028")
		case '\u2029':
			quoted.WriteString("\\u2029")
		default:
			if r < 0x20 || r == 0x7f {
				quoted.WriteString(fmt.Sprintf("\\x%02x", r))
			} else {
				quoted.WriteRune(r)
			}
		}
	}
	quoted.WriteByte('"')
	return quoted.String()
}

func (languageConfig) CommentSingle() config.OpenClose {
	return config.OpenClose{Open: "//", Close: "\n", SpaceAfterOpenWhenNeeded: true}
}

func (languageConfig) CommentMulti() config.OpenClose {
	return config.OpenClose{Open: "/*\n", Close: "\n*/\n"}
}
