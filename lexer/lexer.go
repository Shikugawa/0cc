package lexer

import (
	"errors"
	"fmt"
)

// Token Types
type TokenType int

const (
	INTVALUE TokenType = iota
	OPERATOR           = iota
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

func (tokenizer *Tokenizer) forward() error {
	if tokenizer.pointer >= len(tokenizer.code) {
		return errors.New("can't forward more")
	}
	tokenizer.pointer++
	tokenizer.currentByte = tokenizer.code[tokenizer.pointer]
	return nil
}

func (tokenizer *Tokenizer) backword() error {
	if tokenizer.pointer <= -1 {
		return errors.New("can't forward more")
	}
	tokenizer.pointer--
	tokenizer.currentByte = tokenizer.code[tokenizer.pointer]
	return nil
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
	return symbol
}

func (tokenizer *Tokenizer) Tokenize() []Token {
	tokenList := []Token{}
	for {
		if tokenizer.pointer >= len(tokenizer.code)-1 {
			break
		}
		tokenizer.forward()
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
		case ' ':
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
