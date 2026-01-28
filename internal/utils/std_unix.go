//go:build linux

package utils

import (
	"github.com/jupiterrider/ffi"
)

func LoadStdLib() (ffi.Lib, error) {
	return ffi.Load("libc.so.6")
}
