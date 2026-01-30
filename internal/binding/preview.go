package binding

import (
	"github.com/jupiterrider/ffi"
	"unsafe"
)

var (
	// SD_API const char* sd_preview_name(enum preview_t preview);
	previewNameFun ffi.Fun

	// SD_API enum preview_t str_to_preview(const char* str);
	strToPreviewFun ffi.Fun
)

func LoadPreviewFuns(lib ffi.Lib) error {
	var err error
	// SD_API const char* sd_preview_name(enum preview_t preview);
	previewNameFun, err = lib.Prep("sd_preview_name", &ffi.TypePointer, &ffi.TypeSint32)
	if err != nil {
		return err
	}

	// SD_API enum preview_t str_to_preview(const char* str);
	strToPreviewFun, err = lib.Prep("str_to_preview", &ffi.TypeSint32, &ffi.TypePointer)
	if err != nil {
		return err
	}

	return nil
}

func PreviewName(preview int32) *byte {
	var result *byte
	previewNameFun.Call(unsafe.Pointer(&result), unsafe.Pointer(&preview))
	return result
}

func StrToPreview(str *byte) int32 {
	var result int32
	strToPreviewFun.Call(unsafe.Pointer(&result), unsafe.Pointer(&str))
	return result
}
