/*
 *  Copyright (c) 2025-2026 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package typescript

import "go.osspkg.com/gogen/types"

// Import creates an import declaration from a module.
func (v *Tokens) Import(module string, clause ...types.Token) *Tokens {
	return v.Join(&importDeclaration{module: module, clause: clause})
}

// Import creates an import declaration from a module.
func Import(module string, clause ...types.Token) *Tokens { return create().Import(module, clause...) }

// ImportType creates a type-only import declaration.
func (v *Tokens) ImportType(module string, clause ...types.Token) *Tokens {
	return v.Join(&importDeclaration{module: module, clause: clause, typeOnly: true})
}

// ImportType creates a type-only import declaration.
func ImportType(module string, clause ...types.Token) *Tokens {
	return create().ImportType(module, clause...)
}

// ImportNames creates a named import or export specifier list.
func (v *Tokens) ImportNames(specifiers ...types.Token) *Tokens {
	return v.Join(&importNames{specifiers: specifiers})
}

// ImportNames creates a named import or export specifier list.
func ImportNames(specifiers ...types.Token) *Tokens { return create().ImportNames(specifiers...) }

// ImportNamespace creates a namespace import declaration.
func (v *Tokens) ImportNamespace(alias, module string) *Tokens {
	return v.Import(module, rawWord("* as").Join(identifier(alias)))
}

// ImportNamespace creates a namespace import declaration.
func ImportNamespace(alias, module string) *Tokens { return create().ImportNamespace(alias, module) }

// Export appends the export keyword.
func (v *Tokens) Export() *Tokens { return v.Join(keyword("export")) }

// Export appends the export keyword.
func Export() *Tokens { return create().Export() }

// ExportDefault creates a default export declaration.
func (v *Tokens) ExportDefault(declaration types.Token) *Tokens {
	return v.Join(&exportDeclaration{declaration: declaration, defaultExport: true})
}

// ExportDefault creates a default export declaration.
func ExportDefault(declaration types.Token) *Tokens { return create().ExportDefault(declaration) }

// ExportNames creates a named export declaration.
func (v *Tokens) ExportNames(specifiers ...types.Token) *Tokens {
	return v.Join(&exportNames{specifiers: specifiers})
}

// ExportNames creates a named export declaration.
func ExportNames(specifiers ...types.Token) *Tokens { return create().ExportNames(specifiers...) }

// ExportFrom creates a re-export declaration from a module.
func (v *Tokens) ExportFrom(module string, specifiers ...types.Token) *Tokens {
	return v.Join(&exportNames{specifiers: specifiers, module: module})
}

// ExportFrom creates a re-export declaration from a module.
func ExportFrom(module string, specifiers ...types.Token) *Tokens {
	return create().ExportFrom(module, specifiers...)
}

// From appends a from clause with a quoted module path.
func (v *Tokens) From(module string) *Tokens { return v.Join(keyword("from"), text(module)) }

// Type appends the type-alias declaration keyword.
func (v *Tokens) Type() *Tokens { return v.Join(keyword("type")) }

// Type appends the type-alias declaration keyword.
func Type() *Tokens { return create().Type() }

// Interface appends the interface declaration keyword.
func (v *Tokens) Interface() *Tokens { return v.Join(keyword("interface")) }

// Interface appends the interface declaration keyword.
func Interface() *Tokens { return create().Interface() }

// Class appends the class declaration keyword.
func (v *Tokens) Class() *Tokens { return v.Join(keyword("class")) }

// Class appends the class declaration keyword.
func Class() *Tokens { return create().Class() }

// Function appends the function declaration keyword.
func (v *Tokens) Function() *Tokens { return v.Join(keyword("function")) }

// Function appends the function declaration keyword.
func Function() *Tokens { return create().Function() }

// Const appends the const declaration keyword.
func (v *Tokens) Const() *Tokens { return v.Join(keyword("const")) }

// Const appends the const declaration keyword.
func Const() *Tokens { return create().Const() }

// Let appends the let declaration keyword.
func (v *Tokens) Let() *Tokens { return v.Join(keyword("let")) }

// Let appends the let declaration keyword.
func Let() *Tokens { return create().Let() }

// Var appends the var declaration keyword.
func (v *Tokens) Var() *Tokens { return v.Join(keyword("var")) }

// Var appends the var declaration keyword.
func Var() *Tokens { return create().Var() }

// Async appends the async modifier.
func (v *Tokens) Async() *Tokens { return v.Join(keyword("async")) }

// Async appends the async modifier.
func Async() *Tokens { return create().Async() }

// Await appends the await keyword.
func (v *Tokens) Await() *Tokens { return v.Join(keyword("await")) }

// Await appends the await keyword.
func Await() *Tokens { return create().Await() }

// Return appends the return statement keyword.
func (v *Tokens) Return() *Tokens { return v.Join(keyword("return")) }

// Return appends the return statement keyword.
func Return() *Tokens { return create().Return() }

// Throw appends the throw statement keyword.
func (v *Tokens) Throw() *Tokens { return v.Join(keyword("throw")) }

// Throw appends the throw statement keyword.
func Throw() *Tokens { return create().Throw() }

// Extends appends the extends keyword.
func (v *Tokens) Extends() *Tokens { return v.Join(keyword("extends")) }

// Extends appends the extends keyword.
func Extends() *Tokens { return create().Extends() }

// Implements appends the implements keyword.
func (v *Tokens) Implements() *Tokens { return v.Join(keyword("implements")) }

// Implements appends the implements keyword.
func Implements() *Tokens { return create().Implements() }

// New appends the new keyword.
func (v *Tokens) New() *Tokens { return v.Join(keyword("new")) }

// New appends the new keyword.
func New() *Tokens { return create().New() }

// This appends the this keyword.
func (v *Tokens) This() *Tokens { return v.Join(keyword("this")) }

// This appends the this keyword.
func This() *Tokens { return create().This() }

// Namespace appends the namespace declaration keyword.
func (v *Tokens) Namespace() *Tokens { return v.Join(keyword("namespace")) }

// Namespace appends the namespace declaration keyword.
func Namespace() *Tokens { return create().Namespace() }
