package lexer

import (
	"regexp"

	"github.com/roronya/sokki/token"
)

var SHIFT_REGEXP = regexp.MustCompile(`^ >`)
var MORESHIFT_REGEXP = regexp.MustCompile(`^ >>`)
var PARAGRAPH_REGEXP = regexp.MustCompile(`^([^\n( >)( >>)]+)`)

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
		size := len([]rune(pr[1])) // グルーピングした箇所が知りたいので1でアクセスする
		l.position += size
		return newToken(token.PARAGRAPH, l.input[l.position-size:l.position])
	}
	// SHIFTがマッチしてしまうので、先にMORESHIFTにマッチしてないか調べる
	msh := MORESHIFT_REGEXP.FindStringSubmatch(s)
	if msh != nil {
		size := len([]rune(msh[0]))
		l.position += size
		return newToken(token.MORESHIFT, l.input[l.position-size:l.position])
	}
	sh := SHIFT_REGEXP.FindStringSubmatch(s)
	if sh != nil {
		size := len([]rune(sh[0])) // グルーピングはしていないので0でアクセスする
		l.position += size
		return newToken(token.SHIFT, l.input[l.position-size:l.position])
	}
	return tok
}

func newToken(tokenType token.TokenType, literal []rune) token.Token {
	return token.Token{Type: tokenType, Literal: string(literal)}
}
