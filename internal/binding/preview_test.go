package binding_test

import (
	"testing"

	"github.com/jupiterrider/ffi"
	"github.com/seasonjs/stable-diffusion/internal/binding"
	"github.com/seasonjs/stable-diffusion/internal/utils"
	"github.com/seasonjs/stable-diffusion/pkg/types"
)

func testSetupPreview(lib ffi.Lib) error {
	return binding.LoadPreviewFuns(lib)
}

func TestPreviewName(t *testing.T) {
	lib := testSetup(t)
	defer testCleanup(lib)

	err := testSetupPreview(lib)
	if err != nil {
		t.Fatal(err)
	}

	// 测试各种Preview值
	previews := []types.Preview{
		types.PREVIEW_NONE,
		types.PREVIEW_PROJ,
		types.PREVIEW_TAE,
		types.PREVIEW_VAE,
	}

	for _, preview := range previews {
		namePtr := binding.PreviewName(preview)
		if namePtr == nil {
			t.Errorf("PreviewName returned nil for preview %d", preview)
			continue
		}

		name := utils.GoString(namePtr)
		if name == "" {
			t.Errorf("PreviewName returned empty string for preview %d", preview)
		}

		t.Logf("Preview %d: %s", preview, name)
	}
}

func TestStrToPreview(t *testing.T) {
	lib := testSetup(t)
	defer testCleanup(lib)

	err := testSetupPreview(lib)
	if err != nil {
		t.Fatal(err)
	}

	// 测试各种字符串值
	testCases := []string{"none", "proj", "tae", "vae", "invalid"}

	for _, str := range testCases {
		strPtr, err := utils.CString(str)
		if err != nil {
			t.Fatal(err)
		}

		preview := binding.StrToPreview(strPtr)
		t.Logf("String '%s': %d", str, preview)
	}
}

func TestPreviewRoundTrip(t *testing.T) {
	lib := testSetup(t)
	defer testCleanup(lib)

	err := testSetupPreview(lib)
	if err != nil {
		t.Fatal(err)
	}

	// 测试round trip: preview -> name -> preview
	previews := []types.Preview{
		types.PREVIEW_NONE,
		types.PREVIEW_PROJ,
		types.PREVIEW_TAE,
		types.PREVIEW_VAE,
	}

	for _, preview := range previews {
		namePtr := binding.PreviewName(preview)
		if namePtr == nil {
			t.Errorf("PreviewName returned nil for preview %d", preview)
			continue
		}

		preview2 := binding.StrToPreview(namePtr)
		t.Logf("Preview %d -> %s -> Preview %d", preview, utils.GoString(namePtr), preview2)
	}
}
