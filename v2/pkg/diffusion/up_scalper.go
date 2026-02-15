package diffusion

import (
	"io"

	"github.com/seasonjs/stable-diffusion/v2/internal/binding"
	"github.com/seasonjs/stable-diffusion/v2/internal/utils"
	"github.com/seasonjs/stable-diffusion/v2/pkg/types"
)

type UpScalperContext struct {
	ctx   binding.UpscalerContext
	modal *Diffusion
}

func (d *Diffusion) NewUpScalperContext(esrganPath string, offloadParamsToCPU bool, direct bool, nThreads int32, tileSize int32) (*UpScalperContext, error) {
	cEsrganPath, err := utils.CString(esrganPath)
	if err != nil {
		return nil, err
	}

	ctx := binding.NewUpscalerCtx(cEsrganPath, offloadParamsToCPU, direct, nThreads, tileSize)
	return &UpScalperContext{
		ctx:   ctx,
		modal: d,
	}, nil
}

func (u *UpScalperContext) Upscale(reader io.Reader, outPutImageType types.ImageType, writer io.Writer) error {
	initImage, err := utils.DecodeToCImage(reader)
	if err != nil {
		return err
	}

	upscaleFactor := binding.GetUpscaleFactor(u.ctx)

	result := binding.Upscale(u.ctx, initImage, uint32(upscaleFactor))
	defer utils.FreeImage(result)

	if err := utils.CImageEncode(result, outPutImageType, writer); err != nil {
		return err
	}

	return nil
}

func (u *UpScalperContext) Close() {
	if u.ctx != 0 {
		binding.FreeUpscalerCtx(u.ctx)
	}
}
