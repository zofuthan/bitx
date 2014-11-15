// +build gccgo

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

// this is apparetly the old way -> func clz(uint64) uint64 __asm__("__clzdi2")

//extern __clzdi2
func clz(uint64) uint64

func Bitlen(x uint64) int {
	if x == 0 {
		return 0
	}
	return 64 - int(clz(x))
}
