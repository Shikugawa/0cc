package x64

import (
	"fmt"
	"strconv"

	"github.com/gocc/ast"
)

type Generator struct {
	Asm string
}

func (g *Generator) generate(node *ast.ASTNode) {
	if node.Kind == ast.UNARY {
		num, _ := strconv.ParseInt(node.Value, 10, 32)
		g.Asm += fmt.Sprintf("  pushq $%d\n", num)
		return
	}

	if node.Left != nil {
		g.generate(node.Left)
	}
	if node.Right != nil {
		g.generate(node.Right)
	}

	switch node.Kind {
	case ast.ADD:
		emitAddInstr(g)
	case ast.SUB:
		emitSubInstr(g)
	case ast.MUL:
		emitSubInstr(g)
	case ast.DIV:
		emitDivInstr(g)
	case ast.LEFT_INEQUALITY:
		emitLeftCompInstr(g)
	case ast.RIGHT_INEQUALITY:
		emitRightCompInstr(g)
	case ast.LEFT_INEQUALITY_EQ:
		emitLeftCompEqInstr(g)
	case ast.RIGHT_INEQUALITY_EQ:
		emitRightCompEqInstr(g)
	default:
		panic(fmt.Sprintf("\"%s\" is not supported", node.Kind))
	}
}

func Init(node *ast.ASTNode) *Generator {
	gen := &Generator{
		Asm: "",
	}
	gen.Asm += fmt.Sprintf(".global _main\n")
	gen.Asm += fmt.Sprintf("_main:\n")
	gen.generate(node)
	gen.Asm += fmt.Sprintf("  popq %%rax\n")
	gen.Asm += fmt.Sprintf("  ret\n")
	return gen
}
