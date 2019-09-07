package ast

// Expression Node Type
type ASTNodeKind int

const (
	UNARY               = iota
	ADD                 = iota
	SUB                 = iota
	MUL                 = iota
	DIV                 = iota
	LEFT_INEQUALITY     = iota
	LEFT_INEQUALITY_EQ  = iota
	RIGHT_INEQUALITY    = iota
	RIGHT_INEQUALITY_EQ = iota
)

type ASTNode struct {
	Kind  ASTNodeKind
	Value string
	Left  *ASTNode
	Right *ASTNode
}
