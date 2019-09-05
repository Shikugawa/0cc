package generator

import (
	"fmt"
	"strconv"

	"github.com/gocc/ast"
)

type Generator struct {
	Asm string
}

func (g *Generator) generate(node *ast.ASTNode) {
	switch node.Kind {
	case ast.ADD:
		if node.Left != nil {
			g.generate(node.Left)
		}
		if node.Right != nil {
			g.generate(node.Right)
		}
		g.Asm += fmt.Sprintf("  popq %%rax\n")
		g.Asm += fmt.Sprintf("  popq %%rcx\n")
		g.Asm += fmt.Sprintf("  addq %%rcx, %%rax\n")
		g.Asm += fmt.Sprintf("  pushq %%rax\n")
	case ast.NUMBER:
		num, _ := strconv.ParseInt(node.Value, 10, 32)
		g.Asm += fmt.Sprintf("  pushq $%d\n", num)
	}
}

func Init(node *ast.ASTNode) *Generator {
	gen := &Generator{
		Asm: "",
	}
	gen.Asm += fmt.Sprintf("  .global main\n")
	gen.Asm += fmt.Sprintf("main:\n")
	gen.generate(node)
	gen.Asm += fmt.Sprintf("  ret\n")
	return gen
}
