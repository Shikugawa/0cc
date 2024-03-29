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

func (p *Parser) lookCurrent(operator string) bool {
	if p.pointer < 0 {
		return false
	}

	if operator == p.tokens[p.pointer].Value {
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

// UnaryExpr ::= (+|-)CastExpr | CastExpr | "(" AddExpr ")"
func (p *Parser) parseUnaryExpr() *ast.ASTNode {
	node := p.parseCastExpr()
	for {
		if p.lookAhead("-") {
			node = &ast.ASTNode{
				Kind:  ast.UNARY,
				Value: "-" + node.Value,
			}
		} else if p.lookCurrent(")") {
			node = p.parseAddExpr()
		} else if p.lookCurrent("(") {
			return node
		} else {
			return node
		}
	}
}

// MulExpr ::= UnaryExpr | MulExpr * UnaryExpr | MulExpr / UnaryExpr
func (p *Parser) parseMulExpr() *ast.ASTNode {
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

// AddExpr ::= MulExpr | AddExpr + MulExpr | AddExpr - MulExpr
func (p *Parser) parseAddExpr() *ast.ASTNode {
	node := p.parseMulExpr()
	for {
		if p.lookAhead("+") {
			node = &ast.ASTNode{
				Kind:  ast.ADD,
				Left:  node,
				Right: p.parseAddExpr(),
			}
		} else if p.lookAhead("-") {
			node = &ast.ASTNode{
				Kind:  ast.SUB,
				Left:  node,
				Right: p.parseAddExpr(),
			}
		} else {
			return node
		}
	}
}

// ShiftExpr ::= AddExpr
func (p *Parser) parseShiftExpr() *ast.ASTNode {
	node := p.parseAddExpr()
	return node
}

// RelationalExpr ::= ShiftExpr
//	| RelationalExpr < ShiftExpr
//	| RelationalExpr > ShiftExpr
//	| RelationalExpr <= ShiftExpr
//	| RelationalExpr <= ShiftExpr
func (p *Parser) parseRelationalExpr() *ast.ASTNode {
	node := p.parseShiftExpr()
	for {
		if p.lookAhead("<") {
			node = &ast.ASTNode{
				Kind:  ast.RIGHT_INEQUALITY,
				Left:  node,
				Right: p.parseRelationalExpr(),
			}
		} else if p.lookAhead(">") {
			node = &ast.ASTNode{
				Kind:  ast.LEFT_INEQUALITY,
				Left:  node,
				Right: p.parseRelationalExpr(),
			}
		} else if p.lookAhead("<=") {
			node = &ast.ASTNode{
				Kind:  ast.RIGHT_INEQUALITY_EQ,
				Left:  node,
				Right: p.parseRelationalExpr(),
			}
		} else if p.lookAhead(">=") {
			node = &ast.ASTNode{
				Kind:  ast.LEFT_INEQUALITY_EQ,
				Left:  node,
				Right: p.parseRelationalExpr(),
			}
		} else if p.lookAhead("==") {
			node = &ast.ASTNode{
				Kind:  ast.EQUAL,
				Left:  node,
				Right: p.parseRelationalExpr(),
			}
		} else if p.lookAhead("!=") {
			node = &ast.ASTNode{
				Kind:  ast.NOT_EQUAL,
				Left:  node,
				Right: p.parseRelationalExpr(),
			}
		} else {
			return node
		}
	}
}

func (p *Parser) ParseExpr() *ast.ASTNode {
	node := p.parseRelationalExpr()
	return node
}

func Init(tokens []tokenizer.Token) *Parser {
	return &Parser{
		tokens:  tokens,
		pointer: len(tokens),
	}
}
