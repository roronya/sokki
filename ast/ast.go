package ast

import "github.com/roronya/sokki/token"

type Node interface {
	TokenLiteral() string
	String() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Document struct {
	Left   []Paragraph
	Middle []Paragraph
	right  []Paragraph
}

type Paragraph struct {
	Token   token.Token
	Value   string
	Section int
}

func (p *Paragraph) expressionNode() {}
func (p *Paragraph) TokenLiteral() {
	return p.Token.Literal
}
func (p *Paragraph) String() {
	return p.Value
}
