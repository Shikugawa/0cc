package parser

import (
	"github.com/gocc/ast"
	tokenizer "github.com/gocc/lexer"
)

type Parser struct {
	tokens  []tokenizer.Token
	pointer int
}

func (p *Parser) getToken() *tokenizer.Token {
	p.pointer--
	token := &p.tokens[p.pointer]
	return token
}

func (p *Parser) lookAhead(operator string) bool {
	if p.pointer-1 < 0 {
		return false
	}

	if operator == p.tokens[p.pointer-1].Value {
		p.pointer--
		return true
	} else {
		return false
	}
}

func (p *Parser) parseNumber() *ast.ASTNode {
	token := p.getToken()
	node := &ast.ASTNode{
		Kind:  ast.NUMBER,
		Value: token.Value,
	}
	return node
}

func (p *Parser) parseFactor() *ast.ASTNode {
	// Factor = NUMBER
	node := p.parseNumber()
	return node
}

func (p *Parser) parseTerm() *ast.ASTNode {
	// Term = Factor | Factor * Term | Factor / Term
	node := p.parseFactor()
	for {
		if p.lookAhead("*") {
			node = &ast.ASTNode{
				Kind:  ast.MUL,
				Left:  node,
				Right: p.parseTerm(),
			}
		} else if p.lookAhead("/") {
			node = &ast.ASTNode{
				Kind:  ast.DIV,
				Left:  node,
				Right: p.parseTerm(),
			}
		} else {
			return node
		}
	}
}

func (p *Parser) parseExpression() *ast.ASTNode {
	// Expr = Term | Term + Expr | Term - Expr
	node := p.parseTerm()
	for {
		if p.lookAhead("+") {
			node = &ast.ASTNode{
				Kind:  ast.ADD,
				Left:  node,
				Right: p.parseExpression(),
			}
		} else if p.lookAhead("-") {
			node = &ast.ASTNode{
				Kind:  ast.SUB,
				Left:  node,
				Right: p.parseExpression(),
			}
		} else {
			return node
		}
	}
}

func (p *Parser) Parse() *ast.ASTNode {
	node := p.parseExpression()
	return node
}

func Init(tokens []tokenizer.Token) *Parser {
	return &Parser{
		tokens:  tokens,
		pointer: len(tokens),
	}
}
