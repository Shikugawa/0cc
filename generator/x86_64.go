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
	case ast.SUB:
		if node.Left != nil {
			g.generate(node.Left)
		}
		if node.Right != nil {
			g.generate(node.Right)
		}
		g.Asm += fmt.Sprintf("  popq %%rax\n")
		g.Asm += fmt.Sprintf("  popq %%rcx\n")
		g.Asm += fmt.Sprintf("  subq %%rcx, %%rax\n")
		g.Asm += fmt.Sprintf("  pushq %%rax\n")
	case ast.MUL:
		if node.Left != nil {
			g.generate(node.Left)
		}
		if node.Right != nil {
			g.generate(node.Right)
		}
		g.Asm += fmt.Sprintf("  popq %%rax\n")
		g.Asm += fmt.Sprintf("  popq %%rcx\n")
		g.Asm += fmt.Sprintf("  imulq %%rcx, %%rax\n")
		g.Asm += fmt.Sprintf("  pushq %%rax\n")
	case ast.DIV:
		if node.Left != nil {
			g.generate(node.Left)
		}
		if node.Right != nil {
			g.generate(node.Right)
		}
		g.Asm += fmt.Sprintf("  popq %%rax\n")
		g.Asm += fmt.Sprintf("  popq %%rcx\n")
		g.Asm += fmt.Sprintf("  movq $0, %%rdx\n") // 余りが格納されるrdxレジスタを0埋めで初期化
		g.Asm += fmt.Sprintf("  idivq %%rcx\n")
		g.Asm += fmt.Sprintf("  pushq %%rax\n")
	case ast.NUMBER:
		num, _ := strconv.ParseInt(node.Value, 10, 32)
		g.Asm += fmt.Sprintf("  pushq $%d\n", num)
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
