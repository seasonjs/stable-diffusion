package binding

import (
	"unsafe"
	"github.com/ebitengine/purego"
	"github.com/jupiterrider/ffi"
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

var (
	cacheParamsInitFun func(uintptr)
)

type CacheParams struct {
	Mode                     int32
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


func LoadCacheFuns(lib ffi.Lib) error{
	// SD_API void sd_cache_params_init(sd_cache_params_t* cache_params);
	// 这种通过c初始化结构体的操作使用purego 进行
	purego.RegisterLibFunc(&cacheParamsInitFun, lib.Addr, "sd_cache_params_init")

	return nil
}

func CacheParamsInit() CacheParams{
	//这里需要分配到堆上，防止内存发生漂移
	structPtr := new(CacheParams)
	ptr := uintptr(unsafe.Pointer(structPtr))
	cacheParamsInitFun(ptr)
	return *structPtr
}