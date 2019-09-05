package ast

// Expression Node Type
type ASTNodeKind int

const (
	NUMBER ASTNodeKind = iota
	ADD                = iota
	SUB                = iota
	MUL                = iota
	DIV                = iota
)

type ASTNode struct {
	Kind  ASTNodeKind
	Value string
	Left  *ASTNode
	Right *ASTNode
}
