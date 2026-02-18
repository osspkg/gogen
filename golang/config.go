/*
 *  Copyright (c) 2025-2026 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package golang

import (
	cfg "go.osspkg.com/gogen/internal/config"
)

var _ cfg.Config = (*config)(nil)

type config struct{}

func (c config) OperationAvailable(op string) bool {
	switch op {
	case "+", "-", "*", "/", "%", "&", "|", "^", "<<", ">>", "&^", "+=", "-=", "*=", "/=", "%=",
		"&=", "|=", "^=", "<<=", ">>=", "&^=", "&&", "||", "<-", "++", "--", "==", "<", ">", "=", "!", "~", "!=",
		"<=", ">=", ":=", "...", "(", ")", "[", "]", "{", "}", ",", ".", ";", ":":
		return true
	default:
		return false
	}
}

func (config) CommentSingle() cfg.OpenClose {
	return cfg.OpenClose{
		Open:  "//",
		Close: "\n",
	}
}

func (config) CommentMulti() cfg.OpenClose {
	return cfg.OpenClose{
		Open:  "/*\n",
		Close: "\n*/\n",
	}
}
