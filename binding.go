// Copyright (c) seasonjs. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package sd

import "unsafe"

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

type CStableDiffusionCtx struct {
	ctx    uintptr
	cgoCtx unsafe.Pointer
}

type CUpScalerCtx struct {
	ctx    uintptr
	cgoCtx unsafe.Pointer
}

type CLogCallback func(level LogLevel, text string)

type CStableDiffusion interface {
	NewCtx(modelPath string, vaePath string, taesdPath string, loraModelDir string, vaeDecodeOnly bool, vaeTiling bool, freeParamsImmediately bool, nThreads int, wType WType, rngType RNGType, schedule Schedule) *CStableDiffusionCtx
	PredictImage(ctx *CStableDiffusionCtx, prompt string, negativePrompt string, clipSkip int, cfgScale float32, width int, height int, sampleMethod SampleMethod, sampleSteps int, seed int64, batchCount int) []Image
	ImagePredictImage(ctx *CStableDiffusionCtx, img Image, prompt string, negativePrompt string, clipSkip int, cfgScale float32, width int, height int, sampleMethod SampleMethod, sampleSteps int, strength float32, seed int64, batchCount int) []Image
	SetLogCallBack(cb CLogCallback)
	GetSystemInfo() string
	FreeCtx(ctx *CStableDiffusionCtx)

	NewUpscalerCtx(esrganPath string, nThreads int, wType WType) *CUpScalerCtx
	FreeUpscalerCtx(ctx *CUpScalerCtx)
	UpscaleImage(ctx *CUpScalerCtx, img Image, upscaleFactor uint32) Image

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
