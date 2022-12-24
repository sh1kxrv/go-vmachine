package lexer

import "go/token"

type Lexer struct {
	Raw        string
	Pos        int
	Characters []rune
}

func NewLexer(raw string) *Lexer {
	return &Lexer{
		Raw:        raw,
		Pos:        0,
		Characters: []rune(raw),
	}
}

func (l *Lexer) NextToken() {
	var token token.Token
}

func (l *Lexer) NextChar() rune {
	if l.Pos >= len(l.Characters) {
		return 0
	}
	l.Pos++
	return l.Characters[l.Pos-1]
}

func (l *Lexer) CurrentChar() rune {
	return l.Characters[l.Pos]
}

func (l *Lexer) SkipWhitespace() {
	for isWhitespace(l.CurrentChar()) {
		l.NextChar()
	}
}
