// +build !gccgo

// func Bitlen(x Word) (n int)
TEXT ·Bitlen(SB),4,$0
	BSRQ x+0(FP), AX
	JZ Z1
	ADDQ $1, AX
	MOVQ AX, n+8(FP)
	RET

Z1:	MOVQ $0, n+8(FP)
	RET
