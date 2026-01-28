package binding

import (
	"structs"

	"github.com/jupiterrider/ffi"
)

// FFITypeImage 是Image结构体的ffi.Type定义
var FFITypeImage = ffi.NewType(
	&ffi.TypeUint32,  // Width: uint32
	&ffi.TypeUint32,  // Height: uint32
	&ffi.TypeUint32,  // Channel: uint32
	&ffi.TypePointer, // Data: *byte (C中的uint8_t*)
)

type Image struct {
	_       structs.HostLayout
	Width   uint32
	Height  uint32
	Channel uint32
	Data    uintptr
}

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
