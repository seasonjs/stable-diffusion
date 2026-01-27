package binding

import (
	"unsafe"

	"github.com/jupiterrider/ffi"
)

var (
	// SD_API bool preprocess_canny(sd_image_t image,float high_threshold,float low_threshold,float weak,float strong,bool inverse);
	preprocessCannyFun ffi.Fun
)

func LoadPreprocessingFuncs(lib ffi.Lib) error {
	var err error
	// SD_API bool preprocess_canny(sd_image_t image,float high_threshold,float low_threshold,float weak,float strong,bool inverse);
	preprocessCannyFun, err = lib.Prep("preprocess_canny", &ffi.TypeUint32, &ffi.TypeFloat, &ffi.TypeFloat, &ffi.TypeFloat, &ffi.TypeFloat, &ffi.TypeUint8)
	
	if err != nil {
		return err
	}

	return nil
}

func PreprocessCanny(image uintptr, highThreshold float32, lowThreshold float32, weak float32, strong float32, inverse bool) bool {
	var result ffi.Arg
	preprocessCannyFun.Call(unsafe.Pointer(&result), unsafe.Pointer(&image), unsafe.Pointer(&highThreshold), unsafe.Pointer(&lowThreshold), unsafe.Pointer(&weak), unsafe.Pointer(&strong), unsafe.Pointer(&inverse))
	return result.Bool()
}


