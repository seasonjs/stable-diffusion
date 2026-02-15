package binding_test

import (
	"fmt"
	"image"
	"runtime"
	"testing"

	"image/color"
	"image/png"
	"os"

	"github.com/jupiterrider/ffi"
	"github.com/seasonjs/hf-hub/api"
	"github.com/seasonjs/stable-diffusion/v2/internal/binding"
	"github.com/seasonjs/stable-diffusion/v2/internal/utils"
	"github.com/seasonjs/stable-diffusion/v2/pkg/loader"
	"github.com/seasonjs/stable-diffusion/v2/pkg/types"
)

func getLibrary(t *testing.T) (ffi.Lib, error) {
	switch runtime.GOOS {
	case "darwin":
		return loader.LoadLibrary("../../testdata/deps/darwin", "sd-abi")
	case "linux":
		return loader.LoadLibrary("../../testdata/deps/linux", "sd-abi")
	case "windows":
		// 使用CUDA版本的DLL文件
		return loader.LoadLibrary("../../testdata/deps/windows", "sd-abi_cuda12")
		// return loader.LoadLibrary("../../testdata/deps/windows", "stable-diffusion")
	default:
		panic(fmt.Errorf("GOOS=%s is not supported", runtime.GOOS))
	}
}

func testSetup(t *testing.T) ffi.Lib {
	lib, err := getLibrary(t)
	if err != nil {
		t.Fatal("unable to load library", err.Error())
	}
	return lib
}

func testCleanup(lib ffi.Lib) error {
	return lib.Close()
}

func getTestModelPath(t *testing.T) *byte {
	hapi, err := api.NewApi()
	if err != nil {
		t.Fatal(err)
	}

	modelPath, err := hapi.Model("justinpinkney/miniSD").Get("miniSD.ckpt")
	if err != nil {
		t.Fatal(err)
	}

	b, err := utils.CString(modelPath)
	if err != nil {
		t.Fatal(err)
	}

	return b
}

func getTestUpscalerModelPath(t *testing.T) string {
	hapi, err := api.NewApi()
	if err != nil {
		t.Fatal(err)
	}

	modelPath, err := hapi.Model("ai-forever/Real-ESRGAN").Get("RealESRGAN_x4.pth")
	if err != nil {
		t.Fatal(err)
	}

	return modelPath
}

func getDefaultContext(t *testing.T) binding.Context {
	ctxParams := binding.CtxParamsInit()

	ctxParams.ModelPath = getTestModelPath(t)
	ctx := binding.NewCtx(&ctxParams)
	if ctx == 0 {
		t.Fatal("NewCtx return 0")
	}

	return ctx
}

func readTestImageFromFile(t *testing.T, path string) types.Image {
	file, err := os.Open(path)
	if err != nil {
		t.Error(err)
	}
	defer file.Close()
	decode, err := png.Decode(file)
	if err != nil {
		t.Error(err)
	}

	bounds := decode.Bounds()
	width := bounds.Max.X - bounds.Min.X
	height := bounds.Max.Y - bounds.Min.Y
	size := width * height * 3
	img := make([]byte, size)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			idx := ((y-bounds.Min.Y)*width + (x - bounds.Min.X)) * 3
			r, g, b, _ := decode.At(x, y).RGBA()
			// RGBA() 返回 16位值 (0-65535)，需要转换为 8位 (0-255)
			img[idx] = byte(r >> 8)
			img[idx+1] = byte(g >> 8)
			img[idx+2] = byte(b >> 8)
		}
	}
	return types.Image{
		Width:   uint32(width),
		Height:  uint32(height),
		Channel: 3, // RGB 图像有 3 个通道
		Data:    img,
	}
}

func writeTestImageToFile(t *testing.T, byteData []byte, height int, width int, outputPath string) {

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			idx := (y*width + x) * 3
			img.Set(x, y, color.RGBA{
				R: byteData[idx],
				G: byteData[idx+1],
				B: byteData[idx+2],
				A: 255,
			})
		}
	}

	file, err := os.Create(outputPath)
	if err != nil {
		t.Error(err)
	}
	defer file.Close()

	err = png.Encode(file, img)
	if err != nil {
		t.Error(err)
	}
	t.Log("Image saved at", outputPath)
}

func testSetupStd(t *testing.T) ffi.Lib {
	lib, err := utils.LoadStdLib()
	if err != nil {
		t.Fatal(err)
	}

	utils.LoadStdFuns(lib)
	return lib
}
