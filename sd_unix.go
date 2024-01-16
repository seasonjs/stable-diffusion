// Copyright (c) seasonjs. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

//go:build (darwin || linux) && cgo

package sd

import (
	"errors"
	"unsafe"
)

/*
#include <dlfcn.h>

void *Dlopen(const char *libPath) {
    void *handle = dlopen(libPath, RTLD_GLOBAL | RTLD_NOW);
    if (!handle) {
        return NULL;
    }

    return handle;
}

void closeLibrary(void *handle) {
    dlclose(handle);
}
*/
import "C"

func openLibrary(name string) (uintptr, error) {
	handle := C.Dlopen(C.CString(name))
	if handle == nil {
		return 0, errors.New("failed to open library")
	}
	return uintptr(handle), nil
}

func closeLibrary(handle uintptr) error {
	C.closeLibrary(*(*unsafe.Pointer)(unsafe.Pointer(&handle)))
	return nil
}
