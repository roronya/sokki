package evaluator

import (
	"testing"

	"github.com/roronya/sokki/lexer"
	"github.com/roronya/sokki/parser"
)

func TestEvalDocument(t *testing.T) {
	input := `マリア様の庭に集う少女たちが、
今日も天使のような無垢な笑顔で、 >
背の高い門をくぐり抜けていく。 >>

汚れを知らない心身を包むのは、深い色の制服。
スカートのプリーツは乱さないように、 >>
白いセーラーカラーは翻さないように、 >
ゆっくりと歩くのが、ここでのたしなみ。

私立リリアン女学園。ここは乙女の園。
`
	l := lexer.New(input)
	p := parser.New(l)
	ast := p.ParseDocument()

	result := evalDocument(ast)
	if result != `<body>
<div class="left" style="grid-row: 1;">
<p>マリア様の庭に集う少女たちが、</p>
</div>
<div class="left" style="grid-row: 2;">
<p>汚れを知らない心身を包むのは、深い色の制服。</p>
<p>ゆっくりと歩くのが、ここでのたしなみ。</p>
</div>
<div class="left" style="grid-row: 3;">
<p>私立リリアン女学園。ここは乙女の園。</p>
</div>
<div class="middle" style="grid-row: 1;">
<p>今日も天使のような無垢な笑顔で、</p>
</div>
<div class="middle" style="grid-row: 2;">
<p>白いセーラーカラーは翻さないように、</p>
</div>
<div class="middle" style="grid-row: 3;">
</div>
<div class="right" style="grid-row: 1;">
<p>背の高い門をくぐり抜けていく。</p>
</div>
<div class="right" style="grid-row: 2;">
<p>スカートのプリーツは乱さないように、</p>
</div>
<div class="right" style="grid-row: 3;">
</div>
</body>
` {
		t.Errorf("result is invalid. got=%s", result)
	}
}
