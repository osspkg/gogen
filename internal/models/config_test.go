package models

import (
	"bytes"
	"io"
	"strings"
	"testing"

	"go.osspkg.com/gogen/internal/config"
	"go.osspkg.com/gogen/internal/gen"
)

type alternateConfig struct{}

func (alternateConfig) CommentSingle() config.OpenClose {
	return config.OpenClose{Open: "#", Close: "\n", SpaceAfterOpenWhenNeeded: true}
}

func (alternateConfig) CommentMulti() config.OpenClose {
	return config.OpenClose{Open: "/*", Close: "*/"}
}

func (alternateConfig) OperationAvailable(string) bool        { return true }
func (alternateConfig) OperationKind(string) config.TokenKind { return config.TokenDefault }
func (alternateConfig) RawKind(string, bool) config.TokenKind { return config.TokenDefault }
func (alternateConfig) IsIdentifier(string) bool              { return true }
func (alternateConfig) CanEndExpression(word string) bool     { return word != "return" }
func (alternateConfig) QuoteString(value string) string {
	return "'" + strings.ReplaceAll(value, "'", "\\'") + "'"
}

func TestLanguageConfigControlsTextAndComments(t *testing.T) {
	var got bytes.Buffer
	if err := (&Text[alternateConfig]{D: "hello", C: alternateConfig{}}).Render(&got); err != nil {
		t.Fatal(err)
	}
	if want := "'hello'"; got.String() != want {
		t.Fatalf("Text.Render() = %q, want %q", got.String(), want)
	}

	got.Reset()
	if err := (&Comment[alternateConfig]{D: "note", c: alternateConfig{}}).Render(&got); err != nil {
		t.Fatal(err)
	}
	if want := "# note\n"; got.String() != want {
		t.Fatalf("Comment.Render() = %q, want %q", got.String(), want)
	}
}

func TestLanguageConfigControlsIdentifierAndOperatorLayout(t *testing.T) {
	identifier := &Keyword[alternateConfig]{D: "yield", Verify: true, C: alternateConfig{}}
	var got bytes.Buffer
	if err := gen.Render(&got, identifier); err != nil {
		t.Fatal(err)
	}
	if want := "yield"; got.String() != want {
		t.Fatalf("identifier render = %q, want %q", got.String(), want)
	}

	operator := &Operation[alternateConfig]{D: "@", c: alternateConfig{}}
	if got := operator.RenderLayout().First.Kind; got != gen.KindOperator {
		t.Fatalf("operator layout kind = %v, want %v", got, gen.KindOperator)
	}

	var writer io.Writer = io.Discard
	if err := operator.Render(writer); err != nil {
		t.Fatal(err)
	}
}
