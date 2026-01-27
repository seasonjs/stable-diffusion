package binding_test

import (
	"math"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/jupiterrider/ffi"
	"github.com/seasonjs/stable-diffusion/internal/binding"
	"github.com/seasonjs/stable-diffusion/internal/utils"
)

func testSetupContext(lib ffi.Lib) error {
	return binding.LoadContextFuns(lib)
}
func TestLoadContext(t *testing.T) {
	lib := testSetup(t)
	defer testCleanup(lib)
	err := testSetupContext(lib)
	if err != nil {
		t.Fatal(err)
	}
}

func TestCtxParamsInit(t *testing.T) {
	lib := testSetup(t)
	defer testCleanup(lib)
	err := testSetupContext(lib)
	if err != nil {
		t.Fatal(err)
	}

	ctxParams := binding.CtxParamsInit()

	// https://github.com/leejet/stable-diffusion.cpp/blob/master/stable-diffusion.cpp#L2884
	// *sd_ctx_params                         = {};
	// sd_ctx_params->vae_decode_only         = true;
	// sd_ctx_params->free_params_immediately = true;
	// sd_ctx_params->n_threads               = sd_get_num_physical_cores();
	// sd_ctx_params->wtype                   = SD_TYPE_COUNT;
	// sd_ctx_params->rng_type                = CUDA_RNG;
	// sd_ctx_params->sampler_rng_type        = RNG_TYPE_COUNT;
	// sd_ctx_params->prediction              = PREDICTION_COUNT;
	// sd_ctx_params->lora_apply_mode         = LORA_APPLY_AUTO;
	// sd_ctx_params->offload_params_to_cpu   = false;
	// sd_ctx_params->enable_mmap             = false;
	// sd_ctx_params->keep_clip_on_cpu        = false;
	// sd_ctx_params->keep_control_net_on_cpu = false;
	// sd_ctx_params->keep_vae_on_cpu         = false;
	// sd_ctx_params->diffusion_flash_attn    = false;
	// sd_ctx_params->circular_x              = false;
	// sd_ctx_params->circular_y              = false;
	// sd_ctx_params->chroma_use_dit_mask     = true;
	// sd_ctx_params->chroma_use_t5_mask      = false;
	// sd_ctx_params->chroma_t5_mask_pad      = 1;
	// sd_ctx_params->flow_shift              = INFINITY;
	expected := binding.CtxParams{
		VaeDecodeOnly:         true,
		FreeParamsImmediately: true,
		NThreads:              10, // sd_get_num_physical_cores() returns 10
		Wtype:                 40, // SD_TYPE_COUNT = COUNT = 40
		RngType:               1,  // CUDA_RNG
		SamplerRngType:        3,  // RNG_TYPE_COUNT
		Prediction:            6,  // PREDICTION_COUNT
		LoraApplyMode:         0,  // LORA_APPLY_AUTO
		OffloadParamsToCpu:    false,
		EnableMmap:            false,
		KeepClipOnCpu:         false,
		KeepControlNetOnCpu:   false,
		KeepVaeOnCpu:          false,
		DiffusionFlashAttn:    false,
		CircularX:             false,
		CircularY:             false,
		ChromaUseDitMask:      true,
		ChromaUseT5Mask:       false,
		ChromaT5MaskPad:       1,
		FlowShift:             float32(math.Inf(1)), // INFINITY
	}
	if diff := cmp.Diff(expected, ctxParams); diff != "" {
		t.Errorf("数据不匹配 (-期望 +实际):\n%s", diff)
	}

}

func TestCtxParamsToStr(t *testing.T) {
	lib := testSetup(t)
	defer testCleanup(lib)
	err := testSetupContext(lib)
	if err != nil {
		t.Fatal(err)
	}
	ctxParams := binding.CtxParamsInit()
	ctxParamsStr := binding.CtxParamsToStr(&ctxParams)
	str := utils.GoString(ctxParamsStr)
	if str == "" {
		t.Fatal("CtxParamsToStr return empty string")
	}
	t.Logf("%s", str)
}

func TestNewCtx(t *testing.T) {
	lib := testSetup(t)
	defer testCleanup(lib)
	err := testSetupContext(lib)
	if err != nil {
		t.Fatal(err)
	}
	ctx := getDefaultContext(t)
	binding.FreeCtx(ctx)
}