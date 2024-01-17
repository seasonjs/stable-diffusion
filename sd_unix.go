// Copyright (c) seasonjs. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

//go:build cgo && (darwin || linux)

package sd

import (
	"errors"
	"unsafe"
)

/*
#include <dlfcn.h>
*/
import "C"

func openLibrary(name string) (uintptr, error) {
	handle := C.dlopen(C.CString(name))
	if handle == nil {
		return 0, errors.New("failed to open library")
	}
	return uintptr(handle), nil
}

func closeLibrary(handle uintptr) error {
	C.dlclose(*(*unsafe.Pointer)(unsafe.Pointer(&handle)))
	return nil
}
