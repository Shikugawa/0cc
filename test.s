.global _main

_main:
  movl $4, %eax
  movl $3, %ebx
  addl %ebx, %eax
  ret
