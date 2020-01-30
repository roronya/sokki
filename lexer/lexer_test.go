package lexer

import (
	"testing"

	"github.com/roronya/sokki/token"
)

func TestNew(t *testing.T) {
	input := "abc"
	l := New(input)

	if string(l.input) != "abc" {
		t.Fatalf("input is not abc. got=%s", string(l.input))
	}

	if l.position != 0 {
		t.Fatalf("position is not 0. got=%d", l.position)
	}

}

func TestNextToken(t *testing.T) {
	input := `マリア様の庭に集う少女たちが、
今日も天使のような無垢な笑顔で、
背の高い門をくぐり抜けていく。
`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.PARAGRAPH, "マリア様の庭に集う少女たちが、"},
		{token.NEWLINE, "\n"},
		{token.PARAGRAPH, "今日も天使のような無垢な笑顔で、"},
		{token.NEWLINE, "\n"},
		{token.PARAGRAPH, "背の高い門をくぐり抜けていく。"},
		{token.NEWLINE, "\n"},
		{token.EOD, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - Literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}
