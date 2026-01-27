//go:build darwin || linux

package utils

import "golang.org/x/sys/unix"

// CString converts a Go string to a C-style null-terminated byte pointer.
func CString(s string) (*byte, error) {
	return unix.BytePtrFromString(s)
}

// GoString converts a C-style null-terminated byte pointer to a Go string.
func GoString(p *byte) string {
	return unix.BytePtrToString(p)
}