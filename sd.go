// Copyright (c) seasonjs. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package sd

type StableDiffusionOptions struct {
	Threads               int
	VaeDecodeOnly         bool
	FreeParamsImmediately bool
	RngType               RNGType
	Schedule              Schedule
	NegativePrompt        string
	CfgScale              float32
	Width                 int
	Height                int
	SampleMethod          SampleMethod
	SampleSteps           int
	Strength              float32
	Seed                  int64
}

type StableDiffusionModel struct {
	ctx     *CSDCtx
	options *StableDiffusionOptions
}

func NewStableDiffusionAutoModel(options StableDiffusionOptions) {

}

func NewStableDiffusionModel(path string, options StableDiffusionOptions) (*StableDiffusionModel, error) {
	sd, err := NewCStableDiffusion(path)
	if err != nil {
		return nil, err
	}
	ctx := sd.NewStableDiffusionCtx(options.Threads, options.VaeDecodeOnly, options.FreeParamsImmediately, options.RngType)
	return &StableDiffusionModel{
		ctx:     ctx,
		options: &options,
	}, nil
}

func (sd *StableDiffusionModel) LoadFromFile(path string) error {
	sd.ctx.StableDiffusionLoadFromFile(path, sd.options.Schedule)
	return nil
}
