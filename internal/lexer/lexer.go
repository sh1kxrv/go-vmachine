package lexer

import "go-vmachine/internal/token"

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

func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	l.SkipWhitespace()

	switch l.CurrentChar() {
	case 0:
		tok = token.Token{Type: token.EOF, Literal: ""}
	case ',':
		tok = token.NewToken(token.COMMA, string(l.CurrentChar()))
	case '#':
		if !isDigit(l.PeekChar()) {
			l.SkipComment()
			return l.NextToken()
		}
	case '"':
		tok = token.NewToken(token.STRING, l.ReadString())
	default:
		if isDigit(l.CurrentChar()) {
			return l.ReadDecimal()
		}

		tok.Literal = l.ReadIdentifier()
		tok.Type = token.LookupIdent(tok.Literal)
		return tok
	}

	l.NextChar()
	return tok
}

func (l *Lexer) NextChar() rune {
	if l.Pos >= len(l.Characters) {
		return rune(0)
	}
	l.Pos++
	return l.Characters[l.Pos-1]
}

func (l *Lexer) PeekChar() rune {
	if l.Pos+1 >= len(l.Characters) {
		return 0
	}
	return l.Characters[l.Pos+1]
}

func (l *Lexer) CurrentChar() rune {
	if l.Pos >= len(l.Characters) {
		return rune(0)
	}
	return l.Characters[l.Pos]
}

func (l *Lexer) SkipWhitespace() {
	for isWhitespace(l.CurrentChar()) {
		l.NextChar()
	}
}

func (l *Lexer) SkipComment() {
	for l.CurrentChar() != '\n' && l.CurrentChar() != 0 {
		l.NextChar()
	}
	l.SkipWhitespace()
}

func (l *Lexer) ReadNumber() string {
	start := l.Pos
	for isHexDigit(l.CurrentChar()) {
		l.NextChar()
	}
	return string(l.Characters[start:l.Pos])
}

func (l *Lexer) ReadDecimal() token.Token {
	number := l.ReadNumber()
	currentChar := l.CurrentChar()
	if isEmpty(currentChar) || isWhitespace(currentChar) || currentChar == ',' {
		return token.Token{Type: token.INT, Literal: number}
	}
	return token.Token{Type: token.ILLEGAL, Literal: number}
}

func (l *Lexer) ReadIdentifier() string {
	start := l.Pos
	for !isEmpty(l.CurrentChar()) && !isWhitespace(l.CurrentChar()) {
		l.NextChar()
	}
	return string(l.Characters[start:l.Pos])
}

func (l *Lexer) ReadString() string {
	start := l.Pos + 1
	for {
		l.NextChar()
		if l.CurrentChar() == '\\' {
			l.NextChar()
			if l.CurrentChar() == '"' {
				continue
			}
		}

		if l.CurrentChar() == '"' {
			break
		}
	}
	return string(l.Characters[start:l.Pos])
}
