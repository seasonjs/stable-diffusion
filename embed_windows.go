// Copyright (c) seasonjs. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

//go:build windows && amd64

package sd

import (
	_ "embed"
	"golang.org/x/sys/cpu"
)

//go:embed deps/windows/sd-abi_avx2.dll
var libStableDiffusionAvx2 []byte

//go:embed deps/windows/sd-abi_avx.dll
var libStableDiffusionAvx []byte

//go:embed deps/windows/sd-abi_avx512.dll
var libStableDiffusionAvx512 []byte

var libName = "stable-diffusion-*.dll"

func getDl() []byte {
	if cpu.X86.HasAVX512 {
		return libStableDiffusionAvx512
	}

	if cpu.X86.HasAVX2 {
		return libStableDiffusionAvx2
	}

	return libStableDiffusionAvx
}
