package lexer

import (
	"fmt"
)

// Token Types
type TokenType int

const (
	INTVALUE   TokenType = iota
	OPERATOR             = iota
	OP_BRACKET           = iota
	CL_BRACKET           = iota
)

type Token struct {
	Type  TokenType
	Value string
}

type Tokenizer struct {
	code        []byte
	currentByte byte
	pointer     int
}

func (tokenizer *Tokenizer) forward() {
	if tokenizer.pointer == len(tokenizer.code)-1 {
		tokenizer.currentByte = 0x0
		return
	}

	tokenizer.pointer++
	tokenizer.currentByte = tokenizer.code[tokenizer.pointer]
}

func (tokenizer *Tokenizer) backword() {
	if tokenizer.pointer == 0 {
		tokenizer.currentByte = 0x0
		return
	}

	tokenizer.pointer--
	tokenizer.currentByte = tokenizer.code[tokenizer.pointer]
}

func (tokenizer *Tokenizer) getNumber() []byte {
	number := []byte{}
	for {
		if tokenizer.currentByte >= '0' && tokenizer.currentByte <= '9' {
			number = append(number, tokenizer.currentByte)
			tokenizer.forward()
		} else {
			break
		}
	}
	return number
}

func (tokenizer *Tokenizer) getSymbol() []byte {
	symbol := []byte{}
	symbol = append(symbol, tokenizer.currentByte)
	tokenizer.forward()
	return symbol
}

func (tokenizer *Tokenizer) Tokenize() []Token {
	tokenList := []Token{}
	tokenizer.forward()
	for {
		if tokenizer.currentByte == 0x0 {
			break
		}

		switch tokenizer.currentByte {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			num := tokenizer.getNumber()
			token := Token{
				Type:  INTVALUE,
				Value: fmt.Sprintf("%s", num),
			}
			tokenList = append(tokenList, token)
		case '+', '-', '*', '/':
			symbol := tokenizer.getSymbol()
			token := Token{
				Type:  OPERATOR,
				Value: fmt.Sprintf("%s", symbol),
			}
			tokenList = append(tokenList, token)
		case '(':
			symbol := tokenizer.getSymbol()
			token := Token{
				Type:  OP_BRACKET,
				Value: fmt.Sprintf("%s", symbol),
			}
			tokenList = append(tokenList, token)
		case ')':
			symbol := tokenizer.getSymbol()
			token := Token{
				Type:  CL_BRACKET,
				Value: fmt.Sprintf("%s", symbol),
			}
			tokenList = append(tokenList, token)
		case ' ', '\n':
			tokenizer.forward()
			continue
		default:
			panic(fmt.Sprintf("character '%c' is invalid syntax", tokenizer.currentByte))
		}
	}
	return tokenList
}

func Init(code []byte) *Tokenizer {
	tokenzer := &Tokenizer{
		code:    code,
		pointer: -1,
	}
	return tokenzer
}
