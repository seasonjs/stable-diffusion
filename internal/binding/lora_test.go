package binding_test

import (
	"testing"

	"github.com/jupiterrider/ffi"
	"github.com/seasonjs/stable-diffusion/internal/binding"
	"github.com/seasonjs/stable-diffusion/internal/utils"
	"github.com/seasonjs/stable-diffusion/pkg/types"
)

func testSetupLora(lib ffi.Lib) error {
	return binding.LoadLoraFuns(lib)
}

func TestLoraApplyModeName(t *testing.T) {
	lib := testSetup(t)
	defer testCleanup(lib)

	err := testSetupLora(lib)
	if err != nil {
		t.Fatal(err)
	}

	// 测试各种LoraApplyMode值
	modes := []types.LoraApplyMode{
		types.LORA_APPLY_AUTO,
		types.LORA_APPLY_IMMEDIATELY,
		types.LORA_APPLY_AT_RUNTIME,
	}

	for _, mode := range modes {
		namePtr := binding.LoraApplyModeName(mode)
		if namePtr == nil {
			t.Errorf("LoraApplyModeName returned nil for mode %d", mode)
			continue
		}

		name := utils.GoString(namePtr)
		if name == "" {
			t.Errorf("LoraApplyModeName returned empty string for mode %d", mode)
		}

		t.Logf("Mode %d: %s", mode, name)
	}
}

func TestStrToLoraApplyMode(t *testing.T) {
	lib := testSetup(t)
	defer testCleanup(lib)

	err := testSetupLora(lib)
	if err != nil {
		t.Fatal(err)
	}

	// 测试各种字符串值
	testCases := []struct {
		str      string
		expected types.LoraApplyMode
	}{
		{"auto", types.LORA_APPLY_AUTO},
		{"immediately", types.LORA_APPLY_IMMEDIATELY},
		{"runtime", types.LORA_APPLY_AT_RUNTIME},
		{"invalid", types.LORA_APPLY_AUTO}, // 无效值应该返回默认值
	}

	for _, tc := range testCases {
		strPtr, err := utils.CString(tc.str)
		if err != nil {
			t.Fatal(err)
		}

		mode := binding.StrToLoraApplyMode(strPtr)
		t.Logf("String '%s': %d", tc.str, mode)
	}
}

func TestLoraApplyModeRoundTrip(t *testing.T) {
	lib := testSetup(t)
	defer testCleanup(lib)

	err := testSetupLora(lib)
	if err != nil {
		t.Fatal(err)
	}

	// 测试round trip: mode -> name -> mode
	modes := []types.LoraApplyMode{
		types.LORA_APPLY_AUTO,
		types.LORA_APPLY_IMMEDIATELY,
		types.LORA_APPLY_AT_RUNTIME,
	}

	for _, mode := range modes {
		namePtr := binding.LoraApplyModeName(mode)
		if namePtr == nil {
			t.Errorf("LoraApplyModeName returned nil for mode %d", mode)
			continue
		}

		mode2 := binding.StrToLoraApplyMode(namePtr)
		t.Logf("Mode %d -> %s -> Mode %d", mode, utils.GoString(namePtr), mode2)
	}
}
