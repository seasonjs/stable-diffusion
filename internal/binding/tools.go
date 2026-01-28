package binding

import (
	"unsafe"

	"github.com/jupiterrider/ffi"
	"github.com/seasonjs/stable-diffusion/pkg/types"
)

var (
	// SD_API void sd_set_progress_callback(sd_progress_cb_t cb, void* data);
	setProgressCallbackFun ffi.Fun

	// SD_API void sd_set_preview_callback(sd_preview_cb_t cb, enum preview_t mode, int interval, bool denoised, bool noisy, void* data);
	setPreviewCallbackFun ffi.Fun

	// SD_API int32_t sd_get_num_physical_cores();
	getNumPhysicalCoresFun ffi.Fun

	// SD_API const char* sd_get_system_info();
	getSystemInfoFun ffi.Fun

	// SD_API const char* sd_type_name(enum sd_type_t type);
	typeNameFun ffi.Fun

	// SD_API enum sd_type_t str_to_sd_type(const char* str);
	strToSdTypeFun ffi.Fun

	// SD_API const char* sd_rng_type_name(enum rng_type_t rng_type);
	rngTypeNameFun ffi.Fun

	// SD_API enum rng_type_t str_to_rng_type(const char* str);
	strToRngTypeFun ffi.Fun

	// SD_API bool convert(const char* input_path, const char* vae_path, const char* output_path, enum sd_type_t output_type, const char* tensor_type_rules, bool convert_name);
	convertFun ffi.Fun

	// SD_API const char* sd_commit(void);
	commitFun ffi.Fun

	// SD_API const char* sd_version(void);
	versionFun ffi.Fun
)

func LoadToosFuncs(lib ffi.Lib) error {
	var err error
	// // SD_API void sd_set_progress_callback(sd_progress_cb_t cb, void* data);
	setProgressCallbackFun, err = lib.Prep("sd_set_progress_callback", &ffi.TypeVoid, &ffi.TypePointer, &ffi.TypePointer)
	if err != nil {
		return err
	}

	// SD_API void sd_set_preview_callback(sd_preview_cb_t cb, enum preview_t mode, int interval, bool denoised, bool noisy, void* data);
	setPreviewCallbackFun, err = lib.Prep("sd_set_preview_callback", &ffi.TypeVoid, &ffi.TypePointer, &ffi.TypeSint32, &ffi.TypeSint32, &ffi.TypeUint8, &ffi.TypeUint8, &ffi.TypePointer)
	if err != nil {
		return err
	}

	// SD_API int32_t sd_get_num_physical_cores();
	getNumPhysicalCoresFun, err = lib.Prep("sd_get_num_physical_cores", &ffi.TypeSint32)
	if err != nil {
		return err
	}

	// SD_API const char* sd_get_system_info();
	getSystemInfoFun, err = lib.Prep("sd_get_system_info", &ffi.TypePointer)
	if err != nil {
		return err
	}

	// SD_API const char* sd_type_name(enum sd_type_t type);
	typeNameFun, err = lib.Prep("sd_type_name", &ffi.TypePointer, &ffi.TypeSint32)
	if err != nil {
		return err
	}

	// SD_API enum sd_type_t str_to_sd_type(const char* str);
	strToSdTypeFun, err = lib.Prep("str_to_sd_type", &ffi.TypeSint32, &ffi.TypePointer)
	if err != nil {
		return err
	}

	// SD_API const char* sd_rng_type_name(enum rng_type_t rng_type);
	rngTypeNameFun, err = lib.Prep("sd_rng_type_name", &ffi.TypePointer, &ffi.TypeSint32)
	if err != nil {
		return err
	}

	// SD_API enum rng_type_t str_to_rng_type(const char* str);
	strToRngTypeFun, err = lib.Prep("str_to_rng_type", &ffi.TypeSint32, &ffi.TypePointer)
	if err != nil {
		return err
	}

	// SD_API const char* sd_commit(void);
	commitFun, err = lib.Prep("sd_commit", &ffi.TypePointer)
	if err != nil {
		return err
	}

	// SD_API const char* sd_version(void);
	versionFun, err = lib.Prep("sd_version", &ffi.TypePointer)
	if err != nil {
		return err
	}

	// SD_API bool convert(const char* input_path, const char* vae_path, const char* output_path, enum sd_type_t output_type, const char* tensor_type_rules, bool convert_name);
	convertFun, err = lib.Prep("convert", &ffi.TypeUint32, &ffi.TypePointer, &ffi.TypePointer, &ffi.TypePointer, &ffi.TypeSint32, &ffi.TypePointer, &ffi.TypeUint8)
	if err != nil {
		return err
	}

	return nil
}

func SetProgressCallback(callback uintptr) {
	nada := uintptr(0)
	setProgressCallbackFun.Call(nil, unsafe.Pointer(&callback), unsafe.Pointer(&nada))
}

func SetPreviewCallback(callback uintptr, mode types.Preview, interval int, denoised bool, noisy bool) {
	nada := uintptr(0)
	setPreviewCallbackFun.Call(nil, unsafe.Pointer(&callback), unsafe.Pointer(&mode), unsafe.Pointer(&interval), unsafe.Pointer(&denoised), unsafe.Pointer(&noisy), unsafe.Pointer(&nada))
}

func GetNumPhysicalCores() int32 {
	var result int32
	getNumPhysicalCoresFun.Call(unsafe.Pointer(&result))
	return result
}

func GetSystemInfo() *byte {
	var result *byte
	getSystemInfoFun.Call(unsafe.Pointer(&result))
	return result
}

func TypeName(t types.SdType) *byte {
	var result *byte
	typeNameFun.Call(unsafe.Pointer(&result), unsafe.Pointer(&t))
	return result
}

func StrToSdType(str *byte) types.SdType {
	var result int32
	strToSdTypeFun.Call(unsafe.Pointer(&result), unsafe.Pointer(&str))
	return types.SdType(result)
}

func RngTypeName(t types.RNGType) *byte {
	var result *byte
	rngTypeNameFun.Call(unsafe.Pointer(&result), unsafe.Pointer(&t))
	return result
}

func StrToRngType(str *byte) types.RNGType {
	var result int32
	strToRngTypeFun.Call(unsafe.Pointer(&result), unsafe.Pointer(&str))
	return types.RNGType(result)
}

func Convert(inputPath, vaePath, outputPath string, outputType types.SdType, tensorTypeRules string, convertName bool) bool {
	var result ffi.Arg
	convertFun.Call(unsafe.Pointer(&result), unsafe.Pointer(&inputPath), unsafe.Pointer(&vaePath), unsafe.Pointer(&outputPath), unsafe.Pointer(&outputType), unsafe.Pointer(&tensorTypeRules), unsafe.Pointer(&convertName))
	return result.Bool()
}

func Commit() *byte {
	var result *byte
	commitFun.Call(unsafe.Pointer(&result))
	return result
}

func Version() *byte {
	var result *byte
	versionFun.Call(unsafe.Pointer(&result))
	return result
}
