// Copyright (c) seasonjs. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

//go:build linux

package sd

//go:embed deps/linux/libsd-abi.so
var libStableDiffusion []byte

func getDl(gpu bool) []byte {
	if gpu {
		panic("Not support linux. Push request is welcome.")
	}
	return libStableDiffusion
}
