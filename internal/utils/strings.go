package utils

import "unsafe"

type Stringer interface {
	*byte | uintptr
}

// GoString converts a C-style null-terminated byte pointer to a Go string.
func GoString[T Stringer](p T) string {
	switch val := any(p).(type) {
	case *byte:
		// 在这个分支，val 的类型是 *byte
		return goByteString(val)
	case uintptr:
		// 在这个分支，val 的类型是 uintptr
		return goPrtString(val)
	default:
		return ""
	}

}

func goPrtString(c uintptr) string {
	// We take the address and then dereference it to trick go vet from creating a possible misuse of unsafe.Pointer
	ptr := *(*unsafe.Pointer)(unsafe.Pointer(&c))
	if ptr == nil {
		return ""
	}
	var length int
	for {
		if *(*byte)(unsafe.Add(ptr, uintptr(length))) == '\x00' {
			break
		}
		length++
	}
	return unsafe.String((*byte)(ptr), length)
}
