package parser

import (
	"github.com/roronya/sokki/ast"
	"github.com/roronya/sokki/lexer"
	"github.com/roronya/sokki/token"
)

type Parser struct {
	l *lexer.Lexer

	section int

	errors []string

	curToken  token.Token
	peekToken token.Token

	//	prefixParseFns map[token.TokenType]prefixParseFn
	//	infixParseFns  map[token.TokenType]infixParseFn
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:       l,
		errors:  []string{},
		section: 0,
	}

	// 2回nextTokenすることでcurTokenとpeekTokenがセットされる
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseDocument() *ast.Document {
	dcmt := &ast.Document{}
	dcmt.Left = []ast.Paragraph{}
	for p.curToken.Type != token.EOD {
		// TODO: SHIFTやMORESHIFTの対応
		pr := ast.Paragraph{
			Token:   p.curToken,
			Value:   p.curToken.Literal,
			Section: p.section,
		}
		dcmt.Left = append(dcmt.Left, pr)

		p.nextToken()
		p.section++
	}
	return dcmt
}
