// Copyright (c) seasonjs. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package sd

import (
	"encoding/base64"
	"errors"
	"github.com/ebitengine/purego"
	"runtime"
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
	cNewSdTxt2imgOptions      = "new_sd_txt2img_options"
	cSetTxt2imgPrompt         = "set_txt2img_prompt"
	cSetTxt2imgNegativePrompt = "set_txt2img_negative_prompt"
	cSetTxt2imgCfgScale       = "set_txt2img_cfg_scale"
	cSetTxt2imgSize           = "set_txt2img_size"
	cSetTxt2imgSampleMethod   = "set_txt2img_sample_method"
	cSetTxt2imgSampleSteps    = "set_txt2img_sample_steps"
	cSetTxt2imgSeed           = "set_txt2img_seed"

	cNewSdImg2imgOptions      = "new_sd_img2img_options"
	cSetImg2imgInitImg        = "set_img2img_init_img"
	cSetImg2imgPrompt         = "set_img2img_prompt"
	cSetImg2imgNegativePrompt = "set_img2img_negative_prompt"
	cSetImg2imgCfgScale       = "set_img2img_cfg_scale"
	cSetImg2imgSize           = "set_img2img_size"
	cSetImg2imgSampleMethod   = "set_img2img_sample_method"
	cSetImg2imgSampleSteps    = "set_img2img_sample_steps"
	cSetImg2imgStrength       = "set_img2img_strength"
	cSetImg2imgSeed           = "set_img2img_seed"

	cCreateStableDiffusion  = "create_stable_diffusion"
	cDestroyStableDiffusion = "destroy_stable_diffusion"
	cLoadFromFile           = "load_from_file"
	cTxt2img                = "txt2img"
	cImg2img                = "img2img"

	cSetStableDiffusionLogLevel   = "set_stable_diffusion_log_level"
	cGetStableDiffusionSystemInfo = "get_stable_diffusion_system_info"
	cFreeBuffer                   = "free_buffer"
)

type CStableDiffusion struct {
	libSd uintptr

	cNewSdTxt2imgOptions      func() uintptr
	cSetTxt2imgPrompt         func(options uintptr, prompt string)
	cSetTxt2imgNegativePrompt func(options uintptr, negativePrompt string)
	cSetTxt2imgCfgScale       func(options uintptr, scale float32)
	cSetTxt2imgSize           func(options uintptr, width, height int)
	cSetTxt2imgSampleMethod   func(options uintptr, sampleMethod string)
	cSetTxt2imgSampleSteps    func(options uintptr, sampleSteps int)
	cSetTxt2imgSeed           func(options uintptr, seed int64)

	cNewSdImg2imgOptions      func() uintptr
	cSetImg2imgInitImg        func(options uintptr, base64Str string)
	cSetImg2imgPrompt         func(options uintptr, prompt string)
	cSetImg2imgNegativePrompt func(options uintptr, negativePrompt string)
	cSetImg2imgCfgScale       func(options uintptr, cfgScale float32)
	cSetImg2imgSize           func(options uintptr, width, height int)
	cSetImg2imgSampleMethod   func(options uintptr, sampleMethod string)
	cSetImg2imgSampleSteps    func(options uintptr, sampleSteps int)
	cSetImg2imgStrength       func(options uintptr, strength float32)
	cSetImg2imgSeed           func(options uintptr, seed int64)

	cCreateStableDiffusion  func(nThreads int, vaeDecodeOnly bool, freeParamsImmediately bool, loraModelDir string, rngType string) uintptr
	cDestroyStableDiffusion func(sd uintptr)
	cLoadFromFile           func(sd uintptr, path string, schedule string)
	cTxt2img                func(sd uintptr, options uintptr, byteSize *int64) *byte
	cImg2img                func(sd uintptr, options uintptr, byte2 *int64) *byte

	cSetStableDiffusionLogLevel   func(level string)
	cGetStableDiffusionSystemInfo func() string

	cFreeBuffer func(buffer uintptr)
}

func NewCStableDiffusion(libraryPath string) (*CStableDiffusion, error) {
	libSd, err := openLibrary(libraryPath)
	if err != nil {
		return nil, err
	}
	var (
		newSdTxt2imgOptions      func() uintptr
		setTxt2imgPrompt         func(options uintptr, prompt string)
		setTxt2imgNegativePrompt func(options uintptr, negativePrompt string)
		setTxt2imgCfgScale       func(options uintptr, cfgScale float32)
		setTxt2imgSize           func(options uintptr, width, height int)
		setTxt2imgSampleMethod   func(options uintptr, sampleMethod string)
		setTxt2imgSampleSteps    func(options uintptr, sampleSteps int)
		setTxt2imgSeed           func(options uintptr, seed int64)

		newSdImg2imgOptions      func() uintptr
		setImg2imgInitImg        func(options uintptr, base64Str string)
		setImg2imgPrompt         func(options uintptr, prompt string)
		setImg2imgNegativePrompt func(options uintptr, negativePrompt string)
		setImg2imgCfgScale       func(options uintptr, cfgScale float32)
		setImg2imgSize           func(options uintptr, width, height int)
		setImg2imgSampleMethod   func(options uintptr, sampleMethod string)
		setImg2imgSampleSteps    func(options uintptr, sampleSteps int)
		setImg2imgStrength       func(options uintptr, strength float32)
		setImg2imgSeed           func(options uintptr, seed int64)

		createStableDiffusion  func(nThreads int, vaeDecodeOnly bool, freeParamsImmediately bool, loraModelDir string, rngType string) uintptr
		destroyStableDiffusion func(sd uintptr)
		loadFromFile           func(sd uintptr, path string, schedule string)
		txt2img                func(sd uintptr, options uintptr, byteSize *int64) *byte
		img2img                func(sd uintptr, options uintptr, byte2 *int64) *byte

		setStableDiffusionLogLevel   func(level string)
		getStableDiffusionSystemInfo func() string

		freeBuffer func(buffer uintptr)
	)
	purego.RegisterLibFunc(&newSdTxt2imgOptions, libSd, cNewSdTxt2imgOptions)
	purego.RegisterLibFunc(&setTxt2imgPrompt, libSd, cSetTxt2imgPrompt)
	purego.RegisterLibFunc(&setTxt2imgNegativePrompt, libSd, cSetTxt2imgNegativePrompt)
	purego.RegisterLibFunc(&setTxt2imgCfgScale, libSd, cSetTxt2imgCfgScale)
	purego.RegisterLibFunc(&setTxt2imgSize, libSd, cSetTxt2imgSize)
	purego.RegisterLibFunc(&setTxt2imgSampleMethod, libSd, cSetTxt2imgSampleMethod)
	purego.RegisterLibFunc(&setTxt2imgSampleSteps, libSd, cSetTxt2imgSampleSteps)
	purego.RegisterLibFunc(&setTxt2imgSeed, libSd, cSetTxt2imgSeed)

	purego.RegisterLibFunc(&newSdImg2imgOptions, libSd, cNewSdImg2imgOptions)
	purego.RegisterLibFunc(&setImg2imgInitImg, libSd, cSetImg2imgInitImg)
	purego.RegisterLibFunc(&setImg2imgPrompt, libSd, cSetImg2imgPrompt)
	purego.RegisterLibFunc(&setImg2imgNegativePrompt, libSd, cSetImg2imgNegativePrompt)
	purego.RegisterLibFunc(&setImg2imgCfgScale, libSd, cSetImg2imgCfgScale)
	purego.RegisterLibFunc(&setImg2imgSize, libSd, cSetImg2imgSize)
	purego.RegisterLibFunc(&setImg2imgSampleMethod, libSd, cSetImg2imgSampleMethod)
	purego.RegisterLibFunc(&setImg2imgSampleSteps, libSd, cSetImg2imgSampleSteps)
	purego.RegisterLibFunc(&setImg2imgStrength, libSd, cSetImg2imgStrength)
	purego.RegisterLibFunc(&setImg2imgSeed, libSd, cSetImg2imgSeed)

	purego.RegisterLibFunc(&createStableDiffusion, libSd, cCreateStableDiffusion)
	purego.RegisterLibFunc(&destroyStableDiffusion, libSd, cDestroyStableDiffusion)
	purego.RegisterLibFunc(&loadFromFile, libSd, cLoadFromFile)
	purego.RegisterLibFunc(&txt2img, libSd, cTxt2img)
	purego.RegisterLibFunc(&img2img, libSd, cImg2img)

	purego.RegisterLibFunc(&setStableDiffusionLogLevel, libSd, cSetStableDiffusionLogLevel)
	purego.RegisterLibFunc(&getStableDiffusionSystemInfo, libSd, cGetStableDiffusionSystemInfo)

	purego.RegisterLibFunc(&freeBuffer, libSd, cFreeBuffer)

	return &CStableDiffusion{
		libSd,

		newSdTxt2imgOptions,
		setTxt2imgPrompt,
		setTxt2imgNegativePrompt,
		setTxt2imgCfgScale,
		setTxt2imgSize,
		setTxt2imgSampleMethod,
		setTxt2imgSampleSteps,
		setTxt2imgSeed,

		newSdImg2imgOptions,
		setImg2imgInitImg,
		setImg2imgPrompt,
		setImg2imgNegativePrompt,
		setImg2imgCfgScale,
		setImg2imgSize,
		setImg2imgSampleMethod,
		setImg2imgSampleSteps,
		setImg2imgStrength,
		setImg2imgSeed,

		createStableDiffusion,
		destroyStableDiffusion,
		loadFromFile,
		txt2img,
		img2img,

		setStableDiffusionLogLevel,
		getStableDiffusionSystemInfo,
		freeBuffer,
	}, nil
}

func (cSD *CStableDiffusion) NewStableDiffusionCtx(nThreads int, vaeDecodeOnly bool, freeParamsImmediately bool, loraModelDir string, rngType RNGType) *CSDCtx {
	ctx := cSD.cCreateStableDiffusion(nThreads, vaeDecodeOnly, freeParamsImmediately, loraModelDir, string(rngType))
	return &CSDCtx{ctx: ctx, csd: cSD}
}

func (cSD *CStableDiffusion) StableDiffusionSetLogLevel(level SDLogLevel) {
	cSD.cSetStableDiffusionLogLevel(string(level))
}

func (cSD *CStableDiffusion) StableDiffusionGetSystemInfo() string {
	return cSD.cGetStableDiffusionSystemInfo()
}

type CSDCtx struct {
	csd *CStableDiffusion
	ctx uintptr
}

func (c *CSDCtx) StableDiffusionLoadFromFile(path string, schedule Schedule) {
	c.csd.cLoadFromFile(c.ctx, path, string(schedule))
}

func (c *CSDCtx) StableDiffusionTextToImage(prompt string, negativePrompt string, cfgScale float32, width int, height int, sampleMethod SampleMethod, sampleSteps int, seed int64) ([]byte, error) {
	if width <= 0 {
		return nil, errors.New("width must be greater than 0")
	}
	if height <= 0 {
		return nil, errors.New("height must be greater than 0")
	}
	options := c.csd.cNewSdImg2imgOptions()
	c.csd.cSetTxt2imgPrompt(options, prompt)
	c.csd.cSetTxt2imgNegativePrompt(options, negativePrompt)
	c.csd.cSetTxt2imgCfgScale(options, cfgScale)
	c.csd.cSetTxt2imgSize(options, width, height)
	c.csd.cSetTxt2imgSampleMethod(options, string(sampleMethod))
	c.csd.cSetTxt2imgSampleSteps(options, sampleSteps)
	c.csd.cSetTxt2imgSeed(options, seed)
	minSize := int64(width * height * 3)
	var size int64
	runtime.KeepAlive(size)
	output := c.csd.cTxt2img(c.ctx, options, &size)
	if size < minSize {
		size = minSize
	}
	data := unsafe.Slice(output, size)
	size = 0
	c.csd.cFreeBuffer(uintptr(unsafe.Pointer(output)))
	return data, nil
}

func (c *CSDCtx) StableDiffusionImageToImage(initImg []byte, prompt string, negativePrompt string, cfgScale float32, width int, height int, sampleMethod SampleMethod, sampleSteps int, strength float32, seed int64) ([]byte, error) {
	if width <= 0 || width%64 != 0 {
		return nil, errors.New("width must be greater than 0 and must be a multiple of 64")
	}
	if height <= 0 || height%64 != 0 {
		return nil, errors.New("height must be greater than 0 and must be a multiple of 64")
	}
	options := c.csd.cNewSdImg2imgOptions()
	c.csd.cSetImg2imgInitImg(options, base64.StdEncoding.EncodeToString(initImg))
	c.csd.cSetImg2imgPrompt(options, prompt)
	c.csd.cSetImg2imgNegativePrompt(options, negativePrompt)
	c.csd.cSetImg2imgCfgScale(options, cfgScale)
	c.csd.cSetImg2imgSize(options, width, height)
	c.csd.cSetImg2imgSampleMethod(options, string(sampleMethod))
	c.csd.cSetImg2imgSampleSteps(options, sampleSteps)
	c.csd.cSetImg2imgStrength(options, strength)
	c.csd.cSetImg2imgSeed(options, seed)
	minSize := int64(width * height * 3)
	var size int64
	runtime.KeepAlive(size)
	output := c.csd.cImg2img(c.ctx, options, &size)
	if size < minSize {
		size = minSize
	}
	data := unsafe.Slice(output, size)
	size = 0
	c.csd.cFreeBuffer(uintptr(unsafe.Pointer(output)))
	return data, nil
}

func (c *CSDCtx) Close() {
	if c.ctx != 0 {
		c.csd.cDestroyStableDiffusion(c.ctx)
	}
	c.ctx = 0

}
