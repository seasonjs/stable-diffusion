package binding_test

import (
	"testing"
	"unsafe"

	"github.com/jupiterrider/ffi"
	"github.com/seasonjs/stable-diffusion/internal/binding"
	"github.com/seasonjs/stable-diffusion/internal/utils"
)

func testSetupUpScaler(lib ffi.Lib) error {
	return binding.LoadUpsScalerFuncs(lib)
}

func getDefaultUpscalerContext(t *testing.T) binding.UpscalerContext {

	esrganPath := getTestUpscalerModelPath(t)
	offloadParamsToCPU := false
	direct := false
	nThreads := int32(4)
	tileSize := int32(128)
	cEsrganPath,err:=utils.CString(esrganPath)
	if err!= nil {
		t.Fatal(err)
	}
	// 创建超分辨率上下文
	upscalerCtx := binding.NewUpscalerCtx(cEsrganPath, offloadParamsToCPU, direct, nThreads, tileSize)

	return upscalerCtx
}

func TestNewAndFreeUpscalerCtx(t *testing.T) {
	lib := testSetup(t)
	defer testCleanup(lib)

	err := testSetupUpScaler(lib)
	if err != nil {
		t.Fatal(err)
	}
	upscalerCtx := getDefaultUpscalerContext(t)
	defer binding.FreeUpscalerCtx(upscalerCtx)
	t.Log("NewUpscalerCtx and FreeUpscalerCtx called successfully")
}

func TestGetUpscaleFactor(t *testing.T) {
	lib := testSetup(t)
	defer testCleanup(lib)

	err := testSetupUpScaler(lib)
	if err != nil {
		t.Fatal(err)
	}

	// 创建超分辨率上下文
	upscalerCtx := getDefaultUpscalerContext(t)
	defer binding.FreeUpscalerCtx(upscalerCtx)

	// 调用GetUpscaleFactor函数，确保它能被调用而不崩溃
	factor := binding.GetUpscaleFactor(upscalerCtx)
	t.Logf("GetUpscaleFactor returned: %d", factor)
}

func TestUpscale(t *testing.T) {
	lib := testSetup(t)
	defer testCleanup(lib)

	err := testSetupUpScaler(lib)
	if err != nil {
		t.Fatal(err)
	}

	// 创建超分辨率上下文
	upscalerCtx := getDefaultUpscalerContext(t)
	defer binding.FreeUpscalerCtx(upscalerCtx)

	inputImage := readTestImageFromFile(t, "../../testdata/images/love_cat0.png")

	var upscaleFactor uint32 = 4

	data := unsafe.SliceData(inputImage.Data)
	ci := binding.Image{
		Width:   inputImage.Width,
		Height:  inputImage.Height,
		Channel: inputImage.Channel,
		Data:    uintptr(unsafe.Pointer(&data)),
	}
	// 调用Upscale函数，确保它能被调用而不崩溃
	result := binding.Upscale(upscalerCtx, ci, upscaleFactor)

	goImage:= utils.GoImage(result)
	writeTestImageToFile(t, goImage.Data, int(goImage.Height), int(goImage.Width), "../../testdata/images/upscaler.png")
}
