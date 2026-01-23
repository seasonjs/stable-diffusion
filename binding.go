// Copyright (c) seasonjs. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package sd

import (
	"unsafe"

	"github.com/ebitengine/purego"
	"github.com/jupiterrider/ffi"
)

type RNGType int32

type SampleMethod int32

type Scheduler int32

type Prediction int32

type SdType int32

type LogLevel int32

type Preview int32

type LoraApplyMode int32

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
	N_SAMPLE_METHODS
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
	N_SCHEDULES
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

// C结构体对应的Go结构体定义
type CacheMode int

const (
	SD_CACHE_DISABLED CacheMode = iota
	SD_CACHE_EASYCACHE
	SD_CACHE_UCACHE
	SD_CACHE_DBCACHE
	SD_CACHE_TAYLORSEER
	SD_CACHE_CACHE_DIT
)

type TilingParams struct {
	Enabled       bool
	TileSizeX     int32
	TileSizeY     int32
	TargetOverlap float32
	RelSizeX      float32
	RelSizeY      float32
}

type Embedding struct {
	Name *byte
	Path *byte
}

type CtxParams struct {
	ModelPath                   *byte
	ClipLPath                   *byte
	ClipGPath                   *byte
	ClipVisionPath              *byte
	T5xxlPath                   *byte
	LlmPath                     *byte
	LlmVisionPath               *byte
	DiffusionModelPath          *byte
	HighNoiseDiffusionModelPath *byte
	VaePath                     *byte
	TaesdPath                   *byte
	ControlNetPath              *byte
	Embeddings                  uintptr
	EmbeddingCount              uint32
	PhotoMakerPath              *byte
	TensorTypeRules             *byte
	VaeDecodeOnly               bool
	FreeParamsImmediately       bool
	NThreads                    int32
	Wtype                       int32
	RngType                     int32
	SamplerRngType              int32
	Prediction                  int32
	LoraApplyMode               int32
	OffloadParamsToCpu          bool
	EnableMmap                  bool
	KeepClipOnCpu               bool
	KeepControlNetOnCpu         bool
	KeepVaeOnCpu                bool
	DiffusionFlashAttn          bool
	TaePreviewOnly              bool
	DiffusionConvDirect         bool
	VaeConvDirect               bool
	CircularX                   bool
	CircularY                   bool
	ForceSdxlVaeConvScale       bool
	ChromaUseDitMask            bool
	ChromaUseT5Mask             bool
	ChromaT5MaskPad             int32
	QwenImageZeroCondT          bool
	FlowShift                   float32
}

type SlgParams struct {
	Layers     *int32
	LayerCount uintptr
	LayerStart float32
	LayerEnd   float32
	Scale      float32
}

type GuidanceParams struct {
	TxtCfg            float32
	ImgCfg            float32
	DistilledGuidance float32
	Slg               SlgParams
}

type SampleParams struct {
	Guidance          GuidanceParams
	Scheduler         Scheduler
	SampleMethod      SampleMethod
	SampleSteps       int32
	Eta               float32
	ShiftedTimestep   int32
	CustomSigmas      *float32
	CustomSigmasCount int32
}

type PmParams struct {
	IdImages      *Image
	IdImagesCount int32
	IdEmbedPath   *byte
	StyleStrength float32
}

type CacheParams struct {
	Mode                     CacheMode
	ReuseThreshold           float32
	StartPercent             float32
	EndPercent               float32
	ErrorDecayRate           float32
	UseRelativeThreshold     bool
	ResetErrorOnCompute      bool
	FnComputeBlocks          int32
	BnComputeBlocks          int32
	ResidualDiffThreshold    float32
	MaxWarmupSteps           int32
	MaxCachedSteps           int32
	MaxContinuousCachedSteps int32
	TaylorseerNDerivatives   int32
	TaylorseerSkipInterval   int32
	ScmMask                  *byte
	ScmPolicyDynamic         bool
}

type Lora struct {
	IsHighNoise bool
	Multiplier  float32
	Path        *byte
}

type ImgGenParams struct {
	Loras              *Lora
	LoraCount          uint32
	Prompt             *byte
	NegativePrompt     *byte
	ClipSkip           int32
	InitImage          Image
	RefImages          *Image
	RefImagesCount     int32
	AutoResizeRefImage bool
	IncreaseRefIndex   bool
	MaskImage          Image
	Width              int32
	Height             int32
	SampleParams       SampleParams
	Strength           float32
	Seed               int64
	BatchCount         int32
	ControlImage       Image
	ControlStrength    float32
	PmParams           PmParams
	VaeTilingParams    TilingParams
	Cache              CacheParams
}

type VidGenParams struct {
	Loras                 *Lora
	LoraCount             uint32
	Prompt                *byte
	NegativePrompt        *byte
	ClipSkip              int32
	InitImage             Image
	EndImage              Image
	ControlFrames         *Image
	ControlFramesSize     int32
	Width                 int32
	Height                int32
	SampleParams          SampleParams
	HighNoiseSampleParams SampleParams
	MoeBoundary           float32
	Strength              float32
	Seed                  int64
	VideoFrames           int32
	VaceStrength          float32
	VaeTilingParams       TilingParams
	Cache                 CacheParams
}

type CStableDiffusionCtx struct {
	ctx uintptr
}

type CUpScalerCtx struct {
	ctx uintptr
}

// CStableDiffusion 是stable-diffusion.h API的Go语言封装接口
// 为所有结构体定义ffi.Type，确保正确的内存布局和参数传递

// FFITypeImage 是Image结构体的ffi.Type定义
var FFITypeImage = ffi.NewType(
	&ffi.TypeUint32,  // Width: uint32
	&ffi.TypeUint32,  // Height: uint32
	&ffi.TypeUint32,  // Channel: uint32
	&ffi.TypePointer, // Data: *byte (C中的uint8_t*)
)

// FFITypeEmbedding 是Embedding结构体的ffi.Type定义
var FFITypeEmbedding = ffi.NewType(
	&ffi.TypePointer, // Name: *byte
	&ffi.TypePointer, // Path: *byte
)

// FFITypeCtxParams 是CtxParams结构体的ffi.Type定义
var FFITypeCtxParams = ffi.NewType(
	&ffi.TypePointer, // ModelPath: *byte
	&ffi.TypePointer, // ClipLPath: *byte
	&ffi.TypePointer, // ClipGPath: *byte
	&ffi.TypePointer, // ClipVisionPath: *byte
	&ffi.TypePointer, // T5xxlPath: *byte
	&ffi.TypePointer, // LlmPath: *byte
	&ffi.TypePointer, // LlmVisionPath: *byte
	&ffi.TypePointer, // DiffusionModelPath: *byte
	&ffi.TypePointer, // HighNoiseDiffusionModelPath: *byte
	&ffi.TypePointer, // VaePath: *byte
	&ffi.TypePointer, // TaesdPath: *byte
	&ffi.TypePointer, // ControlNetPath: *byte
	&ffi.TypePointer, // Embeddings: *Embedding (需要使用Embedding的ffi.Type)
	&ffi.TypeUint32,  // EmbeddingCount: uint32
	&ffi.TypePointer, // PhotoMakerPath: *byte
	&ffi.TypePointer, // TensorTypeRules: *byte
	&ffi.TypeUint8,   // VaeDecodeOnly: bool (C bool is uint8)
	&ffi.TypeUint8,   // FreeParamsImmediately: bool
	&ffi.TypeSint32,  // NThreads: int32
	&ffi.TypeSint32,  // Wtype: SdType (enum is int32)
	&ffi.TypeSint32,  // RngType: RNGType (enum is int32)
	&ffi.TypeSint32,  // SamplerRngType: RNGType
	&ffi.TypeSint32,  // Prediction: Prediction (enum is int32)
	&ffi.TypeSint32,  // LoraApplyMode: LoraApplyMode (enum is int32)
	&ffi.TypeUint8,   // OffloadParamsToCpu: bool
	&ffi.TypeUint8,   // EnableMmap: bool
	&ffi.TypeUint8,   // KeepClipOnCpu: bool
	&ffi.TypeUint8,   // KeepControlNetOnCpu: bool
	&ffi.TypeUint8,   // KeepVaeOnCpu: bool
	&ffi.TypeUint8,   // DiffusionFlashAttn: bool
	&ffi.TypeUint8,   // TaePreviewOnly: bool
	&ffi.TypeUint8,   // DiffusionConvDirect: bool
	&ffi.TypeUint8,   // VaeConvDirect: bool
	&ffi.TypeUint8,   // CircularX: bool
	&ffi.TypeUint8,   // CircularY: bool
	&ffi.TypeUint8,   // ForceSdxlVaeConvScale: bool
	&ffi.TypeUint8,   // ChromaUseDitMask: bool
	&ffi.TypeUint8,   // ChromaUseT5Mask: bool
	&ffi.TypeSint32,  // ChromaT5MaskPad: int32
	&ffi.TypeUint8,   // QwenImageZeroCondT: bool
	&ffi.TypeFloat,   // FlowShift: float32
)

// FFITypeLora 是Lora结构体的ffi.Type定义
var FFITypeLora = ffi.NewType(
	&ffi.TypeUint8,   // IsHighNoise: bool
	&ffi.TypeFloat,   // Multiplier: float32
	&ffi.TypePointer, // Path: *byte
)

// FFITypePmParams 是PmParams结构体的ffi.Type定义
var FFITypePmParams = ffi.NewType(
	&ffi.TypePointer, // IdImages: *Image (使用指针，因为C中是数组)
	&ffi.TypeSint32,  // IdImagesCount: int32
	&ffi.TypePointer, // IdEmbedPath: *byte
	&ffi.TypeFloat,   // StyleStrength: float32
)

// FFITypeTilingParams 是TilingParams结构体的ffi.Type定义
var FFITypeTilingParams = ffi.NewType(
	&ffi.TypeUint8,  // Enabled: bool
	&ffi.TypeSint32, // TileSizeX: int32
	&ffi.TypeSint32, // TileSizeY: int32
	&ffi.TypeFloat,  // TargetOverlap: float32
	&ffi.TypeFloat,  // RelSizeX: float32
	&ffi.TypeFloat,  // RelSizeY: float32
)

// FFITypeSampleParams 是SampleParams结构体的ffi.Type定义
var FFITypeSampleParams = ffi.NewType(
	&ffi.TypeFloat,   // Guidance.TxtCfg: float32
	&ffi.TypeFloat,   // Guidance.ImgCfg: float32
	&ffi.TypeFloat,   // Guidance.DistilledGuidance: float32
	&ffi.TypePointer, // Guidance.Slg.Layers: *int32
	&ffi.TypePointer, // Guidance.Slg.LayerCount: uintptr (使用*byte替代uintptr)
	&ffi.TypeFloat,   // Guidance.Slg.LayerStart: float32
	&ffi.TypeFloat,   // Guidance.Slg.LayerEnd: float32
	&ffi.TypeFloat,   // Guidance.Slg.Scale: float32
	&ffi.TypeSint32,  // Scheduler: Scheduler (enum is int32)
	&ffi.TypeSint32,  // SampleMethod: SampleMethod (enum is int32)
	&ffi.TypeSint32,  // SampleSteps: int32
	&ffi.TypeFloat,   // Eta: float32
	&ffi.TypeSint32,  // ShiftedTimestep: int32
	&ffi.TypePointer, // CustomSigmas: *float32
	&ffi.TypeSint32,  // CustomSigmasCount: int32
)

// FFITypeCacheParams 是CacheParams结构体的ffi.Type定义
var FFITypeCacheParams = ffi.NewType(
	&ffi.TypeSint32,  // Mode: CacheMode (enum is int32)
	&ffi.TypeFloat,   // ReuseThreshold: float32
	&ffi.TypeFloat,   // StartPercent: float32
	&ffi.TypeFloat,   // EndPercent: float32
	&ffi.TypeFloat,   // ErrorDecayRate: float32
	&ffi.TypeUint8,   // UseRelativeThreshold: bool
	&ffi.TypeUint8,   // ResetErrorOnCompute: bool
	&ffi.TypeSint32,  // FnComputeBlocks: int32
	&ffi.TypeSint32,  // BnComputeBlocks: int32
	&ffi.TypeFloat,   // ResidualDiffThreshold: float32
	&ffi.TypeSint32,  // MaxWarmupSteps: int32
	&ffi.TypeSint32,  // MaxCachedSteps: int32
	&ffi.TypeSint32,  // MaxContinuousCachedSteps: int32
	&ffi.TypeSint32,  // TaylorseerNDerivatives: int32
	&ffi.TypeSint32,  // TaylorseerSkipInterval: int32
	&ffi.TypePointer, // ScmMask: *byte
	&ffi.TypeUint8,   // ScmPolicyDynamic: bool
)

// FFITypeImgGenParams 是ImgGenParams结构体的ffi.Type定义
var FFITypeImgGenParams = ffi.NewType(
	&ffi.TypePointer,     // Loras: *Lora
	&ffi.TypeUint32,      // LoraCount: uint32
	&ffi.TypePointer,     // Prompt: *byte
	&ffi.TypePointer,     // NegativePrompt: *byte
	&ffi.TypeSint32,      // ClipSkip: int32
	&FFITypeImage,        // InitImage: Image
	&ffi.TypePointer,     // RefImages: *Image
	&ffi.TypeSint32,      // RefImagesCount: int32
	&ffi.TypeUint8,       // AutoResizeRefImage: bool
	&ffi.TypeUint8,       // IncreaseRefIndex: bool
	&FFITypeImage,        // MaskImage: Image
	&ffi.TypeSint32,      // Width: int32
	&ffi.TypeSint32,      // Height: int32
	&FFITypeSampleParams, // SampleParams: SampleParams
	&ffi.TypeFloat,       // Strength: float32
	&ffi.TypeUint64,      // Seed: int64
	&ffi.TypeSint32,      // BatchCount: int32
	&FFITypeImage,        // ControlImage: Image
	&ffi.TypeFloat,       // ControlStrength: float32
	&FFITypePmParams,     // PmParams: PmParams
	&FFITypeTilingParams, // VaeTilingParams: TilingParams
	&FFITypeCacheParams,  // Cache: CacheParams
)

// FFITypeVidGenParams 是VidGenParams结构体的ffi.Type定义
var FFITypeVidGenParams = ffi.NewType(
	&ffi.TypePointer, // Loras: *Lora
	&ffi.TypeUint32,  // LoraCount: uint32
	&ffi.TypePointer, // Prompt: *byte
	&ffi.TypePointer, // NegativePrompt: *byte
	&ffi.TypeSint32,  // ClipSkip: int32
	&FFITypeImage,    // InitImage: Image
	&FFITypeImage,    // EndImage: Image
	&ffi.TypePointer, // ControlFrames: *Image
	&ffi.TypeSint32,  // ControlFramesSize: int32
	&ffi.TypeSint32,  // Width: int32
	&ffi.TypeSint32,  // Height: int32
	&ffi.TypePointer, // SampleParams: *SampleParams
	&ffi.TypePointer, // HighNoiseSampleParams: *SampleParams
	&ffi.TypeFloat,   // MoeBoundary: float32
	&ffi.TypeFloat,   // Strength: float32
	&ffi.TypeUint64,  // Seed: int64
	&ffi.TypeSint32,  // VideoFrames: int32
	&ffi.TypeFloat,   // VaceStrength: float32
	&ffi.TypePointer, // VaeTilingParams: *TilingParams
	&ffi.TypePointer, // Cache: *CacheParams
)

type CStableDiffusion interface {
	// 上下文管理
	NewCtx(params *CtxParams) *CStableDiffusionCtx
	FreeCtx(ctx *CStableDiffusionCtx)

	// 参数初始化
	CtxParamsInit() CtxParams
	SampleParamsInit(params *SampleParams)
	ImgGenParamsInit(params *ImgGenParams)
	VidGenParamsInit(params *VidGenParams)
	CacheParamsInit(params *CacheParams)

	// 参数转换为字符串
	CtxParamsToStr(params *CtxParams) string
	SampleParamsToStr(params *SampleParams) string
	ImgGenParamsToStr(params *ImgGenParams)

	// 生成功能
	GenerateImage(ctx *CStableDiffusionCtx, params *ImgGenParams) []Image
	GenerateVideo(ctx *CStableDiffusionCtx, params *VidGenParams) ([]Image, int)

	// 回调设置
	SetLogCallback(cb LogCallback)
	SetProgressCallback(cb ProgressCallback)
	SetPreviewCallback(cb PreviewCallback, mode Preview, interval int, denoised bool, noisy bool)

	// 辅助函数
	GetNumPhysicalCores() int32
	GetSystemInfo() string
	GetCommit() string
	GetVersion() string

	// 类型转换和名称获取
	SdTypeName(sdType SdType) string
	StrToSdType(str string) SdType
	RngTypeName(rngType RNGType) string
	StrToRngType(str string) RNGType
	SampleMethodName(sampleMethod SampleMethod) string
	StrToSampleMethod(str string) SampleMethod
	SchedulerName(scheduler Scheduler) string
	StrToScheduler(str string) Scheduler
	PredictionName(prediction Prediction) string
	StrToPrediction(str string) Prediction
	PreviewName(preview Preview) string
	StrToPreview(str string) Preview
	LoraApplyModeName(mode LoraApplyMode) string
	StrToLoraApplyMode(str string) LoraApplyMode

	// 默认值获取
	GetDefaultSampleMethod(ctx *CStableDiffusionCtx) SampleMethod
	GetDefaultScheduler(ctx *CStableDiffusionCtx, sampleMethod SampleMethod) Scheduler

	// 升频功能
	NewUpscalerCtx(esrganPath string, offloadParamsToCpu bool, direct bool, nThreads int, tileSize int) *CUpScalerCtx
	FreeUpscalerCtx(ctx *CUpScalerCtx)
	Upscale(ctx *CUpScalerCtx, inputImage Image, upscaleFactor uint32) Image
	GetUpscaleFactor(ctx *CUpScalerCtx) int

	// 其他功能
	Convert(inputPath, vaePath, outputPath string, outputType SdType, tensorTypeRules string, convertName bool) bool
	PreprocessCanny(image Image, highThreshold, lowThreshold, weak, strong float32, inverse bool) bool

	// 关闭资源
	Close() error
}

// LogCallback 日志回调函数类型
type LogCallback func(level LogLevel, text string)

// ProgressCallback 进度回调函数类型
type ProgressCallback func(step, steps int, time float32)

// PreviewCallback 预览回调函数类型
type PreviewCallback func(step, frameCount int, frames []Image, isNoisy bool)

type Image struct {
	Width   uint32
	Height  uint32
	Channel uint32
	Data    *byte
}

type CStableDiffusionImpl struct {
	libSd ffi.Lib

	// SD_API void sd_set_log_callback(sd_log_cb_t sd_log_cb, void* data);
	setLogCallback ffi.Fun

	// SD_API void sd_set_progress_callback(sd_progress_cb_t cb, void* data);
	setProgressCallback ffi.Fun

	// SD_API void sd_set_preview_callback(sd_preview_cb_t cb, enum preview_t mode, int interval, bool denoised, bool noisy, void* data);
	setPreviewCallback ffi.Fun

	// SD_API int32_t sd_get_num_physical_cores();
	getNumPhysicalCores ffi.Fun

	// SD_API const char* sd_get_system_info();
	getSystemInfo ffi.Fun

	// SD_API const char* sd_type_name(enum sd_type_t type);
	typeName ffi.Fun

	// SD_API enum sd_type_t str_to_sd_type(const char* str);
	strToSdType ffi.Fun

	// SD_API const char* sd_rng_type_name(enum rng_type_t rng_type);
	rngTypeName ffi.Fun

	// SD_API enum rng_type_t str_to_rng_type(const char* str);
	strToRngType ffi.Fun

	// SD_API const char* sd_sample_method_name(enum sample_method_t sample_method);
	sampleMethodName ffi.Fun

	// SD_API enum sample_method_t str_to_sample_method(const char* str);
	strToSampleMethod ffi.Fun

	// SD_API const char* sd_scheduler_name(enum scheduler_t scheduler);
	schedulerName ffi.Fun

	// SD_API enum scheduler_t str_to_scheduler(const char* str);
	strToScheduler ffi.Fun

	// SD_API const char* sd_prediction_name(enum prediction_t prediction);
	predictionName ffi.Fun

	// SD_API enum prediction_t str_to_prediction(const char* str);
	strToPrediction ffi.Fun

	// SD_API const char* sd_preview_name(enum preview_t preview);
	previewName ffi.Fun

	// SD_API enum preview_t str_to_preview(const char* str);
	strToPreview ffi.Fun

	// SD_API const char* sd_lora_apply_mode_name(enum lora_apply_mode_t mode);
	loraApplyModeName ffi.Fun

	// SD_API enum lora_apply_mode_t str_to_lora_apply_mode(const char* str);
	strToLoraApplyMode ffi.Fun

	// SD_API void sd_cache_params_init(sd_cache_params_t* cache_params);
	cacheParamsInit ffi.Fun

	// SD_API void sd_ctx_params_init(sd_ctx_params_t* sd_ctx_params);
	// 这里因为C语言端进行了初始化：
	//  *sd_ctx_params                         = {};
	// 所以导致内存结构通过ffi无法对齐，所以需要使用更系统级别的调用，让c正确写入
	ctxParamsInit  func(uintptr) 

	// SD_API char* sd_ctx_params_to_str(const sd_ctx_params_t* sd_ctx_params);
	ctxParamsToStr ffi.Fun

	// SD_API sd_ctx_t* new_sd_ctx(const sd_ctx_params_t* sd_ctx_params);
	newCtx ffi.Fun

	// SD_API void free_sd_ctx(sd_ctx_t* sd_ctx);
	freeCtx ffi.Fun

	// SD_API void sd_sample_params_init(sd_sample_params_t* sample_params);
	sampleParamsInit ffi.Fun

	// SD_API char* sd_sample_params_to_str(const sd_sample_params_t* sample_params);
	sampleParamsToStr ffi.Fun

	// SD_API enum sample_method_t sd_get_default_sample_method(const sd_ctx_t* sd_ctx);
	getDefaultSampleMethod ffi.Fun

	// SD_API enum scheduler_t sd_get_default_scheduler(const sd_ctx_t* sd_ctx, enum sample_method_t sample_method);
	getDefaultScheduler ffi.Fun

	// SD_API void sd_img_gen_params_init(sd_img_gen_params_t* sd_img_gen_params);
	imgGenParamsInit ffi.Fun

	// SD_API char* sd_img_gen_params_to_str(const sd_img_gen_params_t* sd_img_gen_params);
	imgGenParamsToStr ffi.Fun

	// SD_API sd_image_t* generate_image(sd_ctx_t* sd_ctx, const sd_img_gen_params_t* sd_img_gen_params);
	generateImage ffi.Fun

	// SD_API void sd_vid_gen_params_init(sd_vid_gen_params_t* sd_vid_gen_params);
	vidGenParamsInit ffi.Fun

	// SD_API sd_image_t* generate_video(sd_ctx_t* sd_ctx, const sd_vid_gen_params_t* sd_vid_gen_params, int* num_frames_out);
	generateVideo ffi.Fun

	// SD_API upscaler_ctx_t* new_upscaler_ctx(const char* esrgan_path, bool offload_params_to_cpu, bool direct, int n_threads, int tile_size);
	newUpscalerCtx ffi.Fun

	// SD_API void free_upscaler_ctx(upscaler_ctx_t* upscaler_ctx);
	freeUpscalerCtx ffi.Fun

	// SD_API sd_image_t* upscale(upscaler_ctx_t* upscaler_ctx, sd_image_t input_image, uint32_t upscale_factor);
	upscale ffi.Fun

	// SD_API int get_upscale_factor(upscaler_ctx_t* upscaler_ctx);
	getUpscaleFactor ffi.Fun

	// SD_API bool convert(const char* input_path, const char* vae_path, const char* output_path, enum sd_type_t output_type, const char* tensor_type_rules, bool convert_name);
	convert ffi.Fun

	// SD_API const char* sd_commit(void);
	commit ffi.Fun

	// SD_API const char* sd_version(void);
	version ffi.Fun
}

func NewCStableDiffusion(libraryPath string) (*CStableDiffusionImpl, error) {
	// 使用ffi包加载动态库
	libSd, err := ffi.Load(libraryPath)

	if err != nil {
		return nil, err
	}

	impl := CStableDiffusionImpl{
		libSd: libSd,
	}

	// 注册所有C函数
	// SD_API void sd_set_log_callback(sd_log_cb_t sd_log_cb, void* data);
	impl.setLogCallback, err = libSd.Prep("sd_set_log_callback", &ffi.TypeVoid, &ffi.TypePointer, &ffi.TypePointer)
	if err != nil {
		return nil, err
	}

	// SD_API void sd_set_progress_callback(sd_progress_cb_t cb, void* data);
	impl.setProgressCallback, err = libSd.Prep("sd_set_progress_callback", &ffi.TypeVoid, &ffi.TypePointer, &ffi.TypePointer)
	if err != nil {
		return nil, err
	}

	// SD_API void sd_set_preview_callback(sd_preview_cb_t cb, enum preview_t mode, int interval, bool denoised, bool noisy, void* data);
	impl.setPreviewCallback, err = libSd.Prep("sd_set_preview_callback", &ffi.TypeVoid, &ffi.TypePointer, &ffi.TypeSint32, &ffi.TypeSint32, &ffi.TypeUint8, &ffi.TypeUint8, &ffi.TypePointer)
	if err != nil {
		return nil, err
	}

	// SD_API int32_t sd_get_num_physical_cores();
	impl.getNumPhysicalCores, err = libSd.Prep("sd_get_num_physical_cores", &ffi.TypeSint32)
	if err != nil {
		return nil, err
	}

	// SD_API const char* sd_get_system_info();
	impl.getSystemInfo, err = libSd.Prep("sd_get_system_info", &ffi.TypePointer)
	if err != nil {
		return nil, err
	}

	// SD_API const char* sd_type_name(enum sd_type_t type);
	impl.typeName, err = libSd.Prep("sd_type_name", &ffi.TypePointer, &ffi.TypeSint32)
	if err != nil {
		return nil, err
	}

	// SD_API enum sd_type_t str_to_sd_type(const char* str);
	impl.strToSdType, err = libSd.Prep("str_to_sd_type", &ffi.TypeSint32, &ffi.TypePointer)
	if err != nil {
		return nil, err
	}

	// SD_API const char* sd_rng_type_name(enum rng_type_t rng_type);
	impl.rngTypeName, err = libSd.Prep("sd_rng_type_name", &ffi.TypePointer, &ffi.TypeSint32)
	if err != nil {
		return nil, err
	}

	// SD_API enum rng_type_t str_to_rng_type(const char* str);
	impl.strToRngType, err = libSd.Prep("str_to_rng_type", &ffi.TypeSint32, &ffi.TypePointer)
	if err != nil {
		return nil, err
	}

	// SD_API const char* sd_sample_method_name(enum sample_method_t sample_method);
	impl.sampleMethodName, err = libSd.Prep("sd_sample_method_name", &ffi.TypePointer, &ffi.TypeSint32)
	if err != nil {
		return nil, err
	}

	// SD_API enum sample_method_t str_to_sample_method(const char* str);
	impl.strToSampleMethod, err = libSd.Prep("str_to_sample_method", &ffi.TypeSint32, &ffi.TypePointer)
	if err != nil {
		return nil, err
	}

	// SD_API const char* sd_scheduler_name(enum scheduler_t scheduler);
	impl.schedulerName, err = libSd.Prep("sd_scheduler_name", &ffi.TypePointer, &ffi.TypeSint32)
	if err != nil {
		return nil, err
	}

	// SD_API enum scheduler_t str_to_scheduler(const char* str);
	impl.strToScheduler, err = libSd.Prep("str_to_scheduler", &ffi.TypeSint32, &ffi.TypePointer)
	if err != nil {
		return nil, err
	}

	// SD_API const char* sd_prediction_name(enum prediction_t prediction);
	impl.predictionName, err = libSd.Prep("sd_prediction_name", &ffi.TypePointer, &ffi.TypeSint32)
	if err != nil {
		return nil, err
	}

	// SD_API enum prediction_t str_to_prediction(const char* str);
	impl.strToPrediction, err = libSd.Prep("str_to_prediction", &ffi.TypeSint32, &ffi.TypePointer)
	if err != nil {
		return nil, err
	}

	// SD_API const char* sd_preview_name(enum preview_t preview);
	impl.previewName, err = libSd.Prep("sd_preview_name", &ffi.TypePointer, &ffi.TypeSint32)
	if err != nil {
		return nil, err
	}

	// SD_API enum preview_t str_to_preview(const char* str);
	impl.strToPreview, err = libSd.Prep("str_to_preview", &ffi.TypeSint32, &ffi.TypePointer)
	if err != nil {
		return nil, err
	}

	// SD_API const char* sd_lora_apply_mode_name(enum lora_apply_mode_t mode);
	impl.loraApplyModeName, err = libSd.Prep("sd_lora_apply_mode_name", &ffi.TypePointer, &ffi.TypeSint32)
	if err != nil {
		return nil, err
	}

	// SD_API enum lora_apply_mode_t str_to_lora_apply_mode(const char* str);
	impl.strToLoraApplyMode, err = libSd.Prep("str_to_lora_apply_mode", &ffi.TypeSint32, &ffi.TypePointer)
	if err != nil {
		return nil, err
	}

	// SD_API void sd_cache_params_init(sd_cache_params_t* cache_params);
	impl.cacheParamsInit, err = libSd.Prep("sd_cache_params_init", &ffi.TypeVoid, &FFITypeCacheParams)
	if err != nil {
		return nil, err
	}

	// SD_API void sd_ctx_params_init(sd_ctx_params_t* sd_ctx_params);
	// 这种通过c初始化结构体的操作使用purego 进行
	purego.RegisterLibFunc(&impl.ctxParamsInit, libSd.Addr, "sd_ctx_params_init")
	if err != nil {
		return nil, err
	}

	// SD_API char* sd_ctx_params_to_str(const sd_ctx_params_t* sd_ctx_params);
	impl.ctxParamsToStr, err = libSd.Prep("sd_ctx_params_to_str", &ffi.TypePointer, &FFITypeCtxParams)
	if err != nil {
		return nil, err
	}

	// SD_API sd_ctx_t* new_sd_ctx(const sd_ctx_params_t* sd_ctx_params);
	impl.newCtx, err = libSd.Prep("new_sd_ctx", &ffi.TypePointer, &FFITypeCtxParams)
	if err != nil {
		return nil, err
	}

	// SD_API void free_sd_ctx(sd_ctx_t* sd_ctx);
	impl.freeCtx, err = libSd.Prep("free_sd_ctx", &ffi.TypeVoid, &ffi.TypePointer)
	if err != nil {
		return nil, err
	}

	// SD_API void sd_sample_params_init(sd_sample_params_t* sample_params);
	impl.sampleParamsInit, err = libSd.Prep("sd_sample_params_init", &ffi.TypeVoid, &FFITypeSampleParams)
	if err != nil {
		return nil, err
	}

	// SD_API char* sd_sample_params_to_str(const sd_sample_params_t* sample_params);
	impl.sampleParamsToStr, err = libSd.Prep("sd_sample_params_to_str", &ffi.TypePointer, &FFITypeSampleParams)
	if err != nil {
		return nil, err
	}

	// SD_API enum sample_method_t sd_get_default_sample_method(const sd_ctx_t* sd_ctx);
	impl.getDefaultSampleMethod, err = libSd.Prep("sd_get_default_sample_method", &ffi.TypeSint32, &ffi.TypePointer)
	if err != nil {
		return nil, err
	}

	// SD_API enum scheduler_t sd_get_default_scheduler(const sd_ctx_t* sd_ctx, enum sample_method_t sample_method);
	impl.getDefaultScheduler, err = libSd.Prep("sd_get_default_scheduler", &ffi.TypeSint32, &ffi.TypePointer, &ffi.TypeSint32)
	if err != nil {
		return nil, err
	}

	// SD_API void sd_img_gen_params_init(sd_img_gen_params_t* sd_img_gen_params);
	impl.imgGenParamsInit, err = libSd.Prep("sd_img_gen_params_init", &ffi.TypeVoid, &FFITypeImgGenParams)
	if err != nil {
		return nil, err
	}

	// SD_API char* sd_img_gen_params_to_str(const sd_img_gen_params_t* sd_img_gen_params);
	impl.imgGenParamsToStr, err = libSd.Prep("sd_img_gen_params_to_str", &ffi.TypePointer, &FFITypeImgGenParams)
	if err != nil {
		return nil, err
	}

	// SD_API sd_image_t* generate_image(sd_ctx_t* sd_ctx, const sd_img_gen_params_t* sd_img_gen_params);
	impl.generateImage, err = libSd.Prep("generate_image", &ffi.TypePointer, &ffi.TypePointer, &FFITypeImgGenParams)
	if err != nil {
		return nil, err
	}

	// SD_API void sd_vid_gen_params_init(sd_vid_gen_params_t* sd_vid_gen_params);
	impl.vidGenParamsInit, err = libSd.Prep("sd_vid_gen_params_init", &ffi.TypeVoid, &FFITypeVidGenParams)
	if err != nil {
		return nil, err
	}

	// SD_API sd_image_t* generate_video(sd_ctx_t* sd_ctx, const sd_vid_gen_params_t* sd_vid_gen_params, int* num_frames_out);
	impl.generateVideo, err = libSd.Prep("generate_video", &ffi.TypePointer, &ffi.TypePointer, &FFITypeVidGenParams, &ffi.TypePointer)
	if err != nil {
		return nil, err
	}

	// SD_API upscaler_ctx_t* new_upscaler_ctx(const char* esrgan_path, bool offload_params_to_cpu, bool direct, int n_threads, int tile_size);
	impl.newUpscalerCtx, err = libSd.Prep("new_upscaler_ctx", &ffi.TypePointer, &ffi.TypePointer, &ffi.TypeUint8, &ffi.TypeUint8, &ffi.TypeSint32, &ffi.TypeSint32)
	if err != nil {
		return nil, err
	}

	// SD_API void free_upscaler_ctx(upscaler_ctx_t* upscaler_ctx);
	impl.freeUpscalerCtx, err = libSd.Prep("free_upscaler_ctx", &ffi.TypeVoid, &ffi.TypePointer)
	if err != nil {
		return nil, err
	}

	// SD_API sd_image_t* upscale(upscaler_ctx_t* upscaler_ctx, sd_image_t input_image, uint32_t upscale_factor);
	impl.upscale, err = libSd.Prep("upscale", &ffi.TypePointer, &ffi.TypePointer, &ffi.TypePointer, &ffi.TypeUint32)
	if err != nil {
		return nil, err
	}

	// SD_API int get_upscale_factor(upscaler_ctx_t* upscaler_ctx);
	impl.getUpscaleFactor, err = libSd.Prep("get_upscale_factor", &ffi.TypeSint32, &ffi.TypePointer)
	if err != nil {
		return nil, err
	}

	// SD_API bool convert(const char* input_path, const char* vae_path, const char* output_path, enum sd_type_t output_type, const char* tensor_type_rules, bool convert_name);
	impl.convert, err = libSd.Prep("convert", &ffi.TypeUint8, &ffi.TypePointer, &ffi.TypePointer, &ffi.TypePointer, &ffi.TypeSint32, &ffi.TypePointer, &ffi.TypeUint8)
	if err != nil {
		return nil, err
	}

	// SD_API const char* sd_commit(void);
	impl.commit, err = libSd.Prep("sd_commit", &ffi.TypePointer)
	if err != nil {
		return nil, err
	}

	// SD_API const char* sd_version(void);
	impl.version, err = libSd.Prep("sd_version", &ffi.TypePointer)
	if err != nil {
		return nil, err
	}

	return &impl, nil
}

// NewCtx 创建新的稳定扩散上下文
func (c *CStableDiffusionImpl) NewCtx(params *CtxParams) *CStableDiffusionCtx {
	var result uintptr
	c.newCtx.Call(unsafe.Pointer(&result), unsafe.Pointer(params))
	return &CStableDiffusionCtx{
		ctx: result,
	}
}

// FreeCtx 释放稳定扩散上下文
func (c *CStableDiffusionImpl) FreeCtx(ctx *CStableDiffusionCtx) {
	c.freeCtx.Call(nil, unsafe.Pointer(&ctx.ctx))
}

// CtxParamsInit 初始化上下文参数
func (c *CStableDiffusionImpl) CtxParamsInit() CtxParams {
	//这里需要分配到堆上，防止内存发生漂移
	structPtr := new(CtxParams)
	ptr := uintptr(unsafe.Pointer(structPtr))
	c.ctxParamsInit(ptr)
	return *structPtr
}

// SampleParamsInit 初始化采样参数
func (c *CStableDiffusionImpl) SampleParamsInit(params *SampleParams) {
	c.sampleParamsInit.Call(nil, unsafe.Pointer(params))
}

// ImgGenParamsInit 初始化图像生成参数
func (c *CStableDiffusionImpl) ImgGenParamsInit(params *ImgGenParams) {
	c.imgGenParamsInit.Call(nil, unsafe.Pointer(params))
}

// VidGenParamsInit 初始化视频生成参数
func (c *CStableDiffusionImpl) VidGenParamsInit(params *VidGenParams) {
	c.vidGenParamsInit.Call(nil, unsafe.Pointer(params))
}

// CacheParamsInit 初始化缓存参数
func (c *CStableDiffusionImpl) CacheParamsInit(params *CacheParams) {
	c.cacheParamsInit.Call(nil, unsafe.Pointer(params))
}

// CtxParamsToStr 将上下文参数转换为字符串
func (c *CStableDiffusionImpl) CtxParamsToStr(params *CtxParams) string {
	var result uintptr
	c.ctxParamsToStr.Call(unsafe.Pointer(&result), unsafe.Pointer(params))
	return goString(result)
}

// SampleParamsToStr 将采样参数转换为字符串
func (c *CStableDiffusionImpl) SampleParamsToStr(params *SampleParams) string {
	var result uintptr
	c.sampleParamsToStr.Call(unsafe.Pointer(&result), unsafe.Pointer(params))
	return goString(result)
}

// ImgGenParamsToStr 将图像生成参数转换为字符串
func (c *CStableDiffusionImpl) ImgGenParamsToStr(params *ImgGenParams) {
	var result uintptr
	c.imgGenParamsToStr.Call(unsafe.Pointer(&result), unsafe.Pointer(params))
	// 注意：根据接口声明，该方法不需要返回值
	_ = goString(result)
}

// GenerateImage 生成图像
func (c *CStableDiffusionImpl) GenerateImage(ctx *CStableDiffusionCtx, params *ImgGenParams) []Image {
	var result uintptr
	c.generateImage.Call(unsafe.Pointer(&result), unsafe.Pointer(&ctx.ctx), unsafe.Pointer(params))
	ptr := *(*unsafe.Pointer)(unsafe.Pointer(&result))
	if ptr == nil {
		return nil
	}
	img := (*Image)(ptr)
	return unsafe.Slice(img, int(params.BatchCount))
}

// GenerateVideo 生成视频
func (c *CStableDiffusionImpl) GenerateVideo(ctx *CStableDiffusionCtx, params *VidGenParams) ([]Image, int) {
	var result uintptr
	var numFrames int
	c.generateVideo.Call(unsafe.Pointer(&result), unsafe.Pointer(&ctx.ctx), unsafe.Pointer(params), unsafe.Pointer(&numFrames))
	// 注意：这里需要根据实际情况转换为Image切片
	// 目前暂时返回空切片，需要进一步实现
	return []Image{}, numFrames
}

// SetLogCallback 设置日志回调
func (c *CStableDiffusionImpl) SetLogCallback(cb LogCallback) {
	nada := uintptr(0)
	cCallback := purego.NewCallback(func(level int32, text uintptr, data uintptr) uintptr {
		cb(LogLevel(level), goString(text))
		return 0
	})
	c.setLogCallback.Call(nil, unsafe.Pointer(&cCallback), unsafe.Pointer(&nada))
}

// SetProgressCallback 设置进度回调
func (c *CStableDiffusionImpl) SetProgressCallback(cb ProgressCallback) {
	// 使用purego.NewCallback创建C兼容的回调函数
	cCallback := purego.NewCallback(func(step int32, steps int32, time float32, data uintptr) uintptr {
		cb(int(step), int(steps), time)
		return 0
	})
	// 直接传递回调函数地址，而不是地址的地址
	c.setProgressCallback.Call(nil, unsafe.Pointer(uintptr(cCallback)), unsafe.Pointer(nil))
}

// SetPreviewCallback 设置预览回调
func (c *CStableDiffusionImpl) SetPreviewCallback(cb PreviewCallback, mode Preview, interval int, denoised bool, noisy bool) {
	// 使用purego.NewCallback创建C兼容的回调函数
	cCallback := purego.NewCallback(func(step int32, frameCount int32, frames uintptr, isNoisy uint8, data uintptr) uintptr {
		// 注意：这里需要根据实际情况转换为Image切片
		// 目前暂时使用空切片，需要进一步实现
		cb(int(step), int(frameCount), []Image{}, isNoisy != 0)
		return 0
	})
	// 直接传递回调函数地址，而不是地址的地址
	c.setPreviewCallback.Call(nil, unsafe.Pointer(uintptr(cCallback)), unsafe.Pointer(&mode), unsafe.Pointer(&interval), unsafe.Pointer(&denoised), unsafe.Pointer(&noisy), unsafe.Pointer(nil))
}

// GetNumPhysicalCores 获取物理核心数量
func (c *CStableDiffusionImpl) GetNumPhysicalCores() int32 {
	var result int32
	c.getNumPhysicalCores.Call(unsafe.Pointer(&result))
	return result
}

// GetSystemInfo 获取系统信息
func (c *CStableDiffusionImpl) GetSystemInfo() string {
	var result uintptr
	c.getSystemInfo.Call(unsafe.Pointer(&result))
	return goString(result)
}

// GetCommit 获取提交信息
func (c *CStableDiffusionImpl) GetCommit() string {
	var result uintptr
	c.commit.Call(unsafe.Pointer(&result))
	return goString(result)
}

// GetVersion 获取版本信息
func (c *CStableDiffusionImpl) GetVersion() string {
	var result uintptr
	c.version.Call(unsafe.Pointer(&result))
	return goString(result)
}

// SdTypeName 获取SD类型名称
func (c *CStableDiffusionImpl) SdTypeName(sdType SdType) string {
	var result uintptr
	c.typeName.Call(unsafe.Pointer(&result), unsafe.Pointer(&sdType))
	return goString(result)
}

// StrToSdType 将字符串转换为SD类型
func (c *CStableDiffusionImpl) StrToSdType(str string) SdType {
	var result int32
	c.strToSdType.Call(unsafe.Pointer(&result), unsafe.Pointer(&str))
	return SdType(result)
}

// RngTypeName 获取RNG类型名称
func (c *CStableDiffusionImpl) RngTypeName(rngType RNGType) string {
	var result uintptr
	c.rngTypeName.Call(unsafe.Pointer(&result), unsafe.Pointer(&rngType))
	return goString(result)
}

// StrToRngType 将字符串转换为RNG类型
func (c *CStableDiffusionImpl) StrToRngType(str string) RNGType {
	var result int32
	c.strToRngType.Call(unsafe.Pointer(&result), unsafe.Pointer(&str))
	return RNGType(result)
}

// SampleMethodName 获取采样方法名称
func (c *CStableDiffusionImpl) SampleMethodName(sampleMethod SampleMethod) string {
	var result uintptr
	c.sampleMethodName.Call(unsafe.Pointer(&result), unsafe.Pointer(&sampleMethod))
	return goString(result)
}

// StrToSampleMethod 将字符串转换为采样方法
func (c *CStableDiffusionImpl) StrToSampleMethod(str string) SampleMethod {
	var result int32
	c.strToSampleMethod.Call(unsafe.Pointer(&result), unsafe.Pointer(&str))
	return SampleMethod(result)
}

// SchedulerName 获取调度器名称
func (c *CStableDiffusionImpl) SchedulerName(scheduler Scheduler) string {
	var result uintptr
	c.schedulerName.Call(unsafe.Pointer(&result), unsafe.Pointer(&scheduler))
	return goString(result)
}

// StrToScheduler 将字符串转换为调度器
func (c *CStableDiffusionImpl) StrToScheduler(str string) Scheduler {
	var result int32
	c.strToScheduler.Call(unsafe.Pointer(&result), unsafe.Pointer(&str))
	return Scheduler(result)
}

// PredictionName 获取预测类型名称
func (c *CStableDiffusionImpl) PredictionName(prediction Prediction) string {
	var result uintptr
	c.predictionName.Call(unsafe.Pointer(&result), unsafe.Pointer(&prediction))
	return goString(result)
}

// StrToPrediction 将字符串转换为预测类型
func (c *CStableDiffusionImpl) StrToPrediction(str string) Prediction {
	var result int32
	c.strToPrediction.Call(unsafe.Pointer(&result), unsafe.Pointer(&str))
	return Prediction(result)
}

// PreviewName 获取预览类型名称
func (c *CStableDiffusionImpl) PreviewName(preview Preview) string {
	var result uintptr
	c.previewName.Call(unsafe.Pointer(&result), unsafe.Pointer(&preview))
	return goString(result)
}

// StrToPreview 将字符串转换为预览类型
func (c *CStableDiffusionImpl) StrToPreview(str string) Preview {
	var result int32
	c.strToPreview.Call(unsafe.Pointer(&result), unsafe.Pointer(&str))
	return Preview(result)
}

// LoraApplyModeName 获取Lora应用模式名称
func (c *CStableDiffusionImpl) LoraApplyModeName(mode LoraApplyMode) string {
	var result uintptr
	c.loraApplyModeName.Call(unsafe.Pointer(&result), unsafe.Pointer(&mode))
	return goString(result)
}

// StrToLoraApplyMode 将字符串转换为Lora应用模式
func (c *CStableDiffusionImpl) StrToLoraApplyMode(str string) LoraApplyMode {
	var result int32
	c.strToLoraApplyMode.Call(unsafe.Pointer(&result), unsafe.Pointer(&str))
	return LoraApplyMode(result)
}

// GetDefaultSampleMethod 获取默认采样方法
func (c *CStableDiffusionImpl) GetDefaultSampleMethod(ctx *CStableDiffusionCtx) SampleMethod {
	var result int32
	c.getDefaultSampleMethod.Call(unsafe.Pointer(&result), unsafe.Pointer(&ctx.ctx))
	return SampleMethod(result)
}

// GetDefaultScheduler 获取默认调度器
func (c *CStableDiffusionImpl) GetDefaultScheduler(ctx *CStableDiffusionCtx, sampleMethod SampleMethod) Scheduler {
	var result int32
	c.getDefaultScheduler.Call(unsafe.Pointer(&result), unsafe.Pointer(&ctx.ctx), unsafe.Pointer(&sampleMethod))
	return Scheduler(result)
}

// NewUpscalerCtx 创建新的升频器上下文
func (c *CStableDiffusionImpl) NewUpscalerCtx(esrganPath string, offloadParamsToCpu bool, direct bool, nThreads int, tileSize int) *CUpScalerCtx {
	var result uintptr
	offload := uint8(0)
	if offloadParamsToCpu {
		offload = 1
	}
	directFlag := uint8(0)
	if direct {
		directFlag = 1
	}
	c.newUpscalerCtx.Call(unsafe.Pointer(&result), unsafe.Pointer(&esrganPath), unsafe.Pointer(&offload), unsafe.Pointer(&directFlag), unsafe.Pointer(&nThreads), unsafe.Pointer(&tileSize))
	return &CUpScalerCtx{
		ctx: result,
	}
}

// FreeUpscalerCtx 释放升频器上下文
func (c *CStableDiffusionImpl) FreeUpscalerCtx(ctx *CUpScalerCtx) {
	c.freeUpscalerCtx.Call(nil, unsafe.Pointer(&ctx.ctx))
}

// Upscale 升频图像
func (c *CStableDiffusionImpl) Upscale(ctx *CUpScalerCtx, inputImage Image, upscaleFactor uint32) Image {
	var result uintptr
	c.upscale.Call(unsafe.Pointer(&result), unsafe.Pointer(&ctx.ctx), unsafe.Pointer(&inputImage), unsafe.Pointer(&upscaleFactor))
	// 注意：这里需要根据实际情况转换为Image
	// 目前暂时返回空Image，需要进一步实现
	return Image{}
}

// GetUpscaleFactor 获取升频因子
func (c *CStableDiffusionImpl) GetUpscaleFactor(ctx *CUpScalerCtx) int {
	var result int32
	c.getUpscaleFactor.Call(unsafe.Pointer(&result), unsafe.Pointer(&ctx.ctx))
	return int(result)
}

// Convert 转换模型
func (c *CStableDiffusionImpl) Convert(inputPath, vaePath, outputPath string, outputType SdType, tensorTypeRules string, convertName bool) bool {
	var result uint8
	convertNameFlag := uint8(0)
	if convertName {
		convertNameFlag = 1
	}
	c.convert.Call(unsafe.Pointer(&result), unsafe.Pointer(&inputPath), unsafe.Pointer(&vaePath), unsafe.Pointer(&outputPath), unsafe.Pointer(&outputType), unsafe.Pointer(&tensorTypeRules), unsafe.Pointer(&convertNameFlag))
	return result != 0
}

// PreprocessCanny 预处理Canny边缘检测
func (c *CStableDiffusionImpl) PreprocessCanny(image Image, highThreshold, lowThreshold, weak, strong float32, inverse bool) bool {
	// 注意：这里需要调用对应的C函数，目前暂时返回false，需要进一步实现
	// 由于没有找到对应的ffi.Fun字段，可能需要添加到CStableDiffusionImpl结构体中
	return false
}

// Close 关闭资源
func (c *CStableDiffusionImpl) Close() error {
	// 注意：这里需要释放资源，目前暂时返回nil，需要进一步实现
	return nil
}

func (c *CStableDiffusionImpl) UpscaleImage(ctx *CUpScalerCtx, img Image, upscaleFactor uint32) Image {

	var result Image
	upscalerPtr := *(*unsafe.Pointer)(unsafe.Pointer(&ctx.ctx))
	ciPtr := uintptr(unsafe.Pointer(&img))

	c.upscale.Call(unsafe.Pointer(&result), unsafe.Pointer(&upscalerPtr), unsafe.Pointer(&ciPtr), unsafe.Pointer(&upscaleFactor))

	// ptr := *(*unsafe.Pointer)(unsafe.Pointer(&result))
	// if ptr == nil {
	// 	return Image{}
	// }
	// cimg := (*cImage)(ptr)
	// dataPtr := *(*unsafe.Pointer)(unsafe.Pointer(&cimg.data))
	// return Image{
	// 	Width:   cimg.width,
	// 	Height:  cimg.height,
	// 	Channel: cimg.channel,
	// 	Data:    unsafe.Slice((*byte)(dataPtr), cimg.channel*cimg.width*cimg.height),
	// }
	return result
}

// hasSuffix tests whether the string s ends with suffix.
func hasSuffix(s, suffix string) bool {
	return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
}

// CString converts a go string to *byte that can be passed to C code.
func CString(name string) *byte {
	if hasSuffix(name, "\x00") {
		return &(*(*[]byte)(unsafe.Pointer(&name)))[0]
	}
	b := make([]byte, len(name)+1)
	copy(b, name)
	return &b[0]
}

// GoString copies a null-terminated char* to a Go string.
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
