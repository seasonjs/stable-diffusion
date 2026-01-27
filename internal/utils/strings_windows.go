//go:build windows

package utils

import "golang.org/x/sys/windows"

// CString converts a go string to *byte that can be passed to C code.
func CString(s string) (*byte, error) {
	return windows.BytePtrFromString(s)
}

// GoString copies a null-terminated char* to a Go string.
func GoString(p *byte) string {
	return windows.BytePtrToString(p)
}

