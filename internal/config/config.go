/*
 *  Copyright (c) 2025 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package config

type Config interface {
	CommentSingle() OpenClose
	CommentMulti() OpenClose
	OperationAvailable(op string) bool
}

type OpenClose struct {
	Open, Close string
}
