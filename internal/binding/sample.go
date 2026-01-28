package binding

import (
	"github.com/ebitengine/purego"
	"structs"
	"unsafe"

	"github.com/jupiterrider/ffi"
	"github.com/seasonjs/stable-diffusion/pkg/types"
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

type SlgParams struct {
	_          structs.HostLayout
	Layers     *int32
	LayerCount int64
	LayerStart float32
	LayerEnd   float32
	Scale      float32
}

type GuidanceParams struct {
	_                 structs.HostLayout
	TxtCfg            float32
	ImgCfg            float32
	DistilledGuidance float32
	Slg               SlgParams
}

type SampleParams struct {
	_                 structs.HostLayout
	Guidance          GuidanceParams
	Scheduler         int32
	SampleMethod      int32
	SampleSteps       int32
	Eta               float32
	ShiftedTimestep   int32
	CustomSigmas      *float32
	CustomSigmasCount int32
}

var (
	// SD_API const char* sd_sample_method_name(enum sample_method_t sample_method);
	sampleMethodNameFun ffi.Fun

	// SD_API enum sample_method_t str_to_sample_method(const char* str);
	strToSampleMethodFun ffi.Fun

	// SD_API void sd_sample_params_init(sd_sample_params_t* sample_params);
	sampleParamsInitFun func(uintptr)
	// sampleParamsInitFun ffi.Fun

	// SD_API char* sd_sample_params_to_str(const sd_sample_params_t* sample_params);
	sampleParamsToStrFun ffi.Fun

	// SD_API enum sample_method_t sd_get_default_sample_method(const sd_ctx_t* sd_ctx);
	getDefaultSampleMethodFun ffi.Fun
)

func LoadSampleFuns(lib ffi.Lib) error {
	var err error

	// SD_API const char* sd_sample_method_name(enum sample_method_t sample_method);
	sampleMethodNameFun, err = lib.Prep("sd_sample_method_name", &ffi.TypePointer, &ffi.TypeSint32)
	if err != nil {
		return err
	}

	// SD_API enum sample_method_t str_to_sample_method(const char* str);
	strToSampleMethodFun, err = lib.Prep("str_to_sample_method", &ffi.TypeSint32, &ffi.TypePointer)
	if err != nil {
		return err
	}

	// SD_API void sd_sample_params_init(sd_sample_params_t* sample_params);
	// 这种通过c初始化结构体的操作使用purego 进行注册
	purego.RegisterLibFunc(&sampleParamsInitFun, lib.Addr, "sd_sample_params_init")
	// sampleParamsInitFun, err = lib.Prep("sd_sample_params_init", &ffi.TypeVoid, &ffi.TypePointer)
	// if err != nil {
	// 	return err
	// }

	// SD_API char* sd_sample_params_to_str(const sd_sample_params_t* sample_params);
	sampleParamsToStrFun, err = lib.Prep("sd_sample_params_to_str", &ffi.TypePointer, &ffi.TypePointer)
	if err != nil {
		return err
	}

	// SD_API enum sample_method_t sd_get_default_sample_method(const sd_ctx_t* sd_ctx);
	getDefaultSampleMethodFun, err = lib.Prep("sd_get_default_sample_method", &ffi.TypeSint32, &ffi.TypePointer)
	if err != nil {
		return err
	}

	return nil
}

func SampleMethodName(sampleMethod types.SampleMethod) *byte {
	var result *byte
	sampleMethodNameFun.Call(unsafe.Pointer(&result), unsafe.Pointer(&sampleMethod))
	return result
}

func StrToSampleMethod(str *byte) types.SampleMethod {
	var result int32
	strToSampleMethodFun.Call(unsafe.Pointer(&result), unsafe.Pointer(&str))
	return types.SampleMethod(result)
}

func SampleParamsInit() SampleParams {
	//这里需要分配到堆上，防止内存发生漂移
	structPtr := new(SampleParams)
	ptr := uintptr(unsafe.Pointer(structPtr))
	sampleParamsInitFun(ptr)
	return *structPtr
}

func SampleParamsToStr(sampleParams *SampleParams) *byte {
	var result *byte
	sampleParamsToStrFun.Call(unsafe.Pointer(&result), unsafe.Pointer(sampleParams))
	return result
}

func GetDefaultSampleMethod(ctx Context) types.SampleMethod {
	var result int32
	getDefaultSampleMethodFun.Call(unsafe.Pointer(&result), unsafe.Pointer(&ctx))
	return types.SampleMethod(result)
}
