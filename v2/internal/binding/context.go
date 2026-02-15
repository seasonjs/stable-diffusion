package binding

import (
	"structs"
	"unsafe"

	"github.com/ebitengine/purego"
	"github.com/jupiterrider/ffi"
)

type CtxParams struct {
	_                           structs.HostLayout
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
	&ffi.TypePointer, // Embeddings: *Embedding
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

var (
	// SD_API void sd_ctx_params_init(sd_ctx_params_t* sd_ctx_params);
	// 这里因为C语言端进行了初始化：
	//  *sd_ctx_params                         = {};
	// 所以导致内存结构通过ffi无法对齐，所以需要使用更系统级别的调用，让c正确写入
	ctxParamsInitFun func(uintptr)

	// SD_API char* sd_ctx_params_to_str(const sd_ctx_params_t* sd_ctx_params);
	ctxParamsToStrFun ffi.Fun

	// SD_API sd_ctx_t* new_sd_ctx(const sd_ctx_params_t* sd_ctx_params);
	newCtxFun ffi.Fun

	// SD_API void free_sd_ctx(sd_ctx_t* sd_ctx);
	freeCtxFun ffi.Fun
)

type Context uintptr

func LoadContextFuns(lib ffi.Lib) error {
	var err error
	// SD_API void sd_ctx_params_init(sd_ctx_params_t* sd_ctx_params);
	// 这种通过c初始化结构体的操作使用purego 进行
	purego.RegisterLibFunc(&ctxParamsInitFun, lib.Addr, "sd_ctx_params_init")

	// SD_API char* sd_ctx_params_to_str(const sd_ctx_params_t* sd_ctx_params);
	ctxParamsToStrFun, err = lib.Prep("sd_ctx_params_to_str", &ffi.TypePointer, &FFITypeCtxParams)
	if err != nil {
		return err
	}

	// SD_API sd_ctx_t* new_sd_ctx(const sd_ctx_params_t* sd_ctx_params);
	newCtxFun, err = lib.Prep("new_sd_ctx", &ffi.TypePointer, &FFITypeCtxParams)
	if err != nil {
		return err
	}

	// SD_API void free_sd_ctx(sd_ctx_t* sd_ctx);
	freeCtxFun, err = lib.Prep("free_sd_ctx", &ffi.TypeVoid, &ffi.TypePointer)
	if err != nil {
		return err
	}

	return nil

}

// CtxParamsInit 初始化上下文参数
func CtxParamsInit() CtxParams {
	structPtr := new(CtxParams)
	ptr := uintptr(unsafe.Pointer(structPtr))
	ctxParamsInitFun(ptr)
	return *structPtr
}

// CtxParamsToStr 将上下文参数转换为字符串
func CtxParamsToStr(params *CtxParams) *byte {
	var result *byte
	ctxParamsToStrFun.Call(unsafe.Pointer(&result), unsafe.Pointer(params))
	return result
}

// NewCtx 创建新的稳定扩散上下文
func NewCtx(params *CtxParams) Context {
	var result uintptr
	newCtxFun.Call(unsafe.Pointer(&result), unsafe.Pointer(params))
	return Context(result)
}

// FreeCtx 释放稳定扩散上下文
func FreeCtx(ctx Context) {
	freeCtxFun.Call(nil, unsafe.Pointer(&ctx))
}
