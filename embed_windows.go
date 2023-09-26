// Copyright (c) seasonjs. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

//go:build windows && amd64

package sd

import (
	_ "embed"
	"golang.org/x/sys/cpu"
)

//go:embed deps/windows/stable-diffusion-avx2.dll
var libStableDiffusionAvx []byte

var libName = "stable-diffusion-*.dll"

func getDl() []byte {
	if cpu.X86.HasAVX2 {
		return libStableDiffusionAvx
	}
	panic("Automatic loading of dynamic library failed, please use `NewStableDiffusionModel` method load manually. ")
	return nil
}
