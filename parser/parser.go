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
	dcmt := &ast.Document{
		Sections: []*ast.Section{},
	}
	id := 0
	for p.curToken.Type != token.EOD {
		s := p.parseSection(id)
		if s != nil {
			dcmt.Sections = append(dcmt.Sections, s)
		}
		p.nextToken()
		id++
	}
	return dcmt
}

func (p *Parser) parseSection(id int) *ast.Section {
	// TODO: skip NEWLINE
	s := &ast.Section{
		Id:     id,
		Left:   []*ast.Paragraph{},
		Middle: []*ast.Paragraph{},
		Right:  []*ast.Paragraph{},
	}
	for p.curToken.Type != token.EOD && p.curToken.Type != token.NEWLINE {
		// TODO: SHIFTやMORESHIFTの対応
		if p.curToken.Type == token.PARAGRAPH {
			pr := &ast.Paragraph{
				Token: p.curToken,
				Value: p.curToken.Literal,
			}
			s.Left = append(s.Left, pr)
			p.nextToken()

			if p.curToken.Type == token.NEWLINE {
				p.nextToken()
			}
		}
	}

	return s
}
