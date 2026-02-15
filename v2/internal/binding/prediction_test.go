package binding_test

import (
	"testing"

	"github.com/jupiterrider/ffi"
	"github.com/seasonjs/stable-diffusion/v2/internal/binding"
	"github.com/seasonjs/stable-diffusion/v2/internal/utils"
	"github.com/seasonjs/stable-diffusion/v2/pkg/types"
)

func testSetupPrediction(lib ffi.Lib) error {
	return binding.LoadPredictionFuns(lib)
}

func TestPredictionName(t *testing.T) {
	lib := testSetup(t)
	defer testCleanup(lib)

	err := testSetupPrediction(lib)
	if err != nil {
		t.Fatal(err)
	}

	// 测试各种Prediction值
	predictions := []types.Prediction{
		types.EPS_PRED,
		types.V_PRED,
		types.EDM_V_PRED,
		types.FLOW_PRED,
		types.FLUX_FLOW_PRED,
		types.FLUX2_FLOW_PRED,
	}

	for _, prediction := range predictions {
		namePtr := binding.PredictionName(int32(prediction))
		if namePtr == nil {
			t.Errorf("PredictionName returned nil for prediction %d", prediction)
			continue
		}

		name := utils.GoString(namePtr)
		if name == "" {
			t.Errorf("PredictionName returned empty string for prediction %d", prediction)
		}

		t.Logf("Prediction %d: %s", prediction, name)
	}
}

func TestStrToPrediction(t *testing.T) {
	lib := testSetup(t)
	defer testCleanup(lib)

	err := testSetupPrediction(lib)
	if err != nil {
		t.Fatal(err)
	}

	// 测试各种字符串值
	testCases := []string{"eps", "v", "edm_v", "sd3_flow", "flux_flow", "flux2_flow", "invalid"}

	for _, str := range testCases {
		strPtr, err := utils.CString(str)
		if err != nil {
			t.Fatal(err)
		}

		prediction := binding.StrToPrediction(strPtr)
		t.Logf("String '%s': %d", str, prediction)
	}
}

func TestPredictionRoundTrip(t *testing.T) {
	lib := testSetup(t)
	defer testCleanup(lib)

	err := testSetupPrediction(lib)
	if err != nil {
		t.Fatal(err)
	}

	// 测试round trip: prediction -> name -> prediction
	predictions := []types.Prediction{
		types.EPS_PRED,
		types.V_PRED,
		types.EDM_V_PRED,
		types.FLOW_PRED,
		types.FLUX_FLOW_PRED,
		types.FLUX2_FLOW_PRED,
	}

	for _, prediction := range predictions {
		namePtr := binding.PredictionName(int32(prediction))
		if namePtr == nil {
			t.Errorf("PredictionName returned nil for prediction %d", prediction)
			continue
		}

		prediction2 := binding.StrToPrediction(namePtr)
		t.Logf("Prediction %d -> %s -> Prediction %d", prediction, utils.GoString(namePtr), prediction2)
	}
}
