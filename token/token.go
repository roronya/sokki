package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	EOD       = "EOD" // End Of Document
	NEWLINE   = "NEWLINE"
	STRING    = "STRING"
	SHIFT     = " >"
	MORESHIFT = " >>"
)
