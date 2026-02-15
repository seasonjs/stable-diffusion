// Copyright (c) seasonjs. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package binding

// import (
// 	"unsafe"

// 	"github.com/ebitengine/purego"
// 	"github.com/jupiterrider/ffi"
// )

// import (
// 	"github.com/seasonjs/stable-diffusion/pkg/types"
// )

// C结构体对应的Go结构体定义

// type SlgParams struct {
// 	Layers     *int32
// 	LayerCount uintptr
// 	LayerStart float32
// 	LayerEnd   float32
// 	Scale      float32
// }

// type GuidanceParams struct {
// 	TxtCfg            float32
// 	ImgCfg            float32
// 	DistilledGuidance float32
// 	Slg               SlgParams
// }

// type SampleParams struct {
// 	Guidance          GuidanceParams
// 	Scheduler         types.Scheduler
// 	SampleMethod      types.SampleMethod
// 	SampleSteps       int32
// 	Eta               float32
// 	ShiftedTimestep   int32
// 	CustomSigmas      *float32
// 	CustomSigmasCount int32
// }

// type CStableDiffusion interface {
// 	// 上下文管理
// 	NewCtx(params *CtxParams) *CStableDiffusionCtx
// 	FreeCtx(ctx *CStableDiffusionCtx)

// 	// 参数初始化
// 	CtxParamsInit() CtxParams
// 	SampleParamsInit() SampleParams
// 	ImgGenParamsInit() ImgGenParams
// 	VidGenParamsInit() VidGenParams
// 	CacheParamsInit() CacheParams

// 	// 参数转换为字符串
// 	CtxParamsToStr(params *CtxParams) string
// 	SampleParamsToStr(params *SampleParams) string
// 	ImgGenParamsToStr(params *ImgGenParams)

// 	// 生成功能
// 	GenerateImage(ctx *CStableDiffusionCtx, params *ImgGenParams) []Image
// 	GenerateVideo(ctx *CStableDiffusionCtx, params *VidGenParams) ([]Image, int)

// 	// 回调设置
// 	SetLogCallback(cb LogCallback)
// 	SetProgressCallback(cb ProgressCallback)
// 	SetPreviewCallback(cb PreviewCallback, mode Preview, interval int, denoised bool, noisy bool)

// 	// 辅助函数
// 	GetNumPhysicalCores() int32
// 	GetSystemInfo() string
// 	GetCommit() string
// 	GetVersion() string

// 	// 类型转换和名称获取
// 	SdTypeName(sdType SdType) string
// 	StrToSdType(str string) SdType
// 	RngTypeName(rngType RNGType) string
// 	StrToRngType(str string) RNGType
// 	SampleMethodName(sampleMethod SampleMethod) string
// 	StrToSampleMethod(str string) SampleMethod
// 	SchedulerName(scheduler Scheduler) string
// 	StrToScheduler(str string) Scheduler
// 	PredictionName(prediction Prediction) string
// 	StrToPrediction(str string) Prediction
// 	PreviewName(preview Preview) string
// 	StrToPreview(str string) Preview
// 	LoraApplyModeName(mode LoraApplyMode) string
// 	StrToLoraApplyMode(str string) LoraApplyMode

// 	// 默认值获取
// 	GetDefaultSampleMethod(ctx *CStableDiffusionCtx) SampleMethod
// 	GetDefaultScheduler(ctx *CStableDiffusionCtx, sampleMethod SampleMethod) Scheduler

// 	// 升频功能
// 	NewUpscalerCtx(esrganPath string, offloadParamsToCpu bool, direct bool, nThreads int, tileSize int) *CUpScalerCtx
// 	FreeUpscalerCtx(ctx *CUpScalerCtx)
// 	Upscale(ctx *CUpScalerCtx, inputImage Image, upscaleFactor uint32) Image
// 	GetUpscaleFactor(ctx *CUpScalerCtx) int

// 	// 其他功能
// 	Convert(inputPath, vaePath, outputPath string, outputType SdType, tensorTypeRules string, convertName bool) bool
// 	PreprocessCanny(image Image, highThreshold, lowThreshold, weak, strong float32, inverse bool) bool

// 	// 关闭资源
// 	Close() error
// }

type CStableDiffusionImpl struct {
	// libSd ffi.Lib

	// // SD_API void sd_set_log_callback(sd_log_cb_t sd_log_cb, void* data);
	// setLogCallback ffi.Fun

	// // SD_API void sd_set_progress_callback(sd_progress_cb_t cb, void* data);
	// setProgressCallback ffi.Fun

	// // SD_API void sd_set_preview_callback(sd_preview_cb_t cb, enum preview_t mode, int interval, bool denoised, bool noisy, void* data);
	// setPreviewCallback ffi.Fun

	// // SD_API int32_t sd_get_num_physical_cores();
	// getNumPhysicalCores ffi.Fun

	// // SD_API const char* sd_get_system_info();
	// getSystemInfo ffi.Fun

	// // SD_API const char* sd_type_name(enum sd_type_t type);
	// typeName ffi.Fun

	// // SD_API enum sd_type_t str_to_sd_type(const char* str);
	// strToSdType ffi.Fun

	// // SD_API const char* sd_rng_type_name(enum rng_type_t rng_type);
	// rngTypeName ffi.Fun

	// // SD_API enum rng_type_t str_to_rng_type(const char* str);
	// strToRngType ffi.Fun

	// SD_API const char* sd_sample_method_name(enum sample_method_t sample_method);
	// sampleMethodName ffi.Fun

	// // SD_API enum sample_method_t str_to_sample_method(const char* str);
	// strToSampleMethod ffi.Fun

	// SD_API const char* sd_scheduler_name(enum scheduler_t scheduler);
	// schedulerName ffi.Fun

	// // SD_API enum scheduler_t str_to_scheduler(const char* str);
	// strToScheduler ffi.Fun

	// // SD_API const char* sd_prediction_name(enum prediction_t prediction);
	// predictionName ffi.Fun

	// // SD_API enum prediction_t str_to_prediction(const char* str);
	// strToPrediction ffi.Fun

	// // SD_API const char* sd_preview_name(enum preview_t preview);
	// previewName ffi.Fun

	// // SD_API enum preview_t str_to_preview(const char* str);
	// strToPreview ffi.Fun

	// // SD_API const char* sd_lora_apply_mode_name(enum lora_apply_mode_t mode);
	// loraApplyModeName ffi.Fun

	// // SD_API enum lora_apply_mode_t str_to_lora_apply_mode(const char* str);
	// strToLoraApplyMode ffi.Fun

	// // SD_API void sd_cache_params_init(sd_cache_params_t* cache_params);
	// cacheParamsInit func(uintptr)

	// SD_API void sd_ctx_params_init(sd_ctx_params_t* sd_ctx_params);
	// 这里因为C语言端进行了初始化：
	//  *sd_ctx_params                         = {};
	// 所以导致内存结构通过ffi无法对齐，所以需要使用更系统级别的调用，让c正确写入
	// ctxParamsInit func(uintptr)

	// // SD_API char* sd_ctx_params_to_str(const sd_ctx_params_t* sd_ctx_params);
	// ctxParamsToStr ffi.Fun

	// // SD_API sd_ctx_t* new_sd_ctx(const sd_ctx_params_t* sd_ctx_params);
	// newCtx ffi.Fun

	// // SD_API void free_sd_ctx(sd_ctx_t* sd_ctx);
	// freeCtx ffi.Fun

	// SD_API void sd_sample_params_init(sd_sample_params_t* sample_params);
	// sampleParamsInit func(uintptr)

	// // SD_API char* sd_sample_params_to_str(const sd_sample_params_t* sample_params);
	// sampleParamsToStr ffi.Fun

	// // SD_API enum sample_method_t sd_get_default_sample_method(const sd_ctx_t* sd_ctx);
	// getDefaultSampleMethod ffi.Fun

	// // SD_API enum scheduler_t sd_get_default_scheduler(const sd_ctx_t* sd_ctx, enum sample_method_t sample_method);
	// getDefaultScheduler ffi.Fun

	// // SD_API void sd_img_gen_params_init(sd_img_gen_params_t* sd_img_gen_params);
	// imgGenParamsInit func(uintptr)

	// // SD_API char* sd_img_gen_params_to_str(const sd_img_gen_params_t* sd_img_gen_params);
	// imgGenParamsToStr ffi.Fun

	// // SD_API sd_image_t* generate_image(sd_ctx_t* sd_ctx, const sd_img_gen_params_t* sd_img_gen_params);
	// generateImage ffi.Fun

	// // SD_API void sd_vid_gen_params_init(sd_vid_gen_params_t* sd_vid_gen_params);
	// vidGenParamsInit func(uintptr)

	// // SD_API sd_image_t* generate_video(sd_ctx_t* sd_ctx, const sd_vid_gen_params_t* sd_vid_gen_params, int* num_frames_out);
	// generateVideo ffi.Fun

	// SD_API upscaler_ctx_t* new_upscaler_ctx(const char* esrgan_path, bool offload_params_to_cpu, bool direct, int n_threads, int tile_size);
	// newUpscalerCtx ffi.Fun

	// // SD_API void free_upscaler_ctx(upscaler_ctx_t* upscaler_ctx);
	// freeUpscalerCtx ffi.Fun

	// // SD_API sd_image_t* upscale(upscaler_ctx_t* upscaler_ctx, sd_image_t input_image, uint32_t upscale_factor);
	// upscale ffi.Fun

	// // SD_API int get_upscale_factor(upscaler_ctx_t* upscaler_ctx);
	// getUpscaleFactor ffi.Fun

	// // SD_API bool convert(const char* input_path, const char* vae_path, const char* output_path, enum sd_type_t output_type, const char* tensor_type_rules, bool convert_name);
	// convert ffi.Fun

	// SD_API bool preprocess_canny(sd_image_t image,float high_threshold,float low_threshold,float weak,float strong,bool inverse);
	// preprocessCanny ffi.Fun

	// // SD_API const char* sd_commit(void);
	// commit ffi.Fun

	// // SD_API const char* sd_version(void);
	// version ffi.Fun
}

func NewCStableDiffusion(libraryPath string) (*CStableDiffusionImpl, error) {
	// // 使用ffi包加载动态库
	// libSd, err := ffi.Load(libraryPath)

	// if err != nil {
	// 	return nil, err
	// }

	// impl := CStableDiffusionImpl{
	// 	libSd: libSd,
	// }

	// 注册所有C函数
	// // SD_API void sd_set_log_callback(sd_log_cb_t sd_log_cb, void* data);
	// impl.setLogCallback, err = libSd.Prep("sd_set_log_callback", &ffi.TypeVoid, &ffi.TypePointer, &ffi.TypePointer)
	// if err != nil {
	// 	return nil, err
	// }

	// // SD_API void sd_set_progress_callback(sd_progress_cb_t cb, void* data);
	// impl.setProgressCallback, err = libSd.Prep("sd_set_progress_callback", &ffi.TypeVoid, &ffi.TypePointer, &ffi.TypePointer)
	// if err != nil {
	// 	return nil, err
	// }

	// SD_API void sd_set_preview_callback(sd_preview_cb_t cb, enum preview_t mode, int interval, bool denoised, bool noisy, void* data);
	// impl.setPreviewCallback, err = libSd.Prep("sd_set_preview_callback", &ffi.TypeVoid, &ffi.TypePointer, &ffi.TypeSint32, &ffi.TypeSint32, &ffi.TypeUint8, &ffi.TypeUint8, &ffi.TypePointer)
	// if err != nil {
	// 	return nil, err
	// }

	// // SD_API int32_t sd_get_num_physical_cores();
	// impl.getNumPhysicalCores, err = libSd.Prep("sd_get_num_physical_cores", &ffi.TypeSint32)
	// if err != nil {
	// 	return nil, err
	// }

	// SD_API const char* sd_get_system_info();
	// impl.getSystemInfo, err = libSd.Prep("sd_get_system_info", &ffi.TypePointer)
	// if err != nil {
	// 	return nil, err
	// }

	// SD_API const char* sd_type_name(enum sd_type_t type);
	// impl.typeName, err = libSd.Prep("sd_type_name", &ffi.TypePointer, &ffi.TypeSint32)
	// if err != nil {
	// 	return nil, err
	// }

	// // SD_API enum sd_type_t str_to_sd_type(const char* str);
	// impl.strToSdType, err = libSd.Prep("str_to_sd_type", &ffi.TypeSint32, &ffi.TypePointer)
	// if err != nil {
	// 	return nil, err
	// }

	// SD_API const char* sd_rng_type_name(enum rng_type_t rng_type);
	// impl.rngTypeName, err = libSd.Prep("sd_rng_type_name", &ffi.TypePointer, &ffi.TypeSint32)
	// if err != nil {
	// 	return nil, err
	// }

	// // SD_API enum rng_type_t str_to_rng_type(const char* str);
	// impl.strToRngType, err = libSd.Prep("str_to_rng_type", &ffi.TypeSint32, &ffi.TypePointer)
	// if err != nil {
	// 	return nil, err
	// }

	// SD_API const char* sd_sample_method_name(enum sample_method_t sample_method);
	// impl.sampleMethodName, err = libSd.Prep("sd_sample_method_name", &ffi.TypePointer, &ffi.TypeSint32)
	// if err != nil {
	// 	return nil, err
	// }

	// // SD_API enum sample_method_t str_to_sample_method(const char* str);
	// impl.strToSampleMethod, err = libSd.Prep("str_to_sample_method", &ffi.TypeSint32, &ffi.TypePointer)
	// if err != nil {
	// 	return nil, err
	// }

	// //SD_API const char* sd_scheduler_name(enum scheduler_t scheduler);
	// impl.schedulerName, err = libSd.Prep("sd_scheduler_name", &ffi.TypePointer, &ffi.TypeSint32)
	// if err != nil {
	// 	return nil, err
	// }

	// // SD_API enum scheduler_t str_to_scheduler(const char* str);
	// impl.strToScheduler, err = libSd.Prep("str_to_scheduler", &ffi.TypeSint32, &ffi.TypePointer)
	// if err != nil {
	// 	return nil, err
	// }

	// // SD_API const char* sd_prediction_name(enum prediction_t prediction);
	// impl.predictionName, err = libSd.Prep("sd_prediction_name", &ffi.TypePointer, &ffi.TypeSint32)
	// if err != nil {
	// 	return nil, err
	// }

	// // SD_API enum prediction_t str_to_prediction(const char* str);
	// impl.strToPrediction, err = libSd.Prep("str_to_prediction", &ffi.TypeSint32, &ffi.TypePointer)
	// if err != nil {
	// 	return nil, err
	// }

	// // SD_API const char* sd_preview_name(enum preview_t preview);
	// impl.previewName, err = libSd.Prep("sd_preview_name", &ffi.TypePointer, &ffi.TypeSint32)
	// if err != nil {
	// 	return nil, err
	// }

	// // SD_API enum preview_t str_to_preview(const char* str);
	// impl.strToPreview, err = libSd.Prep("str_to_preview", &ffi.TypeSint32, &ffi.TypePointer)
	// if err != nil {
	// 	return nil, err
	// }

	// // SD_API const char* sd_lora_apply_mode_name(enum lora_apply_mode_t mode);
	// impl.loraApplyModeName, err = libSd.Prep("sd_lora_apply_mode_name", &ffi.TypePointer, &ffi.TypeSint32)
	// if err != nil {
	// 	return nil, err
	// }

	// // SD_API enum lora_apply_mode_t str_to_lora_apply_mode(const char* str);
	// impl.strToLoraApplyMode, err = libSd.Prep("str_to_lora_apply_mode", &ffi.TypeSint32, &ffi.TypePointer)
	// if err != nil {
	// 	return nil, err
	// }

	// // SD_API void sd_cache_params_init(sd_cache_params_t* cache_params);
	// // 这种通过c初始化结构体的操作使用purego 进行
	// purego.RegisterLibFunc(&impl.cacheParamsInit, libSd.Addr, "sd_cache_params_init")

	// SD_API void sd_ctx_params_init(sd_ctx_params_t* sd_ctx_params);
	// 这种通过c初始化结构体的操作使用purego 进行
	// purego.RegisterLibFunc(&impl.ctxParamsInit, libSd.Addr, "sd_ctx_params_init")

	// // SD_API char* sd_ctx_params_to_str(const sd_ctx_params_t* sd_ctx_params);
	// impl.ctxParamsToStr, err = libSd.Prep("sd_ctx_params_to_str", &ffi.TypePointer, &FFITypeCtxParams)
	// if err != nil {
	// 	return nil, err
	// }

	// // SD_API sd_ctx_t* new_sd_ctx(const sd_ctx_params_t* sd_ctx_params);
	// impl.newCtx, err = libSd.Prep("new_sd_ctx", &ffi.TypePointer, &FFITypeCtxParams)
	// if err != nil {
	// 	return nil, err
	// }

	// // SD_API void free_sd_ctx(sd_ctx_t* sd_ctx);
	// impl.freeCtx, err = libSd.Prep("free_sd_ctx", &ffi.TypeVoid, &ffi.TypePointer)
	// if err != nil {
	// 	return nil, err
	// }

	// SD_API void sd_sample_params_init(sd_sample_params_t* sample_params);
	// 这种通过c初始化结构体的操作使用purego 进行
	// purego.RegisterLibFunc(&impl.sampleParamsInit, libSd.Addr, "sd_sample_params_init")

	// // SD_API char* sd_sample_params_to_str(const sd_sample_params_t* sample_params);
	// impl.sampleParamsToStr, err = libSd.Prep("sd_sample_params_to_str", &ffi.TypePointer, &FFITypeSampleParams)
	// if err != nil {
	// 	return nil, err
	// }

	// // SD_API enum sample_method_t sd_get_default_sample_method(const sd_ctx_t* sd_ctx);
	// impl.getDefaultSampleMethod, err = libSd.Prep("sd_get_default_sample_method", &ffi.TypeSint32, &ffi.TypePointer)
	// if err != nil {
	// 	return nil, err
	// }

	// // SD_API enum scheduler_t sd_get_default_scheduler(const sd_ctx_t* sd_ctx, enum sample_method_t sample_method);
	// impl.getDefaultScheduler, err = libSd.Prep("sd_get_default_scheduler", &ffi.TypeSint32, &ffi.TypePointer, &ffi.TypeSint32)
	// if err != nil {
	// 	return nil, err
	// }

	// SD_API void sd_img_gen_params_init(sd_img_gen_params_t* sd_img_gen_params);
	// purego.RegisterLibFunc(&impl.imgGenParamsInit, libSd.Addr, "sd_img_gen_params_init")

	// // SD_API char* sd_img_gen_params_to_str(const sd_img_gen_params_t* sd_img_gen_params);
	// impl.imgGenParamsToStr, err = libSd.Prep("sd_img_gen_params_to_str", &ffi.TypePointer, &FFITypeImgGenParams)
	// if err != nil {
	// 	return nil, err
	// }

	// // SD_API sd_image_t* generate_image(sd_ctx_t* sd_ctx, const sd_img_gen_params_t* sd_img_gen_params);
	// impl.generateImage, err = libSd.Prep("generate_image", &ffi.TypePointer, &ffi.TypePointer, &FFITypeImgGenParams)
	// if err != nil {
	// 	return nil, err
	// }

	// // SD_API void sd_vid_gen_params_init(sd_vid_gen_params_t* sd_vid_gen_params);
	// purego.RegisterLibFunc(&impl.vidGenParamsInit, libSd.Addr, "sd_vid_gen_params_init")

	// // SD_API sd_image_t* generate_video(sd_ctx_t* sd_ctx, const sd_vid_gen_params_t* sd_vid_gen_params, int* num_frames_out);
	// impl.generateVideo, err = libSd.Prep("generate_video", &ffi.TypePointer, &ffi.TypePointer, &FFITypeVidGenParams, &ffi.TypePointer)
	// if err != nil {
	// 	return nil, err
	// }

	// // SD_API upscaler_ctx_t* new_upscaler_ctx(const char* esrgan_path, bool offload_params_to_cpu, bool direct, int n_threads, int tile_size);
	// impl.newUpscalerCtx, err = libSd.Prep("new_upscaler_ctx", &ffi.TypePointer, &ffi.TypePointer, &ffi.TypeUint8, &ffi.TypeUint8, &ffi.TypeSint32, &ffi.TypeSint32)
	// if err != nil {
	// 	return nil, err
	// }

	// // SD_API void free_upscaler_ctx(upscaler_ctx_t* upscaler_ctx);
	// impl.freeUpscalerCtx, err = libSd.Prep("free_upscaler_ctx", &ffi.TypeVoid, &ffi.TypePointer)
	// if err != nil {
	// 	return nil, err
	// }

	// // SD_API sd_image_t* upscale(upscaler_ctx_t* upscaler_ctx, sd_image_t input_image, uint32_t upscale_factor);
	// impl.upscale, err = libSd.Prep("upscale", &ffi.TypePointer, &ffi.TypePointer, &ffi.TypePointer, &ffi.TypeUint32)
	// if err != nil {
	// 	return nil, err
	// }

	// // SD_API int get_upscale_factor(upscaler_ctx_t* upscaler_ctx);
	// impl.getUpscaleFactor, err = libSd.Prep("get_upscale_factor", &ffi.TypeSint32, &ffi.TypePointer)
	// if err != nil {
	// 	return nil, err
	// }

	// SD_API bool convert(const char* input_path, const char* vae_path, const char* output_path, enum sd_type_t output_type, const char* tensor_type_rules, bool convert_name);
	// impl.convert, err = libSd.Prep("convert", &ffi.TypeUint32, &ffi.TypePointer, &ffi.TypePointer, &ffi.TypePointer, &ffi.TypeSint32, &ffi.TypePointer, &ffi.TypeUint8)
	// if err != nil {
	// 	return nil, err
	// }

	// // SD_API bool preprocess_canny(sd_image_t image,float high_threshold,float low_threshold,float weak,float strong,bool inverse);
	// impl.preprocessCanny, err = libSd.Prep("preprocess_canny", &ffi.TypeUint32, &ffi.TypeFloat, &ffi.TypeFloat, &ffi.TypeFloat, &ffi.TypeFloat, &ffi.TypeUint8)
	// if err != nil {
	// 	return nil, err
	// }

	// // SD_API const char* sd_commit(void);
	// impl.commit, err = libSd.Prep("sd_commit", &ffi.TypePointer)
	// if err != nil {
	// 	return nil, err
	// }

	// // SD_API const char* sd_version(void);
	// impl.version, err = libSd.Prep("sd_version", &ffi.TypePointer)
	// if err != nil {
	// 	return nil, err
	// }

	// return &impl, nil
	return nil, nil
}

// NewCtx 创建新的稳定扩散上下文
// func (c *CStableDiffusionImpl) NewCtx(params *CtxParams) *CStableDiffusionCtx {
// 	var result uintptr
// 	c.newCtx.Call(unsafe.Pointer(&result), unsafe.Pointer(params))
// 	return &CStableDiffusionCtx{
// 		ctx: result,
// 	}
// }

// FreeCtx 释放稳定扩散上下文
// func (c *CStableDiffusionImpl) FreeCtx(ctx *CStableDiffusionCtx) {
// 	c.freeCtx.Call(nil, unsafe.Pointer(&ctx.ctx))
// }

// CtxParamsInit 初始化上下文参数
// func (c *CStableDiffusionImpl) CtxParamsInit() CtxParams {
// 	//这里需要分配到堆上，防止内存发生漂移
// 	structPtr := new(CtxParams)
// 	ptr := uintptr(unsafe.Pointer(structPtr))
// 	c.ctxParamsInit(ptr)
// 	return *structPtr
// }

// SampleParamsInit 初始化采样参数
// func (c *CStableDiffusionImpl) SampleParamsInit() SampleParams {
// 	//这里需要分配到堆上，防止内存发生漂移
// 	structPtr := new(SampleParams)
// 	ptr := uintptr(unsafe.Pointer(structPtr))
// 	c.sampleParamsInit(ptr)
// 	return *structPtr
// }

// // ImgGenParamsInit 初始化图像生成参数
// func (c *CStableDiffusionImpl) ImgGenParamsInit() ImgGenParams {
// 	//这里需要分配到堆上，防止内存发生漂移
// 	structPtr := new(ImgGenParams)
// 	ptr := uintptr(unsafe.Pointer(structPtr))
// 	c.imgGenParamsInit(ptr)
// 	return *structPtr
// }

// // CacheParamsInit 初始化缓存参数
// func (c *CStableDiffusionImpl) CacheParamsInit() CacheParams {
// 	//这里需要分配到堆上，防止内存发生漂移
// 	structPtr := new(CacheParams)
// 	ptr := uintptr(unsafe.Pointer(structPtr))
// 	c.cacheParamsInit(ptr)
// 	return *structPtr
// }

// CtxParamsToStr 将上下文参数转换为字符串
// func (c *CStableDiffusionImpl) CtxParamsToStr(params *CtxParams) string {
// 	var result uintptr
// 	c.ctxParamsToStr.Call(unsafe.Pointer(&result), unsafe.Pointer(params))
// 	return goString(result)
// }

// // SampleParamsToStr 将采样参数转换为字符串
// func (c *CStableDiffusionImpl) SampleParamsToStr(params *SampleParams) string {
// 	var result uintptr
// 	c.sampleParamsToStr.Call(unsafe.Pointer(&result), unsafe.Pointer(params))
// 	return goString(result)
// }

// // ImgGenParamsToStr 将图像生成参数转换为字符串
// func (c *CStableDiffusionImpl) ImgGenParamsToStr(params *ImgGenParams) {
// 	var result uintptr
// 	c.imgGenParamsToStr.Call(unsafe.Pointer(&result), unsafe.Pointer(params))
// 	// 注意：根据接口声明，该方法不需要返回值
// 	_ = goString(result)
// }

// // GenerateImage 生成图像
// func (c *CStableDiffusionImpl) GenerateImage(ctx *CStableDiffusionCtx, params *ImgGenParams) []Image {
// 	var result uintptr
// 	c.generateImage.Call(unsafe.Pointer(&result), unsafe.Pointer(&ctx.ctx), unsafe.Pointer(params))
// 	ptr := *(*unsafe.Pointer)(unsafe.Pointer(&result))
// 	if ptr == nil {
// 		return nil
// 	}
// 	img := (*Image)(ptr)
// 	return unsafe.Slice(img, int(params.BatchCount))
// }

// // GenerateVideo 生成视频
// func (c *CStableDiffusionImpl) GenerateVideo(ctx *CStableDiffusionCtx, params *VidGenParams) ([]Image, int) {
// 	var result uintptr
// 	var numFrames int
// 	c.generateVideo.Call(unsafe.Pointer(&result), unsafe.Pointer(&ctx.ctx), unsafe.Pointer(params), unsafe.Pointer(&numFrames))
// 	// 注意：这里需要根据实际情况转换为Image切片
// 	// 目前暂时返回空切片，需要进一步实现
// 	return []Image{}, numFrames
// }

// SetLogCallback 设置日志回调
// func (c *CStableDiffusionImpl) SetLogCallback(cb LogCallback) {
// 	nada := uintptr(0)
// 	cCallback := purego.NewCallback(func(level int32, text uintptr, data uintptr) uintptr {
// 		cb(types.LogLevel(level), goString(text))
// 		return 0
// 	})
// 	c.setLogCallback.Call(nil, unsafe.Pointer(&cCallback), unsafe.Pointer(&nada))
// }

// SetProgressCallback 设置进度回调
// func (c *CStableDiffusionImpl) SetProgressCallback(cb ProgressCallback) {
// 	// 使用purego.NewCallback创建C兼容的回调函数
// 	cCallback := purego.NewCallback(func(step int32, steps int32, time float32, data uintptr) uintptr {
// 		cb(int(step), int(steps), time)
// 		return 0
// 	})
// 	// 直接传递回调函数地址，而不是地址的地址
// 	c.setProgressCallback.Call(nil, unsafe.Pointer(cCallback), unsafe.Pointer(nil))
// }

// SetPreviewCallback 设置预览回调
// func (c *CStableDiffusionImpl) SetPreviewCallback(cb PreviewCallback, mode types.Preview, interval int, denoised bool, noisy bool) {
// 	// 使用purego.NewCallback创建C兼容的回调函数
// 	cCallback := purego.NewCallback(func(step int32, frameCount int32, frames uintptr, isNoisy uint8, data uintptr) uintptr {
// 		// 注意：这里需要根据实际情况转换为Image切片
// 		// 目前暂时使用空切片，需要进一步实现
// 		cb(int(step), int(frameCount), []Image{}, isNoisy != 0)
// 		return 0
// 	})
// 	// 直接传递回调函数地址，而不是地址的地址
// 	c.setPreviewCallback.Call(nil, unsafe.Pointer(cCallback), unsafe.Pointer(&mode), unsafe.Pointer(&interval), unsafe.Pointer(&denoised), unsafe.Pointer(&noisy), unsafe.Pointer(nil))
// }

// // GetNumPhysicalCores 获取物理核心数量
// func (c *CStableDiffusionImpl) GetNumPhysicalCores() int32 {
// 	var result int32
// 	c.getNumPhysicalCores.Call(unsafe.Pointer(&result))
// 	return result
// }

// GetSystemInfo 获取系统信息
// func (c *CStableDiffusionImpl) GetSystemInfo() string {
// 	var result uintptr
// 	c.getSystemInfo.Call(unsafe.Pointer(&result))
// 	return goString(result)
// }

// GetCommit 获取提交信息
// func (c *CStableDiffusionImpl) GetCommit() string {
// 	var result uintptr
// 	c.commit.Call(unsafe.Pointer(&result))
// 	return goString(result)
// }

// GetVersion 获取版本信息
// func (c *CStableDiffusionImpl) GetVersion() string {
// 	var result uintptr
// 	c.version.Call(unsafe.Pointer(&result))
// 	return goString(result)
// }

// // SdTypeName 获取SD类型名称
// func (c *CStableDiffusionImpl) SdTypeName(sdType types.SdType) string {
// 	var result uintptr
// 	c.typeName.Call(unsafe.Pointer(&result), unsafe.Pointer(&sdType))
// 	return goString(result)
// }

// // StrToSdType 将字符串转换为SD类型
// func (c *CStableDiffusionImpl) StrToSdType(str string) types.SdType {
// 	var result int32
// 	c.strToSdType.Call(unsafe.Pointer(&result), unsafe.Pointer(&str))
// 	return types.SdType(result)
// }

// // RngTypeName 获取RNG类型名称
// func (c *CStableDiffusionImpl) RngTypeName(rngType types.RNGType) string {
// 	var result uintptr
// 	c.rngTypeName.Call(unsafe.Pointer(&result), unsafe.Pointer(&rngType))
// 	return goString(result)
// }

// // StrToRngType 将字符串转换为RNG类型
// func (c *CStableDiffusionImpl) StrToRngType(str string) types.RNGType {
// 	var result int32
// 	c.strToRngType.Call(unsafe.Pointer(&result), unsafe.Pointer(&str))
// 	return types.RNGType(result)
// }

// SampleMethodName 获取采样方法名称
// func (c *CStableDiffusionImpl) SampleMethodName(sampleMethod types.SampleMethod) string {
// 	var result uintptr
// 	c.sampleMethodName.Call(unsafe.Pointer(&result), unsafe.Pointer(&sampleMethod))
// 	return goString(result)
// }

// // StrToSampleMethod 将字符串转换为采样方法
// func (c *CStableDiffusionImpl) StrToSampleMethod(str string) types.SampleMethod {
// 	var result int32
// 	c.strToSampleMethod.Call(unsafe.Pointer(&result), unsafe.Pointer(&str))
// 	return types.SampleMethod(result)
// }

// SchedulerName 获取调度器名称
// func (c *CStableDiffusionImpl) SchedulerName(scheduler types.Scheduler) string {
// 	var result uintptr
// 	c.schedulerName.Call(unsafe.Pointer(&result), unsafe.Pointer(&scheduler))
// 	return goString(result)
// }

// // StrToScheduler 将字符串转换为调度器
// func (c *CStableDiffusionImpl) StrToScheduler(str string) types.Scheduler {
// 	var result int32
// 	c.strToScheduler.Call(unsafe.Pointer(&result), unsafe.Pointer(&str))
// 	return types.Scheduler(result)
// }

// PredictionName 获取预测类型名称
// func (c *CStableDiffusionImpl) PredictionName(prediction types.Prediction) string {
// 	var result uintptr
// 	c.predictionName.Call(unsafe.Pointer(&result), unsafe.Pointer(&prediction))
// 	return goString(result)
// }

// // StrToPrediction 将字符串转换为预测类型
// func (c *CStableDiffusionImpl) StrToPrediction(str string) types.Prediction {
// 	var result int32
// 	c.strToPrediction.Call(unsafe.Pointer(&result), unsafe.Pointer(&str))
// 	return types.Prediction(result)
// }

// // PreviewName 获取预览类型名称
// func (c *CStableDiffusionImpl) PreviewName(preview types.Preview) string {
// 	var result uintptr
// 	c.previewName.Call(unsafe.Pointer(&result), unsafe.Pointer(&preview))
// 	return goString(result)
// }

// // StrToPreview 将字符串转换为预览类型
// func (c *CStableDiffusionImpl) StrToPreview(str string) types.Preview {
// 	var result int32
// 	c.strToPreview.Call(unsafe.Pointer(&result), unsafe.Pointer(&str))
// 	return types.Preview(result)
// }

// // LoraApplyModeName 获取Lora应用模式名称
// func (c *CStableDiffusionImpl) LoraApplyModeName(mode types.LoraApplyMode) string {
// 	var result uintptr
// 	c.loraApplyModeName.Call(unsafe.Pointer(&result), unsafe.Pointer(&mode))
// 	return goString(result)
// }

// // StrToLoraApplyMode 将字符串转换为Lora应用模式
// func (c *CStableDiffusionImpl) StrToLoraApplyMode(str string) types.LoraApplyMode {
// 	var result int32
// 	c.strToLoraApplyMode.Call(unsafe.Pointer(&result), unsafe.Pointer(&str))
// 	return types.LoraApplyMode(result)
// }

// // GetDefaultSampleMethod 获取默认采样方法
// func (c *CStableDiffusionImpl) GetDefaultSampleMethod(ctx *CStableDiffusionCtx) types.SampleMethod {
// 	var result int32
// 	c.getDefaultSampleMethod.Call(unsafe.Pointer(&result), unsafe.Pointer(&ctx.ctx))
// 	return types.SampleMethod(result)
// }

// // GetDefaultScheduler 获取默认调度器
// func (c *CStableDiffusionImpl) GetDefaultScheduler(ctx *CStableDiffusionCtx, sampleMethod types.SampleMethod) types.Scheduler {
// 	var result int32
// 	c.getDefaultScheduler.Call(unsafe.Pointer(&result), unsafe.Pointer(&ctx.ctx), unsafe.Pointer(&sampleMethod))
// 	return types.Scheduler(result)
// }

// NewUpscalerCtx 创建新的升频器上下文
// func (c *CStableDiffusionImpl) NewUpscalerCtx(esrganPath string, offloadParamsToCpu bool, direct bool, nThreads int, tileSize int) *CUpScalerCtx {
// 	var result uintptr
// 	offload := uint8(0)
// 	if offloadParamsToCpu {
// 		offload = 1
// 	}
// 	directFlag := uint8(0)
// 	if direct {
// 		directFlag = 1
// 	}
// 	c.newUpscalerCtx.Call(unsafe.Pointer(&result), unsafe.Pointer(&esrganPath), unsafe.Pointer(&offload), unsafe.Pointer(&directFlag), unsafe.Pointer(&nThreads), unsafe.Pointer(&tileSize))
// 	return &CUpScalerCtx{
// 		ctx: result,
// 	}
// }

// FreeUpscalerCtx 释放升频器上下文
// func (c *CStableDiffusionImpl) FreeUpscalerCtx(ctx *CUpScalerCtx) {
// 	c.freeUpscalerCtx.Call(nil, unsafe.Pointer(&ctx.ctx))
// }

// // Upscale 升频图像
// func (c *CStableDiffusionImpl) Upscale(ctx *CUpScalerCtx, inputImage Image, upscaleFactor uint32) Image {
// 	var result uintptr
// 	c.upscale.Call(unsafe.Pointer(&result), unsafe.Pointer(&ctx.ctx), unsafe.Pointer(&inputImage), unsafe.Pointer(&upscaleFactor))
// 	// 注意：这里需要根据实际情况转换为Image
// 	// 目前暂时返回空Image，需要进一步实现
// 	return Image{}
// }

// // GetUpscaleFactor 获取升频因子
// func (c *CStableDiffusionImpl) GetUpscaleFactor(ctx *CUpScalerCtx) int {
// 	var result int32
// 	c.getUpscaleFactor.Call(unsafe.Pointer(&result), unsafe.Pointer(&ctx.ctx))
// 	return int(result)
// }

// Convert 转换模型
// func (c *CStableDiffusionImpl) Convert(inputPath, vaePath, outputPath string, outputType types.SdType, tensorTypeRules string, convertName bool) bool {
// 	var result ffi.Arg
// 	c.convert.Call(unsafe.Pointer(&result), unsafe.Pointer(&inputPath), unsafe.Pointer(&vaePath), unsafe.Pointer(&outputPath), unsafe.Pointer(&outputType), unsafe.Pointer(&tensorTypeRules), unsafe.Pointer(&convertName))
// 	return result.Bool()
// }

// PreprocessCanny 预处理Canny边缘检测
// func (c *CStableDiffusionImpl) PreprocessCanny(image Image, highThreshold, lowThreshold, weak, strong float32, inverse bool) bool {
// 	var result ffi.Arg
// 	c.preprocessCanny.Call(unsafe.Pointer(&result), unsafe.Pointer(&image), unsafe.Pointer(&highThreshold), unsafe.Pointer(&lowThreshold), unsafe.Pointer(&weak), unsafe.Pointer(&strong), unsafe.Pointer(&inverse))
// 	return result.Bool()
// }

// Close 关闭资源
// func (c *CStableDiffusionImpl) Close(ctx *CStableDiffusionCtx) error {
// 	c.freeCtx.Call(nil, unsafe.Pointer(&ctx.ctx))
// 	return nil
// }

// func (c *CStableDiffusionImpl) UpscaleImage(ctx *CUpScalerCtx, img Image, upscaleFactor uint32) Image {

// 	var result Image
// 	upscalerPtr := *(*unsafe.Pointer)(unsafe.Pointer(&ctx.ctx))
// 	ciPtr := uintptr(unsafe.Pointer(&img))

// 	c.upscale.Call(unsafe.Pointer(&result), unsafe.Pointer(&upscalerPtr), unsafe.Pointer(&ciPtr), unsafe.Pointer(&upscaleFactor))

// 	// ptr := *(*unsafe.Pointer)(unsafe.Pointer(&result))
// 	// if ptr == nil {
// 	// 	return Image{}
// 	// }
// 	// cimg := (*cImage)(ptr)
// 	// dataPtr := *(*unsafe.Pointer)(unsafe.Pointer(&cimg.data))
// 	// return Image{
// 	// 	Width:   cimg.width,
// 	// 	Height:  cimg.height,
// 	// 	Channel: cimg.channel,
// 	// 	Data:    unsafe.Slice((*byte)(dataPtr), cimg.channel*cimg.width*cimg.height),
// 	// }
// 	return result
// }
