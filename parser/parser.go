package parser

import (
	"github.com/roronya/sokki/lexer"
	"github.com/roronya/sokki/token"
)

type Parser struct {
	l *lexer.Lexer

	section int

	errors []string

	curToken  token.Token
	peekToken token.Token

	prefixParseFns map[token.TokenType]prefixParseFn
	infixParseFns  map[token.TokenType]infixParseFn
}

func New(l *lexer.Lexer) *Parser {
	return &Parase{
		l:      l,
		errors: []string(),
	}
}
