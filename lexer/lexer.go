package lexer

import (
	"strings"

	"github.com/roronya/sokki/token"
)

type Lexer struct {
	input    []rune
	position int
}

func New(input string) *Lexer {
	trimmed := strings.TrimSpace(input) // 行末の判定を簡単にするためにtrimしておく
	l := &Lexer{input: []rune(trimmed), position: 0}
	return l
}

func (l *Lexer) NextToken() token.Token {
	if l.position >= len(l.input) {
		return newToken(token.EOD, "")
	}
	if l.input[l.position] == '\n' {
		l.position++
		return newToken(token.NEWLINE, "\n")
	}

	// 行単位で処理すると楽なので、現在のポジションから改行までを取得する
	row := strings.SplitN(string(l.input[l.position:]), "\n", 2)[0]
	//fmt.Printf("row: %#v\n", row)
	//fmt.Printf("position: %#v\n", l.position)
	//fmt.Printf("%#v\n", string(l.input[l.position]))
	var tok token.Token
	switch {
	case l.position >= len(l.input):
		tok = newToken(token.EOD, "")
	case l.input[l.position] == '\n':
		tok = newToken(token.NEWLINE, "\n")
		l.position++
	case row == " >>":
		tok = newToken(token.MORESHIFT, " >>")
		l.position += 3
	case row == " >":
		tok = newToken(token.SHIFT, " >")
		l.position += 2
	case strings.HasSuffix(row, " >>"):
		tok = newToken(token.PARAGRAPH, row[:len(row)-3])
		l.position += len([]rune(row)) - 3
	case strings.HasSuffix(row, " >"):
		tok = newToken(token.PARAGRAPH, row[:len(row)-2])
		l.position += len([]rune(row)) - 2
	default:
		tok = newToken(token.PARAGRAPH, row)
		l.position += len([]rune(row))
	}

	return tok
}

func newToken(tokenType token.TokenType, literal string) token.Token {
	return token.Token{Type: tokenType, Literal: literal}
}
