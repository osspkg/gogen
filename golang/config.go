/*
 *  Copyright (c) 2025 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package golang

import (
	config2 "go.osspkg.com/gogen/internal/config"
)

var _ config2.Config = (*config)(nil)

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

func (config) CommentSingle() config2.OpenClose {
	return config2.OpenClose{
		Open:  "//",
		Close: "\n",
	}
}

func (config) CommentMulti() config2.OpenClose {
	return config2.OpenClose{
		Open:  "/*\n",
		Close: "\n*/\n",
	}
}
