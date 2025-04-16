/*
 *  Copyright (c) 2025 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package golang

import "go.osspkg.com/gogen/types"

func Comment(arg string) *Tokens {
	return newStmt().Comment(arg)
}

func Line() *Tokens {
	return newStmt().Line()
}

func Return() *Tokens {
	return newStmt().Return()
}

func Nil() *Tokens {
	return newStmt().Nil()
}

func Var() *Tokens {
	return newStmt().Var()
}

func Text(arg string) *Tokens {
	return newStmt().Text(arg)
}

func Id(arg string) *Tokens {
	return newStmt().Id(arg)
}

func Package(arg string) *Tokens {
	return newStmt().Package(arg)
}

func Import(name, module string) *Tokens {
	return newStmt().Import(name, module)
}

func Func() *Tokens {
	return newStmt().Func()
}

func Params(args ...types.Token) *Tokens {
	return newStmt().Params(args...)
}

func List(args ...types.Token) *Tokens {
	return newStmt().List(args...)
}

func Defer() *Tokens {
	return newStmt().__defer()
}

func Go() *Tokens {
	return newStmt().__go()
}
