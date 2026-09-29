/*
 *  Copyright (c) 2025-2026 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package golang

import "go.osspkg.com/gogen/internal/models"

// If appends the if keyword to the token sequence.
func (v *Tokens) If() *Tokens {
	*v = append(*v, &models.Keyword{D: "if"})
	return v
}

// If creates a token sequence containing the if keyword.
func If() *Tokens {
	return create().If()
}

// Else appends the else keyword to the token sequence.
func (v *Tokens) Else() *Tokens {
	*v = append(*v, &models.Keyword{D: "else"})
	return v
}

// ElseIf appends an else-if clause keyword to the token sequence.
func (v *Tokens) ElseIf() *Tokens {
	*v = append(*v, &models.Keyword{D: "else if"})
	return v
}

// Default appends the default keyword for a switch clause to the token sequence.
func (v *Tokens) Default() *Tokens {
	*v = append(*v, &models.Keyword{D: "default"})
	return v
}

// Default creates a token sequence containing the default keyword for a switch clause.
func Default() *Tokens {
	return create().Default()
}

// Fallthrough appends the fallthrough keyword to the token sequence.
func (v *Tokens) Fallthrough() *Tokens {
	*v = append(*v, &models.Keyword{D: "fallthrough"})
	return v
}

// Break appends the break keyword to the token sequence.
func (v *Tokens) Break() *Tokens {
	*v = append(*v, &models.Keyword{D: "break"})
	return v
}

// Case appends the case keyword to the token sequence.
func (v *Tokens) Case() *Tokens {
	*v = append(*v, &models.Keyword{D: "case"})
	return v
}

// Case creates a token sequence containing the case keyword.
func Case() *Tokens {
	return create().Case()
}

// Continue appends the continue keyword to the token sequence.
func (v *Tokens) Continue() *Tokens {
	*v = append(*v, &models.Keyword{D: "continue"})
	return v
}

// Goto appends the goto keyword to the token sequence.
func (v *Tokens) Goto() *Tokens {
	*v = append(*v, &models.Keyword{D: "goto"})
	return v
}

// For appends the for keyword to the token sequence.
func (v *Tokens) For() *Tokens {
	*v = append(*v, &models.Keyword{D: "for"})
	return v
}

// For creates a token sequence containing the for keyword.
func For() *Tokens {
	return create().For()
}

// Range appends the range keyword to the token sequence.
func (v *Tokens) Range() *Tokens {
	*v = append(*v, &models.Keyword{D: "range"})
	return v
}

// Select creates a token sequence containing the select keyword.
func Select() *Tokens {
	return create().Join(&models.Keyword{D: "select"})
}

// Switch creates a token sequence containing the switch keyword.
func Switch() *Tokens {
	return create().Join(&models.Keyword{D: "switch"})
}
