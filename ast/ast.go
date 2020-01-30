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
	Right  []Paragraph
}

func (d *Document) expressionNode() {}
func (d *Document) TokenLiteral() string {
	// TODO
	return ""
}
func (d *Document) String() string {
	// TODO
	return ""
}

type Paragraph struct {
	Token   token.Token
	Value   string
	Section int
}

func (pr *Paragraph) expressionNode() {}
func (pr *Paragraph) TokenLiteral() string {
	return pr.Token.Literal
}
func (pr *Paragraph) String() string {
	return pr.Value
}
