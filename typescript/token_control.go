/*
 *  Copyright (c) 2025-2026 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package typescript

import "go.osspkg.com/gogen/types"

// If starts an if statement.
func (v *Tokens) If() *Tokens { return v.Join(rawWord("if ")) }

// If starts an if statement.
func If() *Tokens { return create().If() }

// Else appends an else branch.
func (v *Tokens) Else() *Tokens { return v.Join(keyword("else")) }

// Else appends an else branch.
func Else() *Tokens { return create().Else() }

// ElseIf starts an else-if branch.
func (v *Tokens) ElseIf() *Tokens { return v.Join(rawWord("else if ")) }

// ElseIf starts an else-if branch.
func ElseIf() *Tokens { return create().ElseIf() }

// For starts a for statement.
func (v *Tokens) For() *Tokens { return v.Join(rawWord("for ")) }

// For starts a for statement.
func For() *Tokens { return create().For() }

// ForOf creates a for-of loop with a const binding.
func (v *Tokens) ForOf(name string, iterable types.Token) *Tokens {
	return v.Join(rawWord("for (const"), identifier(name), keyword("of"), iterable, symbol(")"))
}

// ForOf creates a for-of loop with a const binding.
func ForOf(name string, iterable types.Token) *Tokens { return create().ForOf(name, iterable) }

// ForIn creates a for-in loop with a const binding.
func (v *Tokens) ForIn(name string, object types.Token) *Tokens {
	return v.Join(rawWord("for (const"), identifier(name), keyword("in"), object, symbol(")"))
}

// ForIn creates a for-in loop with a const binding.
func ForIn(name string, object types.Token) *Tokens { return create().ForIn(name, object) }

// While starts a while loop.
func (v *Tokens) While() *Tokens { return v.Join(rawWord("while ")) }

// While starts a while loop.
func While() *Tokens { return create().While() }

// Switch starts a switch statement.
func (v *Tokens) Switch() *Tokens { return v.Join(rawWord("switch ")) }

// Switch starts a switch statement.
func Switch() *Tokens { return create().Switch() }

// Case starts a case clause.
func (v *Tokens) Case() *Tokens { return v.Join(keyword("case")) }

// Case starts a case clause.
func Case() *Tokens { return create().Case() }

// Default starts a default clause.
func (v *Tokens) Default() *Tokens { return v.Join(keyword("default")) }

// Default starts a default clause.
func Default() *Tokens { return create().Default() }

// Break appends a break statement keyword.
func (v *Tokens) Break() *Tokens { return v.Join(keyword("break")) }

// Break appends a break statement keyword.
func Break() *Tokens { return create().Break() }

// Continue appends a continue statement keyword.
func (v *Tokens) Continue() *Tokens { return v.Join(keyword("continue")) }

// Continue appends a continue statement keyword.
func Continue() *Tokens { return create().Continue() }

// Try starts a try statement.
func (v *Tokens) Try() *Tokens { return v.Join(keyword("try")) }

// Try starts a try statement.
func Try() *Tokens { return create().Try() }

// Catch starts a catch clause.
func (v *Tokens) Catch() *Tokens { return v.Join(rawWord("catch ")) }

// Catch starts a catch clause.
func Catch() *Tokens { return create().Catch() }

// Finally starts a finally clause.
func (v *Tokens) Finally() *Tokens { return v.Join(keyword("finally")) }

// Finally starts a finally clause.
func Finally() *Tokens { return create().Finally() }

// Do starts a do-while loop.
func (v *Tokens) Do() *Tokens { return v.Join(keyword("do")) }

// Do starts a do-while loop.
func Do() *Tokens { return create().Do() }

// Semicolon appends a statement terminator.
func (v *Tokens) Semicolon() *Tokens { return v.Join(operation(";")) }

// Semicolon appends a statement terminator.
func Semicolon() *Tokens { return create().Semicolon() }

// Comma appends a comma separator.
func (v *Tokens) Comma() *Tokens { return v.Join(operation(",")) }

// Comma appends a comma separator.
func Comma() *Tokens { return create().Comma() }

// Colon appends a colon separator.
func (v *Tokens) Colon() *Tokens { return v.Join(operation(":")) }

// Colon appends a colon separator.
func Colon() *Tokens { return create().Colon() }

// Optional appends an optional-property marker.
func (v *Tokens) Optional() *Tokens { return v.Join(postfix("?")) }

// Optional appends an optional-property marker.
func Optional() *Tokens { return create().Optional() }

// NonNull appends a non-null assertion marker.
func (v *Tokens) NonNull() *Tokens { return v.Join(postfix("!")) }

// NonNull appends a non-null assertion marker.
func NonNull() *Tokens { return create().NonNull() }

// Spread appends the spread operator.
func (v *Tokens) Spread() *Tokens { return v.Join(operation("...")) }

// Spread appends the spread operator.
func Spread() *Tokens { return create().Spread() }

// Arrow appends the arrow-function operator.
func (v *Tokens) Arrow() *Tokens { return v.Join(operation("=>")) }

// Arrow appends the arrow-function operator.
func Arrow() *Tokens { return create().Arrow() }

// As appends the as keyword for aliases or type assertions.
func (v *Tokens) As() *Tokens { return v.Join(keyword("as")) }

// As appends the as keyword for aliases or type assertions.
func As() *Tokens { return create().As() }

// Of appends the of keyword.
func (v *Tokens) Of() *Tokens { return v.Join(keyword("of")) }

// Of appends the of keyword.
func Of() *Tokens { return create().Of() }
