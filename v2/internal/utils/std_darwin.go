//go:build darwin

package utils

import (
	"github.com/jupiterrider/ffi"
)

func LoadStdLib() (ffi.Lib, error) {
	return ffi.Load("libSystem.B.dylib")
}
