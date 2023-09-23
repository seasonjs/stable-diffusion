package stable_diffusion

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"runtime"
	"testing"
)

func getLibrary() string {
	switch runtime.GOOS {
	case "darwin":
		return "./deps/darwin/libstable-diffusion-arm64.dylib"
	case "linux":
		return "./deps/linux/libstable-diffusion.so"
	case "windows":
		return "./deps/windows/stable-diffusion_avx2_x64.dll"
	default:
		panic(fmt.Errorf("GOOS=%s is not supported", runtime.GOOS))
	}
}

// int n_threads = -1;
// std::string mode = TXT2IMG;
// std::string model_path;
// std::string output_path = "output.png";
// std::string init_img;
// std::string prompt;
// std::string negative_prompt;
// float cfg_scale = 7.0f;
// int w = 512;
// int h = 512;
// SampleMethod sample_method = EULER_A;
// Schedule schedule = DEFAULT;
// int sample_steps = 20;
// float strength = 0.75f;
// RNGType rng_type = CUDA_RNG;
// int64_t seed = 42;
// bool verbose = false;

func writeToFile(byteData []byte, height int, width int) {

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

	file, err := os.Create("./data/output.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	err = png.Encode(file, img)
	if err != nil {
		panic(err)
	}

	println("Image saved as output.png")
}

func TestSD(t *testing.T) {
	sd, err := NewCStableDiffusion(getLibrary())
	if err != nil {
		t.Log(err)
	}
	info := sd.cGetStableDiffusionSystemInfo()
	t.Log(info)
	ctx := sd.NewStableDiffusionCtx(20, true, true, STD_DEFAULT_RNG)
	defer ctx.Close()
	ctx.StableDiffusionLoadFromFile("./data/sd_v1-4_ggml_Q5.bin", DEFAULT)
	data := ctx.StableDiffusionTextToImage("a lovely cat", "", 7.0, 128, 128, EULER_A, 20, 42)
	writeToFile(data, 128, 128)
}
