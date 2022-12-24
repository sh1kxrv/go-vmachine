package lexer

func isWhitespace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r'
}

func isEmpty(ch rune) bool {
	return ch == 0
}
