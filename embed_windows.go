// Copyright (c) seasonjs. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

//go:build windows && amd64

package stable_diffusion

import _ "embed"

//go:embed deps/windows/stable-diffusion_avx2_x64.dll
var libStableDiffusionAvx []byte

var libName = "stable-diffusion-*.dll"

func getDl() []byte {
	//TODO: support x86
	return libStableDiffusionAvx
}
