/*
 *  Copyright (c) 2025-2026 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package typescript

import (
	"fmt"
	"io"

	"go.osspkg.com/gogen/internal/gen"
	"go.osspkg.com/gogen/types"
)

type importDeclaration struct {
	module   string
	clause   []types.Token
	typeOnly bool
}

func (v *importDeclaration) Render(w io.Writer) error {
	parts := []types.Token{keyword("import")}
	if v.typeOnly {
		parts = append(parts, keyword("type"))
	}
	if len(v.clause) > 0 {
		parts = append(parts, v.clause[0])
		for _, clause := range v.clause[1:] {
			parts = append(parts, symbol(","), clause)
		}
		parts = append(parts, keyword("from"))
	}
	parts = append(parts, text(v.module), symbol(";"))
	return gen.Render(w, parts)
}

func (v *importDeclaration) RenderLayout() gen.Layout {
	last := gen.Style{Kind: gen.KindSemicolon, Text: ";"}
	return gen.Layout{First: wordStyle("import"), Last: last}
}

type importNames struct{ specifiers []types.Token }

func (v *importNames) Render(w io.Writer) error {
	if _, err := io.WriteString(w, "{"); err != nil {
		return err
	}
	if len(v.specifiers) > 0 {
		if _, err := io.WriteString(w, " "); err != nil {
			return err
		}
		for index, specifier := range v.specifiers {
			if index > 0 {
				if _, err := io.WriteString(w, ", "); err != nil {
					return err
				}
			}
			if err := specifier.Render(w); err != nil {
				return err
			}
		}
		if _, err := io.WriteString(w, " "); err != nil {
			return err
		}
	}
	_, err := io.WriteString(w, "}")
	return err
}

func (v *importNames) RenderLayout() gen.Layout {
	return gen.Layout{
		First: gen.Style{Kind: gen.KindBlockOpen, Text: "{"},
		Last:  gen.Style{Kind: gen.KindBlockClose, Text: "}"},
	}
}

type exportDeclaration struct {
	declaration   types.Token
	defaultExport bool
}

func (v *exportDeclaration) Render(w io.Writer) error {
	if v.declaration == nil {
		return fmt.Errorf("export declaration is nil")
	}
	parts := []types.Token{keyword("export")}
	if v.defaultExport {
		parts = append(parts, keyword("default"))
	}
	parts = append(parts, v.declaration)
	return gen.Render(w, parts)
}

func (v *exportDeclaration) RenderLayout() gen.Layout {
	layout := gen.LayoutOf([]types.Token{v.declaration})
	return gen.Layout{First: wordStyle("export"), Last: layout.Last}
}

type exportNames struct {
	specifiers []types.Token
	module     string
}

func (v *exportNames) Render(w io.Writer) error {
	parts := []types.Token{keyword("export"), &importNames{specifiers: v.specifiers}}
	if v.module != "" {
		parts = append(parts, keyword("from"), text(v.module))
	}
	parts = append(parts, symbol(";"))
	return gen.Render(w, parts)
}

func (v *exportNames) RenderLayout() gen.Layout {
	return gen.Layout{First: wordStyle("export"), Last: gen.Style{Kind: gen.KindSemicolon, Text: ";"}}
}
