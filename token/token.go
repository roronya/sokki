package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	EOD     = "EOD" // End Of Document
	NEWLINE = "NEWLINE"

	// 識別子 + リテラル
	PARAGRAPH = "PARAGRAPH"

	// 演算子
	SHIFT     = " >"
	MORESHIFT = " >>"
)
