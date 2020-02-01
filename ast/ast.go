package ast

import (
	"bytes"

	"github.com/roronya/sokki/token"
)

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
	Sections []*Section
}

func (d *Document) expressionNode() {}
func (d *Document) TokenLiteral() string {
	if len(d.Sections) > 0 {
		return d.Sections[0].TokenLiteral()
	}
	return ""
}
func (d *Document) String() string {
	var out bytes.Buffer

	for _, s := range d.Sections {
		out.WriteString(s.String())
	}

	return out.String()
}

type Section struct {
	Id     int
	Left   []*Paragraph
	Middle []*Paragraph
	Right  []*Paragraph
}

func (s *Section) expressionNode() {}
func (s *Section) TokenLiteral() string {
	if len(s.Left) > 0 {
		return s.Left[0].TokenLiteral()
	}
	return ""
}
func (s *Section) String() string {
	var out bytes.Buffer
	for _, p := range s.Left {
		out.WriteString(p.String())
	}

	return out.String()
}

type Paragraph struct {
	Token token.Token
	Value string
}

func (pr *Paragraph) expressionNode() {}
func (pr *Paragraph) TokenLiteral() string {
	return pr.Token.Literal
}
func (pr *Paragraph) String() string {
	return pr.Value
}

type SuffixExpression struct {
	Token    token.Token
	Operator string
	Left     Expression
}

func (se *SuffixExpression) expressionNode() {}
func (se *SuffixExpression) TokenLiteral() string {
	return se.Token.Literal
}
func (se *SuffixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(se.Left.String())
	out.WriteString(se.Operator)
	out.WriteString(")")

	return out.String()
}
