// Copyright (c) seasonjs. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

//go:build linux

package sd

// TODO support linux
var libName = "libstable-diffusion-*.so"

func getDl() []byte {
	panic("Automatic loading of dynamic library failed, please use `NewStableDiffusionModel` method load manually. ")
	return nil
}
