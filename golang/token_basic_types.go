/*
 *  Copyright (c) 2025-2026 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package golang

import (
	"go.osspkg.com/gogen/internal/models"
	"go.osspkg.com/gogen/types"
)

func (v *Tokens) Struct() *Tokens {
	*v = append(*v, &models.Keyword{D: "struct"})
	return v
}

func Struct() *Tokens {
	return create().Struct()
}

func (v *Tokens) Interface() *Tokens {
	*v = append(*v, &models.Keyword{D: "interface"})
	return v
}

func Interface() *Tokens {
	return create().Interface()
}

func (v *Tokens) Any() *Tokens {
	*v = append(*v, &models.Keyword{D: "any"})
	return v
}

func Any() *Tokens {
	return create().Any()
}

func (v *Tokens) Nil() *Tokens {
	*v = append(*v, &models.Keyword{D: "nil"})
	return v
}

func Nil() *Tokens {
	return create().Nil()
}

func (v *Tokens) Chan() *Tokens {
	*v = append(*v, &models.Keyword{D: "chan"})
	return v
}

func Chan() *Tokens {
	return create().Chan()
}

func (v *Tokens) Uint8() *Tokens {
	*v = append(*v, &models.Keyword{D: "uint8"})
	return v
}

func Uint8() *Tokens {
	return create().Uint8()
}

func (v *Tokens) Uint16() *Tokens {
	*v = append(*v, &models.Keyword{D: "uint16"})
	return v
}

func Uint16() *Tokens {
	return create().Uint16()
}

func (v *Tokens) Uint32() *Tokens {
	*v = append(*v, &models.Keyword{D: "uint32"})
	return v
}

func Uint32() *Tokens {
	return create().Uint32()
}

func (v *Tokens) Uint64() *Tokens {
	*v = append(*v, &models.Keyword{D: "uint64"})
	return v
}

func Uint64() *Tokens {
	return create().Uint64()
}

func (v *Tokens) Int8() *Tokens {
	*v = append(*v, &models.Keyword{D: "int8"})
	return v
}

func Int8() *Tokens {
	return create().Int8()
}

func (v *Tokens) Int16() *Tokens {
	*v = append(*v, &models.Keyword{D: "int16"})
	return v
}

func Int16() *Tokens {
	return create().Int16()
}

func (v *Tokens) Int32() *Tokens {
	*v = append(*v, &models.Keyword{D: "int32"})
	return v
}

func Int32() *Tokens {
	return create().Int32()
}

func (v *Tokens) Int64() *Tokens {
	*v = append(*v, &models.Keyword{D: "int64"})
	return v
}

func Int64() *Tokens {
	return create().Int64()
}

func (v *Tokens) Float32() *Tokens {
	*v = append(*v, &models.Keyword{D: "float32"})
	return v
}

func Float32() *Tokens {
	return create().Float32()
}

func (v *Tokens) Float64() *Tokens {
	*v = append(*v, &models.Keyword{D: "float64"})
	return v
}

func Float64() *Tokens {
	return create().Float64()
}

func (v *Tokens) Complex64() *Tokens {
	*v = append(*v, &models.Keyword{D: "complex64"})
	return v
}

func Complex64() *Tokens {
	return create().Complex64()
}

func (v *Tokens) Complex128() *Tokens {
	*v = append(*v, &models.Keyword{D: "complex128"})
	return v
}

func Complex128() *Tokens {
	return create().Complex128()
}

func (v *Tokens) Byte() *Tokens {
	*v = append(*v, &models.Keyword{D: "byte"})
	return v
}

func Byte() *Tokens {
	return create().Byte()
}

func (v *Tokens) Rune() *Tokens {
	*v = append(*v, &models.Keyword{D: "rune"})
	return v
}

func Rune() *Tokens {
	return create().Rune()
}

func (v *Tokens) Uint() *Tokens {
	*v = append(*v, &models.Keyword{D: "uint"})
	return v
}

func Uint() *Tokens {
	return create().Uint()
}

func (v *Tokens) Int() *Tokens {
	*v = append(*v, &models.Keyword{D: "int"})
	return v
}

func Int() *Tokens {
	return create().Int()
}

func (v *Tokens) Uintptr() *Tokens {
	*v = append(*v, &models.Keyword{D: "uintptr"})
	return v
}

func Uintptr() *Tokens {
	return create().Uintptr()
}

func (v *Tokens) String() *Tokens {
	*v = append(*v, &models.Keyword{D: "string"})
	return v
}

func String() *Tokens {
	return create().String()
}

func (v *Tokens) Bool() *Tokens {
	*v = append(*v, &models.Keyword{D: "bool"})
	return v
}

func Bool() *Tokens {
	return create().Bool()
}

func (v *Tokens) Error() *Tokens {
	*v = append(*v, &models.Keyword{D: "error"})
	return v
}

func Error() *Tokens {
	return create().Error()
}

func (v *Tokens) Map(key, val types.Token) *Tokens {
	return v.Join(
		&models.Raw{D: "map["},
		&models.Raw{T: key},
		&models.Raw{D: "]"},
		&models.Raw{T: val},
	)
}

func Map(key, val types.Token) *Tokens {
	return create().Map(key, val)
}
