package evaluator

import (
	"fmt"

	"github.com/roronya/sokki/ast"
)

const TEMPLATE = `<html>
<head>
<style>
body {
display: grid;
grid-template-columns: 50%% 1fr 1fr;
}
p {
margin: 0;
}
div {
padding: 16px;
}
.left {
grid-column: 1;
}
.middle {
grid-column: 2;
background-color: #ECECEC;
}
.right {
grid-column: 3;
}
</style>
%s
</html>`

func Eval(node ast.Node) string {
	switch node := node.(type) {
	case *ast.Document:
		body := evalDocument(node)
		return fmt.Sprintf(TEMPLATE, body)
	case *ast.Paragraph:
		return fmt.Sprintf("<p>%s</p>\n", node.Value)
	}
	return ""
}

func evalDocument(node *ast.Document) string {
	// HTMLの上から順番に選択されるので、各カラムを混ぜずに固めておく
	left := ""
	middle := ""
	right := ""
	for _, s := range node.Sections {
		// gridは0でなく1から指定するから、0始まりのs.Idを+1する
		left += fmt.Sprintf("<div class=\"left\" style=\"grid-row: %d;\">\n", s.Id+1)
		for _, pr := range s.Left {
			left += Eval(pr)
		}
		left += "</div>\n"

		middle += fmt.Sprintf("<div class=\"middle\" style=\"grid-row: %d;\">\n", s.Id+1)
		for _, pr := range s.Middle {
			middle += Eval(pr)
		}
		middle += "</div>\n"

		right += fmt.Sprintf("<div class=\"right\" style=\"grid-row: %d;\">\n", s.Id+1)
		for _, pr := range s.Right {
			right += Eval(pr)
		}
		right += "</div>\n"
	}
	body := "<body>\n"
	body += left
	body += middle
	body += right
	body += "</body>\n"
	return body
}
