package evaluator

import (
	"testing"

	"github.com/roronya/sokki/lexer"
	"github.com/roronya/sokki/parser"
)

func TestEval(t *testing.T) {
	input := "私立リリアン女学園。ここは乙女の園。"
	l := lexer.New(input)
	p := parser.New(l)
	ast := p.ParseDocument()

	result := Eval(ast)
	if result != `<html>
<body>
<section>
<p>私立リリアン女学園。ここは乙女の園。</p>
</section>
<section>

</section>
<section>

</section>
</body>
</html>
` {
		t.Errorf("result is invalid. got=%s", result)
	}
}
