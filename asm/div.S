.global _main

_main:
  movq $0, %rdx
  movq $4, %rax
  movq $2, %rcx
  idivq %rcx
  ret
