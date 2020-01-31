package evaluator

import (
	"fmt"

	"github.com/roronya/sokki/ast"
)

func Eval(node ast.Node) string {
	switch node := node.(type) {
	case *ast.Document:
		return evalDocument(node)
	case *ast.Paragraph:
		return fmt.Sprintf("<p>%s</p>", node.Value)
	}
	return ""
}

func evalDocument(node *ast.Document) string {
	html := `<html>
<head>
<style>
body {
display: grid;
grid-template-columns: 1fr 1fr 1fr;
}
</style>
<body>
`
	for _, s := range node.Left {
		html += Eval(&s)
	}
	for _, s := range node.Middle {
		html += Eval(&s)
	}
	for _, s := range node.Right {
		html += Eval(&s)
	}
	html += "</body></html>"
	return html
}
