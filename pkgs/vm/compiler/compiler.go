package compiler

import (
	"go-vmachine/pkgs/vm/lexer"
	"go-vmachine/pkgs/vm/token"
)

type Compiler struct {
	Lexer        *lexer.Lexer
	Bytecode     []byte
	CurrentToken token.Token
	PeekToken    token.Token
}

func NewCompiler(lex *lexer.Lexer) *Compiler {
	return &Compiler{
		Lexer:    lex,
		Bytecode: []byte{},
	}
}

func (c *Compiler) Compile() {
	c.NextToken()
	c.NextToken()

	for c.CurrentToken.Type != token.EOF {
		switch c.CurrentToken.Type {
		case token.ADD:
		case token.SUB:

			// todo
		}
		c.NextToken()
	}
}

func (c *Compiler) NextToken() {
	c.CurrentToken = c.PeekToken
	c.PeekToken = c.Lexer.NextToken()
}

func (c *Compiler) MathOperation(operation int) {

}

// Bytecode helpers
func (c *Compiler) WriteByte(b byte) {
	c.Bytecode = append(c.Bytecode, b)
}
