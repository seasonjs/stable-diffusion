package types

import (
	"errors"
)

var (
	ContextParamsEmptyError = errors.New("context params is empty")
)

type ContextParams struct {
	ModelPath                   string
	ClipLPath                   string
	ClipGPath                   string
	ClipVisionPath              string
	T5xxlPath                   string
	LlmPath                     string
	LlmVisionPath               string
	DiffusionModelPath          string
	HighNoiseDiffusionModelPath string
	VaePath                     string
	TaesdPath                   string
	ControlNetPath              string
	Embeddings                  []Embedding
	EmbeddingCount              uint32
	PhotoMakerPath              string
	TensorTypeRules             string
	VaeDecodeOnly               bool
	FreeParamsImmediately       bool
	NThreads                    int32
	Wtype                       SdType
	RngType                     RNGType
	SamplerRngType              RNGType
	Prediction                  Prediction
	LoraApplyMode               LoraApplyMode
	OffloadParamsToCpu          bool
	EnableMmap                  bool
	KeepClipOnCpu               bool
	KeepControlNetOnCpu         bool
	KeepVaeOnCpu                bool
	DiffusionFlashAttn          bool
	TaePreviewOnly              bool
	DiffusionConvDirect         bool
	VaeConvDirect               bool
	CircularX                   bool
	CircularY                   bool
	ForceSdxlVaeConvScale       bool
	ChromaUseDitMask            bool
	ChromaUseT5Mask             bool
	ChromaT5MaskPad             int32
	QwenImageZeroCondT          bool
	FlowShift                   float32
}

