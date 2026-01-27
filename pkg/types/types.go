package types

type RNGType int32

type SampleMethod int32

type Scheduler int32

type Prediction int32

type SdType int32

type LogLevel int32

type Preview int32

type LoraApplyMode int32

type CacheMode int

const (
	SD_CACHE_DISABLED CacheMode = iota
	SD_CACHE_EASYCACHE
	SD_CACHE_UCACHE
	SD_CACHE_DBCACHE
	SD_CACHE_TAYLORSEER
	SD_CACHE_CACHE_DIT
)

const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
)

const (
	STD_DEFAULT_RNG RNGType = iota
	CUDA_RNG
	CPU_RNG
	RNG_TYPE_COUNT
)

const (
	EULER SampleMethod = iota
	EULER_A
	HEUN
	DPM2
	DPMPP2S_A
	DPMPP2M
	DPMPP2Mv2
	IPNDM
	IPNDM_V
	LCM
	DDIM_TRAILING
	TCD
	SAMPLE_METHOD_COUNT
)

const (
	DISCRETE Scheduler = iota
	KARRAS
	EXPONENTIAL
	AYS
	GITS
	SGM_UNIFORM
	SIMPLE
	SMOOTHSTEP
	KL_OPTIMAL
	LCM_SCHEDULER
	SCHEDULER_COUNT
)

const (
	EPS_PRED Prediction = iota
	V_PRED
	EDM_V_PRED
	FLOW_PRED
	FLUX_FLOW_PRED
	FLUX2_FLOW_PRED
	PREDICTION_COUNT
)

const (
	PREVIEW_NONE Preview = iota
	PREVIEW_PROJ
	PREVIEW_TAE
	PREVIEW_VAE
	PREVIEW_COUNT
)

const (
	LORA_APPLY_AUTO LoraApplyMode = iota
	LORA_APPLY_IMMEDIATELY
	LORA_APPLY_AT_RUNTIME
	LORA_APPLY_MODE_COUNT
)

const (
	F32     SdType = 0
	F16            = 1
	Q4_0           = 2
	Q4_1           = 3
	Q5_0           = 6
	Q5_1           = 7
	Q8_0           = 8
	Q8_1           = 9
	Q2_K           = 10
	Q3_K           = 11
	Q4_K           = 12
	Q5_K           = 13
	Q6_K           = 14
	Q8_K           = 15
	IQ2_XXS        = 16
	IQ2_XS         = 17
	IQ3_XXS        = 18
	IQ1_S          = 19
	IQ4_NL         = 20
	IQ3_S          = 21
	IQ2_S          = 22
	IQ4_XS         = 23
	I8             = 24
	I16            = 25
	I32            = 26
	I64            = 27
	F64            = 28
	IQ1_M          = 29
	BF16           = 30
	TQ1_0          = 34
	TQ2_0          = 35
	MXFP4          = 39
	COUNT          = 40 // don't use this when specifying a type
)

type Image struct {
	Width   uint32
	Height  uint32
	Channel uint32
	Data    []byte
}

// ProgressCallback 进度回调函数类型
type ProgressCallback func(step, steps int, time float32)

// PreviewCallback 预览回调函数类型
type PreviewCallback func(step, frameCount int, frames []Image, isNoisy bool)
