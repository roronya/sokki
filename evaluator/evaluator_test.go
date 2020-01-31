package evaluator

import (
	"testing"

	"github.com/roronya/sokki/lexer"
	"github.com/roronya/sokki/parser"
)

func TestEval(t *testing.T) {
	input := `マリア様の庭に集う乙女達が、
今日も天使のような無垢な笑顔で、
背の高い門をくぐり抜けていく。`
	l := lexer.New(input)
	p := parser.New(l)
	ast := p.ParseDocument()

	result := Eval(ast)
	if result != `<html>
<head>
<style>
body {
display: grid;
grid-template-columns: 1fr 1fr 1fr;
}
</style>
<body>
<div style="grid-column: 0; grid-row: 0;">
<p>マリア様の庭に集う乙女達が、</p>
<p>今日も天使のような無垢な笑顔で、</p>
<p>背の高い門をくぐり抜けていく。</p>
</div>
<div style="grid-column: 1; grid-row: 0;">
</div>
<div style="grid-column: 2; grid-row: 0;">
</div>
</body>
</html>` {
		t.Errorf("result is invalid. got=%s", result)
	}
}
