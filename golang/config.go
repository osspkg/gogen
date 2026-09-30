/*
 *  Copyright (c) 2025-2026 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package golang

import (
	"errors"
	"regexp"
	"strconv"
	"strings"

	cfg "go.osspkg.com/gogen/internal/config"
	"go.osspkg.com/gogen/internal/models"
	"go.osspkg.com/gogen/types"
)

var _ cfg.Config = config{}

type config struct{}

var errInvalidStructTag = errors.New("invalid struct tags: expected key/value pairs")

var identifierPattern = regexp.MustCompile(`(?i)^[a-z][0-9a-z_]*$`)

func (config) OperationAvailable(op string) bool {
	switch op {
	case "+", "-", "*", "/", "%", "&", "|", "^", "<<", ">>", "&^", "+=", "-=", "*=", "/=", "%=",
		"&=", "|=", "^=", "<<=", ">>=", "&^=", "&&", "||", "<-", "++", "--", "==", "<", ">", "=", "!", "~", "!=",
		"<=", ">=", ":=", "...", "(", ")", "[", "]", "{", "}", ",", ".", ";", ":":
		return true
	default:
		return false
	}
}

func (config) OperationKind(op string) cfg.TokenKind {
	switch op {
	case ".":
		return cfg.TokenDot
	case ",":
		return cfg.TokenComma
	case ":":
		return cfg.TokenColon
	case ";":
		return cfg.TokenSemicolon
	case "(":
		return cfg.TokenOpenParen
	case ")":
		return cfg.TokenCloseParen
	case "[":
		return cfg.TokenOpenSquare
	case "]":
		return cfg.TokenCloseSquare
	case "{":
		return cfg.TokenBlockOpen
	case "}":
		return cfg.TokenBlockClose
	case "++", "--", "...":
		return cfg.TokenPostfixOperator
	case "+", "-", "*", "&", "!", "^", "~", "<-":
		return cfg.TokenPrefixOperator
	default:
		return cfg.TokenDefault
	}
}

func (config) RawKind(text string, verified bool) cfg.TokenKind {
	if verified {
		return cfg.TokenDefault
	}
	switch {
	case text == "]":
		return cfg.TokenCloseSquare
	case text == "map[" || strings.HasPrefix(text, "["):
		return cfg.TokenTypePrefix
	default:
		return cfg.TokenDefault
	}
}

func (config) IsIdentifier(text string) bool {
	return identifierPattern.MatchString(text)
}

func (config) CanEndExpression(word string) bool {
	switch strings.TrimSpace(word) {
	case "break", "case", "const", "continue", "defer", "else", "fallthrough", "for", "func", "go", "goto", "if", "import", "package", "return", "select", "switch", "type", "var":
		return false
	default:
		return true
	}
}

func (config) QuoteString(value string) string {
	return strconv.Quote(value)
}

func (config) CommentSingle() cfg.OpenClose {
	return cfg.OpenClose{Open: "//", Close: "\n", SpaceAfterOpenWhenNeeded: true}
}

func (config) CommentMulti() cfg.OpenClose {
	return cfg.OpenClose{Open: "/*\n", Close: "\n*/\n"}
}

func (config) structTag(tags []string) (string, error) {
	if len(tags)%2 != 0 {
		return "", errInvalidStructTag
	}
	if len(tags) == 0 {
		return "", nil
	}

	var tag strings.Builder
	for i := 0; i < len(tags); i += 2 {
		if i > 0 {
			tag.WriteByte(' ')
		}
		tag.WriteString(tags[i])
		tag.WriteByte(':')
		tag.WriteString(strconv.Quote(tags[i+1]))
	}

	value := tag.String()
	if strings.ContainsAny(value, "`\r\n") {
		return strconv.Quote(value), nil
	}
	return "`" + value + "`", nil
}

func keyword(value string) *models.Keyword[config] {
	return &models.Keyword[config]{C: config{}, D: value}
}

func identifier(value string) *models.Keyword[config] {
	return &models.Keyword[config]{C: config{}, D: value, Verify: true}
}

func rawKeyword(value string) *models.Keyword[config] {
	return &models.Keyword[config]{C: config{}, D: value, Raw: true}
}

func rawToken(value string) *models.Raw[config] {
	return &models.Raw[config]{C: config{}, D: value}
}

func verifiedRawToken(value string) *models.Raw[config] {
	return &models.Raw[config]{C: config{}, D: value, Verify: true}
}

func rawTokenOf(token types.Token) *models.Raw[config] {
	return &models.Raw[config]{C: config{}, T: token}
}

func textToken(value string) *models.Text[config] {
	return &models.Text[config]{C: config{}, D: value}
}
