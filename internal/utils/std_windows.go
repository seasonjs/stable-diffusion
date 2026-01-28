//go:build windows

package utils

import (
	"github.com/jupiterrider/ffi"
)

func LoadStdLib() (ffi.Lib, error) {
	libc, err := ffi.Load("ucrtbase.dll")
	if err != nil {
		libc, err = ffi.Load("msvcrt.dll")
		if err != nil {
			panic(err)
		}
	}
	return libc, nil
}
