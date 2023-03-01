package compiler

import (
	"encoding/binary"
	"fmt"
	"go-vmachine/pkgs/vm/lexer"
	"go-vmachine/pkgs/vm/token"
	"strconv"
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

func (c *Compiler) Dump() {
	for _, b := range c.Bytecode {
		println(fmt.Sprintf("0x%02x", b))
	}
}

func (c *Compiler) Compile() []byte {
	c.NextToken()
	c.NextToken()

	for c.CurrentToken.Type != token.EOF {
		// println(fmt.Sprintf("CurrentToken: %s, %s", c.CurrentToken.Type, c.PeekToken.Type))
		switch c.CurrentToken.Type {
		// Todo: More arithmetic operations
		case token.ADD, token.SUB, token.MUL, token.DIV:
			c.MathOperation(TokenToOpcode(c.CurrentToken))
		case token.INT:
			// Todo: Handle int token
		default:
			panic("Unknown token")
		}
		c.NextToken()
	}
	return c.Bytecode
}

func (c *Compiler) NextToken() {
	c.CurrentToken = c.PeekToken
	c.PeekToken = c.Lexer.NextToken()
}

func (c *Compiler) MathOperation(operation byte) {
	// Example:
	/*
		add 0x10 0x10 # result: 0x20
	*/
	if !c.Expect(token.INT) {
		panic("Expecting int token")
	}

	// Skip math operation token
	// add | math operation token
	c.NextToken()

	// 0x10 | int token
	if !c.Expect(token.INT) {
		panic("Expecting int token")
	}

	v1, err := strconv.ParseInt(c.CurrentToken.Literal, 0, 64)
	if err != nil {
		panic(err)
	}
	v2, err := strconv.ParseInt(c.PeekToken.Literal, 0, 64)
	if err != nil {
		panic(err)
	}

	c.NextToken()

	c.WriteBytes(operation)
	c.WriteNumber(v1)
	c.WriteNumber(v2)
}

// Helpers
func (c *Compiler) WriteBytes(b ...byte) {
	c.Bytecode = append(c.Bytecode, b...)
}

func (c *Compiler) WriteNumber(number int64) {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, uint64(number))
	c.WriteBytes(byte(len(b)))
	c.WriteBytes(b...)
}

func (c *Compiler) Expect(expected token.TokenType) bool {
	return c.PeekToken.Type == expected
}
