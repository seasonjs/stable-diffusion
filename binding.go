package stable_diffusion

import (
	"github.com/ebitengine/purego"
)

const (
	cNewSdTxt2imgOptions      = "new_sd_txt2img_options"
	cSetTxt2imgPrompt         = "set_txt2img_prompt"
	cSetTxt2imgNegativePrompt = "set_txt2img_negative_prompt"
	cSetTxt2imgCfgScale       = "set_txt2img_cfg_scale"
	cSetTxt2imgSize           = "set_txt2img_size"
	cSetTxt2imgSampleMethod   = "set_txt2img_sample_method"
	cSetTxt2imgSampleSteps    = "set_txt2img_sample_steps"
	cSetTxt2imgStrength       = "set_txt2img_strength"
	cSetTxt2imgSeed           = "set_txt2img_seed"

	cNewSdImg2imgOptions      = "new_sd_img2img_options"
	cSetImg2imgInitImg        = "set_img2img_init_img"
	cSetImg2imgPrompt         = "set_img2img_prompt"
	cSetImg2imgNegativePrompt = "set_img2img_negative_prompt"
	cSetImg2imgCfgScale       = "set_img2img_cfg_scale"
	cSetImg2imgSize           = "set_img2img_size"
	cSetImg2imgSampleMethod   = "set_img2img_sample_method"
	cSetImg2imgSampleSteps    = "set_img2img_sample_steps"
	cSetImg2imgStrength       = "set_img2img_strength"
	cSetImg2imgSeed           = "set_img2img_seed"

	cCreateStableDiffusion  = "create_stable_diffusion"
	cDestroyStableDiffusion = "destroy_stable_diffusion"
	cLoadFromFile           = "load_from_file"
	cTxt2img                = "txt2img"
	cImg2img                = "img2img"

	cSetStableDiffusionLogLevel   = "set_stable_diffusion_log_level"
	cGetStableDiffusionSystemInfo = "get_stable_diffusion_system_info"
)

type CStableDiffusion interface {
	CreateStableDiffusion() *CSDCtx
	StableDiffusionLoadFromFile(ctx *CSDCtx, path string)
	StableDiffusionTextToImg(ctx *CSDCtx)
	StableDiffusionImgToImg(ctx *CSDCtx)
	StableDiffusionSetLogLevel()
	StableDiffusionGetSystemInfo()
}

type CStableDiffusionImpl struct {
	libSd uintptr

	cNewSdTxt2imgOptions      func() uintptr
	cSetTxt2imgPrompt         func(options uintptr, prompt string)
	cSetTxt2imgNegativePrompt func(options uintptr, prompt string)
	cSetTxt2imgCfgScale       func(options uintptr, scale float32)
	cSetTxt2imgSize           func(options uintptr, size int)
	cSetTxt2imgSampleMethod   func(options uintptr, method int)
	cSetTxt2imgSampleSteps    func(options uintptr, steps int)
	cSetTxt2imgStrength       func(options uintptr, strength float32)
	cSetTxt2imgSeed           func(options uintptr, seed int)

	cNewSdImg2imgOptions      func() uintptr
	cSetImg2imgInitImg        func(options uintptr, initImg string)
	cSetImg2imgPrompt         func(options uintptr, prompt string)
	cSetImg2imgNegativePrompt func(options uintptr, prompt string)
	cSetImg2imgCfgScale       func(options uintptr, scale float32)
	cSetImg2imgSize           func(options uintptr, size int)
	cSetImg2imgSampleMethod   func(options uintptr, method int)
	cSetImg2imgSampleSteps    func(options uintptr, steps int)
	cSetImg2imgStrength       func(options uintptr, strength float32)
	cSetImg2imgSeed           func(options uintptr, seed int)

	cCreateStableDiffusion  func(options uintptr) uintptr
	cDestroyStableDiffusion func(sd uintptr)
	cLoadFromFile           func(sd uintptr, path string)
	cTxt2img                func(sd uintptr, options uintptr)
	cImg2img                func(sd uintptr, options uintptr)

	cSetStableDiffusionLogLevel   func(level int)
	cGetStableDiffusionSystemInfo func() string
}

type CSDCtx struct {
	ctx uintptr
}

func NewCStableDiffusion(libraryPath string) (*CStableDiffusionImpl, error) {
	libSd, err := openLibrary(libraryPath)
	if err != nil {
		return nil, err
	}
	var (
		newSdTxt2imgOptions      func() uintptr
		setTxt2imgPrompt         func(options uintptr, prompt string)
		setTxt2imgNegativePrompt func(options uintptr, prompt string)
		setTxt2imgCfgScale       func(options uintptr, scale float32)
		setTxt2imgSize           func(options uintptr, size int)
		setTxt2imgSampleMethod   func(options uintptr, method int)
		setTxt2imgSampleSteps    func(options uintptr, steps int)
		setTxt2imgStrength       func(options uintptr, strength float32)
		setTxt2imgSeed           func(options uintptr, seed int)

		newSdImg2imgOptions      func() uintptr
		setImg2imgInitImg        func(options uintptr, initImg string)
		setImg2imgPrompt         func(options uintptr, prompt string)
		setImg2imgNegativePrompt func(options uintptr, prompt string)
		setImg2imgCfgScale       func(options uintptr, scale float32)
		setImg2imgSize           func(options uintptr, size int)
		setImg2imgSampleMethod   func(options uintptr, method int)
		setImg2imgSampleSteps    func(options uintptr, steps int)
		setImg2imgStrength       func(options uintptr, strength float32)
		setImg2imgSeed           func(options uintptr, seed int)

		createStableDiffusion  func(options uintptr) uintptr
		destroyStableDiffusion func(sd uintptr)
		loadFromFile           func(sd uintptr, path string)
		txt2img                func(sd uintptr, options uintptr)
		img2img                func(sd uintptr, options uintptr)

		setStableDiffusionLogLevel   func(level int)
		getStableDiffusionSystemInfo func() string
	)
	purego.RegisterLibFunc(&newSdTxt2imgOptions, libSd, cNewSdTxt2imgOptions)
	purego.RegisterLibFunc(&setTxt2imgPrompt, libSd, cSetTxt2imgPrompt)
	purego.RegisterLibFunc(&setTxt2imgNegativePrompt, libSd, cSetTxt2imgNegativePrompt)
	purego.RegisterLibFunc(&setTxt2imgCfgScale, libSd, cSetTxt2imgCfgScale)
	purego.RegisterLibFunc(&setTxt2imgSize, libSd, cSetTxt2imgSize)
	purego.RegisterLibFunc(&setTxt2imgSampleMethod, libSd, cSetTxt2imgSampleMethod)
	purego.RegisterLibFunc(&setTxt2imgSampleSteps, libSd, cSetTxt2imgSampleSteps)
	purego.RegisterLibFunc(&setTxt2imgStrength, libSd, cSetTxt2imgStrength)
	purego.RegisterLibFunc(&setTxt2imgSeed, libSd, cSetTxt2imgSeed)

	purego.RegisterLibFunc(&newSdImg2imgOptions, libSd, cNewSdImg2imgOptions)
	purego.RegisterLibFunc(&setImg2imgInitImg, libSd, cSetImg2imgInitImg)
	purego.RegisterLibFunc(&setImg2imgPrompt, libSd, cSetImg2imgPrompt)
	purego.RegisterLibFunc(&setImg2imgNegativePrompt, libSd, cSetImg2imgNegativePrompt)
	purego.RegisterLibFunc(&setImg2imgCfgScale, libSd, cSetImg2imgCfgScale)
	purego.RegisterLibFunc(&setImg2imgSize, libSd, cSetImg2imgSize)
	purego.RegisterLibFunc(&setImg2imgSampleMethod, libSd, cSetImg2imgSampleMethod)
	purego.RegisterLibFunc(&setImg2imgSampleSteps, libSd, cSetImg2imgSampleSteps)
	purego.RegisterLibFunc(&setImg2imgStrength, libSd, cSetImg2imgStrength)
	purego.RegisterLibFunc(&setImg2imgSeed, libSd, cSetImg2imgSeed)

	purego.RegisterLibFunc(&createStableDiffusion, libSd, cCreateStableDiffusion)
	purego.RegisterLibFunc(&destroyStableDiffusion, libSd, cDestroyStableDiffusion)
	purego.RegisterLibFunc(&loadFromFile, libSd, cLoadFromFile)
	purego.RegisterLibFunc(&txt2img, libSd, cTxt2img)
	purego.RegisterLibFunc(&img2img, libSd, cImg2img)

	purego.RegisterLibFunc(&setStableDiffusionLogLevel, libSd, cSetStableDiffusionLogLevel)
	purego.RegisterLibFunc(&getStableDiffusionSystemInfo, libSd, cGetStableDiffusionSystemInfo)

	return &CStableDiffusionImpl{
		libSd,

		newSdTxt2imgOptions,
		setTxt2imgPrompt,
		setTxt2imgNegativePrompt,
		setTxt2imgCfgScale,
		setTxt2imgSize,
		setTxt2imgSampleMethod,
		setTxt2imgSampleSteps,
		setTxt2imgStrength,
		setTxt2imgSeed,

		newSdImg2imgOptions,
		setImg2imgInitImg,
		setImg2imgPrompt,
		setImg2imgNegativePrompt,
		setImg2imgCfgScale,
		setImg2imgSize,
		setImg2imgSampleMethod,
		setImg2imgSampleSteps,
		setImg2imgStrength,
		setImg2imgSeed,

		createStableDiffusion,
		destroyStableDiffusion,
		loadFromFile,
		txt2img,
		img2img,

		setStableDiffusionLogLevel,
		getStableDiffusionSystemInfo,
	}, nil
}

func (cSD *CStableDiffusionImpl) CreateStableDiffusion() *CSDCtx {
	ctx := cSD.cCreateStableDiffusion(0)
	return &CSDCtx{ctx: ctx}
}

func (cSD *CStableDiffusionImpl) StableDiffusionLoadFromFile(ctx *CSDCtx, path string) {
	cSD.cLoadFromFile(ctx.ctx, path)
}

func (cSD *CStableDiffusionImpl) StableDiffusionTextToImg(ctx *CSDCtx) {
	options := cSD.cNewSdImg2imgOptions()
	cSD.cTxt2img(ctx.ctx, options)
}

func (cSD *CStableDiffusionImpl) StableDiffusionImgToImg(ctx *CSDCtx) {
	options := cSD.cNewSdImg2imgOptions()
	cSD.cTxt2img(ctx.ctx, options)
}

func (cSD *CStableDiffusionImpl) StableDiffusionSetLogLevel() {

}

func (cSD *CStableDiffusionImpl) StableDiffusionGetSystemInfo() {

}
