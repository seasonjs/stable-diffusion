// Copyright (c) seasonjs. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

//go:build cgo

package sd

/*
#cgo CFLAGS: -I deps
//todo: avoid warning
#ifdef  __WIN32__
#include "windows.h"
#define dlsym GetProcAddress
#else
#include <dlfcn.h>
#endif

#include <stdlib.h>
#include "deps/stable-diffusion.h"

extern goLogCallback(int level, const char* text, void* data);

typedef sd_ctx_t* (*new_sd_ctx_t)(const char* model_path,
								 const char* vae_path,
								 const char* taesd_path,
								 const char* lora_model_dir,
								 bool vae_decode_only,
								 bool vae_tiling,
								 bool free_params_immediately,
								 int n_threads,
								 enum sd_type_t wtype,
								 enum rng_type_t rng_type,
								 enum schedule_t s);

typedef sd_image_t* (*txt2img_t)(sd_ctx_t* sd_ctx,
                           const char* prompt,
                           const char* negative_prompt,
                           int clip_skip,
                           float cfg_scale,
                           int width,
                           int height,
                           enum sample_method_t sample_method,
                           int sample_steps,
                           int64_t seed,
                           int batch_count);

typedef sd_image_t* (*img2img_t)(sd_ctx_t* sd_ctx,
                           sd_image_t init_image,
                           const char* prompt,
                           const char* negative_prompt,
                           int clip_skip,
                           float cfg_scale,
                           int width,
                           int height,
                           enum sample_method_t sample_method,
                           int sample_steps,
                           float strength,
                           int64_t seed,
                           int batch_count);

typedef upscaler_ctx_t* (*new_upscaler_ctx_t)(const char* esrgan_path,
                                        int n_threads,
                                        enum sd_type_t wtype);



typedef sd_image_t (*upscale_t)(upscaler_ctx_t* upscaler_ctx, sd_image_t input_image, uint32_t upscale_factor);

typedef	void (*sd_set_log_callback_t)(sd_log_cb_t sd_log_cb, void* data);

sd_ctx_t *new_sd_ctx_func(void* handle,
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
	if (!handle) {
		return NULL;
	}
	new_sd_ctx_t new_sd_ctx =(new_sd_ctx_t) dlsym(handle, "new_sd_ctx");
	if (!new_sd_ctx) {
		return NULL;
	}
	return new_sd_ctx(modelPath, vaePath, taesdPath, loraModelDir, vaeDecodeOnly, vaeTiling, freeParamsImmediately, nThreads, wType, rngType, schedule);
}

sd_image_t *txt2img_func(void* handle,
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
	if (!handle) {
		return NULL;
	}

	txt2img_t txt2img = (txt2img_t)dlsym(handle, "txt2img");
	if (!txt2img) {
		return NULL;
	}
	return txt2img(ctx, prompt, negativePrompt, clipSkip, cfgScale, width, height, sampleMethod, sampleSteps, seed, batchCount);
}

sd_image_t *img2img_func(void* handle,
				   void *ctx,
				   sd_image_t img,
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
	if (!handle) {
		return NULL;
	}
	img2img_t img2img =(img2img_t)dlsym(handle, "img2img");
	if (!img2img) {
		return NULL;
	}
	return img2img(ctx, img, prompt, negativePrompt, clipSkip, cfgScale, width, height, sampleMethod, sampleSteps, strength, seed, batchCount);
}

void free_sd_ctx_func(void* handle, sd_ctx_t *ctx)
{
	if (!handle) {
		return;
	}
	void (*free_sd_ctx)(sd_ctx_t*) = dlsym(handle, "free_sd_ctx");
	if (!free_sd_ctx) {
		return;
	}
	if (ctx) {
	    free_sd_ctx(ctx);
    }
}

upscaler_ctx_t *new_upscaler_ctx_func(void* handle,
							const char *esrganPath,
							int nThreads,
							int wType)
{
	if (!handle) {
		return NULL;
	}
	new_upscaler_ctx_t new_upscaler_ctx =(new_upscaler_ctx_t) dlsym(handle, "new_upscaler_ctx");
	if (!new_upscaler_ctx) {
		return NULL;
	}
	return new_upscaler_ctx(esrganPath, nThreads, wType);
}

void free_upscaler_ctx_func(void* handle, upscaler_ctx_t *ctx)
{
	if (!handle) {
		return;
	}
	void (*free_upscaler_ctx)(upscaler_ctx_t*) = dlsym(handle, "free_upscaler_ctx");
	if (!free_upscaler_ctx) {
		return;
	}
	if (ctx) {
		free_upscaler_ctx(ctx);
	}
}

sd_image_t upscale_func(void* handle,
				   upscaler_ctx_t *ctx,
				   sd_image_t img,
                   int upscaleFactor)
{
	sd_image_t empty_image;
	empty_image.width=0;
	empty_image.height=0;
	empty_image.channel=0;
	empty_image.data=NULL;
	if (!handle) {
		return empty_image;
	}
	upscale_t upscale=(upscale_t) dlsym(handle, "upscale");
	if (!upscale) {
		return empty_image;
	}
	return upscale(ctx, img, upscaleFactor);
}

const char *sd_get_system_info_func(void* handle)
{
	if (!handle) {
		return NULL;
	}
	const char * (*sd_get_system_info)() = dlsym(handle, "sd_get_system_info");
	if (!sd_get_system_info) {
		return NULL;
	}
	return sd_get_system_info();
}

void sd_set_log_callback_func (void* handle, sd_log_cb_t sd_log_cb)
{
	if (!handle) {
		return;
	}
	sd_set_log_callback_t sd_set_log_callback = dlsym(handle, "sd_set_log_callback");
	if (!sd_set_log_callback) {
		return;
	}
	sd_set_log_callback(sd_log_cb, NULL);
}
*/
import "C"

import (
	"runtime"
	"unsafe"
)

var logCallback CLogCallback

type CStableDiffusionImpl struct {
	libSd unsafe.Pointer
}

func NewCStableDiffusion(libraryPath string) (*CStableDiffusionImpl, error) {
	libSd, err := openLibrary(libraryPath)
	if err != nil {
		return nil, err
	}
	return &CStableDiffusionImpl{libSd: *(*unsafe.Pointer)(unsafe.Pointer(&libSd))}, nil
}

func (s *CStableDiffusionImpl) NewCtx(modelPath string, vaePath string, taesdPath string, loraModelDir string, vaeDecodeOnly bool, vaeTiling bool, freeParamsImmediately bool, nThreads int, wType WType, rngType RNGType, schedule Schedule) *CStableDiffusionCtx {
	ctx := C.new_sd_ctx_func(s.libSd, C.CString(modelPath), C.CString(vaePath), C.CString(taesdPath), C.CString(loraModelDir), C.bool(vaeDecodeOnly), C.bool(vaeTiling), C.bool(freeParamsImmediately), C.int(nThreads), C.int(wType), C.int(rngType), C.int(schedule))
	return &CStableDiffusionCtx{
		cgoCtx: unsafe.Pointer(ctx),
	}
}

func (s *CStableDiffusionImpl) PredictImage(ctx *CStableDiffusionCtx, prompt string, negativePrompt string, clipSkip int, cfgScale float32, width int, height int, sampleMethod SampleMethod, sampleSteps int, seed int64, batchCount int) []Image {
	images := C.txt2img_func(s.libSd, ctx.cgoCtx, C.CString(prompt), C.CString(negativePrompt), C.int(clipSkip), C.float(cfgScale), C.int(width), C.int(height), C.int(sampleMethod), C.int(sampleSteps), C.longlong(seed), C.int(batchCount))
	return convertCArrayToGoSlice(images, batchCount)
}

func (s *CStableDiffusionImpl) ImagePredictImage(ctx *CStableDiffusionCtx, img Image, prompt string, negativePrompt string, clipSkip int, cfgScale float32, width int, height int, sampleMethod SampleMethod, sampleSteps int, strength float32, seed int64, batchCount int) []Image {
	images := C.img2img_func(s.libSd, ctx.cgoCtx, convertGoStructToCStruct(img), C.CString(prompt), C.CString(negativePrompt), C.int(clipSkip), C.float(cfgScale), C.int(width), C.int(height), C.int(sampleMethod), C.int(sampleSteps), C.float(strength), C.longlong(seed), C.int(batchCount))
	return convertCArrayToGoSlice(images, batchCount)
}

func (s *CStableDiffusionImpl) SetLogCallBack(cb CLogCallback) {
	logCallback = cb
	C.sd_set_log_callback_func(s.libSd, C.sd_log_cb_t(goLogCallback))
}

func (s *CStableDiffusionImpl) GetSystemInfo() string {
	info := C.sd_get_system_info_func(s.libSd)
	return C.GoString((*C.char)(info))
}

func (s *CStableDiffusionImpl) FreeCtx(ctx *CStableDiffusionCtx) {
	C.free_sd_ctx_func(s.libSd, ctx.cgoCtx)
	ctx = nil
	runtime.GC()
}

func (s *CStableDiffusionImpl) NewUpscalerCtx(esrganPath string, nThreads int, wType WType) *CUpScalerCtx {
	ctx := C.new_upscaler_ctx_func(s.libSd, C.CString(esrganPath), C.int(nThreads), C.int(wType))
	return &CUpScalerCtx{ctx: uintptr(ctx)}
}

func (s *CStableDiffusionImpl) FreeUpscalerCtx(ctx *CUpScalerCtx) {
	C.free_upscaler_ctx_func(s.libSd, ctx.cgoCtx)
	ctx = nil
	runtime.GC()
}

func (s *CStableDiffusionImpl) UpscaleImage(ctx *CUpScalerCtx, img Image, upscaleFactor uint32) Image {
	result := C.upscale_func(s.libSd, ctx.cgoCtx, convertGoStructToCStruct(img), C.uint(upscaleFactor))
	return convertCStructToGoStruct(result)
}

func (s *CStableDiffusionImpl) Close() error {
	if s.libSd != nil {
		err := closeLibrary(uintptr(s.libSd))
		if err != nil {
			return err
		}
		s.libSd = nil
	}
	return nil
}

func convertCStructToGoStruct(cStruct C.sd_image_t) Image {
	defer C.free(unsafe.Pointer(cStruct.data))
	defer C.free(unsafe.Pointer(&cStruct))
	goImage := Image{
		Width:   uint32(cStruct.width),
		Height:  uint32(cStruct.height),
		Channel: uint32(cStruct.channel),
		Data:    C.GoBytes(unsafe.Pointer(cStruct.data), C.int(cStruct.width*cStruct.height*cStruct.channel)),
	}
	return goImage
}

func convertCArrayToGoSlice(cArray *C.sd_image_t, length int) []Image {
	defer C.free(unsafe.Pointer(cArray))
	cSlice := unsafe.Slice(cArray, length)
	goSlice := make([]Image, length)
	for i, cStruct := range cSlice {
		goSlice[i] = convertCStructToGoStruct(cStruct)
	}
	return goSlice
}

func convertGoStructToCStruct(goStruct Image) C.sd_image_t {
	cStruct := C.sd_image_t{
		width:   C.uint32_t(goStruct.Width),
		height:  C.uint32_t(goStruct.Height),
		channel: C.uint32_t(goStruct.Channel),
		data:    nil,
	}

	if len(goStruct.Data) > 0 {
		cStruct.data = (*C.uint8_t)(C.CBytes(goStruct.Data))
	}
	return cStruct
}

//export goLogCallback
func goLogCallback(level C.enum_sd_log_level_t, text *C.char, data unsafe.Pointer) {
	goMessage := C.GoString(text)
	if logCallback != nil {
		logCallback(LogLevel(level), goMessage)
	}
}
