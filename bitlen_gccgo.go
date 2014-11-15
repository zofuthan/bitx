// +build gccgo

package bithacks

// this is apparetly the old way -> func clz(uint64) uint64 __asm__("__clzdi2")

//extern __clzdi2
func clz(uint64) uint64

func Bitlen(x uint64) int {
	if x == 0 {
		return 0
	}
	return 64 - int(clz(x))
}
