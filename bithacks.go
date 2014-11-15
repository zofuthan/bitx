// Copyright (c) 2014 Dataence, LLC. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package bithacks

func PowerOfTwo(n int) bool {
	return n != 0 && (n&(n-1)) == 0
}

func RoundUpPowerOfTwo32(n int32) int32 {
	n--
	n |= n >> 1
	n |= n >> 2
	n |= n >> 4
	n |= n >> 8
	n |= n >> 16
	n++

	return n
}

func RoundUpPowerOfTwo64(n int64) int64 {
	n--
	n |= n >> 1
	n |= n >> 2
	n |= n >> 4
	n |= n >> 8
	n |= n >> 16
	n |= n >> 32
	n++

	return n
}

// Finalization mix - force all bits of a hash block to avalanche
// Copied from murmur32: https://code.google.com/p/smhasher/wiki/MurmurHash3
func Fmix32(h uint32) uint32 {
	h ^= h >> 16
	h *= 0x85ebca6b
	h ^= h >> 13
	h *= 0xc2b2ae35
	h ^= h >> 16

	return h
}

// Copied from murmur32: https://code.google.com/p/smhasher/wiki/MurmurHash3
func Fmix64(h uint64) uint64 {
	h ^= h >> 33
	h *= 0xff51afd7ed558ccd
	h ^= h >> 33
	h *= 0xc4ceb9fe1a85ec53
	h ^= h >> 33

	return h
}

// The bswap32() and bswap64() functions return a byte order swapped integer.
// On big endian systems, the number is converted to little endian byte order.
// On little endian systems, the number is converted to big endian byte order.
func Bswap32(x uint32) uint32 {
	// Copied from netbsd's bswap32.c
	return ((x << 24) & 0xff000000) |
		((x << 8) & 0x00ff0000) |
		((x >> 8) & 0x0000ff00) |
		((x >> 24) & 0x000000ff)
}

func Bswap64(x uint64) uint64 {
	// Copied from netbsd's bswap64.c
	return ((x << 56) & 0xff00000000000000) |
		((x << 40) & 0x00ff000000000000) |
		((x << 24) & 0x0000ff0000000000) |
		((x << 8) & 0x000000ff00000000) |
		((x >> 8) & 0x00000000ff000000) |
		((x >> 24) & 0x0000000000ff0000) |
		((x >> 40) & 0x000000000000ff00) |
		((x >> 56) & 0x00000000000000ff)
}

// Copied from http://www.hackersdelight.org/hdcodetxt/nlz.c.txt - nlz1a
// Number of leading zeros, version 1a
func Nlz1a(x uint32) int {
	var n uint32 = 0
	if x <= 0 {
		return int((^x >> 26) & 32)
	}

	n = 1

	if (x >> 16) == 0 {
		n = n + 16
		x = x << 16
	}
	if (x >> 24) == 0 {
		n = n + 8
		x = x << 8
	}
	if (x >> 28) == 0 {
		n = n + 4
		x = x << 4
	}
	if (x >> 30) == 0 {
		n = n + 2
		x = x << 2
	}
	n = n - (x >> 31)
	return int(n)
}

// Number of leading zeros, version 2
func Nlz2(x uint32) int {
	var y uint32
	var n uint32 = 32

	y = x >> 16
	if y != 0 {
		n = n - 16
		x = y
	}
	y = x >> 8
	if y != 0 {
		n = n - 8
		x = y
	}
	y = x >> 4
	if y != 0 {
		n = n - 4
		x = y
	}
	y = x >> 2
	if y != 0 {
		n = n - 2
		x = y
	}
	y = x >> 1
	if y != 0 {
		return int(n - 2)
	}
	return int(n - x)
}

func LeadingBit(x uint32) int {
	//return 32 - int32(nlz1a(x))
	return int(Bitlen(uint64(x)))
}

// This is defined in util_{amd64,386}.s, copied from pkg/math/big/arith_{amd64/386}.s
func Bitlen(x uint64) int
