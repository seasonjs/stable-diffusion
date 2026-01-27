package binding


import (
	"unsafe"
	"github.com/ebitengine/purego"
	"github.com/jupiterrider/ffi"
)

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


var (
	// SD_API void sd_vid_gen_params_init(sd_vid_gen_params_t* sd_vid_gen_params);
	vidGenParamsInit func(uintptr)

	// SD_API sd_image_t* generate_video(sd_ctx_t* sd_ctx, const sd_vid_gen_params_t* sd_vid_gen_params, int* num_frames_out);
	generateVideo ffi.Fun
)

func LoadVideoGenFuns(lib ffi.Lib) error {
	var err error
	
	// SD_API void sd_vid_gen_params_init(sd_vid_gen_params_t* sd_vid_gen_params);
	purego.RegisterLibFunc(&vidGenParamsInit, lib.Addr, "sd_vid_gen_params_init")
	
	// SD_API sd_image_t* generate_video(sd_ctx_t* sd_ctx, const sd_vid_gen_params_t* sd_vid_gen_params, int* num_frames_out);
	generateVideo, err = lib.Prep("generate_video", &ffi.TypePointer, &ffi.TypePointer, &ffi.TypeSint32)
	if err != nil {
		return err
	}
	return nil
}

func VideoGenParamsInit() VidGenParams {
	//这里需要分配到堆上，防止内存发生漂移
	structPtr := new(VidGenParams)
	ptr := uintptr(unsafe.Pointer(structPtr))
	vidGenParamsInit(ptr)
	return *structPtr
}

func GenerateVideo(params VidGenParams) uintptr {
	var result uintptr
	generateVideo.Call(unsafe.Pointer(&result), unsafe.Pointer(&params))
	return result
}