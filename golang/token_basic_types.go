/*
 *  Copyright (c) 2025-2026 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package golang

import (
	"go.osspkg.com/gogen/types"
)

// Struct appends the struct type keyword to the token sequence.
func (v *Tokens) Struct() *Tokens {
	*v = append(*v, keyword("struct"))
	return v
}

// Struct creates a token sequence containing the struct type keyword.
func Struct() *Tokens {
	return create().Struct()
}

// Interface appends the interface type keyword to the token sequence.
func (v *Tokens) Interface() *Tokens {
	*v = append(*v, keyword("interface"))
	return v
}

// Interface creates a token sequence containing the interface type keyword.
func Interface() *Tokens {
	return create().Interface()
}

// Any appends the predeclared any type to the token sequence.
func (v *Tokens) Any() *Tokens {
	*v = append(*v, keyword("any"))
	return v
}

// Any creates a token sequence containing the predeclared any type.
func Any() *Tokens {
	return create().Any()
}

// Nil appends the predeclared nil value to the token sequence.
func (v *Tokens) Nil() *Tokens {
	*v = append(*v, keyword("nil"))
	return v
}

// Nil creates a token sequence containing the predeclared nil value.
func Nil() *Tokens {
	return create().Nil()
}

// Chan appends the chan type keyword to the token sequence.
func (v *Tokens) Chan() *Tokens {
	*v = append(*v, keyword("chan"))
	return v
}

// Chan creates a token sequence containing the chan type keyword.
func Chan() *Tokens {
	return create().Chan()
}

// Uint8 appends the predeclared uint8 type to the token sequence.
func (v *Tokens) Uint8() *Tokens {
	*v = append(*v, keyword("uint8"))
	return v
}

// Uint8 creates a token sequence containing the predeclared uint8 type.
func Uint8() *Tokens {
	return create().Uint8()
}

// Uint16 appends the predeclared uint16 type to the token sequence.
func (v *Tokens) Uint16() *Tokens {
	*v = append(*v, keyword("uint16"))
	return v
}

// Uint16 creates a token sequence containing the predeclared uint16 type.
func Uint16() *Tokens {
	return create().Uint16()
}

// Uint32 appends the predeclared uint32 type to the token sequence.
func (v *Tokens) Uint32() *Tokens {
	*v = append(*v, keyword("uint32"))
	return v
}

// Uint32 creates a token sequence containing the predeclared uint32 type.
func Uint32() *Tokens {
	return create().Uint32()
}

// Uint64 appends the predeclared uint64 type to the token sequence.
func (v *Tokens) Uint64() *Tokens {
	*v = append(*v, keyword("uint64"))
	return v
}

// Uint64 creates a token sequence containing the predeclared uint64 type.
func Uint64() *Tokens {
	return create().Uint64()
}

// Int8 appends the predeclared int8 type to the token sequence.
func (v *Tokens) Int8() *Tokens {
	*v = append(*v, keyword("int8"))
	return v
}

// Int8 creates a token sequence containing the predeclared int8 type.
func Int8() *Tokens {
	return create().Int8()
}

// Int16 appends the predeclared int16 type to the token sequence.
func (v *Tokens) Int16() *Tokens {
	*v = append(*v, keyword("int16"))
	return v
}

// Int16 creates a token sequence containing the predeclared int16 type.
func Int16() *Tokens {
	return create().Int16()
}

// Int32 appends the predeclared int32 type to the token sequence.
func (v *Tokens) Int32() *Tokens {
	*v = append(*v, keyword("int32"))
	return v
}

// Int32 creates a token sequence containing the predeclared int32 type.
func Int32() *Tokens {
	return create().Int32()
}

// Int64 appends the predeclared int64 type to the token sequence.
func (v *Tokens) Int64() *Tokens {
	*v = append(*v, keyword("int64"))
	return v
}

// Int64 creates a token sequence containing the predeclared int64 type.
func Int64() *Tokens {
	return create().Int64()
}

// Float32 appends the predeclared float32 type to the token sequence.
func (v *Tokens) Float32() *Tokens {
	*v = append(*v, keyword("float32"))
	return v
}

// Float32 creates a token sequence containing the predeclared float32 type.
func Float32() *Tokens {
	return create().Float32()
}

// Float64 appends the predeclared float64 type to the token sequence.
func (v *Tokens) Float64() *Tokens {
	*v = append(*v, keyword("float64"))
	return v
}

// Float64 creates a token sequence containing the predeclared float64 type.
func Float64() *Tokens {
	return create().Float64()
}

// Complex64 appends the predeclared complex64 type to the token sequence.
func (v *Tokens) Complex64() *Tokens {
	*v = append(*v, keyword("complex64"))
	return v
}

// Complex64 creates a token sequence containing the predeclared complex64 type.
func Complex64() *Tokens {
	return create().Complex64()
}

// Complex128 appends the predeclared complex128 type to the token sequence.
func (v *Tokens) Complex128() *Tokens {
	*v = append(*v, keyword("complex128"))
	return v
}

// Complex128 creates a token sequence containing the predeclared complex128 type.
func Complex128() *Tokens {
	return create().Complex128()
}

// Byte appends the predeclared byte alias to the token sequence.
func (v *Tokens) Byte() *Tokens {
	*v = append(*v, keyword("byte"))
	return v
}

// Byte creates a token sequence containing the predeclared byte alias.
func Byte() *Tokens {
	return create().Byte()
}

// Rune appends the predeclared rune alias to the token sequence.
func (v *Tokens) Rune() *Tokens {
	*v = append(*v, keyword("rune"))
	return v
}

// Rune creates a token sequence containing the predeclared rune alias.
func Rune() *Tokens {
	return create().Rune()
}

// Uint appends the predeclared uint type to the token sequence.
func (v *Tokens) Uint() *Tokens {
	*v = append(*v, keyword("uint"))
	return v
}

// Uint creates a token sequence containing the predeclared uint type.
func Uint() *Tokens {
	return create().Uint()
}

// Int appends the predeclared int type to the token sequence.
func (v *Tokens) Int() *Tokens {
	*v = append(*v, keyword("int"))
	return v
}

// Int creates a token sequence containing the predeclared int type.
func Int() *Tokens {
	return create().Int()
}

// Uintptr appends the predeclared uintptr type to the token sequence.
func (v *Tokens) Uintptr() *Tokens {
	*v = append(*v, keyword("uintptr"))
	return v
}

// Uintptr creates a token sequence containing the predeclared uintptr type.
func Uintptr() *Tokens {
	return create().Uintptr()
}

// String appends the predeclared string type to the token sequence.
func (v *Tokens) String() *Tokens {
	*v = append(*v, keyword("string"))
	return v
}

// String creates a token sequence containing the predeclared string type.
func String() *Tokens {
	return create().String()
}

// Bool appends the predeclared bool type to the token sequence.
func (v *Tokens) Bool() *Tokens {
	*v = append(*v, keyword("bool"))
	return v
}

// Bool creates a token sequence containing the predeclared bool type.
func Bool() *Tokens {
	return create().Bool()
}

// Error appends the predeclared error interface to the token sequence.
func (v *Tokens) Error() *Tokens {
	*v = append(*v, keyword("error"))
	return v
}

// Error creates a token sequence containing the predeclared error interface.
func Error() *Tokens {
	return create().Error()
}

// Map appends a map type with the supplied key and value types to the token sequence.
func (v *Tokens) Map(key, val types.Token) *Tokens {
	return v.Join(
		rawToken("map["),
		rawTokenOf(key),
		rawToken("]"),
		rawTokenOf(val),
	)
}

// Map creates a token sequence containing a map type with the supplied key and value types.
func Map(key, val types.Token) *Tokens {
	return create().Map(key, val)
}
