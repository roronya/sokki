package object

import "fmt"

type ObjectType string

const (
	STRING_OBJ = "STRING"
)

type Object interface {
	Type() ObjectType
	Inspect() string
}

type Paragraph struct {
	Value string
}

func (pr *Paragraph) Type() string    { return STRING_OBJ }
func (pr *Paragraph) Inspect() string { return fmt.Sprintf("<p>%s</p>", pr.Value) }
