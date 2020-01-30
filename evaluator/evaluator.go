package evaluator

import (
	"fmt"
	"strings"

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
	left := []string{}
	middle := []string{}
	right := []string{}
	for _, pr := range node.Left {
		left = append(left, Eval(&pr))
	}
	for _, pr := range node.Middle {
		middle = append(middle, Eval(&pr))
	}
	for _, pr := range node.Right {
		right = append(right, Eval(&pr))
	}
	leftJoined := strings.Join(left, "")
	middleJoined := strings.Join(middle, "")
	rightJoined := strings.Join(right, "")
	return fmt.Sprintf(`<html>
<body>
<section>
%s
</section>
<section>
%s
</section>
<section>
%s
</section>
</body>
</html>
`, leftJoined, middleJoined, rightJoined)
}
