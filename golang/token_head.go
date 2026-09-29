/*
 *  Copyright (c) 2025-2026 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package golang

// Package appends a package declaration followed by a line break to the token sequence.
func (v *Tokens) Package(arg string) *Tokens {
	return v.Join(
		create().Join(
			keyword("package"),
			keyword(arg),
		),
	).Line()
}

// Package creates a token sequence containing a package declaration followed by a line break.
func Package(arg string) *Tokens {
	return create().Package(arg)
}

//------------------------------------------------------

// Import appends a single import declaration with an alias and path to the token sequence. name is the local import name and module is the quoted import path.
func (v *Tokens) Import(name, module string) *Tokens {
	return v.Join(
		create().Join(
			keyword("import"),
			identifier(name),
			textToken(module),
		),
	).Line()
}

// Import creates a token sequence containing a single import declaration with an alias and path. name is the local import name and module is the quoted import path.
func Import(name, module string) *Tokens {
	return create().Import(name, module)
}
