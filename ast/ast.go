package ast

// Expression Node Type
type ASTNodeType int

const (
	NUMBER = iota
	ADD    = iota
	SUB    = iota
	MUL    = iota
	DIV    = iota
)

type ASTNode struct {
	Kind  ASTNodeType
	Value string
	Left  *ASTNode
	Right *ASTNode
}
