/*
 *  Copyright (c) 2025-2026 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package python

import "go.osspkg.com/gogen/types"

// If appends an if header; call Block to add its suite.
func (v *Tokens) If(condition types.Token) *Tokens { return v.Join(keyword("if"), condition) }

// If creates an if header.
func If(condition types.Token) *Tokens { return create().If(condition) }

// Elif appends an elif header.
func (v *Tokens) Elif(condition types.Token) *Tokens { return v.Join(keyword("elif"), condition) }

// Elif creates an elif header.
func Elif(condition types.Token) *Tokens { return create().Elif(condition) }

// Else appends an else header.
func (v *Tokens) Else() *Tokens { return v.Join(keyword("else")) }

// Else creates an else header.
func Else() *Tokens { return create().Else() }

// ForEach appends a for-in header.
func (v *Tokens) ForEach(target, iterable types.Token) *Tokens {
	return v.Join(keyword("for"), target, keyword("in"), iterable)
}

// ForEach creates a for-in header.
func ForEach(target, iterable types.Token) *Tokens { return create().ForEach(target, iterable) }

// AsyncForEach appends an async for-in header.
func (v *Tokens) AsyncForEach(target, iterable types.Token) *Tokens {
	return v.Join(keyword("async"), keyword("for"), target, keyword("in"), iterable)
}

// AsyncForEach creates an async for-in header.
func AsyncForEach(target, iterable types.Token) *Tokens {
	return create().AsyncForEach(target, iterable)
}

// While appends a while header.
func (v *Tokens) While(condition types.Token) *Tokens { return v.Join(keyword("while"), condition) }

// While creates a while header.
func While(condition types.Token) *Tokens { return create().While(condition) }

// With appends a with header for one or more context expressions.
func (v *Tokens) With(contexts ...types.Token) *Tokens {
	return v.Join(keyword("with"), list(contexts...))
}

// With creates a with header for one or more context expressions.
func With(contexts ...types.Token) *Tokens { return create().With(contexts...) }

// AsyncWith appends an async with header.
func (v *Tokens) AsyncWith(contexts ...types.Token) *Tokens {
	return v.Join(keyword("async"), keyword("with"), list(contexts...))
}

// AsyncWith creates an async with header.
func AsyncWith(contexts ...types.Token) *Tokens { return create().AsyncWith(contexts...) }

// Try appends a try header.
func (v *Tokens) Try() *Tokens { return v.Join(keyword("try")) }

// Try creates a try header.
func Try() *Tokens { return create().Try() }

// Except appends an except header, optionally with an exception type.
func (v *Tokens) Except(exception ...types.Token) *Tokens {
	v.Join(keyword("except"))
	if len(exception) > 0 {
		v.Join(list(exception...))
	}
	return v
}

// Except creates an except header, optionally with an exception type.
func Except(exception ...types.Token) *Tokens { return create().Except(exception...) }

// Finally appends a finally header.
func (v *Tokens) Finally() *Tokens { return v.Join(keyword("finally")) }

// Finally creates a finally header.
func Finally() *Tokens { return create().Finally() }

// Match appends a match header.
func (v *Tokens) Match(subject types.Token) *Tokens { return v.Join(keyword("match"), subject) }

// Match creates a match header.
func Match(subject types.Token) *Tokens { return create().Match(subject) }

// Case appends a case header.
func (v *Tokens) Case(pattern types.Token) *Tokens { return v.Join(keyword("case"), pattern) }

// Case creates a case header.
func Case(pattern types.Token) *Tokens { return create().Case(pattern) }

// Return appends the return keyword.
func (v *Tokens) Return() *Tokens { return v.Join(keyword("return")) }

// Return creates a return statement token sequence.
func Return() *Tokens { return create().Return() }

// Yield appends the yield keyword.
func (v *Tokens) Yield() *Tokens { return v.Join(keyword("yield")) }

// Yield creates a yield expression token sequence.
func Yield() *Tokens { return create().Yield() }

// Raise appends the raise keyword.
func (v *Tokens) Raise() *Tokens { return v.Join(keyword("raise")) }

// Raise creates a raise statement token sequence.
func Raise() *Tokens { return create().Raise() }

// Pass appends the pass keyword.
func (v *Tokens) Pass() *Tokens { return v.Join(keyword("pass")) }

// Pass creates a pass statement token sequence.
func Pass() *Tokens { return create().Pass() }

// Break appends the break keyword.
func (v *Tokens) Break() *Tokens { return v.Join(keyword("break")) }

// Break creates a break statement token sequence.
func Break() *Tokens { return create().Break() }

// Continue appends the continue keyword.
func (v *Tokens) Continue() *Tokens { return v.Join(keyword("continue")) }

// Continue creates a continue statement token sequence.
func Continue() *Tokens { return create().Continue() }

// Assert appends the assert keyword.
func (v *Tokens) Assert() *Tokens { return v.Join(keyword("assert")) }

// Assert creates an assert statement token sequence.
func Assert() *Tokens { return create().Assert() }

// Del appends the del keyword.
func (v *Tokens) Del() *Tokens { return v.Join(keyword("del")) }

// Del creates a del statement token sequence.
func Del() *Tokens { return create().Del() }

// Global appends the global keyword.
func (v *Tokens) Global() *Tokens { return v.Join(keyword("global")) }

// Global creates a global declaration token sequence.
func Global() *Tokens { return create().Global() }

// Nonlocal appends the nonlocal keyword.
func (v *Tokens) Nonlocal() *Tokens { return v.Join(keyword("nonlocal")) }

// Nonlocal creates a nonlocal declaration token sequence.
func Nonlocal() *Tokens { return create().Nonlocal() }

// Await appends the await keyword.
func (v *Tokens) Await() *Tokens { return v.Join(keyword("await")) }

// Await creates an await expression token sequence.
func Await() *Tokens { return create().Await() }

// As appends the as keyword.
func (v *Tokens) As() *Tokens { return v.Join(keyword("as")) }

// As creates an as-clause token sequence.
func As() *Tokens { return create().As() }

// Arrow appends the function return annotation arrow.
func (v *Tokens) Arrow() *Tokens { return v.Join(operation("->")) }

// Arrow creates the function return annotation arrow.
func Arrow() *Tokens { return create().Arrow() }

// Colon appends a colon punctuation token.
func (v *Tokens) Colon() *Tokens { return v.Join(operation(":")) }

// Colon creates a colon punctuation token.
func Colon() *Tokens { return create().Colon() }

// Comma appends a comma punctuation token.
func (v *Tokens) Comma() *Tokens { return v.Join(operation(",")) }

// Comma creates a comma punctuation token.
func Comma() *Tokens { return create().Comma() }
