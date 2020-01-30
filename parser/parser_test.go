package parser

import (
	"testing"

	"github.com/roronya/sokki/lexer"
	"github.com/roronya/sokki/token"
)

func TestLeftParagraph(t *testing.T) {
	input := `私立リリアン女学園。ここは乙女の園`

	l := lexer.New(input)
	p := New(l)

	dcmt := p.ParseDocument()
	if len(dcmt.Left) != 1 {
		t.Fatalf("document.left does not contain 1 paragraph. got=%d",
			len(dcmt.Left))
	}

	pr := dcmt.Left[0]
	if pr.Token.Type != token.PARAGRAPH {
		t.Errorf("pr not *ast.Paragraph. got=%T", dcmt.Left[0])
	}

	val := pr.Value
	if val != "私立リリアン女学園。ここは乙女の園" {
		t.Errorf("pr.Value not %s. got=%s", "私立リリアン女学園。ここは乙女の園", pr.Value)
	}

}
