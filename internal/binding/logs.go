package binding

import (
	"unsafe"

	"github.com/jupiterrider/ffi"
)

var (
	// SD_API void sd_set_log_callback(sd_log_cb_t sd_log_cb, void* data);
	setLogCallback ffi.Fun
)

func LoadLogFuns(lib ffi.Lib) error {
	var err error

	setLogCallback, err = lib.Prep("sd_set_log_callback", &ffi.TypeVoid, &ffi.TypePointer, &ffi.TypePointer)

	if err != nil {
		return err
	}

	return nil
}

func SetLogCallback(callback uintptr) {
	nada := uintptr(0)
	setLogCallback.Call(nil, unsafe.Pointer(&callback), unsafe.Pointer(&nada))
}
