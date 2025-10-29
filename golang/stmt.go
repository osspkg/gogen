/*
 *  Copyright (c) 2025 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package golang

import (
	"io"

	"go.osspkg.com/gogen/internal/gen"
	"go.osspkg.com/gogen/internal/models"
	"go.osspkg.com/gogen/types"
)

type Tokens []types.Token

func newStmt() *Tokens {
	return &Tokens{}
}

func (v *Tokens) Render(w io.Writer) error {
	for _, token := range *v {
		if err := gen.Render(w, token); err != nil {
			return err
		}
	}
	return nil
}

func (v *Tokens) Unwrap() []types.Token {
	return *v
}

func (v *Tokens) _add(args ...types.Token) *Tokens {
	t := Tokens(args)
	*v = append(*v, &t)
	return v
}

func (v *Tokens) Add(args ...types.Token) *Tokens {
	*v = append(*v, args...)
	return v
}

/////////////////////////////////////////////////////////////////

func (v *Tokens) Package(arg string) *Tokens {
	return v._add(
		&models.Keyword{D: "package"},
		&models.Keyword{D: arg},
	)
}

func (v *Tokens) Import(name, module string) *Tokens {
	return v._add(
		&models.Keyword{D: "import"},
		&models.Keyword{D: name, Verify: true},
		&models.Text{D: module},
	)
}

/////////////////////////////////////////////////////////////////

func (v *Tokens) Func() *Tokens {
	*v = append(*v, &models.Keyword{D: "func"})
	return v
}

func (v *Tokens) Return() *Tokens {
	*v = append(*v, &models.Keyword{D: "return"})
	return v
}

func (v *Tokens) Params(args ...types.Token) *Tokens {
	*v = append(*v, &models.Params{D: args, Brace: true})
	return v
}

func (v *Tokens) List(args ...types.Token) *Tokens {
	*v = append(*v, &models.Params{D: args, Brace: false})
	return v
}

func (v *Tokens) Call(args ...types.Token) *Tokens {
	return v.Params(args...)
}

func (v *Tokens) Block(args ...types.Token) *Tokens {
	*v = append(*v, &models.Block{D: args})
	return v
}

/////////////////////////////////////////////////////////////////

func (v *Tokens) Comment(arg string) *Tokens {
	*v = append(*v, &models.Comment[config]{D: arg})
	return v
}

func (v *Tokens) Line() *Tokens {
	*v = append(*v, &models.Letter{D: "\n"})
	return v
}

func (v *Tokens) Text(arg string) *Tokens {
	*v = append(*v, &models.Text{D: arg})
	return v
}

func (v *Tokens) Id(arg string) *Tokens {
	*v = append(*v, &models.Keyword{D: arg, Verify: true})
	return v
}

func (v *Tokens) Nil() *Tokens {
	*v = append(*v, &models.Keyword{D: "nil"})
	return v
}

func (v *Tokens) Type() *Tokens {
	*v = append(*v, &models.Keyword{D: "type"})
	return v
}

func (v *Tokens) Var() *Tokens {
	*v = append(*v, &models.Keyword{D: "var"})
	return v
}

func (v *Tokens) Const() *Tokens {
	*v = append(*v, &models.Keyword{D: "const"})
	return v
}

/////////////////////////////////////////////////////////////////

func (v *Tokens) Op(arg string) *Tokens {
	*v = append(*v, &models.Operation[config]{D: arg})
	return v
}

/////////////////////////////////////////////////////////////////

func (v *Tokens) If() *Tokens {
	*v = append(*v, &models.Keyword{D: "if"})
	return v
}

func (v *Tokens) Else() *Tokens {
	*v = append(*v, &models.Keyword{D: "else"})
	return v
}

func (v *Tokens) Default() *Tokens {
	*v = append(*v, &models.Keyword{D: "default"})
	return v
}

func (v *Tokens) Fallthrough() *Tokens {
	*v = append(*v, &models.Keyword{D: "fallthrough"})
	return v
}

func (v *Tokens) Break() *Tokens {
	*v = append(*v, &models.Keyword{D: "break"})
	return v
}

func (v *Tokens) Case() *Tokens {
	*v = append(*v, &models.Keyword{D: "case"})
	return v
}

func (v *Tokens) Continue() *Tokens {
	*v = append(*v, &models.Keyword{D: "continue"})
	return v
}

func (v *Tokens) Goto() *Tokens {
	*v = append(*v, &models.Keyword{D: "goto"})
	return v
}

/////////////////////////////////////////////////////////////////

/////////////////////////////////////////////////////////////////

func (v *Tokens) Struct() *Tokens {
	*v = append(*v, &models.Keyword{D: "struct"})
	return v
}

func (v *Tokens) Interface() *Tokens {
	*v = append(*v, &models.Keyword{D: "interface"})
	return v
}

func (v *Tokens) Any() *Tokens {
	*v = append(*v, &models.Keyword{D: "any"})
	return v
}

func (v *Tokens) Chan() *Tokens {
	*v = append(*v, &models.Keyword{D: "chan"})
	return v
}

func (v *Tokens) Uint8() *Tokens {
	*v = append(*v, &models.Keyword{D: "uint8"})
	return v
}

func (v *Tokens) Uint16() *Tokens {
	*v = append(*v, &models.Keyword{D: "uint16"})
	return v
}

func (v *Tokens) Uint32() *Tokens {
	*v = append(*v, &models.Keyword{D: "uint32"})
	return v
}

func (v *Tokens) Uint64() *Tokens {
	*v = append(*v, &models.Keyword{D: "uint64"})
	return v
}

func (v *Tokens) Int8() *Tokens {
	*v = append(*v, &models.Keyword{D: "int8"})
	return v
}

func (v *Tokens) Int16() *Tokens {
	*v = append(*v, &models.Keyword{D: "int16"})
	return v
}

func (v *Tokens) Int32() *Tokens {
	*v = append(*v, &models.Keyword{D: "int32"})
	return v
}

func (v *Tokens) Int64() *Tokens {
	*v = append(*v, &models.Keyword{D: "int64"})
	return v
}

func (v *Tokens) Float32() *Tokens {
	*v = append(*v, &models.Keyword{D: "float32"})
	return v
}

func (v *Tokens) Float64() *Tokens {
	*v = append(*v, &models.Keyword{D: "float64"})
	return v
}

func (v *Tokens) Complex64() *Tokens {
	*v = append(*v, &models.Keyword{D: "complex64"})
	return v
}

func (v *Tokens) Complex128() *Tokens {
	*v = append(*v, &models.Keyword{D: "complex128"})
	return v
}

func (v *Tokens) Byte() *Tokens {
	*v = append(*v, &models.Keyword{D: "byte"})
	return v
}

func (v *Tokens) Rune() *Tokens {
	*v = append(*v, &models.Keyword{D: "rune"})
	return v
}

func (v *Tokens) Uint() *Tokens {
	*v = append(*v, &models.Keyword{D: "uint"})
	return v
}

func (v *Tokens) Int() *Tokens {
	*v = append(*v, &models.Keyword{D: "int"})
	return v
}

func (v *Tokens) Uintptr() *Tokens {
	*v = append(*v, &models.Keyword{D: "uintptr"})
	return v
}

func (v *Tokens) String() *Tokens {
	*v = append(*v, &models.Keyword{D: "string"})
	return v
}

func (v *Tokens) Bool() *Tokens {
	*v = append(*v, &models.Keyword{D: "bool"})
	return v
}

func (v *Tokens) Error() *Tokens {
	*v = append(*v, &models.Keyword{D: "error"})
	return v
}

/////////////////////////////////////////////////////////////////

//
//func (v *Tokens) For() *Tokens {
//	*v = append(*v, &models.Keyword{D: "for"})
//	return v
//}

func (v *Tokens) __defer() *Tokens {
	*v = append(*v, &models.Keyword{D: "defer"})
	return v
}

func (v *Tokens) __go() *Tokens {
	*v = append(*v, &models.Keyword{D: "go"})
	return v
}

//
//func (v *Tokens) Map() *Tokens {
//	*v = append(*v, &models.Keyword{D: "map"})
//	return v
//}

//
//func (v *Tokens) Range() *Tokens {
//	*v = append(*v, &models.Keyword{D: "range"})
//	return v
//}

//func (v *Tokens) Select() *Tokens {
//	*v = append(*v, &models.Keyword{D: "select"})
//	return v
//}

//func (v *Tokens) Switch() *Tokens {
//	*v = append(*v, &models.Keyword{D: "switch"})
//	return v
//}

//func (v *Tokens) Array(size int) *Tokens {
//	*v = append(*v, &models.Keyword{D: "bool"})
//	return v
//}
