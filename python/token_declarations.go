/*
 *  Copyright (c) 2025-2026 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package python

import (
	"fmt"
	"io"

	"go.osspkg.com/gogen/internal/gen"
	"go.osspkg.com/gogen/types"
)

// Import appends an import statement for a dotted module path.
func (v *Tokens) Import(module string) *Tokens { return v.Join(keyword("import"), raw(module)) }

// Import creates an import statement for a dotted module path.
func Import(module string) *Tokens { return create().Import(module) }

// ImportAs appends an import statement with an alias.
func (v *Tokens) ImportAs(module, alias string) *Tokens {
	return v.Join(keyword("import"), raw(module), keyword("as"), identifier(alias))
}

// ImportAs creates an import statement with an alias.
func ImportAs(module, alias string) *Tokens { return create().ImportAs(module, alias) }

// FromImport appends a from-import statement with optional names.
func (v *Tokens) FromImport(module string, names ...types.Token) *Tokens {
	v.Join(keyword("from"), raw(module), keyword("import"))
	if len(names) > 0 {
		v.Join(list(names...))
	}
	return v
}

// FromImport creates a from-import statement with optional names.
func FromImport(module string, names ...types.Token) *Tokens {
	return create().FromImport(module, names...)
}

// Def appends a function declaration header; call Block to add its suite.
func (v *Tokens) Def(name string) *Tokens { return v.Join(keyword("def"), identifier(name)) }

// Def creates a function declaration header.
func Def(name string) *Tokens { return create().Def(name) }

// AsyncDef appends an asynchronous function declaration header.
func (v *Tokens) AsyncDef(name string) *Tokens {
	return v.Join(keyword("async"), keyword("def"), identifier(name))
}

// AsyncDef creates an asynchronous function declaration header.
func AsyncDef(name string) *Tokens { return create().AsyncDef(name) }

// Class appends a class declaration header with optional base classes.
func (v *Tokens) Class(name string, bases ...types.Token) *Tokens {
	v.Join(keyword("class"), identifier(name))
	if len(bases) > 0 {
		v.Join(brackets(bases...))
	}
	return v
}

// Class creates a class declaration header with optional base classes.
func Class(name string, bases ...types.Token) *Tokens { return create().Class(name, bases...) }

// Decorator appends a decorator expression followed by a line break.
func (v *Tokens) Decorator(expression types.Token) *Tokens {
	return v.Join(&decorator{expression: expression})
}

// Decorator creates a decorator expression followed by a line break.
func Decorator(expression types.Token) *Tokens { return create().Decorator(expression) }

type decorator struct{ expression types.Token }

func (v *decorator) Render(w io.Writer) error {
	if v.expression == nil {
		return fmt.Errorf("decorator expression is nil")
	}
	if err := writeString(w, "@"); err != nil {
		return err
	}
	if err := v.expression.Render(w); err != nil {
		return err
	}
	err := writeString(w, "\n")
	return err
}

func (v *decorator) RenderLayout() gen.Layout {
	return gen.Layout{First: gen.Style{Kind: gen.KindPrefixOperator}, Last: gen.Style{Kind: gen.KindLine}}
}
