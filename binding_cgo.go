// Copyright (c) seasonjs. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

//go:build cgo

package sd

/*
#cgo CFLAGS: -I deps
#include "deps/stable-diffusion.h"

void *new_sd_ctx_func(void* handle,
					  const char *modelPath,
					  const char *vaePath,
					  const char *taesdPath,
					  const char *loraModelDir,
					  bool vaeDecodeOnly,
                      bool vaeTiling,
					  bool freeParamsImmediately,
					  int nThreads,
					  int wType,
					  int rngType,
					  int schedule)
{
	if (!libHandle) {
		return NULL;
	}
	void (*new_sd_ctx)() = dlsym(handle, "new_sd_ctx");
	if (!new_sd_ctx) {
		return NULL;
	}
	return new_sd_ctx(modelPath, vaePath, taesdPath, loraModelDir, vaeDecodeOnly, vaeTiling, freeParamsImmediately, nThreads, wType, rngType, schedule);
}

void *txt2img_func(void* handle,
				   void *ctx,
				   const char *prompt,
				   const char *negativePrompt,
				   int clipSkip,
				   float cfgScale,
				   int width,
				   int height,
				   int sampleMethod,
				   int sampleSteps,
				   long long seed,
				   int batchCount)
{
	if (!libHandle) {
		return NULL;
	}
	void (*txt2img)() = dlsym(handle, "txt2img");
	if (!txt2img) {
		return NULL;
	}
	return txt2img(ctx, prompt, negativePrompt, clipSkip, cfgScale, width, height, sampleMethod, sampleSteps, seed, batchCount);
}

void *img2img_func(void* handle,
				   void *ctx,
				   void *img,
				   const char *prompt,
				   const char *negativePrompt,
				   int clipSkip,
				   float cfgScale,
				   int width,
				   int height,
				   int sampleMethod,
                   int sampleSteps,
                   float strength,
				   long long seed,
				   int batchCount)
{
	if (!libHandle) {
		return NULL;
	}
	void (*img2img)() = dlsym(handle, "img2img");
	if (!img2img) {
		return NULL;
	}
	return img2img(ctx, img, prompt, negativePrompt, clipSkip, cfgScale, width, height, sampleMethod, sampleSteps, strength, seed, batchCount);

}

void free_sd_ctx_func(void* handle, void *ctx)
{
	if (!libHandle) {
		return;
	}
	void (*free_sd_ctx)() = dlsym(handle, "free_sd_ctx");
	if (!free_sd_ctx) {
		return;
	}
	if (ctx) {
	    free_sd_ctx(ctx);
    }
}

void *new_upscaler_ctx_func(void* handle,
							const char *esrganPath,
							int nThreads,
							int wType)
{
	if (!libHandle) {
		return NULL;
	}
	void (*new_upscaler_ctx)() = dlsym(handle, "new_upscaler_ctx");
	if (!new_upscaler_ctx) {
		return NULL;
	}
	return new_upscaler_ctx(esrganPath, nThreads, wType);
}

void free_upscaler_ctx_func(void* handle, void *ctx)
{
	if (!libHandle) {
		return;
	}
	void (*free_upscaler_ctx)() = dlsym(handle, "free_upscaler_ctx");
	if (!free_upscaler_ctx) {
		return;
	}
	if (ctx) {
		free_upscaler_ctx(ctx);
	}
}

void *upscale_func(void* handle,
				   void *ctx,
				   void *img,
				   unsigned int upscaleFactor)
{
	if (!libHandle) {
		return NULL;
	}
	void (*upscale)() = dlsym(handle, "upscale");
	if (!upscale) {
		return NULL;
	}
	return upscale(ctx, img, upscaleFactor);
}

void *sd_get_system_info_func(void* handle)
{
	if (!libHandle) {
		return NULL;
	}
	void (*sd_get_system_info)() = dlsym(handle, "sd_get_system_info");
	if (!sd_get_system_info) {
		return NULL;
	}
	return sd_get_system_info();
}

void sd_set_log_callback_func(void* handle, void *callback, void *data)
{
	if (!libHandle) {
		return;
	}
	void (*sd_set_log_callback)() = dlsym(handle, "sd_set_log_callback");
	if (!sd_set_log_callback) {
		return;
	}
	sd_set_log_callback(callback, data);
}
*/
import "C"
import (
	"runtime"
	"unsafe"
)

type CStableDiffusionImpl struct {
	libSd uintptr
}

func NewCStableDiffusion(libraryPath string) (*CStableDiffusionImpl, error) {
	libSd, err := openLibrary(libraryPath)
	if err != nil {
		return nil, err
	}
	return &CStableDiffusionImpl{libSd: libSd}, nil
}

func (s *CStableDiffusionImpl) NewCtx(modelPath string, vaePath string, taesdPath string, loraModelDir string, vaeDecodeOnly bool, vaeTiling bool, freeParamsImmediately bool, nThreads int, wType WType, rngType RNGType, schedule Schedule) *CStableDiffusionCtx {
	ctx := C.new_sd_ctx(C.CString(modelPath), C.CString(vaePath), C.CString(taesdPath), C.CString(loraModelDir), C.bool(vaeDecodeOnly), C.bool(vaeTiling), C.bool(freeParamsImmediately), C.int(nThreads), C.int(wType), C.int(rngType), C.int(schedule))
	return &CStableDiffusionCtx{
		ctx: uintptr(ctx),
	}
}

func (s *CStableDiffusionImpl) PredictImage(ctx *CStableDiffusionCtx, prompt string, negativePrompt string, clipSkip int, cfgScale float32, width int, height int, sampleMethod SampleMethod, sampleSteps int, seed int64, batchCount int) []Image {
	images := C.txt2img_func(s.libSd, ctx.ctx, C.CString(prompt), C.CString(negativePrompt), C.int(clipSkip), C.float(cfgScale), C.int(width), C.int(height), C.int(sampleMethod), C.int(sampleSteps), C.longlong(seed), C.int(batchCount))
	return goImageSlice(uintptr(images), batchCount)
}

func (s *CStableDiffusionImpl) ImagePredictImage(ctx *CStableDiffusionCtx, img Image, prompt string, negativePrompt string, clipSkip int, cfgScale float32, width int, height int, sampleMethod SampleMethod, sampleSteps int, strength float32, seed int64, batchCount int) []Image {
	images := C.img2img_func(s.libSd, ctx.ctx, unsafe.Pointer(&img), C.CString(prompt), C.CString(negativePrompt), C.int(clipSkip), C.float(cfgScale), C.int(width), C.int(height), C.int(sampleMethod), C.int(sampleSteps), C.float(strength), C.longlong(seed), C.int(batchCount))
	return goImageSlice(uintptr(images), batchCount)
}

func (s *CStableDiffusionImpl) SetLogCallBack(cb CLogCallback) {
	C.sd_set_log_callback_func(s.libSd, unsafe.Pointer(&cb), nil)
	panic("check me")
}

func (s *CStableDiffusionImpl) GetSystemInfo() string {
	info := C.sd_get_system_info_func(s.libSd)
	return C.GoString((*C.char)(info))
}

func (s *CStableDiffusionImpl) FreeCtx(ctx *CStableDiffusionCtx) {
	C.free_sd_ctx_func(s.libSd, ctx.ctx)
	ctx = nil
	runtime.GC()
}

func (s *CStableDiffusionImpl) NewUpscalerCtx(esrganPath string, nThreads int, wType WType) *CUpScalerCtx {
	ctx := C.new_upscaler_ctx_func(s.libSd, C.CString(esrganPath), C.int(nThreads), C.int(wType))
	return &CUpScalerCtx{ctx: uintptr(ctx)}
}

func (s *CStableDiffusionImpl) FreeUpscalerCtx(ctx *CUpScalerCtx) {
	C.free_upscaler_ctx_func(s.libSd, ctx.ctx)
	ctx = nil
	runtime.GC()
}

func (s *CStableDiffusionImpl) UpscaleImage(ctx *CUpScalerCtx, img Image, upscaleFactor uint32) Image {
	imgPtr := unsafe.Pointer(&img)
	C.upscale_func(s.libSd, ctx.ctx, imgPtr, C.uint(upscaleFactor))
	panic("check me")
}

func (s *CStableDiffusionImpl) Close() error {
	if s.libSd != 0 {
		err := closeLibrary(s.libSd)
		if err != nil {
			return err
		}
		s.libSd = 0
	}
	return nil
}
