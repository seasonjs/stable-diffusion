package binding

import (
	"unsafe"
	"github.com/ebitengine/purego"
	"github.com/jupiterrider/ffi"
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

type PmParams struct {
	IdImages      *Image
	IdImagesCount int32
	IdEmbedPath   *byte
	StyleStrength float32
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

var (
		// SD_API void sd_img_gen_params_init(sd_img_gen_params_t* sd_img_gen_params);
	imgGenParamsInitFun func(uintptr)

	// SD_API char* sd_img_gen_params_to_str(const sd_img_gen_params_t* sd_img_gen_params);
	imgGenParamsToStrFun ffi.Fun

	// SD_API sd_image_t* generate_image(sd_ctx_t* sd_ctx, const sd_img_gen_params_t* sd_img_gen_params);
	generateImageFun ffi.Fun
)


func LoadImgGenFuns(lib ffi.Lib) error {
	var err error

	// SD_API void sd_img_gen_params_init(sd_img_gen_params_t* sd_img_gen_params);
	purego.RegisterLibFunc(&imgGenParamsInitFun, lib.Addr, "sd_img_gen_params_init")

	// SD_API char* sd_img_gen_params_to_str(const sd_img_gen_params_t* sd_img_gen_params);
	imgGenParamsToStrFun, err = lib.Prep("sd_img_gen_params_to_str", &ffi.TypePointer, &FFITypeImgGenParams)
	if err != nil {
		return err
	}

	// SD_API sd_image_t* generate_image(sd_ctx_t* sd_ctx, const sd_img_gen_params_t* sd_img_gen_params);
	generateImageFun, err = lib.Prep("generate_image", &ffi.TypePointer, &ffi.TypePointer, &FFITypeImgGenParams)
	if err != nil {
		return err
	}

	return nil
}

// ImgGenParamsInit 初始化图像生成参数
func ImgGenParamsInit() ImgGenParams {
	//这里需要分配到堆上，防止内存发生漂移
	structPtr := new(ImgGenParams)
	ptr := uintptr(unsafe.Pointer(structPtr))
	imgGenParamsInitFun(ptr)
	return *structPtr
}

func ImgGenParamsToStr(params ImgGenParams) *byte {
	var result *byte
	imgGenParamsToStrFun.Call(unsafe.Pointer(&result), unsafe.Pointer(&params))
	return result
}

func GenerateImage(ctx Context, params *ImgGenParams) uintptr {
	var result uintptr
	generateImageFun.Call(unsafe.Pointer(&result), unsafe.Pointer(&ctx), unsafe.Pointer(params))
	return result
}
