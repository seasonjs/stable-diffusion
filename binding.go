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

type GgmlType string

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
	T_DEFAULT GgmlType = "DEFAULT"
	F32                = "F32"
	F16                = "F16"
	Q4_0               = "Q4_0"
	Q4_1               = "Q4_1"
	Q5_0               = "Q5_0"
	Q5_1               = "Q5_1"
	Q8_0               = "Q8_0"
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
	StableDiffusionInit(nThreads int, vaeDecodeOnly bool, taesdPath string, freeParamsImmediately bool, loraModelDir string, rngType RNGType) *StableDiffusionCtx
	StableDiffusionLoadFromFile(ctx *StableDiffusionCtx, filePath string, vaePath string, wtype GgmlType, schedule Schedule)
	StableDiffusionPredictImage(ctx *StableDiffusionCtx, params *StableDiffusionFullParams, prompt string) []byte
	StableDiffusionImagePredictImage(ctx *StableDiffusionCtx, params *StableDiffusionFullParams, initImage []byte, prompt string) []byte
	StableDiffusionSetLogLevel(level SDLogLevel)
	StableDiffusionGetSystemInfo() string
	StableDiffusionFree(ctx *StableDiffusionCtx)
	StableDiffusionFreeBuffer(buffer uintptr)
	StableDiffusionFreeFullParams(params *StableDiffusionFullParams)
}

type cImage struct {
	width   uint32
	height  uint32
	channel uint32
	data    uintptr
}

type Image struct {
	width   uint32
	height  uint32
	channel uint32
	data    []byte
}

type CStableDiffusionImpl struct {
	libSd            uintptr
	newSdCtx         func(modelPath string, vaePath string, taesdPath string, loraModelDir string, vaeDecodeOnly bool, vaeTiling bool, freeParamsImmediately bool, nThreads int, wtype int, rngType int, schedule int) uintptr
	sdSetLogCallback func(callback func(level int, text uintptr, data uintptr) uintptr, data uintptr)
	txt2img          func(ctx uintptr, prompt string, negativePrompt string, clipSkip int, cfgScale float32, width int, height int, sampleMethod int, sampleSteps int, seed int64, batchCount int) uintptr
	img2img          func(ctx uintptr, img uintptr, prompt string, negativePrompt string, clipSkip int, cfgScale float32, width int, height int, sampleMethod int, sampleSteps int, strength float32, seed int64, batchCount int) uintptr
	freeSdCtx        func(ctx uintptr)
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

func GoString(c uintptr) string {
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

func GoImageSlice(c uintptr, size int) []Image {
	ptr := *(*unsafe.Pointer)(unsafe.Pointer(&c))
	if ptr == nil {
		return nil
	}
	img := (*cImage)(ptr)
	goImages := make([]Image, 0, size)
	imgSlice := unsafe.Slice(img, size)
	for _, image := range imgSlice {
		var gImg Image
		gImg.channel = image.channel
		gImg.width = image.width
		gImg.height = image.height
		dataPtr := *(*unsafe.Pointer)(unsafe.Pointer(&image.data))
		if ptr != nil {
			gImg.data = unsafe.Slice((*byte)(dataPtr), image.channel*image.width*image.height)
		}
		goImages = append(goImages, gImg)
	}

	return goImages
}
