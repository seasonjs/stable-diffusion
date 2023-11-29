// Copyright (c) seasonjs. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package sd

import (
	"github.com/ebitengine/purego"
	"unsafe"
)

type SDLogLevel string

type RNGType string

type SampleMethod string

type Schedule string

const (
	DEBUG SDLogLevel = "DEBUG"
	INFO             = "INFO"
	WARN             = "WARN"
	ERROR            = "ERROR"
)

const (
	STD_DEFAULT_RNG RNGType = "STD_DEFAULT_RNG"
	CUDA_RNG                = "CUDA_RNG"
)

const (
	EULER_A          SampleMethod = "EULER_A"
	EULER                         = "EULER"
	HEUN                          = "HEUN"
	DPM2                          = "DPM2"
	DPMPP2S_A                     = "DPMPP2S_A"
	DPMPP2M                       = "DPMPP2M"
	DPMPP2Mv2                     = "DPMPP2Mv2"
	LCM                           = "LCM"
	N_SAMPLE_METHODS              = "N_SAMPLE_METHODS"
)

const (
	DEFAULT     Schedule = "DEFAULT"
	DISCRETE             = "DISCRETE"
	KARRAS               = "KARRAS"
	N_SCHEDULES          = "N_SCHEDULES"
)

const (
	cStableDiffusionFullDefaultParamsRef        = "stable_diffusion_full_default_params_ref"
	cStableDiffusionFullParamsSetNegativePrompt = "stable_diffusion_full_params_set_negative_prompt"
	cStableDiffusionFullParamsSetCfgScale       = "stable_diffusion_full_params_set_cfg_scale"
	cStableDiffusionFullParamsSetWidth          = "stable_diffusion_full_params_set_width"
	cStableDiffusionFullParamsSetHeight         = "stable_diffusion_full_params_set_height"
	cStableDiffusionFullParamsSetSampleMethod   = "stable_diffusion_full_params_set_sample_method"
	cStableDiffusionFullParamsSetSampleSteps    = "stable_diffusion_full_params_set_sample_steps"
	cStableDiffusionFullParamsSetSeed           = "stable_diffusion_full_params_set_seed"
	cStableDiffusionFullParamsSetBatchCount     = "stable_diffusion_full_params_set_batch_count"
	cStableDiffusionFullParamsSetStrength       = "stable_diffusion_full_params_set_strength"

	cStableDiffusionInit              = "stable_diffusion_init"
	cStableDiffusionLoadFromFile      = "stable_diffusion_load_from_file"
	cStableDiffusionPredictImage      = "stable_diffusion_predict_image"
	cStableDiffusionImagePredictImage = "stable_diffusion_image_predict_image"
	cStableDiffusionSetLogLevel       = "stable_diffusion_set_log_level"
	cStableDiffusionGetSystemInfo     = "stable_diffusion_get_system_info"
	cStableDiffusionFree              = "stable_diffusion_free"
	cStableDiffusionFreeBuffer        = "stable_diffusion_free_buffer"
	cStableDiffusionFreeFullParams    = "stable_diffusion_free_full_params"
)

type StableDiffusionFullParams struct {
	params         uintptr
	negativePrompt string
	cfgScale       float32
	width          int
	height         int
	sampleMethod   SampleMethod
	sampleSteps    int
	seed           int64
	batchCount     int
	strength       float32
}

type StableDiffusionCtx struct {
	ctx uintptr
}

type CStableDiffusion interface {
	StableDiffusionFullDefaultParamsRef() *StableDiffusionFullParams
	StableDiffusionFullParamsSetNegativePrompt(params *StableDiffusionFullParams, negativePrompt string)
	StableDiffusionFullParamsSetCfgScale(params *StableDiffusionFullParams, cfgScale float32)
	StableDiffusionFullParamsSetWidth(params *StableDiffusionFullParams, width int)
	StableDiffusionFullParamsSetHeight(params *StableDiffusionFullParams, height int)
	StableDiffusionFullParamsSetSampleMethod(params *StableDiffusionFullParams, sampleMethod SampleMethod)
	StableDiffusionFullParamsSetSampleSteps(params *StableDiffusionFullParams, sampleSteps int)
	StableDiffusionFullParamsSetSeed(params *StableDiffusionFullParams, seed int64)
	StableDiffusionFullParamsSetBatchCount(params *StableDiffusionFullParams, batchCount int)
	StableDiffusionFullParamsSetStrength(params *StableDiffusionFullParams, strength float32)

	StableDiffusionInit(nThreads int, vaeDecodeOnly bool, freeParamsImmediately bool, loraModelDir string, rngType RNGType) *StableDiffusionCtx
	StableDiffusionLoadFromFile(ctx *StableDiffusionCtx, filePath string, schedule Schedule)
	StableDiffusionPredictImage(ctx *StableDiffusionCtx, params *StableDiffusionFullParams, prompt string) []byte
	StableDiffusionImagePredictImage(ctx *StableDiffusionCtx, params *StableDiffusionFullParams, initImage []byte, prompt string) []byte
	StableDiffusionSetLogLevel(level SDLogLevel)
	StableDiffusionGetSystemInfo() string
	StableDiffusionFree(ctx *StableDiffusionCtx)
	StableDiffusionFreeBuffer(buffer uintptr)
	StableDiffusionFreeFullParams(params *StableDiffusionFullParams)
}

type CStableDiffusionImpl struct {
	libSd uintptr

	cStableDiffusionFullDefaultParamsRef        func() uintptr
	cStableDiffusionFullParamsSetNegativePrompt func(params uintptr, negative_prompt string)
	cStableDiffusionFullParamsSetCfgScale       func(params uintptr, cfg_scale float32)
	cStableDiffusionFullParamsSetWidth          func(params uintptr, width int)
	cStableDiffusionFullParamsSetHeight         func(params uintptr, height int)
	cStableDiffusionFullParamsSetSampleMethod   func(params uintptr, sampleMethod string)
	cStableDiffusionFullParamsSetSampleSteps    func(params uintptr, sampleSteps int)
	cStableDiffusionFullParamsSetSeed           func(params uintptr, seed int64)
	cStableDiffusionFullParamsSetBatchCount     func(params uintptr, batchCount int)
	cStableDiffusionFullParamsSetStrength       func(params uintptr, strength float32)

	cStableDiffusionInit              func(nThreads int, vaeDecodeOnly bool, freeParamsImmediately bool, loraModelDir string, rngType string) uintptr
	cStableDiffusionLoadFromFile      func(ctx uintptr, filePath string, schedule string)
	cStableDiffusionPredictImage      func(ctx uintptr, params uintptr, prompt string) *byte
	cStableDiffusionImagePredictImage func(ctx uintptr, params uintptr, initImage *byte, prompt string) *byte
	cStableDiffusionSetLogLevel       func(level string)
	cStableDiffusionGetSystemInfo     func() string
	cStableDiffusionFree              func(options uintptr)
	cStableDiffusionFreeBuffer        func(options uintptr)
	cStableDiffusionFreeFullParams    func(options uintptr)
}

func NewCStableDiffusion(libraryPath string) (CStableDiffusion, error) {
	libSd, err := openLibrary(libraryPath)
	if err != nil {
		return nil, err
	}
	var (
		stableDiffusionFullDefaultParamsRef        func() uintptr
		stableDiffusionFullParamsSetNegativePrompt func(params uintptr, negative_prompt string)
		stableDiffusionFullParamsSetCfgScale       func(params uintptr, cfg_scale float32)
		stableDiffusionFullParamsSetWidth          func(params uintptr, width int)
		stableDiffusionFullParamsSetHeight         func(params uintptr, height int)
		stableDiffusionFullParamsSetSampleMethod   func(params uintptr, sampleMethod string)
		stableDiffusionFullParamsSetSampleSteps    func(params uintptr, sampleSteps int)
		stableDiffusionFullParamsSetSeed           func(params uintptr, seed int64)
		stableDiffusionFullParamsSetBatchCount     func(params uintptr, batchCount int)
		stableDiffusionFullParamsSetStrength       func(params uintptr, strength float32)

		stableDiffusionInit              func(nThreads int, vaeDecodeOnly bool, freeParamsImmediately bool, loraModelDir string, rngType string) uintptr
		stableDiffusionLoadFromFile      func(ctx uintptr, filePath string, schedule string)
		stableDiffusionPredictImage      func(ctx uintptr, params uintptr, prompt string) *byte
		stableDiffusionImagePredictImage func(ctx uintptr, params uintptr, initImage *byte, prompt string) *byte
		stableDiffusionSetLogLevel       func(level string)
		stableDiffusionGetSystemInfo     func() string
		stableDiffusionFree              func(options uintptr)
		stableDiffusionFreeBuffer        func(options uintptr)
		stableDiffusionFreeFullParams    func(options uintptr)
	)
	purego.RegisterLibFunc(&stableDiffusionFullDefaultParamsRef, libSd, cStableDiffusionFullDefaultParamsRef)
	purego.RegisterLibFunc(&stableDiffusionFullParamsSetNegativePrompt, libSd, cStableDiffusionFullParamsSetNegativePrompt)
	purego.RegisterLibFunc(&stableDiffusionFullParamsSetCfgScale, libSd, cStableDiffusionFullParamsSetCfgScale)
	purego.RegisterLibFunc(&stableDiffusionFullParamsSetWidth, libSd, cStableDiffusionFullParamsSetWidth)
	purego.RegisterLibFunc(&stableDiffusionFullParamsSetHeight, libSd, cStableDiffusionFullParamsSetHeight)
	purego.RegisterLibFunc(&stableDiffusionFullParamsSetSampleMethod, libSd, cStableDiffusionFullParamsSetSampleMethod)
	purego.RegisterLibFunc(&stableDiffusionFullParamsSetSampleSteps, libSd, cStableDiffusionFullParamsSetSampleSteps)
	purego.RegisterLibFunc(&stableDiffusionFullParamsSetSeed, libSd, cStableDiffusionFullParamsSetSeed)
	purego.RegisterLibFunc(&stableDiffusionFullParamsSetBatchCount, libSd, cStableDiffusionFullParamsSetBatchCount)
	purego.RegisterLibFunc(&stableDiffusionFullParamsSetStrength, libSd, cStableDiffusionFullParamsSetStrength)

	purego.RegisterLibFunc(&stableDiffusionInit, libSd, cStableDiffusionInit)
	purego.RegisterLibFunc(&stableDiffusionLoadFromFile, libSd, cStableDiffusionLoadFromFile)
	purego.RegisterLibFunc(&stableDiffusionPredictImage, libSd, cStableDiffusionPredictImage)
	purego.RegisterLibFunc(&stableDiffusionImagePredictImage, libSd, cStableDiffusionImagePredictImage)
	purego.RegisterLibFunc(&stableDiffusionSetLogLevel, libSd, cStableDiffusionSetLogLevel)
	purego.RegisterLibFunc(&stableDiffusionGetSystemInfo, libSd, cStableDiffusionGetSystemInfo)
	purego.RegisterLibFunc(&stableDiffusionFree, libSd, cStableDiffusionFree)
	purego.RegisterLibFunc(&stableDiffusionFreeBuffer, libSd, cStableDiffusionFreeBuffer)
	purego.RegisterLibFunc(&stableDiffusionFreeFullParams, libSd, cStableDiffusionFreeFullParams)

	return &CStableDiffusionImpl{
		libSd,
		stableDiffusionFullDefaultParamsRef,
		stableDiffusionFullParamsSetNegativePrompt,
		stableDiffusionFullParamsSetCfgScale,
		stableDiffusionFullParamsSetWidth,
		stableDiffusionFullParamsSetHeight,
		stableDiffusionFullParamsSetSampleMethod,
		stableDiffusionFullParamsSetSampleSteps,
		stableDiffusionFullParamsSetSeed,
		stableDiffusionFullParamsSetBatchCount,
		stableDiffusionFullParamsSetStrength,

		stableDiffusionInit,
		stableDiffusionLoadFromFile,
		stableDiffusionPredictImage,
		stableDiffusionImagePredictImage,
		stableDiffusionSetLogLevel,
		stableDiffusionGetSystemInfo,
		stableDiffusionFree,
		stableDiffusionFreeBuffer,
		stableDiffusionFreeFullParams,
	}, nil
}

func (c *CStableDiffusionImpl) StableDiffusionFullDefaultParamsRef() *StableDiffusionFullParams {
	return &StableDiffusionFullParams{
		params: c.cStableDiffusionFullDefaultParamsRef(),
	}
}

func (c *CStableDiffusionImpl) StableDiffusionFullParamsSetNegativePrompt(params *StableDiffusionFullParams, negativePrompt string) {
	c.cStableDiffusionFullParamsSetNegativePrompt(params.params, negativePrompt)
	params.negativePrompt = negativePrompt
}

func (c *CStableDiffusionImpl) StableDiffusionFullParamsSetCfgScale(params *StableDiffusionFullParams, cfgScale float32) {
	c.cStableDiffusionFullParamsSetCfgScale(params.params, cfgScale)
	params.cfgScale = cfgScale
}

func (c *CStableDiffusionImpl) StableDiffusionFullParamsSetWidth(params *StableDiffusionFullParams, width int) {
	c.cStableDiffusionFullParamsSetWidth(params.params, width)
	params.width = width
}

func (c *CStableDiffusionImpl) StableDiffusionFullParamsSetHeight(params *StableDiffusionFullParams, height int) {
	c.cStableDiffusionFullParamsSetHeight(params.params, height)
	params.height = height
}

func (c *CStableDiffusionImpl) StableDiffusionFullParamsSetSampleMethod(params *StableDiffusionFullParams, sampleMethod SampleMethod) {
	c.cStableDiffusionFullParamsSetSampleMethod(params.params, string(sampleMethod))
	params.sampleMethod = sampleMethod
}

func (c *CStableDiffusionImpl) StableDiffusionFullParamsSetSampleSteps(params *StableDiffusionFullParams, sampleSteps int) {
	c.cStableDiffusionFullParamsSetSampleSteps(params.params, sampleSteps)
	params.sampleSteps = sampleSteps
}

func (c *CStableDiffusionImpl) StableDiffusionFullParamsSetSeed(params *StableDiffusionFullParams, seed int64) {
	c.cStableDiffusionFullParamsSetSeed(params.params, seed)
	params.seed = seed
}

func (c *CStableDiffusionImpl) StableDiffusionFullParamsSetBatchCount(params *StableDiffusionFullParams, batchCount int) {
	c.cStableDiffusionFullParamsSetBatchCount(params.params, batchCount)
	params.batchCount = batchCount
}

func (c *CStableDiffusionImpl) StableDiffusionFullParamsSetStrength(params *StableDiffusionFullParams, strength float32) {
	c.cStableDiffusionFullParamsSetStrength(params.params, strength)
	params.strength = strength
}

func (c *CStableDiffusionImpl) StableDiffusionInit(nThreads int, vaeDecodeOnly bool, freeParamsImmediately bool, loraModelDir string, rngType RNGType) *StableDiffusionCtx {
	return &StableDiffusionCtx{
		ctx: c.cStableDiffusionInit(nThreads, vaeDecodeOnly, freeParamsImmediately, loraModelDir, string(rngType)),
	}
}

func (c *CStableDiffusionImpl) StableDiffusionLoadFromFile(ctx *StableDiffusionCtx, filePath string, schedule Schedule) {
	c.cStableDiffusionLoadFromFile(ctx.ctx, filePath, string(schedule))
}

func (c *CStableDiffusionImpl) StableDiffusionPredictImage(ctx *StableDiffusionCtx, params *StableDiffusionFullParams, prompt string) []byte {
	data := c.cStableDiffusionPredictImage(ctx.ctx, params.params, prompt)
	return unsafe.Slice(data, params.width*params.height*3*params.batchCount)
}

func (c *CStableDiffusionImpl) StableDiffusionImagePredictImage(ctx *StableDiffusionCtx, params *StableDiffusionFullParams, initImage []byte, prompt string) []byte {
	data := c.cStableDiffusionImagePredictImage(ctx.ctx, params.params, &initImage[0], prompt)
	return unsafe.Slice(data, params.width*params.height*3)
}

func (c *CStableDiffusionImpl) StableDiffusionSetLogLevel(level SDLogLevel) {
	c.cStableDiffusionSetLogLevel(string(level))
}

func (c *CStableDiffusionImpl) StableDiffusionGetSystemInfo() string {
	return c.cStableDiffusionGetSystemInfo()
}

func (c *CStableDiffusionImpl) StableDiffusionFree(ctx *StableDiffusionCtx) {
	c.cStableDiffusionFree(ctx.ctx)
}

func (c *CStableDiffusionImpl) StableDiffusionFreeBuffer(buffer uintptr) {
	c.cStableDiffusionFreeBuffer(buffer)
}

func (c *CStableDiffusionImpl) StableDiffusionFreeFullParams(params *StableDiffusionFullParams) {
	c.cStableDiffusionFreeFullParams(params.params)
}
