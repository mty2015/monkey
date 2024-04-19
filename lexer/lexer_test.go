package lexer

import (
	"testing"

	"github.com/mty2015/monkey/token"
)

func TestNextToken(t *testing.T) {
	input := `
	=;(),+{}
	let 你 = 5;
	`
	expectedTokens := []token.Token{
		{Type: token.ASSIGN, Literal: "="},
		{Type: token.SEMICOLON, Literal: ";"},
		{Type: token.LPAREN, Literal: "("},
		{Type: token.RPAREN, Literal: ")"},
		{Type: token.COMMA, Literal: ","},
		{Type: token.PLUS, Literal: "+"},
		{Type: token.LBRACE, Literal: "{"},
		{Type: token.RBRACE, Literal: "}"},
		{Type: token.LET, Literal: "let"},
		{Type: token.IDENT, Literal: "你"},
		{Type: token.ASSIGN, Literal: "="},
		{Type: token.INT, Literal: "5"},
		{Type: token.SEMICOLON, Literal: ";"},
		{Type: token.EOF, Literal: ""},
	}

	l := New(input)

	for i, expectedToken := range expectedTokens {
		tok := l.NextToken()

		if tok.Type != expectedToken.Type {
			t.Errorf("Test case %d: expected token type %q, got %q", i, expectedToken.Type, tok.Type)
			return
		}

		if tok.Literal != expectedToken.Literal {
			t.Errorf("Test case %d: expected token literal %q, got %q", i, expectedToken.Literal, tok.Literal)
			return
		}
	}
}
