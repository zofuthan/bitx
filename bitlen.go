// +build !gccgo,!amd64,!386,!arm

// ((NOT gccgo) AND (NOT amd64) AND (NOT 386) AND (NOT ARM))
package bithacks

func Bitlen(x uint64) int {
	return 32 - int(Nlz1a(uint32(x)))
}
