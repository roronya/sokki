package parser

import (
	"testing"

	"github.com/roronya/sokki/lexer"
	"github.com/roronya/sokki/token"
)

func TestLeftParagraph(t *testing.T) {
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
	p := New(l)

	dcmt := p.ParseDocument()
	if len(dcmt.Sections) != 3 {
		t.Fatalf("document.Sections does not contain 1 section. got=%d",
			len(dcmt.Sections))
	}

	s := dcmt.Sections[0]
	if len(s.Left) != 3 {
		t.Fatalf("s.Left does not contain 3 section. got=%d",
			len(s.Left))
	}
	if len(s.Middle) != 0 {
		t.Fatalf("s.Middle does not contain 0 section. got=%d",
			len(s.Middle))
	}
	if len(s.Right) != 0 {
		t.Fatalf("s.Right does not contain 0 section. got=%d",
			len(s.Right))
	}

	expected := []string{
		"マリア様の庭に集う少女たちが、",
		"今日も天使のような無垢な笑顔で、",
		"背の高い門をくぐり抜けていく。",
	}
	for i, e := range expected {
		pr := s.Left[i]
		if pr.Token.Type != token.PARAGRAPH {
			t.Errorf("pr is not PARAGRAPH. got=%T", pr)
		}

		val := pr.Value
		if val != e {
			t.Errorf("pr.Value not %s. got=%s", e, pr.Value)
		}
	}
}
