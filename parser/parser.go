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

func (p *Parser) parseCastExpr() *ast.ASTNode {
	token := p.getToken()
	node := &ast.ASTNode{
		Kind:  ast.UNARY,
		Value: token.Value,
	}
	return node
}

func (p *Parser) parseUnaryExpr() *ast.ASTNode {
	// UnaryExpr ::= (+|-)CastExpr | CastExpr
	node := p.parseCastExpr()
	if p.lookAhead("-") {
		node = &ast.ASTNode{
			Kind:  ast.UNARY,
			Value: "-" + node.Value,
		}
	}
	return node
}

func (p *Parser) parseMulExpr() *ast.ASTNode {
	// Mul ::= UnaryExpr | Mul * UnaryExpr | Mul / UnaryExpr
	node := p.parseUnaryExpr()
	for {
		if p.lookAhead("*") {
			node = &ast.ASTNode{
				Kind:  ast.MUL,
				Left:  node,
				Right: p.parseMulExpr(),
			}
		} else if p.lookAhead("/") {
			node = &ast.ASTNode{
				Kind:  ast.DIV,
				Left:  node,
				Right: p.parseMulExpr(),
			}
		} else {
			return node
		}
	}
}

func (p *Parser) parseAddExpression() *ast.ASTNode {
	// Expr ::= Mul | Expr + Mul | Expr - Mul
	node := p.parseMulExpr()
	for {
		if p.lookAhead("+") {
			node = &ast.ASTNode{
				Kind:  ast.ADD,
				Left:  node,
				Right: p.parseAddExpression(),
			}
		} else if p.lookAhead("-") {
			node = &ast.ASTNode{
				Kind:  ast.SUB,
				Left:  node,
				Right: p.parseAddExpression(),
			}
		} else {
			return node
		}
	}
}

func (p *Parser) Parse() *ast.ASTNode {
	node := p.parseAddExpression()
	return node
}

func Init(tokens []tokenizer.Token) *Parser {
	return &Parser{
		tokens:  tokens,
		pointer: len(tokens),
	}
}
