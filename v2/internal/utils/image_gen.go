package utils

import (
	"github.com/seasonjs/stable-diffusion/v2/internal/binding"
	"github.com/seasonjs/stable-diffusion/v2/pkg/types"
)

func ImgGenParamsToC(params *types.ImgGenParams) (*binding.ImgGenParams, error) {
	if params == nil {
		return nil, types.ImgGenParamsEmptyError
	}
	cParams := binding.ImgGenParamsInit()
	//TODO: 完成ImgGenParams的转换
	return &cParams, nil
}
