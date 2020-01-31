package evaluator

import (
	"testing"

	"github.com/roronya/sokki/lexer"
	"github.com/roronya/sokki/parser"
)

func TestEval(t *testing.T) {
	input := `マリア様の庭に集う少女たちが、
今日も天使のような無垢な笑顔で、
背の高い門をくぐり抜けていく。

汚れを知らない心身を包むのは、深い色の制服。
スカートのプリーツは乱さないように、
白いセーラーカラーは翻さないように、
ゆっくりと歩くのが、ここでのたしなみ。

私立リリアン女学園。ここは乙女の園。
`
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
<p>マリア様の庭に集う少女たちが、</p>
<p>今日も天使のような無垢な笑顔で、</p>
<p>背の高い門をくぐり抜けていく。</p>
</div>
<div style="grid-column: 0; grid-row: 1;">
<p>汚れを知らない心身を包むのは、深い色の制服。</p>
<p>スカートのプリーツは乱さないように、</p>
<p>白いセーラーカラーは翻さないように、</p>
<p>ゆっくりと歩くのが、ここでのたしなみ。</p>
</div>
<div style="grid-column: 0; grid-row: 2;">
<p>私立リリアン女学園。ここは乙女の園。</p>
</div>
</body>
</html>` {
		t.Errorf("result is invalid. got=%s", result)
	}
}
