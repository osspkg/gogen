/*
 *  Copyright (c) 2025-2026 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package models

import (
	"go.osspkg.com/gogen/internal/config"
	"go.osspkg.com/gogen/internal/gen"
)

func layoutKind(kind config.TokenKind) gen.Kind {
	switch kind {
	case config.TokenDot:
		return gen.KindDot
	case config.TokenComma:
		return gen.KindComma
	case config.TokenColon:
		return gen.KindColon
	case config.TokenSemicolon:
		return gen.KindSemicolon
	case config.TokenOpenParen:
		return gen.KindOpenParen
	case config.TokenCloseParen:
		return gen.KindCloseParen
	case config.TokenOpenSquare:
		return gen.KindOpenSquare
	case config.TokenCloseSquare:
		return gen.KindCloseBracket
	case config.TokenBlockOpen:
		return gen.KindBlockOpen
	case config.TokenBlockClose:
		return gen.KindBlockClose
	case config.TokenTypePrefix:
		return gen.KindTypePrefix
	case config.TokenPrefixOperator:
		return gen.KindUnaryOperator
	case config.TokenPostfixOperator:
		return gen.KindPostfixOperator
	default:
		return gen.KindOperator
	}
}
