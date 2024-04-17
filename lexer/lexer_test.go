package lexer

import (
	"testing"

	"github.com/mty2015/monkey/token"
)

func TestNextToken(t *testing.T) {
	input := `=;(),+{}`
	expectedTokens := []token.Token{
		{Type: token.ASSIGN, Literal: "="},
		{Type: token.SEMICOLON, Literal: ";"},
		{Type: token.LPAREN, Literal: "("},
		{Type: token.RPAREN, Literal: ")"},
		{Type: token.COMMA, Literal: ","},
		{Type: token.PLUS, Literal: "+"},
		{Type: token.LBRACE, Literal: "{"},
		{Type: token.RBRACE, Literal: "}"},
	}

	l := New(input)

	for i, expectedToken := range expectedTokens {
		tok := l.NextToken()

		if tok.Type != expectedToken.Type {
			t.Errorf("Test case %d: expected token type %q, got %q", i, expectedToken.Type, tok.Type)
		}

		if tok.Literal != expectedToken.Literal {
			t.Errorf("Test case %d: expected token literal %q, got %q", i, expectedToken.Literal, tok.Literal)
		}
	}
}
