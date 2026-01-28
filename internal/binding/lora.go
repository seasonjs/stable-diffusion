package binding

import (
	"unsafe"

	"github.com/jupiterrider/ffi"
	"github.com/seasonjs/stable-diffusion/pkg/types"
)

// FFITypeLora 是Lora结构体的ffi.Type定义
var FFITypeLora = ffi.NewType(
	&ffi.TypeUint8,   // IsHighNoise: bool
	&ffi.TypeFloat,   // Multiplier: float32
	&ffi.TypePointer, // Path: *byte
)

type Lora struct {
	IsHighNoise bool
	Multiplier  float32
	Path        *byte
}

var (
	// SD_API const char* sd_lora_apply_mode_name(enum lora_apply_mode_t mode);
	loraApplyModeNameFun ffi.Fun

	// SD_API enum lora_apply_mode_t str_to_lora_apply_mode(const char* str);
	strToLoraApplyModeFun ffi.Fun
)

func LoadLoraFuns(lib ffi.Lib) error {
	var err error

	// SD_API const char* sd_lora_apply_mode_name(enum lora_apply_mode_t mode);
	loraApplyModeNameFun, err = lib.Prep("sd_lora_apply_mode_name", &ffi.TypePointer, &ffi.TypeSint32)
	if err != nil {
		return err
	}

	// SD_API enum lora_apply_mode_t str_to_lora_apply_mode(const char* str);
	strToLoraApplyModeFun, err = lib.Prep("str_to_lora_apply_mode", &ffi.TypeSint32, &ffi.TypePointer)
	if err != nil {
		return err
	}

	return nil
}

func LoraApplyModeName(mode types.LoraApplyMode) *byte {
	var result *byte
	loraApplyModeNameFun.Call(unsafe.Pointer(&result), unsafe.Pointer(&mode))
	return result
}

func StrToLoraApplyMode(str *byte) types.LoraApplyMode {
	var result int32
	strToLoraApplyModeFun.Call(unsafe.Pointer(&result), unsafe.Pointer(&str))
	return types.LoraApplyMode(result)
}
