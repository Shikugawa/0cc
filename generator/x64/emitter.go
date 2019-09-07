package x64

import "fmt"

func emitPopBinaryInstr(g *Generator) {
	g.Asm += fmt.Sprintf("  popq %%rax\n")
	g.Asm += fmt.Sprintf("  popq %%rcx\n")
}

func emitAddInstr(g *Generator) {
	emitPopBinaryInstr(g)
	g.Asm += fmt.Sprintf("  addq %%rcx, %%rax\n")
	g.Asm += fmt.Sprintf("  pushq %%rax\n")
}

func emitSubInstr(g *Generator) {
	emitPopBinaryInstr(g)
	g.Asm += fmt.Sprintf("  subq %%rcx, %%rax\n")
	g.Asm += fmt.Sprintf("  pushq %%rax\n")
}

func emitMulInstr(g *Generator) {
	emitPopBinaryInstr(g)
	g.Asm += fmt.Sprintf("  imulq %%rcx, %%rax\n")
	g.Asm += fmt.Sprintf("  pushq %%rax\n")
}

func emitDivInstr(g *Generator) {
	emitPopBinaryInstr(g)
	g.Asm += fmt.Sprintf("  movq $0, %%rdx\n") // 余りが格納されるrdxレジスタを0埋めで初期化
	g.Asm += fmt.Sprintf("  idivq %%rcx\n")
	g.Asm += fmt.Sprintf("  pushq %%rax\n")
}

func emitLeftCompInstr(g *Generator) {
	emitPopBinaryInstr(g)
	g.Asm += fmt.Sprintf("  cmpq %%rcx, %%rax\n")
	g.Asm += fmt.Sprintf("  movq $0, %%rax\n") // raxレジスタの初期化
	g.Asm += fmt.Sprintf("  setnl %%al\n")     // 左項の方が大きかったらraxレジスタに0x01を格納する
	g.Asm += fmt.Sprintf("  movzb %%al, %%rax\n")
	g.Asm += fmt.Sprintf("  pushq %%rax\n")
}

func emitRightCompInstr(g *Generator) {
	emitPopBinaryInstr(g)
	g.Asm += fmt.Sprintf("  cmpq %%rcx, %%rax\n")
	g.Asm += fmt.Sprintf("  movq $0, %%rax\n") // raxレジスタの初期化
	g.Asm += fmt.Sprintf("  setng %%al\n")     // 左項の方が大きかったらraxレジスタに0x01を格納する
	g.Asm += fmt.Sprintf("  movzb %%al, %%rax\n")
	g.Asm += fmt.Sprintf("  pushq %%rax\n")
}

func emitLeftCompEqInstr(g *Generator) {
	emitPopBinaryInstr(g)
	g.Asm += fmt.Sprintf("  cmpq %%rcx, %%rax\n")
	g.Asm += fmt.Sprintf("  movq $0, %%rax\n") // raxレジスタの初期化
	g.Asm += fmt.Sprintf("  setnle %%al\n")    // 左項の方が大きかったらraxレジスタに0x01を格納する
	g.Asm += fmt.Sprintf("  movzb %%al, %%rax\n")
	g.Asm += fmt.Sprintf("  pushq %%rax\n")
}

func emitRightCompEqInstr(g *Generator) {
	emitPopBinaryInstr(g)
	g.Asm += fmt.Sprintf("  cmpq %%rcx, %%rax\n")
	g.Asm += fmt.Sprintf("  movq $0, %%rax\n") // raxレジスタの初期化
	g.Asm += fmt.Sprintf("  setnge %%al\n")    // 左項の方が大きかったらraxレジスタに0x01を格納する
	g.Asm += fmt.Sprintf("  movzb %%al, %%rax\n")
	g.Asm += fmt.Sprintf("  pushq %%rax\n")
}
