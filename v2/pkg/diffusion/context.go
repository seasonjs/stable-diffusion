package diffusion

import (
	"errors"
	"github.com/seasonjs/stable-diffusion/v2/internal/binding"
	"github.com/seasonjs/stable-diffusion/v2/internal/utils"
	"github.com/seasonjs/stable-diffusion/v2/pkg/types"
	"io"
)

type DiffusionContext struct {
	ctx   binding.Context
	modal *Diffusion
}

func (mod *Diffusion) NewContext(ctxParams *types.ContextParams) (*DiffusionContext, error) {
	cParams, err := utils.ContextParamsToC(ctxParams)
	if err != nil {
		return nil, err
	}

	ctx := binding.NewCtx(cParams)

	return &DiffusionContext{
		modal: mod,
		ctx:   ctx,
	}, nil
}

func (d *DiffusionContext) GenerateImage(params *types.ImgGenParams, imageType types.ImageType, writer []io.Writer) error {
	if len(writer) != int(params.BatchCount) {
		return errors.New("writer count not match batch count")
	}

	cParams, err := utils.ImgGenParamsToC(params)
	if err != nil {
		return err
	}

	result := binding.GenerateImage(d.ctx, cParams)
	defer utils.FreeImageSlice(result, int(params.BatchCount))

	if result != 0 {
		images := utils.GoImageSlice(result, int(params.BatchCount))
		for i, image := range images {
			if err := utils.GoImageEncode(image, imageType, writer[i]); err != nil {
				return err
			}
		}
	}
	return nil
}

func (d *DiffusionContext) Close() {
	if d.ctx != 0 {
		binding.FreeCtx(d.ctx)
	}
}
