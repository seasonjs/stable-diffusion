// Copyright (c) seasonjs. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

//go:build linux

package sd

//go:embed deps/linux/libsd-abi.so
var libStableDiffusion []byte

func getDl() []byte {
	return libStableDiffusion
}
