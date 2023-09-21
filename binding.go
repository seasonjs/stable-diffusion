package stable_diffusion

import (
	"fmt"
	"github.com/ebitengine/purego"
)

const (
	new_sd_txt2img_options      = "new_sd_txt2img_options"
	set_txt2img_prompt          = "set_txt2img_prompt"
	set_txt2img_negative_prompt = "set_txt2img_negative_prompt"
	set_txt2img_cfg_scale       = "set_txt2img_cfg_scale"
	set_txt2img_size            = "set_txt2img_size"
	set_txt2img_sample_method   = "set_txt2img_sample_method"
	set_txt2img_sample_steps    = "set_txt2img_sample_steps"
	set_txt2img_strength        = "set_txt2img_strength"
	set_txt2img_seed            = "set_txt2img_seed"

	new_sd_img2img_options      = "new_sd_img2img_options"
	set_img2img_init_img        = "set_img2img_init_img"
	set_img2img_prompt          = "set_img2img_prompt"
	set_img2img_negative_prompt = "set_img2img_negative_prompt"
	set_img2img_cfg_scale       = "set_img2img_cfg_scale"
	set_img2img_size            = "set_img2img_size"
	set_img2img_sample_method   = "set_img2img_sample_method"
	set_img2img_sample_steps    = "set_img2img_sample_steps"
	set_img2img_strength        = "set_img2img_strength"
	set_img2img_seed            = "set_img2img_seed"

	create_stable_diffusion        = "create_stable_diffusion"
	destroy_stable_diffusion       = "destroy_stable_diffusion"
	load_from_file                 = "load_from_file"
	txt2img                        = "txt2img"
	img2img                        = "img2img"
	set_stable_diffusion_log_level = "set_stable_diffusion_log_level"
	cGetStableDiffusionSystemInfo  = "get_stable_diffusion_system_info"
)

type CStableDiffusion interface {
	StableDiffusionLoadFromFile()
	StableDiffusionTextToImg()
	StableDiffusionImgToImg()
}

type CStableDiffusionImpl struct {
	libSd                         uintptr
	cGetStableDiffusionSystemInfo func() string
}

func NewCStableDiffusion(libraryPath string) (*CStableDiffusionImpl, error) {
	libSd, err := openLibrary(libraryPath)
	if err != nil {
		return nil, err
	}
	var getStableDiffusionSystemInfo func() string
	purego.RegisterLibFunc(&getStableDiffusionSystemInfo, libSd, cGetStableDiffusionSystemInfo)
	return &CStableDiffusionImpl{
		libSd,
		getStableDiffusionSystemInfo,
	}, nil
}

func (c *CStableDiffusionImpl) StableDiffusionLoadFromFile() {
	//TODO implement me
	panic("implement me")
}

func (c *CStableDiffusionImpl) StableDiffusionTextToImg() {
	opt := newCStableDiffusionTextToimgOptions(c.libSd)
	fmt.Print(opt)
}

func (c *CStableDiffusionImpl) StableDiffusionImgToImg() {
	opt := newCStableDiffusionImgToImgOptions(c.libSd)
	fmt.Print(opt)
}

type cStableDiffusionTxt2imgOptions struct {
	csdTxt2imgOptions uintptr
	libSd             uintptr
}

func newCStableDiffusionTextToimgOptions(libSd uintptr) *cStableDiffusionTxt2imgOptions {

	return &cStableDiffusionTxt2imgOptions{
		0,
		libSd,
	}
}
func (o *cStableDiffusionTxt2imgOptions) setTxt2imgPrompt(prompt string) {

}

type cStableDiffusionImg2imgOptions struct {
	cSdImg2imgOptions uintptr
	libSd             uintptr
}

func newCStableDiffusionImgToImgOptions(libSd uintptr) *cStableDiffusionImg2imgOptions {
	return &cStableDiffusionImg2imgOptions{
		0,
		libSd,
	}
}

func (o *cStableDiffusionImg2imgOptions) setImg2imgPrompt(prompt string) {

}
