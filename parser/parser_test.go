package parser

import (
	"testing"

	"github.com/roronya/sokki/ast"
	"github.com/roronya/sokki/lexer"
	"github.com/roronya/sokki/token"
)

func TestParagraph(t *testing.T) {
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
	p := New(l)

	dcmt := p.ParseDocument()
	if len(dcmt.Sections) != 3 {
		t.Fatalf("document.Sections does not contain 1 section. got=%d",
			len(dcmt.Sections))
	}

	s := dcmt.Sections[0]
	if len(s.Left) != 1 {
		t.Fatalf("s.Left does not contain 1 section. got=%d",
			len(s.Left))
	}
	if len(s.Middle) != 1 {
		t.Fatalf("s.Middle does not contain 1 section. got=%d",
			len(s.Middle))
	}
	if len(s.Right) != 1 {
		t.Fatalf("s.Right does not contain 1 section. got=%d",
			len(s.Right))
	}

	pr := s.Left[0]
	if !testParagraph(t, pr, "マリア様の庭に集う少女たちが、") {
		return
	}
	pr = s.Middle[0]
	if !testParagraph(t, pr, "今日も天使のような無垢な笑顔で、") {
		return
	}
	pr = s.Right[0]
	if !testParagraph(t, pr, "背の高い門をくぐり抜けていく。") {
		return
	}

	s = dcmt.Sections[1]
	if len(s.Left) != 2 {
		t.Fatalf("s.Left does not contain 1 section. got=%d",
			len(s.Left))
	}
	if len(s.Middle) != 1 {
		t.Fatalf("s.Middle does not contain 1 section. got=%d",
			len(s.Middle))
	}
	if len(s.Right) != 1 {
		t.Fatalf("s.Right does not contain 1 section. got=%d",
			len(s.Right))
	}
	pr = s.Left[0]
	if !testParagraph(t, pr, "汚れを知らない心身を包むのは、深い色の制服。") {
		return
	}
	pr = s.Right[0]
	if !testParagraph(t, pr, "スカートのプリーツは乱さないように、") {
		return
	}
	pr = s.Middle[0]
	if !testParagraph(t, pr, "白いセーラーカラーは翻さないように、") {
		return
	}
	pr = s.Left[1]
	if !testParagraph(t, pr, "ゆっくりと歩くのが、ここでのたしなみ。") {
		return
	}

	s = dcmt.Sections[2]
	if len(s.Left) != 1 {
		t.Fatalf("s.Left does not contain 1 section. got=%d",
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
	pr = s.Left[0]
	if !testParagraph(t, pr, "私立リリアン女学園。ここは乙女の園。") {
		return
	}
}

func TestSkipUntilParagraph(t *testing.T) {
	input := `

aaa

 >>

 >

bbb
`

	l := lexer.New(input)
	p := New(l)
	dcmt := p.ParseDocument()
	if len(dcmt.Sections) != 2 {
		t.Errorf("dcmt.Section does not contain 1 section. got=%d",
			len(dcmt.Sections))
	}

	pr := dcmt.Sections[0].Left[0]
	if !testParagraph(t, pr, "aaa") {
		return
	}

	pr = dcmt.Sections[1].Left[0]
	if !testParagraph(t, pr, "bbb") {
		return
	}
}

func TestStartWithoutParagraph(t *testing.T) {
	input := ` >>aaa
 bbb

 >ccc`

	l := lexer.New(input)
	p := New(l)
	dcmt := p.ParseDocument()
	if len(dcmt.Sections) != 2 {
		t.Errorf("dcmt.Section does not contain 2 section. got=%d",
			len(dcmt.Sections))
	}

	s := dcmt.Sections[0]
	if len(s.Left) != 2 {
		t.Fatalf("s.Left does not contain 2 section. got=%d",
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

	s = dcmt.Sections[1]
	if len(s.Left) != 1 {
		t.Fatalf("s.Left does not contain 1 section. got=%d",
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

}

func testParagraph(t *testing.T, pr *ast.Paragraph, e string) bool {
	if pr.Token.Type != token.PARAGRAPH {
		t.Errorf("pr is not PARAGRAPH. got=%T", pr)
		return false
	}

	val := pr.Value
	if val != e {
		t.Errorf("pr.Value not %s. got=%s", e, pr.Value)
		return false
	}

	return true
}
