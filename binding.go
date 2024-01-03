// Copyright (c) seasonjs. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package sd

import (
	"github.com/ebitengine/purego"
	"unsafe"
)

type LogLevel int

type RNGType int

type SampleMethod int

type Schedule int

type WType int

const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
)

const (
	STD_DEFAULT_RNG RNGType = iota
	CUDA_RNG
)

const (
	EULER_A SampleMethod = iota
	EULER
	HEUN
	DPM2
	DPMPP2S_A
	DPMPP2M
	DPMPP2Mv2
	LCM
	N_SAMPLE_METHODS
)

const (
	DEFAULT Schedule = iota
	DISCRETE
	KARRAS
	N_SCHEDULES
)

const (
	F32   WType = 0
	F16         = 1
	Q4_0        = 2
	Q4_1        = 3
	Q5_0        = 6
	Q5_1        = 7
	Q8_0        = 8
	Q8_1        = 9
	Q2_K        = 10
	Q3_K        = 11
	Q4_K        = 12
	Q5_K        = 13
	Q6_K        = 14
	Q8_K        = 15
	I8          = 16
	I16         = 17
	I32         = 18
	COUNT       = 19 // don't use this when specifying a type
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

type CStableDiffusionCtx struct {
	ctx uintptr
}

type CUpScalerCtx struct {
	ctx uintptr
}

type CLogCallback func(level int, text uintptr)

type CStableDiffusion interface {
	NewCtx(modelPath string, vaePath string, taesdPath string, loraModelDir string, vaeDecodeOnly bool, vaeTiling bool, freeParamsImmediately bool, nThreads int, wType WType, rngType RNGType, schedule Schedule) *CStableDiffusionCtx
	PredictImage(ctx *CStableDiffusionCtx, prompt string, negativePrompt string, clipSkip int, cfgScale float32, width int, height int, sampleMethod SampleMethod, sampleSteps int, seed int64, batchCount int) []Image
	ImagePredictImage(ctx *CStableDiffusionCtx, img Image, prompt string, negativePrompt string, clipSkip int, cfgScale float32, width int, height int, sampleMethod SampleMethod, sampleSteps int, strength float32, seed int64, batchCount int) []Image
	SetLogCallBack(cb CLogCallback)
	GetSystemInfo() string
	FreeCtx(ctx *CStableDiffusionCtx)

	NewUpscalerCtx()
	FreeUpscalerCtx()
	UpscaleImage(ctx *CUpScalerCtx, img Image, upscaleFactor int) []byte
}

type cImage struct {
	width   uint32
	height  uint32
	channel uint32
	data    uintptr
}

type Image struct {
	Width   uint32
	Height  uint32
	Channel uint32
	Data    []byte
}

type CStableDiffusionImpl struct {
	libSd uintptr

	newSdCtx func(modelPath string, vaePath string, taesdPath string, loraModelDir string, vaeDecodeOnly bool, vaeTiling bool, freeParamsImmediately bool, nThreads int, wtype int, rngType int, schedule int) uintptr

	sdSetLogCallback func(callback func(level int, text uintptr, data uintptr) uintptr, data uintptr)

	txt2img func(ctx uintptr, prompt string, negativePrompt string, clipSkip int, cfgScale float32, width int, height int, sampleMethod int, sampleSteps int, seed int64, batchCount int) uintptr

	img2img func(ctx uintptr, img uintptr, prompt string, negativePrompt string, clipSkip int, cfgScale float32, width int, height int, sampleMethod int, sampleSteps int, strength float32, seed int64, batchCount int) uintptr

	freeSdCtx func(ctx uintptr)

	newUpscalerCtx func(esrganPath string, n_threads int, wtype int) uintptr

	freeUpscalerCtx func(ctx uintptr)

	upscale func(ctx uintptr, img uintptr, upscaleFactor uint32) uintptr
}

func NewCStableDiffusion(libraryPath string) (*CStableDiffusionImpl, error) {
	libSd, err := openLibrary(libraryPath)
	if err != nil {
		return nil, err
	}

	impl := CStableDiffusionImpl{}

	purego.RegisterLibFunc(&impl.newSdCtx, libSd, "new_sd_ctx")
	purego.RegisterLibFunc(&impl.sdSetLogCallback, libSd, "sd_set_log_callback")
	purego.RegisterLibFunc(&impl.txt2img, libSd, "txt2img")
	purego.RegisterLibFunc(&impl.img2img, libSd, "img2img")
	purego.RegisterLibFunc(&impl.freeSdCtx, libSd, "free_sd_ctx")

	return &impl, nil
}

func goString(c uintptr) string {
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
	return string(unsafe.Slice((*byte)(ptr), length))
}

func goImageSlice(c uintptr, size int) []Image {
	// We take the address and then dereference it to trick go vet from creating a possible misuse of unsafe.Pointer
	ptr := *(*unsafe.Pointer)(unsafe.Pointer(&c))
	if ptr == nil {
		return nil
	}
	img := (*cImage)(ptr)
	goImages := make([]Image, 0, size)
	imgSlice := unsafe.Slice(img, size)
	for _, image := range imgSlice {
		var gImg Image
		gImg.Channel = image.channel
		gImg.Width = image.width
		gImg.Height = image.height
		dataPtr := *(*unsafe.Pointer)(unsafe.Pointer(&image.data))
		if ptr != nil {
			gImg.Data = unsafe.Slice((*byte)(dataPtr), image.channel*image.width*image.height)
		}
		goImages = append(goImages, gImg)
	}

	return goImages
}
