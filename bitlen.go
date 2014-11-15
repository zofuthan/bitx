// +build !gccgo,!amd64,!386,!arm

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

// ((NOT gccgo) AND (NOT amd64) AND (NOT 386) AND (NOT ARM))
package bithacks

func Bitlen(x uint64) int {
	return 32 - int(Nlz1a(uint32(x)))
}
