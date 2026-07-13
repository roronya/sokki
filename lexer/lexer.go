package lexer

import (
	"strings"

	"github.com/roronya/sokki/token"
)

type Lexer struct {
	input    []rune
	position int
}

func New(input string) *Lexer {
	normalized := strings.ReplaceAll(input, "\r\n", "\n") // Windowsの改行にも対応する
	trimmed := strings.TrimSpace(normalized)              // 行末の判定を簡単にするためにtrimしておく
	l := &Lexer{input: []rune(trimmed), position: 0}
	return l
}

func (l *Lexer) NextToken() token.Token {
	if l.position >= len(l.input) {
		return newToken(token.EOD, "")
	}
	if l.input[l.position] == '\n' {
		l.position++
		return newToken(token.NEWLINE, "\n")
	}

	// 行単位で処理すると楽なので、現在のポジションから改行までを取得する
	// 改行が無い場合は行末までを取得する
	row := l.input[l.position:]
	if i := indexRune(row, '\n'); i >= 0 {
		row = row[:i]
	}

	markerLen, tokType := detectMarker(row)
	switch {
	case markerLen > 0 && markerLen == len(row):
		// 行全体がシフト記号
		l.position += markerLen
		return newToken(tokType, string(row))
	case markerLen > 0:
		// シフト記号の手前までを返す。記号は次のNextTokenでtokenizeされる
		l.position += len(row) - markerLen
		return newToken(token.STRING, unescape(row[:len(row)-markerLen]))
	default:
		l.position += len(row)
		return newToken(token.STRING, unescape(row))
	}
}

// detectMarker は行末のシフト記号を検出し、その長さ(rune数)とトークン種別を返す。
// シフト記号はスペース(半角・全角)に続く1つか2つの '>'(半角・全角)。
// '>' が3つ以上並ぶ場合は記号とみなさず、行全体を本文として扱う。
func detectMarker(row []rune) (int, token.TokenType) {
	k := countTrailingMarkerChars(row)
	if k < 1 || k > 2 {
		return 0, ""
	}
	i := len(row) - k - 1 // '>' の並びの直前
	if i < 0 || !isMarkerSpace(row[i]) {
		return 0, ""
	}
	if k == 2 {
		return 3, token.MORESHIFT
	}
	return 2, token.SHIFT
}

// unescape は行末の「 \>」「 \>>」をエスケープとして解釈し、バックスラッシュを取り除く。
// シフト記号と同じ並びを本文としてそのまま書きたいときに使う。
func unescape(row []rune) string {
	k := countTrailingMarkerChars(row)
	if k < 1 || k > 2 {
		return string(row)
	}
	i := len(row) - k - 1 // '>' の並びの直前
	if i >= 1 && row[i] == '\\' && isMarkerSpace(row[i-1]) {
		return string(row[:i]) + string(row[i+1:])
	}
	return string(row)
}

func countTrailingMarkerChars(row []rune) int {
	n := 0
	for i := len(row) - 1; i >= 0 && isMarkerChar(row[i]); i-- {
		n++
	}
	return n
}

func isMarkerChar(r rune) bool  { return r == '>' || r == '＞' }
func isMarkerSpace(r rune) bool { return r == ' ' || r == '　' }

func indexRune(rs []rune, r rune) int {
	for i, v := range rs {
		if v == r {
			return i
		}
	}
	return -1
}

func newToken(tokenType token.TokenType, literal string) token.Token {
	return token.Token{Type: tokenType, Literal: literal}
}
