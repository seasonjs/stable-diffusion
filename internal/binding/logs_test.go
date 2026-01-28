package binding_test

import (
	"testing"

	"github.com/ebitengine/purego"
	"github.com/jupiterrider/ffi"
	"github.com/seasonjs/stable-diffusion/internal/binding"
	"github.com/seasonjs/stable-diffusion/internal/utils"
)

func setupLogs(lib ffi.Lib) error {
	return binding.LoadLogFuns(lib)
}

func setDefaultLogCall(t *testing.T, lib ffi.Lib) {
	err := setupLogs(lib)
	if err != nil {
		t.Fatal(err)
	}

	callback := purego.NewCallback(func(level int32, text uintptr, data uintptr) uintptr {
		t.Logf("Log: [%d] - %s", level, utils.GoString(text))
		return 0
	})

	binding.SetLogCallback(callback)
}

func TestSetLogCallback(t *testing.T) {
	lib := testSetup(t)
	defer testCleanup(lib)
	setDefaultLogCall(t, lib)
}

func TestSetLogCallbackWithContext(t *testing.T) {
	lib := testSetup(t)
	defer testCleanup(lib)

	err := testSetupContext(lib)
	if err != nil {
		t.Fatal(err)
	}

	err = setupLogs(lib)
	if err != nil {
		t.Fatal(err)
	}

	callback := purego.NewCallback(func(level int32, text uintptr, data uintptr) uintptr {
		t.Logf("Log: [%d] - %s", level, utils.GoString(text))
		return 0
	})

	binding.SetLogCallback(callback)

	ctx := getDefaultContext(t)
	defer binding.FreeCtx(ctx)
}
