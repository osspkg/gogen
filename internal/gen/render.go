/*
 *  Copyright (c) 2025-2026 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by the BSD 3-Clause license that can be found in the LICENSE file.
 */

package gen

import (
	"fmt"
	"io"

	"go.osspkg.com/gogen/types"
)

func Render(w io.Writer, args ...any) error {
	for _, arg := range args {
		if err := drawAny(w, arg); err != nil {
			return err
		}
	}
	return nil
}

func drawAny(w io.Writer, arg any) error {
	switch v := arg.(type) {
	case Unwrap:
		return drawList(w, v.Unwrap())
	case []types.Token:
		return drawList(w, v)
	case types.Token:
		return drawList(w, []types.Token{v})
	case string:
		n, err := io.WriteString(w, v)
		if err != nil {
			return err
		}
		if n != len(v) {
			return io.ErrShortWrite
		}
		return nil
	default:
		fmt.Printf("[X] %T\n", arg)
		return nil
	}
}

func drawList(w io.Writer, list []types.Token) error {
	list = flatten(list)

	var (
		previous Style
		havePrev bool
	)
	for _, token := range list {
		layout := layoutOf(token)
		first := layout.First
		if first.Kind == KindUnaryOperator {
			if !havePrev || !canEndExpression(previous) {
				first.Kind = KindPrefixOperator
				layout.First.Kind = KindPrefixOperator
				if layout.Last.Kind == KindUnaryOperator {
					layout.Last.Kind = KindPrefixOperator
				}
			} else {
				first.Kind = KindOperator
				layout.First.Kind = KindOperator
				if layout.Last.Kind == KindUnaryOperator {
					layout.Last.Kind = KindOperator
				}
			}
		}

		if havePrev {
			if _, err := io.WriteString(w, separator(previous, first)); err != nil {
				return err
			}
		}
		if err := token.Render(w); err != nil {
			return err
		}
		previous = layout.Last
		havePrev = previous.Kind != KindLine
	}
	return nil
}

func flatten(in []types.Token) []types.Token {
	out := make([]types.Token, 0, len(in))
	for _, token := range in {
		if nested, ok := token.(Unwrap); ok {
			out = append(out, flatten(nested.Unwrap())...)
			continue
		}
		out = append(out, token)
	}
	return out
}

func layoutOf(token types.Token) Layout {
	if styled, ok := token.(Styled); ok {
		return styled.RenderLayout()
	}
	unknown := Style{Kind: KindUnknown}
	return Layout{First: unknown, Last: unknown}
}

func separator(previous, next Style) string {
	if previous.Kind == KindLine || next.Kind == KindLine || previous.Kind == KindComment {
		return ""
	}
	if next.Kind == KindComma || next.Kind == KindDot || next.Kind == KindColon || next.Kind == KindSemicolon || next.Kind == KindCloseParen || next.Kind == KindCloseBracket {
		return ""
	}
	if previous.Kind == KindDot || previous.Kind == KindOpenParen || previous.Kind == KindOpenSquare || previous.Kind == KindTypePrefix || previous.Kind == KindCloseBracket {
		return ""
	}
	if next.Kind == KindOpenParen || next.Kind == KindOpenSquare {
		return ""
	}
	if next.Kind == KindBlockOpen {
		return " "
	}
	if previous.Kind == KindBlockOpen {
		return ""
	}
	if previous.Kind == KindComma || previous.Kind == KindColon || previous.Kind == KindOperator {
		return " "
	}
	if previous.Kind == KindPrefixOperator {
		return ""
	}
	if previous.Kind == KindPostfixOperator {
		return ""
	}
	if next.Kind == KindOperator {
		return " "
	}
	if next.Kind == KindPrefixOperator {
		return wordLike(previous)
	}
	if next.Kind == KindPostfixOperator {
		return ""
	}
	if previous.Kind == KindCloseParen || previous.Kind == KindCloseBracket || previous.Kind == KindBlockClose {
		return " "
	}
	if wordLikeKind(previous.Kind) && wordLikeKind(next.Kind) {
		return " "
	}
	if next.Kind == KindSemicolon || previous.Kind == KindSemicolon {
		return " "
	}
	return " "
}

func wordLike(style Style) string {
	if wordLikeKind(style.Kind) {
		return " "
	}
	return ""
}

func wordLikeKind(kind Kind) bool {
	return kind == KindWord || kind == KindLiteral || kind == KindUnknown
}

func canEndExpression(style Style) bool {
	switch style.Kind {
	case KindWord:
		return style.CanEndExpression
	case KindLiteral, KindFragment, KindTypePrefix, KindCloseBracket, KindCloseParen, KindBlockClose, KindPostfixOperator:
		return true
	default:
		return false
	}
}

// LayoutOf returns the outer token styles of a token sequence after unwrapping
// nested token builders.
func LayoutOf(tokens []types.Token) Layout {
	flat := flatten(tokens)
	if len(flat) == 0 {
		return Layout{}
	}
	first := layoutOf(flat[0])
	last := layoutOf(flat[len(flat)-1])
	return Layout{First: first.First, Last: last.Last}
}

func Params(token types.Token) []types.Token {
	if nested, ok := token.(Unwrap); ok {
		return nested.Unwrap()
	}
	return []types.Token{token}
}
