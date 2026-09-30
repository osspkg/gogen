/*
 *  Copyright (c) 2025-2026 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package typescript

import (
	"errors"
	"fmt"
	"io"
	"regexp"
	"strings"

	"go.osspkg.com/gogen/internal/gen"
	"go.osspkg.com/gogen/types"
)

var jsxAttributeName = regexp.MustCompile(`^[A-Za-z_:][A-Za-z0-9_.:-]*$`)

// JSXElement creates a structural TSX element with attributes and children.
func (v *Tokens) JSXElement(tag types.Token, attributes []types.Token, children ...types.Token) *Tokens {
	return v.Join(&jsxElement{tag: tag, attributes: attributes, children: children})
}

// JSXElement creates a structural TSX element with attributes and children.
func JSXElement(tag types.Token, attributes []types.Token, children ...types.Token) *Tokens {
	return create().JSXElement(tag, attributes, children...)
}

// JSXFragment creates a TSX fragment containing child nodes.
func (v *Tokens) JSXFragment(children ...types.Token) *Tokens {
	return v.Join(&jsxFragment{children: children})
}

// JSXFragment creates a TSX fragment containing child nodes.
func JSXFragment(children ...types.Token) *Tokens { return create().JSXFragment(children...) }

// JSXAttribute creates a quoted and escaped TSX string attribute.
func (v *Tokens) JSXAttribute(name, value string) *Tokens {
	return v.Join(&jsxAttribute{name: name, value: value})
}

// JSXAttribute creates a quoted and escaped TSX string attribute.
func JSXAttribute(name, value string) *Tokens { return create().JSXAttribute(name, value) }

// JSXAttributeExpr creates a TSX attribute containing an expression.
func (v *Tokens) JSXAttributeExpr(name string, expression types.Token) *Tokens {
	return v.Join(&jsxAttributeExpression{name: name, expression: expression})
}

// JSXAttributeExpr creates a TSX attribute containing an expression.
func JSXAttributeExpr(name string, expression types.Token) *Tokens {
	return create().JSXAttributeExpr(name, expression)
}

// JSXBooleanAttribute creates a valueless TSX boolean attribute.
func (v *Tokens) JSXBooleanAttribute(name string) *Tokens {
	return v.Join(&jsxBooleanAttribute{name: name})
}

// JSXBooleanAttribute creates a valueless TSX boolean attribute.
func JSXBooleanAttribute(name string) *Tokens { return create().JSXBooleanAttribute(name) }

// JSXSpreadAttribute creates a TSX spread attribute from an expression.
func (v *Tokens) JSXSpreadAttribute(expression types.Token) *Tokens {
	return v.Join(&jsxSpreadAttribute{expression: expression})
}

// JSXSpreadAttribute creates a TSX spread attribute from an expression.
func JSXSpreadAttribute(expression types.Token) *Tokens {
	return create().JSXSpreadAttribute(expression)
}

// JSXText creates an escaped TSX text node.
func (v *Tokens) JSXText(value string) *Tokens { return v.Join(&jsxText{value: value}) }

// JSXText creates an escaped TSX text node.
func JSXText(value string) *Tokens { return create().JSXText(value) }

// JSXExpr creates a TSX expression child.
func (v *Tokens) JSXExpr(expression types.Token) *Tokens {
	return v.Join(&jsxExpression{expression: expression})
}

// JSXExpr creates a TSX expression child.
func JSXExpr(expression types.Token) *Tokens { return create().JSXExpr(expression) }

type jsxElement struct {
	tag        types.Token
	attributes []types.Token
	children   []types.Token
}

func (v *jsxElement) Render(w io.Writer) error {
	if v.tag == nil {
		return errors.New("TSX element has no tag")
	}
	if _, err := io.WriteString(w, "<"); err != nil {
		return err
	}
	if err := v.tag.Render(w); err != nil {
		return err
	}
	for _, attribute := range v.attributes {
		if _, err := io.WriteString(w, " "); err != nil {
			return err
		}
		if err := attribute.Render(w); err != nil {
			return err
		}
	}
	if len(v.children) == 0 {
		_, err := io.WriteString(w, " />")
		return err
	}
	if _, err := io.WriteString(w, ">"); err != nil {
		return err
	}
	for _, child := range v.children {
		if child == nil {
			return errors.New("TSX element has a nil child")
		}
		if err := child.Render(w); err != nil {
			return err
		}
	}
	if _, err := io.WriteString(w, "</"); err != nil {
		return err
	}
	if err := v.tag.Render(w); err != nil {
		return err
	}
	_, err := io.WriteString(w, ">")
	return err
}

func (v *jsxElement) RenderLayout() gen.Layout {
	style := gen.Style{Kind: gen.KindFragment}
	return gen.Layout{First: style, Last: style}
}

type jsxFragment struct{ children []types.Token }

func (v *jsxFragment) Render(w io.Writer) error {
	if _, err := io.WriteString(w, "<>"); err != nil {
		return err
	}
	for _, child := range v.children {
		if child == nil {
			return errors.New("TSX fragment has a nil child")
		}
		if err := child.Render(w); err != nil {
			return err
		}
	}
	_, err := io.WriteString(w, "</>")
	return err
}
func (v *jsxFragment) RenderLayout() gen.Layout { return jsxElementLayout() }

func jsxElementLayout() gen.Layout {
	style := gen.Style{Kind: gen.KindFragment}
	return gen.Layout{First: style, Last: style}
}

type jsxAttribute struct {
	name  string
	value string
}

func (v *jsxAttribute) Render(w io.Writer) error {
	if !jsxAttributeName.MatchString(v.name) {
		return fmt.Errorf("invalid TSX attribute name %q", v.name)
	}
	if _, err := io.WriteString(w, v.name+`="`); err != nil {
		return err
	}
	if _, err := io.WriteString(w, escapeJSXAttribute(v.value)); err != nil {
		return err
	}
	_, err := io.WriteString(w, `"`)
	return err
}
func (v *jsxAttribute) RenderLayout() gen.Layout { return jsxElementLayout() }

type jsxAttributeExpression struct {
	name       string
	expression types.Token
}

func (v *jsxAttributeExpression) Render(w io.Writer) error {
	if !jsxAttributeName.MatchString(v.name) {
		return fmt.Errorf("invalid TSX attribute name %q", v.name)
	}
	if v.expression == nil {
		return fmt.Errorf("TSX attribute %q has no expression", v.name)
	}
	if _, err := io.WriteString(w, v.name+"={"); err != nil {
		return err
	}
	if err := v.expression.Render(w); err != nil {
		return err
	}
	_, err := io.WriteString(w, "}")
	return err
}
func (v *jsxAttributeExpression) RenderLayout() gen.Layout { return jsxElementLayout() }

type jsxBooleanAttribute struct{ name string }

func (v *jsxBooleanAttribute) Render(w io.Writer) error {
	if !jsxAttributeName.MatchString(v.name) {
		return fmt.Errorf("invalid TSX attribute name %q", v.name)
	}
	_, err := io.WriteString(w, v.name)
	return err
}
func (v *jsxBooleanAttribute) RenderLayout() gen.Layout { return jsxElementLayout() }

type jsxSpreadAttribute struct{ expression types.Token }

func (v *jsxSpreadAttribute) Render(w io.Writer) error {
	if v.expression == nil {
		return errors.New("TSX spread attribute has no expression")
	}
	if _, err := io.WriteString(w, "{..."); err != nil {
		return err
	}
	if err := v.expression.Render(w); err != nil {
		return err
	}
	_, err := io.WriteString(w, "}")
	return err
}
func (v *jsxSpreadAttribute) RenderLayout() gen.Layout { return jsxElementLayout() }

type jsxText struct{ value string }

func (v *jsxText) Render(w io.Writer) error {
	_, err := io.WriteString(w, escapeJSXText(v.value))
	return err
}
func (v *jsxText) RenderLayout() gen.Layout { return jsxElementLayout() }

type jsxExpression struct{ expression types.Token }

func (v *jsxExpression) Render(w io.Writer) error {
	if v.expression == nil {
		return errors.New("TSX expression is nil")
	}
	if _, err := io.WriteString(w, "{"); err != nil {
		return err
	}
	if err := v.expression.Render(w); err != nil {
		return err
	}
	_, err := io.WriteString(w, "}")
	return err
}
func (v *jsxExpression) RenderLayout() gen.Layout { return jsxElementLayout() }

func escapeJSXAttribute(value string) string {
	return strings.NewReplacer("&", "&amp;", `"`, "&quot;", "<", "&lt;", ">", "&gt;", "\n", "&#10;", "\r", "&#13;", "\t", "&#9;").Replace(value)
}

func escapeJSXText(value string) string {
	return strings.NewReplacer("&", "&amp;", "<", "&lt;", ">", "&gt;", "{", "&#123;", "}", "&#125;").Replace(value)
}
