.global _main

_main:
  movq $2, %rax
  movq $3, %rcx
  cmpq %rcx, %rax
  setl %al
  ret
