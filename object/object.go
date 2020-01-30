package object

import "fmt"

type ObjectType string

const (
	PARAGRAPH_OBJ = "PARAGRAPH"
)

type Object interface {
	Type() ObjectType
	Inspect() string
}

type Paragraph struct {
	Value string
}

func (pr *Paragraph) Type() string    { return PARAGRAPH_OBJ }
func (pr *Paragraph) Inspect() string { return fmt.Sprintf("<p>%s</p>", pr.Value) }
