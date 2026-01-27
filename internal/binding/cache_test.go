package binding_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/jupiterrider/ffi"
	"github.com/seasonjs/stable-diffusion/internal/binding"
)

func testSetupCache(lib ffi.Lib) error {
	return binding.LoadCacheFuns(lib)
}
func TestCacheParams(t *testing.T) {
	lib := testSetup(t)
	defer testCleanup(lib)

	err := testSetupCache(lib)
	if err != nil {
		t.Fatal(err)
	}

	cacheParams := binding.CacheParamsInit()

	// https://github.com/leejet/stable-diffusion.cpp/blob/master/stable-diffusion.cpp#L2863
	// 	void sd_cache_params_init(sd_cache_params_t* cache_params) {
	//     *cache_params                             = {};
	//     cache_params->mode                        = SD_CACHE_DISABLED;
	//     cache_params->reuse_threshold             = 1.0f;
	//     cache_params->start_percent               = 0.15f;
	//     cache_params->end_percent                 = 0.95f;
	//     cache_params->error_decay_rate            = 1.0f;
	//     cache_params->use_relative_threshold      = true;
	//     cache_params->reset_error_on_compute      = true;
	//     cache_params->Fn_compute_blocks           = 8;
	//     cache_params->Bn_compute_blocks           = 0;
	//     cache_params->residual_diff_threshold     = 0.08f;
	//     cache_params->max_warmup_steps            = 8;
	//     cache_params->max_cached_steps            = -1;
	//     cache_params->max_continuous_cached_steps = -1;
	//     cache_params->taylorseer_n_derivatives    = 1;
	//     cache_params->taylorseer_skip_interval    = 1;
	//     cache_params->scm_mask                    = nullptr;
	//     cache_params->scm_policy_dynamic          = true;
	// }
	expected := binding.CacheParams{
		Mode:                     0, // SD_CACHE_DISABLED
		ReuseThreshold:           1.0,
		StartPercent:             0.15,
		EndPercent:               0.95,
		ErrorDecayRate:           1.0,
		UseRelativeThreshold:     true,
		ResetErrorOnCompute:      true,
		FnComputeBlocks:          8,
		BnComputeBlocks:          0,
		ResidualDiffThreshold:    0.08,
		MaxWarmupSteps:           8,
		MaxCachedSteps:           -1,
		MaxContinuousCachedSteps: -1,
		TaylorseerNDerivatives:   1,
		TaylorseerSkipInterval:   1,
		ScmMask:                  nil, // nullptr
		ScmPolicyDynamic:         true,
	}
	if diff := cmp.Diff(expected, cacheParams); diff != "" {
		t.Errorf("数据不匹配 (-期望 +实际):\n%s", diff)
	}

}
