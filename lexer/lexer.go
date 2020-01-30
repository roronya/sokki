package lexer

import (
	"regexp"

	"github.com/roronya/sokki/token"
)

var SHIFT_REGEXP = regexp.MustCompile(`^ >$`)
var MORESHIFT_REGEXP = regexp.MustCompile(`^ >>$`)
var PARAGRAPH_REGEXP = regexp.MustCompile(`^([^\n]*)`) //TODO SHIFTの対応

type Lexer struct {
	input    []rune
	position int
}

func New(input string) *Lexer {
	l := &Lexer{input: []rune(input), position: 0}
	return l
}

// 各トークンに一致するか正規表現で調べる
// positionは一致した文字列のruneのサイズだけ進める
// newTokenでは進めたサイズだけ減らしてスライスする
func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	if l.position >= len(l.input) {
		return newToken(token.EOD, []rune(""))
	}
	if l.input[l.position] == '\n' {
		l.position++
		return newToken(token.NEWLINE, l.input[l.position-1:l.position])
	}
	s := string(l.input[l.position:])
	pr := PARAGRAPH_REGEXP.FindStringSubmatch(s)
	if pr != nil {
		size := len([]rune(pr[1]))
		l.position += size
		return newToken(token.PARAGRAPH, l.input[l.position-size:l.position])
	}
	// TODO: SHIFTの抜き出し
	return tok
}

func newToken(tokenType token.TokenType, literal []rune) token.Token {
	return token.Token{Type: tokenType, Literal: string(literal)}
}
