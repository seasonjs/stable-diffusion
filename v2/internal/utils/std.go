package utils

import (
	"github.com/ebitengine/purego"
	"github.com/jupiterrider/ffi"
)

var free func(uintptr)

func LoadStdFuns(libc ffi.Lib) error {

	purego.RegisterLibFunc(&free, libc.Addr, "free")

	return nil
}

func Free(ptr uintptr) {
	if ptr != 0 {
		free(ptr)
		ptr = 0
	}
}
