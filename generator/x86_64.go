package generator

import "fmt"

func Generate() {
	fmt.Printf("  .global main\n")
	fmt.Printf("main:\n")
	fmt.Printf("  movq $%d, %%rax\n", 3)
	fmt.Printf("  movq $%d, %%rcx\n", 4)
	fmt.Printf("  ret\n")
}
