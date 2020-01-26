package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	// 識別子 + リテラル
	PARAGRAPH = "PARAGRAPH"

	// 演算子
	SHIFT     = " >"
	MORESHIFT = " >>"

	// FIXME: どう扱うものかよくわからない
	NEWLINE = "\n"
)