package utils

import (
	"github.com/seasonjs/stable-diffusion/v2/internal/binding"
	"github.com/seasonjs/stable-diffusion/v2/pkg/types"
)


func ContextParamsToC(ctxParams *types.ContextParams) (*binding.CtxParams, error) {
	//TODO: 完成ContextParams的转换
	if ctxParams == nil {
		return nil, types.ContextParamsEmptyError
	}

	var err error

	cParams := binding.CtxParamsInit()

	cParams.ModelPath, err = CString(ctxParams.ModelPath)
	if err != nil {
		return nil, err
	}

	cParams.ClipLPath, err = CString(ctxParams.ClipLPath)
	if err != nil {
		return nil, err
	}

	cParams.ClipGPath, err = CString(ctxParams.ClipGPath)
	if err != nil {
		return nil, err
	}

	cParams.ClipVisionPath, err = CString(ctxParams.ClipVisionPath)
	if err != nil {
		return nil, err
	}

	cParams.T5xxlPath, err = CString(ctxParams.T5xxlPath)
	if err != nil {
		return nil, err
	}

	cParams.LlmPath, err = CString(ctxParams.LlmPath)
	if err != nil {
		return nil, err
	}

	cParams.LlmVisionPath, err = CString(ctxParams.LlmVisionPath)		
	if err != nil {
		return nil, err
	}

	cParams.DiffusionModelPath, err = CString(ctxParams.DiffusionModelPath)
	if err != nil {
		return nil, err
	}

	cParams.HighNoiseDiffusionModelPath, err = CString(ctxParams.HighNoiseDiffusionModelPath)
	if err != nil {
		return nil, err
	}

	cParams.VaePath, err = CString(ctxParams.VaePath)		
	if err != nil {
		return nil, err
	}

	cParams.TaesdPath, err = CString(ctxParams.TaesdPath)
	if err != nil {
		return nil, err
	}

	cParams.ControlNetPath, err = CString(ctxParams.ControlNetPath)
	if err != nil {
		return nil, err
	}

	cParams.PhotoMakerPath, err = CString(ctxParams.PhotoMakerPath)
	if err != nil {
		return nil, err
	}

	cParams.TensorTypeRules, err = CString(ctxParams.TensorTypeRules)
	if err != nil {
		return nil, err
	}

	cParams.VaeDecodeOnly = ctxParams.VaeDecodeOnly

	cParams.FreeParamsImmediately = ctxParams.FreeParamsImmediately

	cParams.NThreads = ctxParams.NThreads

	cParams.Wtype = int32(ctxParams.Wtype)

	cParams.RngType = int32(ctxParams.RngType)

	return &cParams, nil
}