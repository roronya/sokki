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
		return fmt.Sprintf("<p>%s</p>\n", node.Value)
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
	// HTMLの上から順番に選択されるので、各カラムを混ぜずに固めておく
	for _, s := range node.Sections {
		html += fmt.Sprintf("<div style=\"grid-column: %d; grid-row: %d;\">\n", 0, s.Id)
		for _, pr := range s.Left {
			html += Eval(pr)
		}
		html += "</div>\n"
	}

	for _, s := range node.Sections {
		html += fmt.Sprintf("<div style=\"grid-column: %d; grid-row: %d;\">\n", 1, s.Id)
		for _, pr := range s.Middle {
			html += Eval(pr)
		}
		html += "</div>\n"
	}

	for _, s := range node.Sections {
		html += fmt.Sprintf("<div style=\"grid-column: %d; grid-row: %d;\">\n", 2, s.Id)
		for _, pr := range s.Right {
			html += Eval(pr)
		}
		html += "</div>\n"
	}
	html += "</body>\n</html>"
	return html
}
