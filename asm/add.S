.global _main

_main:
  pushq $4
  pushq $2
  popq %rax
  popq %rcx
  addq %rcx, %rax
  ret
