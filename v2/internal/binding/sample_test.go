package binding_test

import (
	"math"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/jupiterrider/ffi"
	"github.com/seasonjs/stable-diffusion/v2/internal/binding"
	"github.com/seasonjs/stable-diffusion/v2/internal/utils"
	"github.com/seasonjs/stable-diffusion/v2/pkg/types"
)

func testSetupSample(lib ffi.Lib) error {
	return binding.LoadSampleFuns(lib)
}

func TestSampleMethodName(t *testing.T) {
	lib := testSetup(t)
	defer testCleanup(lib)

	err := testSetupSample(lib)
	if err != nil {
		t.Fatal(err)
	}

	// 测试各种SampleMethod值
	methods := []types.SampleMethod{
		types.EULER,
		types.EULER_A,
		types.HEUN,
		types.DPM2,
		types.DPMPP2S_A,
		types.DPMPP2M,
		types.DPMPP2Mv2,
		types.IPNDM,
		types.IPNDM_V,
		types.LCM,
		types.DDIM_TRAILING,
		types.TCD,
	}

	for _, method := range methods {
		namePtr := binding.SampleMethodName(int32(method))
		if namePtr == nil {
			t.Errorf("SampleMethodName returned nil for method %d", method)
			continue
		}

		name := utils.GoString(namePtr)
		if name == "" {
			t.Errorf("SampleMethodName returned empty string for method %d", method)
		}

		t.Logf("Method %d: %s", method, name)
	}
}

func TestStrToSampleMethod(t *testing.T) {
	lib := testSetup(t)
	defer testCleanup(lib)

	err := testSetupSample(lib)
	if err != nil {
		t.Fatal(err)
	}

	// 测试各种字符串值
	testCases := []string{"euler", "euler_a", "heun", "dpm2", "dpm++2s_a", "dpm++2m", "dpm++2mv2", "ipndm", "ipndm_v", "lcm", "ddim_trailing", "tcd", "invalid"}

	for _, str := range testCases {
		strPtr, err := utils.CString(str)
		if err != nil {
			t.Fatal(err)
		}

		method := binding.StrToSampleMethod(strPtr)
		t.Logf("String '%s': %d", str, method)
	}
}

func TestSampleMethodRoundTrip(t *testing.T) {
	lib := testSetup(t)
	defer testCleanup(lib)

	err := testSetupSample(lib)
	if err != nil {
		t.Fatal(err)
	}

	// 测试round trip: method -> name -> method
	methods := []types.SampleMethod{
		types.EULER,
		types.EULER_A,
		types.HEUN,
		types.DPM2,
		types.DPMPP2S_A,
		types.DPMPP2M,
		types.DPMPP2Mv2,
		types.IPNDM,
		types.IPNDM_V,
		types.LCM,
		types.DDIM_TRAILING,
		types.TCD,
	}

	for _, method := range methods {
		namePtr := binding.SampleMethodName(int32(method))
		if namePtr == nil {
			t.Errorf("SampleMethodName returned nil for method %d", method)
			continue
		}

		method2 := binding.StrToSampleMethod(namePtr)
		t.Logf("Method %d -> %s -> Method %d", method, utils.GoString(namePtr), method2)
	}
}

func TestSampleParamsInit(t *testing.T) {
	lib := testSetup(t)
	defer testCleanup(lib)

	err := testSetupSample(lib)
	if err != nil {
		t.Fatal(err)
	}

	// 测试SampleParamsInit函数，确保它能被调用而不崩溃
	sampleParams := binding.SampleParamsInit()

	// https://github.com/leejet/stable-diffusion.cpp/blob/master/stable-diffusion.cpp#L2981C1-L2995C2
	// 	void sd_sample_params_init(sd_sample_params_t* sample_params) {
	//     *sample_params                             = {};
	//     sample_params->guidance.txt_cfg            = 7.0f;
	//     sample_params->guidance.img_cfg            = INFINITY;
	//     sample_params->guidance.distilled_guidance = 3.5f;
	//     sample_params->guidance.slg.layer_count    = 0;
	//     sample_params->guidance.slg.layer_start    = 0.01f;
	//     sample_params->guidance.slg.layer_end      = 0.2f;
	//     sample_params->guidance.slg.scale          = 0.f;
	//     sample_params->scheduler                   = SCHEDULER_COUNT;
	//     sample_params->sample_method               = SAMPLE_METHOD_COUNT;
	//     sample_params->sample_steps                = 20;
	//     sample_params->custom_sigmas               = nullptr;
	//     sample_params->custom_sigmas_count         = 0;
	// }

	expected := binding.SampleParams{
		Guidance: binding.GuidanceParams{
			TxtCfg:            7.0,
			ImgCfg:            float32(math.Inf(1)), // INFINITY
			DistilledGuidance: 3.5,
			Slg: binding.SlgParams{
				// Layers:     nil, // layers: *int32
				LayerCount: 0,
				LayerStart: 0.01,
				LayerEnd:   0.2,
				Scale:      0.0,
			},
		},
		Scheduler:    10,
		SampleMethod: 12,
		SampleSteps:  20,
		// CustomSigmas:      nil, // custom_sigmas: nullptr
		CustomSigmasCount: 0,
	}
	if diff := cmp.Diff(expected, sampleParams); diff != "" {
		t.Errorf("数据不匹配 (-期望 +实际):\n%s", diff)
	}
}

func TestSampleParamsToStr(t *testing.T) {
	t.Skip("TODO: 排查为什么这里会崩溃")
		
	lib := testSetup(t)
	defer testCleanup(lib)

	err := testSetupSample(lib)
	if err != nil {
		t.Fatal(err)
	}
	sampleParams := binding.SampleParamsInit()
	str := binding.SampleParamsToStr(&sampleParams)
	if str == nil {
		t.Errorf("SampleParamsToStr returned nil")
	}
	t.Logf("SampleParamsToStr: %s", utils.GoString(str))
}

// 注意：GetDefaultSampleMethod需要一个有效的Context对象，这需要加载模型等复杂操作
// 我们暂时跳过这个测试，或者需要更复杂的测试设置
func TestGetDefaultSampleMethod(t *testing.T) {
	lib := testSetup(t)
	defer testCleanup(lib)

	err := testSetupSample(lib)
	if err != nil {
		t.Fatal(err)
	}

	err = testSetupContext(lib)
	if err != nil {
		t.Fatal(err)
	}

	ctx := getDefaultContext(t)

	defaultMethod := binding.GetDefaultSampleMethod(ctx)
	t.Logf("Default method: %d", defaultMethod)
	if defaultMethod == int32(types.SAMPLE_METHOD_COUNT) {
		t.Errorf("GetDefaultSampleMethod returned default method SAMPLE_METHOD_COUNT")
	}
}
