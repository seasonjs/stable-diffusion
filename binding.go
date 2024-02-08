// Copyright (c) seasonjs. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package sd

import (
	"github.com/ebitengine/purego"
	"runtime"
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
	F32     WType = 0
	F16           = 1
	Q4_0          = 2
	Q4_1          = 3
	Q5_0          = 6
	Q5_1          = 7
	Q8_0          = 8
	Q8_1          = 9
	Q2_K          = 10
	Q3_K          = 11
	Q4_K          = 12
	Q5_K          = 13
	Q6_K          = 14
	Q8_K          = 15
	IQ2_XXS       = 16
	I8            = 17
	I16           = 18
	I32           = 19
	AUTO          = 20
)

type CStableDiffusionCtx struct {
	ctx uintptr
}

type CUpScalerCtx struct {
	ctx uintptr
}

type CLogCallback func(level LogLevel, text string)

type CStableDiffusion interface {
	NewCtx(modelPath string, vaePath string, taesdPath string, controlNetPath string, loraModelDir string, embedDir string, vaeDecodeOnly bool, vaeTiling bool, freeParamsImmediately bool, nThreads int, wType WType, rngType RNGType, schedule Schedule, keepControlNetCpu bool) *CStableDiffusionCtx

	PredictImage(ctx *CStableDiffusionCtx, prompt string, negativePrompt string, clipSkip int, cfgScale float32, width int, height int, sampleMethod SampleMethod, sampleSteps int, seed int64, batchCount int, controlCond *Image, controlStrength float32) []Image

	ImagePredictImage(ctx *CStableDiffusionCtx, img Image, prompt string, negativePrompt string, clipSkip int, cfgScale float32, width int, height int, sampleMethod SampleMethod, sampleSteps int, strength float32, seed int64, batchCount int) []Image

	SetLogCallBack(cb CLogCallback)
	GetSystemInfo() string

	Convert(inputPath string, vaePath string, outputPath string, outputType WType) bool

	FreeCtx(ctx *CStableDiffusionCtx)

	NewUpscalerCtx(esrganPath string, nThreads int, wType WType) *CUpScalerCtx
	FreeUpscalerCtx(ctx *CUpScalerCtx)
	UpscaleImage(ctx *CUpScalerCtx, img Image, upscaleFactor uint32) Image

	PreprocessCanny(image Image, width int, height int, highThreshold float32, lowThreshold float32, week float32, strong float32, inverse bool) Image

	Close() error
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

	sdGetSystemInfo func() string

	newSdCtx func(modelPath string,
		vaePath string,
		taesdPath string,
		controlNetPath string,
		loraModelDir string,
		embedDir string,
		vaeDecodeOnly bool,
		vaeTiling bool,
		freeParamsImmediately bool,
		nThreads int,
		wType int,
		rngType int,
		schedule int,
		keepControlNetCpu bool) uintptr

	sdSetLogCallback func(callback func(level int, text uintptr, data uintptr) uintptr, data uintptr)

	txt2img func(
		ctx uintptr,
		prompt string,
		negativePrompt string,
		clipSkip int,
		cfgScale float32,
		width int,
		height int,
		sampleMethod int,
		sampleSteps int,
		seed int64,
		batchCount int,
		controlCond uintptr,
		controlStrength float32) uintptr

	img2img func(
		ctx uintptr,
		img uintptr,
		prompt string,
		negativePrompt string,
		clipSkip int,
		cfgScale float32,
		width int,
		height int,
		sampleMethod int,
		sampleSteps int,
		strength float32,
		seed int64,
		batchCount int) uintptr

	freeSdCtx func(ctx uintptr)

	newUpscalerCtx func(esrganPath string, nThreads int, wtype int) uintptr

	freeUpscalerCtx func(ctx uintptr)

	upscale func(ctx uintptr, img uintptr, upscaleFactor uint32) uintptr

	convert func(inputPath string, vaePath string, outputPath string, outputType int) bool

	preprocessCanny func(img uintptr, width int, height int, highThreshold float32, lowThreshold float32, week float32, strong float32, inverse bool) *byte
}

func NewCStableDiffusion(libraryPath string) (*CStableDiffusionImpl, error) {
	libSd, err := openLibrary(libraryPath)
	if err != nil {
		return nil, err
	}

	impl := CStableDiffusionImpl{}

	purego.RegisterLibFunc(&impl.sdSetLogCallback, libSd, "sd_get_system_info")
	purego.RegisterLibFunc(&impl.newSdCtx, libSd, "new_sd_ctx")
	purego.RegisterLibFunc(&impl.sdSetLogCallback, libSd, "sd_set_log_callback")
	purego.RegisterLibFunc(&impl.txt2img, libSd, "txt2img")
	purego.RegisterLibFunc(&impl.img2img, libSd, "img2img")
	purego.RegisterLibFunc(&impl.freeSdCtx, libSd, "free_sd_ctx")
	purego.RegisterLibFunc(&impl.newUpscalerCtx, libSd, "new_upscaler_ctx")
	purego.RegisterLibFunc(&impl.freeUpscalerCtx, libSd, "free_upscaler_ctx")
	purego.RegisterLibFunc(&impl.upscale, libSd, "upscale")
	purego.RegisterLibFunc(&impl.convert, libSd, "convert")
	purego.RegisterLibFunc(&impl.preprocessCanny, libSd, "preprocess_canny")

	return &impl, nil
}

func (c *CStableDiffusionImpl) NewCtx(modelPath string, vaePath string, taesdPath string, controlNetPath string, loraModelDir string, embedDir string, vaeDecodeOnly bool, vaeTiling bool, freeParamsImmediately bool, nThreads int, wType WType, rngType RNGType, schedule Schedule, keepControlNetCpu bool) *CStableDiffusionCtx {
	ctx := c.newSdCtx(modelPath,
		vaePath,
		taesdPath,
		controlNetPath,
		loraModelDir,
		embedDir,
		vaeDecodeOnly,
		vaeTiling,
		freeParamsImmediately,
		nThreads,
		int(wType),
		int(rngType),
		int(schedule),
		keepControlNetCpu,
	)
	return &CStableDiffusionCtx{
		ctx: ctx,
	}
}

func (c *CStableDiffusionImpl) PredictImage(ctx *CStableDiffusionCtx, prompt string, negativePrompt string, clipSkip int, cfgScale float32, width int, height int, sampleMethod SampleMethod, sampleSteps int, seed int64, batchCount int, controlCond *Image, controlStrength float32) []Image {
	var ci *cImage
	if controlCond != nil {
		ci = &cImage{
			width:   controlCond.Width,
			height:  controlCond.Height,
			channel: controlCond.Channel,
			data:    uintptr(unsafe.Pointer(&controlCond.Data[0])),
		}
	}

	images := c.txt2img(
		ctx.ctx,
		prompt,
		negativePrompt,
		clipSkip,
		cfgScale,
		width,
		height,
		int(sampleMethod),
		sampleSteps,
		seed,
		batchCount,
		uintptr(unsafe.Pointer(ci)),
		controlStrength)

	return goImageSlice(images, batchCount)
}

func (c *CStableDiffusionImpl) ImagePredictImage(ctx *CStableDiffusionCtx, img Image, prompt string, negativePrompt string, clipSkip int, cfgScale float32, width int, height int, sampleMethod SampleMethod, sampleSteps int, strength float32, seed int64, batchCount int) []Image {
	ci := cImage{
		width:   img.Width,
		height:  img.Height,
		channel: img.Channel,
		data:    uintptr(unsafe.Pointer(&img.Data[0])),
	}

	images := c.img2img(ctx.ctx,
		uintptr(unsafe.Pointer(&ci)),
		prompt,
		negativePrompt,
		clipSkip,
		cfgScale,
		width,
		height,
		int(sampleMethod),
		sampleSteps,
		strength,
		seed,
		batchCount,
	)

	return goImageSlice(images, batchCount)
}

func (c *CStableDiffusionImpl) SetLogCallBack(cb CLogCallback) {
	c.sdSetLogCallback(func(level int, text uintptr, data uintptr) uintptr {
		cb(LogLevel(level), goString(text))
		return 0
	}, 0)
}

func (c *CStableDiffusionImpl) GetSystemInfo() string {
	return c.sdGetSystemInfo()
}

func (c *CStableDiffusionImpl) FreeCtx(ctx *CStableDiffusionCtx) {
	ptr := *(*unsafe.Pointer)(unsafe.Pointer(&ctx.ctx))
	if ptr != nil {
		c.freeSdCtx(ctx.ctx)
	}
	ctx = nil
	runtime.GC()
}

func (c *CStableDiffusionImpl) NewUpscalerCtx(esrganPath string, nThreads int, wType WType) *CUpScalerCtx {
	ctx := c.newUpscalerCtx(esrganPath, nThreads, int(wType))

	return &CUpScalerCtx{ctx: ctx}
}

func (c *CStableDiffusionImpl) FreeUpscalerCtx(ctx *CUpScalerCtx) {
	ptr := *(*unsafe.Pointer)(unsafe.Pointer(&ctx.ctx))
	if ptr != nil {
		c.freeUpscalerCtx(ctx.ctx)
	}
	ctx = nil
	runtime.GC()
}

func (c *CStableDiffusionImpl) Close() error {
	if c.libSd != 0 {
		err := closeLibrary(c.libSd)
		return err
	}
	return nil
}

func (c *CStableDiffusionImpl) UpscaleImage(ctx *CUpScalerCtx, img Image, upscaleFactor uint32) Image {
	ci := cImage{
		width:   img.Width,
		height:  img.Height,
		channel: img.Channel,
		data:    uintptr(unsafe.Pointer(&img.Data[0])),
	}
	uptr := c.upscale(ctx.ctx, uintptr(unsafe.Pointer(&ci)), upscaleFactor)
	ptr := *(*unsafe.Pointer)(unsafe.Pointer(&uptr))
	if ptr == nil {
		return Image{}
	}
	cimg := (*cImage)(ptr)
	dataPtr := *(*unsafe.Pointer)(unsafe.Pointer(&cimg.data))
	return Image{
		Width:   cimg.width,
		Height:  cimg.height,
		Channel: cimg.channel,
		Data:    unsafe.Slice((*byte)(dataPtr), cimg.channel*cimg.width*cimg.height),
	}
}

func (c *CStableDiffusionImpl) Convert(inputPath string, vaePath string, outputPath string, outputType WType) bool {
	return c.convert(inputPath, vaePath, outputPath, int(outputType))
}

func (c *CStableDiffusionImpl) PreprocessCanny(image Image, width int, height int, highThreshold float32, lowThreshold float32, week float32, strong float32, inverse bool) Image {
	dataPtr := c.preprocessCanny(uintptr(unsafe.Pointer(&image.Data[0])), width, height, highThreshold, lowThreshold, week, strong, inverse)
	return Image{
		Width:   image.Width,
		Height:  image.Height,
		Channel: image.Channel,
		Data:    unsafe.Slice((*byte)(dataPtr), image.Width*image.Width*image.Height),
	}
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
	return unsafe.String((*byte)(ptr), length)
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
