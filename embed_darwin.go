// Copyright (c) seasonjs. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

//go:build darwin

package stable_diffusion

import (
	_ "embed" // Needed for go:embed
)

//go:embed deps/darwin/libstable-diffusion-arm64.dylib
var libRwkvArm []byte

var libName = "libstable-diffusion-*.dylib"

func getDl() []byte {
	//TODO: support x86
	return libRwkvArm
}
