package compiler

import (
	"encoding/binary"
	"fmt"
	"go-vmachine/pkgs/vm/lexer"
	"go-vmachine/pkgs/vm/opcode"
	"go-vmachine/pkgs/vm/token"
)

type Compiler struct {
	Lexer        *lexer.Lexer
	Bytecode     []byte
	Register     []byte
	CurrentToken token.Token
	PeekToken    token.Token
}

func ExpectedToken(token string) {
	panic(fmt.Sprintf("Expecting token: %s", token))
}

func NewCompiler(lex *lexer.Lexer) *Compiler {
	return &Compiler{
		Lexer:    lex,
		Bytecode: []byte{},
		Register: []byte{},
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
		switch c.CurrentToken.Type {
		// Todo: More arithmetic operations
		case token.ADD, token.SUB, token.MUL, token.DIV:
			c.MathOp(opcode.LookupOpCode(string(c.CurrentToken.Type)))
		case token.REG:
			c.SetRegister()
		case token.READ:
			c.IncludeRead()
			// ...
		case token.INT:
			// ...
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

func (c *Compiler) SetRegister() {
	if !c.Expect(token.INT) {
		ExpectedToken(token.INT)
	}

	c.NextToken()

	if !c.Expect(token.INT) {
		ExpectedToken(token.INT)
	}

	vReg := SafetyParseInt(c.CurrentToken.Literal)

	if vReg > 0xFF {
		panic("Register should be less 0xFF")
	}

	vDefaultValue := SafetyParseInt(c.PeekToken.Literal)

	c.NextToken()

	c.WriteBytes(opcode.REG)
	c.WriteBytes(byte(vReg))
	c.WriteNumber(vDefaultValue)
}

func (c *Compiler) IncludeRead() {
	c.NextToken()
}

func (c *Compiler) MathOp(operation byte) {
	// Example:
	/*
		add 0x10 0x10 # result: 0x20
	*/
	if !c.Expect(token.INT) {
		ExpectedToken(token.INT)
	}

	// Skip math operation token
	// add | math operation token
	c.NextToken()

	// 0x10 | int token
	if !c.Expect(token.INT) {
		ExpectedToken(token.INT)
	}

	v1 := SafetyParseInt(c.CurrentToken.Literal)
	v2 := SafetyParseInt(c.PeekToken.Literal)
	var v3 byte = opcode.NOP

	c.NextToken()

	if c.Expect(token.TO) {
		// Skip prev number and set pointer to token "TO"
		c.NextToken()

		if !c.Expect(token.INT) {
			ExpectedToken(token.INT)
		}

		unsafeReg := SafetyParseInt(c.PeekToken.Literal)
		if unsafeReg >= 0xFF {
			panic("Register should be less 0xFF")
		}

		v3 = byte(unsafeReg)
	}

	c.WriteBytes(operation)
	c.WriteNumber(v1)
	c.WriteNumber(v2)
	c.WriteBytes(v3)
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
