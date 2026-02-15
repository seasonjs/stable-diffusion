package diffusion

import (
	"github.com/ebitengine/purego"
	"github.com/jupiterrider/ffi"
	"github.com/seasonjs/stable-diffusion/v2/internal/binding"
	"github.com/seasonjs/stable-diffusion/v2/internal/utils"
	"github.com/seasonjs/stable-diffusion/v2/pkg/loader"
	"github.com/seasonjs/stable-diffusion/v2/pkg/types"
)

type Diffusion struct {
	stdLib       ffi.Lib
	diffusionLib ffi.Lib
}

func New(dirPath string) (*Diffusion, error) {
	var err error
	d := &Diffusion{}

	d.diffusionLib, err = loader.LoadLibrary(dirPath, "stable-diffusion")
	if err != nil {
		return nil, err
	}

	if err := loadDiffusionFuns(d.diffusionLib); err != nil {
		return nil, err
	}

	d.stdLib, err = utils.LoadStdLib()
	if err != nil {
		return nil, err
	}

	if err := utils.LoadStdFuns(d.stdLib); err != nil {
		return nil, err
	}
	return d, nil
}

func (d *Diffusion) SetLogCallback(callback types.LogCallback) {
	cb := purego.NewCallback(func(level int32, text uintptr, data uintptr) uintptr {
		callback(types.LogLevel(level), utils.GoString(text))
		return 0
	})

	binding.SetLogCallback(cb)
}

func (d *Diffusion) GetNumPhysicalCores() int {
	return int(binding.GetNumPhysicalCores())
}

func (d *Diffusion) GetSystemInfo() string {
	systemInfo := binding.GetSystemInfo()
	return utils.GoString(systemInfo)
}

func(d *Diffusion)  Commit() string {
	result := binding.Commit()
	return utils.GoString(result)
}

func (d *Diffusion) Version() string {
	result := binding.Version()
	return utils.GoString(result)
}

func (d *Diffusion) Close() {
	d.stdLib.Close()
	d.diffusionLib.Close()
}
