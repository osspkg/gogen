/*
 *  Copyright (c) 2025-2026 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package typescript

import (
	"go.osspkg.com/gogen/internal/models"
	"go.osspkg.com/gogen/types"
)

// ID appends a validated TypeScript identifier.
func (v *Tokens) ID(name string) *Tokens { return v.Join(identifier(name)) }

// ID appends a validated TypeScript identifier.
func ID(name string) *Tokens { return create().ID(name) }

// Raw appends source text verbatim.
func (v *Tokens) Raw(source string) *Tokens { return v.Join(raw(source)) }

// Raw appends source text verbatim.
func Raw(source string) *Tokens { return create().Raw(source) }

// Text appends a quoted and escaped JavaScript string literal.
func (v *Tokens) Text(value string) *Tokens { return v.Join(text(value)) }

// Text appends a quoted and escaped JavaScript string literal.
func Text(value string) *Tokens { return create().Text(value) }

// Op appends a supported TypeScript operator.
func (v *Tokens) Op(op string) *Tokens { return v.Join(operation(op)) }

// Op appends a supported TypeScript operator.
func Op(op string) *Tokens { return create().Op(op) }

// Call creates or appends a parenthesized call argument list.
func (v *Tokens) Call(args ...types.Token) *Tokens { return v.Join(brackets(args...)) }

// Call creates or appends a parenthesized call argument list.
func Call(args ...types.Token) *Tokens { return create().Call(args...) }

// Bracket creates or appends a parenthesized token list.
func (v *Tokens) Bracket(args ...types.Token) *Tokens { return v.Join(brackets(args...)) }

// Bracket creates or appends a parenthesized token list.
func Bracket(args ...types.Token) *Tokens { return create().Bracket(args...) }

// List creates or appends a comma-separated token list without delimiters.
func (v *Tokens) List(args ...types.Token) *Tokens { return v.Join(list(args...)) }

// List creates or appends a comma-separated token list without delimiters.
func List(args ...types.Token) *Tokens { return create().List(args...) }

// Index appends an indexed access expression.
func (v *Tokens) Index(index types.Token) *Tokens { return v.Join(square(index)) }

// Index appends an indexed access expression.
func Index(index types.Token) *Tokens { return create().Index(index) }

// TypeArgs creates or appends generic type arguments.
func (v *Tokens) TypeArgs(args ...types.Token) *Tokens { return v.Join(&angleArgs{args: args}) }

// TypeArgs creates or appends generic type arguments.
func TypeArgs(args ...types.Token) *Tokens { return create().TypeArgs(args...) }

// KeyValue creates a key-value member for an object or type literal.
func (v *Tokens) KeyValue(key, value types.Token) *Tokens { return v.Join(keyed(key, value)) }

// KeyValue creates a key-value member for an object or type literal.
func KeyValue(key, value types.Token) *Tokens { return create().KeyValue(key, value) }

// Block creates or appends a brace-delimited block with indentation.
func (v *Tokens) Block(args ...types.Token) *Tokens { return v.Join(block(args...)) }

// Block creates or appends a brace-delimited block with indentation.
func Block(args ...types.Token) *Tokens { return create().Block(args...) }

// Comment appends a TypeScript line comment.
func (v *Tokens) Comment(value string) *Tokens { return v.Join(comment(value)) }

// Comment appends a TypeScript line comment.
func Comment(value string) *Tokens { return create().Comment(value) }

// Line appends an explicit line break.
func (v *Tokens) Line() *Tokens { return v.Join(lineToken{}) }

// Line appends an explicit line break.
func Line() *Tokens { return create().Line() }

// Pkg appends a module-style dotted prefix.
func (v *Tokens) Pkg(module string) *Tokens {
	if module == "" {
		return v
	}
	return v.Join(identifier(module), operation("."))
}

// Pkg appends a module-style dotted prefix.
func Pkg(module string) *Tokens { return create().Pkg(module) }

// ObjectLiteral creates or appends a multiline object literal.
func (v *Tokens) ObjectLiteral(properties ...types.Token) *Tokens {
	return v.Join(&objectLiteral{properties: properties})
}

// ObjectLiteral creates or appends a multiline object literal.
func ObjectLiteral(properties ...types.Token) *Tokens { return create().ObjectLiteral(properties...) }

// ArrayLiteral creates or appends an array literal.
func (v *Tokens) ArrayLiteral(values ...types.Token) *Tokens {
	return v.Join(square(values...))
}

// ArrayLiteral creates or appends an array literal.
func ArrayLiteral(values ...types.Token) *Tokens { return create().ArrayLiteral(values...) }

// Template creates or appends a template string from text and expressions.
func (v *Tokens) Template(parts ...types.Token) *Tokens {
	return v.Join(&templateLiteral{parts: parts})
}

// Template creates or appends a template string from text and expressions.
func Template(parts ...types.Token) *Tokens { return create().Template(parts...) }

// TemplateText creates a token containing escaped template-string text.
func TemplateText(value string) types.Token { return &templateText{value: value} }

// TemplateExpr creates an interpolated template-string expression token.
func TemplateExpr(expression types.Token) types.Token {
	return &templateExpression{expression: expression}
}

// NewCall creates a constructor call expression.
func (v *Tokens) NewCall(target types.Token, args ...types.Token) *Tokens {
	return v.Join(keyword("new"), target, brackets(args...))
}

// NewCall creates a constructor call expression.
func NewCall(target types.Token, args ...types.Token) *Tokens {
	return create().NewCall(target, args...)
}

// Selector appends a property selector.
func (v *Tokens) Selector(name string) *Tokens { return v.Join(operation("."), identifier(name)) }

// Selector appends a property selector.
func Selector(name string) *Tokens { return create().Selector(name) }

// OptionalChain appends an optional property selector.
func (v *Tokens) OptionalChain(name string) *Tokens { return v.Join(operation("?."), identifier(name)) }

// OptionalChain appends an optional property selector.
func OptionalChain(name string) *Tokens { return create().OptionalChain(name) }

// ArrayType creates an array type from its element type.
func (v *Tokens) ArrayType(element types.Token) *Tokens {
	return v.Join(element, &models.SquareBracket{})
}

// ArrayType creates an array type from its element type.
func ArrayType(element types.Token) *Tokens { return create().ArrayType(element) }
