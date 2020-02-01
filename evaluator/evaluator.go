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
	left := ""
	middle := ""
	right := ""
	for _, s := range node.Sections {
		if len(s.Left) > 0 {
			// gridは0でなく1から指定するから、0始まりのs.Idを+1する
			left += fmt.Sprintf("<div style=\"grid-column: %d; grid-row: %d;\">\n", 1, s.Id+1)
			for _, pr := range s.Left {
				left += Eval(pr)
			}
			left += "</div>\n"
		}

		if len(s.Middle) > 0 {
			middle += fmt.Sprintf("<div style=\"grid-column: %d; grid-row: %d;\">\n", 2, s.Id+1)
			for _, pr := range s.Middle {
				middle += Eval(pr)
			}
			middle += "</div>\n"
		}

		if len(s.Right) > 0 {
			right += fmt.Sprintf("<div style=\"grid-column: %d; grid-row: %d;\">\n", 3, s.Id+1)
			for _, pr := range s.Right {
				right += Eval(pr)
			}
			right += "</div>\n"
		}
	}
	html += left
	html += middle
	html += right
	html += "</body>\n</html>"
	return html
}
