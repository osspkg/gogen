/*
 *  Copyright (c) 2025-2026 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package python

import "go.osspkg.com/gogen/types"

// Any appends Any, the typing module's unconstrained type annotation.
func (v *Tokens) Any() *Tokens { return v.ID("Any") }

// Any creates an Any type annotation. Import Any from typing in generated code.
func Any() *Tokens { return create().Any() }

// Bool appends Python's bool type.
func (v *Tokens) Bool() *Tokens { return v.ID("bool") }

// Bool creates Python's bool type.
func Bool() *Tokens { return create().Bool() }

// Bytes appends Python's bytes type.
func (v *Tokens) Bytes() *Tokens { return v.ID("bytes") }

// Bytes creates Python's bytes type.
func Bytes() *Tokens { return create().Bytes() }

// ByteArray appends Python's bytearray type.
func (v *Tokens) ByteArray() *Tokens { return v.ID("bytearray") }

// ByteArray creates Python's bytearray type.
func ByteArray() *Tokens { return create().ByteArray() }

// Complex appends Python's complex type.
func (v *Tokens) Complex() *Tokens { return v.ID("complex") }

// Complex creates Python's complex type.
func Complex() *Tokens { return create().Complex() }

// DictType appends a parameterized dict type such as dict[str, int].
func (v *Tokens) DictType(key, value types.Token) *Tokens {
	return v.ID("dict").Index(List(key, value))
}

// DictType creates a parameterized dict type such as dict[str, int].
func DictType(key, value types.Token) *Tokens { return create().DictType(key, value) }

// Float appends Python's float type.
func (v *Tokens) Float() *Tokens { return v.ID("float") }

// Float creates Python's float type.
func Float() *Tokens { return create().Float() }

// FrozenSetType appends a parameterized frozenset type.
func (v *Tokens) FrozenSetType(element types.Token) *Tokens {
	return v.ID("frozenset").Index(element)
}

// FrozenSetType creates a parameterized frozenset type.
func FrozenSetType(element types.Token) *Tokens { return create().FrozenSetType(element) }

// Int appends Python's int type.
func (v *Tokens) Int() *Tokens { return v.ID("int") }

// Int creates Python's int type.
func Int() *Tokens { return create().Int() }

// ListType appends a parameterized list type such as list[str].
func (v *Tokens) ListType(element types.Token) *Tokens {
	return v.ID("list").Index(element)
}

// ListType creates a parameterized list type such as list[str].
func ListType(element types.Token) *Tokens { return create().ListType(element) }

// MemoryView appends Python's memoryview type.
func (v *Tokens) MemoryView() *Tokens { return v.ID("memoryview") }

// MemoryView creates Python's memoryview type.
func MemoryView() *Tokens { return create().MemoryView() }

// ObjectType appends Python's object type.
func (v *Tokens) ObjectType() *Tokens { return v.ID("object") }

// ObjectType creates Python's object type.
func ObjectType() *Tokens { return create().ObjectType() }

// Range appends Python's range type.
func (v *Tokens) Range() *Tokens { return v.ID("range") }

// Range creates Python's range type.
func Range() *Tokens { return create().Range() }

// SetType appends a parameterized set type such as set[int].
func (v *Tokens) SetType(element types.Token) *Tokens {
	return v.ID("set").Index(element)
}

// SetType creates a parameterized set type such as set[int].
func SetType(element types.Token) *Tokens { return create().SetType(element) }

// Str appends Python's str type.
func (v *Tokens) Str() *Tokens { return v.ID("str") }

// Str creates Python's str type.
func Str() *Tokens { return create().Str() }

// String appends Python's str type as a readable alias for Str.
func (v *Tokens) String() *Tokens { return v.Str() }

// String creates Python's str type as a readable alias for Str.
func String() *Tokens { return create().String() }

// TupleType appends a tuple type. Pass element types for fixed tuples or
// Ellipsis() as the second argument for a variable-length tuple.
func (v *Tokens) TupleType(elements ...types.Token) *Tokens {
	v.ID("tuple")
	if len(elements) > 0 {
		v.Index(List(elements...))
	}
	return v
}

// TupleType creates a tuple type. Pass element types for fixed tuples or
// Ellipsis() as the second argument for a variable-length tuple.
func TupleType(elements ...types.Token) *Tokens { return create().TupleType(elements...) }

// TypeType appends Python's type type.
func (v *Tokens) TypeType() *Tokens { return v.ID("type") }

// TypeType creates Python's type type.
func TypeType() *Tokens { return create().TypeType() }
