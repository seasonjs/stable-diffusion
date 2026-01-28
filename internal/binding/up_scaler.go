package binding

import (
	"github.com/jupiterrider/ffi"
	"unsafe"
)

var (
	// SD_API upscaler_ctx_t* new_upscaler_ctx(const char* esrgan_path, bool offload_params_to_cpu, bool direct, int n_threads, int tile_size);
	newUpscalerCtxFun ffi.Fun

	// SD_API void free_upscaler_ctx(upscaler_ctx_t* upscaler_ctx);
	freeUpscalerCtxFun ffi.Fun

	// SD_API sd_image_t upscale(upscaler_ctx_t* upscaler_ctx, sd_image_t input_image, uint32_t upscale_factor);
	upscaleFun ffi.Fun

	// SD_API int get_upscale_factor(upscaler_ctx_t* upscaler_ctx);
	getUpscaleFactorFun ffi.Fun
)

type UpscalerContext uintptr

func LoadUpsScalerFuncs(lib ffi.Lib) error {
	var err error
	// SD_API upscaler_ctx_t* new_upscaler_ctx(const char* esrgan_path, bool offload_params_to_cpu, bool direct, int n_threads, int tile_size);
	newUpscalerCtxFun, err = lib.Prep("new_upscaler_ctx",
		&ffi.TypePointer,
		&ffi.TypePointer, //esrgan_path
		&ffi.TypeUint8,   //offload_params_to_cpu
		&ffi.TypeUint8,   //direct
		&ffi.TypeSint32,  //n_threads
		&ffi.TypeSint32,  //tile_size
	)

	if err != nil {
		return err
	}

	// SD_API void free_upscaler_ctx(upscaler_ctx_t* upscaler_ctx);
	freeUpscalerCtxFun, err = lib.Prep("free_upscaler_ctx", &ffi.TypeVoid, &ffi.TypePointer)
	if err != nil {
		return err
	}

	// SD_API sd_image_t upscale(upscaler_ctx_t* upscaler_ctx, sd_image_t input_image, uint32_t upscale_factor);
	upscaleFun, err = lib.Prep("upscale", &FFITypeImage, &ffi.TypePointer, &FFITypeImage, &ffi.TypeUint32)
	if err != nil {
		return err
	}

	// SD_API int get_upscale_factor(upscaler_ctx_t* upscaler_ctx);
	getUpscaleFactorFun, err = lib.Prep("get_upscale_factor", &ffi.TypeSint32, &ffi.TypePointer)
	if err != nil {
		return err
	}
	return nil
}

func NewUpscalerCtx(esrganPath *byte, offloadParamsToCPU bool, direct bool, nThreads int32, tileSize int32) UpscalerContext {
	var ctx UpscalerContext
	newUpscalerCtxFun.Call(unsafe.Pointer(&ctx), unsafe.Pointer(&esrganPath), unsafe.Pointer(&offloadParamsToCPU), unsafe.Pointer(&direct), unsafe.Pointer(&nThreads), unsafe.Pointer(&tileSize))
	return ctx
}

func Upscale(ctx UpscalerContext, inputImage Image, upscaleFactor uint32) Image {
	var result Image
	upscaleFun.Call(unsafe.Pointer(&result), unsafe.Pointer(&ctx), unsafe.Pointer(&inputImage), &upscaleFactor)
	return result
}

func GetUpscaleFactor(ctx UpscalerContext) int32 {
	var result int32
	getUpscaleFactorFun.Call(unsafe.Pointer(&result), unsafe.Pointer(&ctx))
	return result
}

func FreeUpscalerCtx(ctx UpscalerContext) {
	freeUpscalerCtxFun.Call(nil, unsafe.Pointer(&ctx))
}
