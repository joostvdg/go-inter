package lexer

import (
	"github.com/joostvdg/go-inter/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `=+(){},;`

	tests := []struct {
		expectedType     token.TokenType
		exptectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	lexer := New(input)
	for i, tt := range tests {
		token := lexer.NextToken()

		if token.Type != tt.expectedType {
			t.Fatalf("test[%d] - tokentype wrong. Expected=%q, got=%q", i, tt.expectedType, token.Type)
		}

		if token.Literal != tt.exptectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. Expected=%q, got %q", i, tt.exptectedLiteral, token.Literal)
		}
	}

}
