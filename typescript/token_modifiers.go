/*
 *  Copyright (c) 2025-2026 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package typescript

// Public appends the public access modifier.
func (v *Tokens) Public() *Tokens { return v.Join(keyword("public")) }

// Public appends the public access modifier.
func Public() *Tokens { return create().Public() }

// Private appends the private access modifier.
func (v *Tokens) Private() *Tokens { return v.Join(keyword("private")) }

// Private appends the private access modifier.
func Private() *Tokens { return create().Private() }

// Protected appends the protected access modifier.
func (v *Tokens) Protected() *Tokens { return v.Join(keyword("protected")) }

// Protected appends the protected access modifier.
func Protected() *Tokens { return create().Protected() }

// Static appends the static modifier.
func (v *Tokens) Static() *Tokens { return v.Join(keyword("static")) }

// Static appends the static modifier.
func Static() *Tokens { return create().Static() }

// Readonly appends the readonly modifier.
func (v *Tokens) Readonly() *Tokens { return v.Join(keyword("readonly")) }

// Readonly appends the readonly modifier.
func Readonly() *Tokens { return create().Readonly() }

// Abstract appends the abstract modifier.
func (v *Tokens) Abstract() *Tokens { return v.Join(keyword("abstract")) }

// Abstract appends the abstract modifier.
func Abstract() *Tokens { return create().Abstract() }

// Override appends the override modifier.
func (v *Tokens) Override() *Tokens { return v.Join(keyword("override")) }

// Override appends the override modifier.
func Override() *Tokens { return create().Override() }

// Constructor appends the constructor keyword.
func (v *Tokens) Constructor() *Tokens { return v.Join(keyword("constructor")) }

// Constructor appends the constructor keyword.
func Constructor() *Tokens { return create().Constructor() }
