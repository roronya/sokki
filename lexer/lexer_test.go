package lexer

import (
	"testing"

	"github.com/roronya/sokki/token"
)

func TestNew(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"abc", "abc"},
		{" abc \n", "abc"},
		{" あいう　\n", "あいう"},
	}
	for _, tt := range tests {
		l := New(tt.input)

		if string(l.input) != tt.expected {
			t.Fatalf("input is not %s. got=%s",
				tt.expected, string(l.input))
		}

		if l.position != 0 {
			t.Fatalf("position is not 0. got=%d",
				l.position)
		}
	}

}

func TestNextToken(t *testing.T) {
	input := `マリア様の庭に集う少女たちが、
今日も天使のような無垢な笑顔で、 >
背の高い門をくぐり抜けていく。 >>

汚れを知らない心身を包むのは、深い色の制服。
スカートのプリーツは乱さないように、 >>
白いセーラーカラーは翻さないように、 >
ゆっくりと歩くのが、ここでのたしなみ。

私立リリアン女学園。ここは乙女の園。
`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.STRING, "マリア様の庭に集う少女たちが、"},
		{token.NEWLINE, "\n"},
		{token.STRING, "今日も天使のような無垢な笑顔で、"},
		{token.SHIFT, " >"},
		{token.NEWLINE, "\n"},
		{token.STRING, "背の高い門をくぐり抜けていく。"},
		{token.MORESHIFT, " >>"},
		{token.NEWLINE, "\n"},
		{token.NEWLINE, "\n"},
		{token.STRING, "汚れを知らない心身を包むのは、深い色の制服。"},
		{token.NEWLINE, "\n"},
		{token.STRING, "スカートのプリーツは乱さないように、"},
		{token.MORESHIFT, " >>"},
		{token.NEWLINE, "\n"},
		{token.STRING, "白いセーラーカラーは翻さないように、"},
		{token.SHIFT, " >"},
		{token.NEWLINE, "\n"},
		{token.STRING, "ゆっくりと歩くのが、ここでのたしなみ。"},
		{token.NEWLINE, "\n"},
		{token.NEWLINE, "\n"},
		{token.STRING, "私立リリアン女学園。ここは乙女の園。"},
		{token.EOD, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - Literal wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}
	}
}

func TestNextTokenWithSpace(t *testing.T) {
	input := ` aaa`
	l := New(input)

	tok := l.NextToken()
	if tok.Type != token.STRING {
		t.Fatalf("Literal wrong. expected=%q, got=%q", " aaa", tok.Literal)
	}
	tok = l.NextToken()
	if tok.Type != token.EOD {
		t.Fatalf("Literal wrong. expected=%q, got=%q", "", tok.Literal)
	}
}

func TestNextTokenWithShiftLiteral(t *testing.T) {
	input := `aaa >>bbb`
	l := New(input)

	tok := l.NextToken()
	if tok.Type != token.STRING {
		t.Fatalf("Literal wrong. expected=%q, got=%q",
			"aaa >>bbb", tok.Literal)
	}
}

// MORESHIFTのtokenizeをバグらせてたから追加した
func TestNextToken2(t *testing.T) {
	input := `hoge
hoge >
hoge >>
`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.STRING, "hoge"},
		{token.NEWLINE, "\n"},
		{token.STRING, "hoge"},
		{token.SHIFT, " >"},
		{token.NEWLINE, "\n"},
		{token.STRING, "hoge"},
		{token.MORESHIFT, " >>"},
		{token.EOD, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - Literal wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}
	}

}
