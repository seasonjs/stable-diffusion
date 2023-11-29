// Copyright (c) seasonjs. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package sd

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
		return "./deps/darwin/libstable-diffusion_arm64.dylib"
	case "linux":
		return "./deps/linux/libstable-diffusion.so"
	case "windows":
		return "./deps/windows/stable-diffusion_avx2.dll"
	default:
		panic(fmt.Errorf("GOOS=%s is not supported", runtime.GOOS))
	}
}

// int n_threads = -1;
// std::string mode = TXT2IMG;
// std::string model_path;
// std::string output_path = "love_cat2.png";
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

func writeToFile(t *testing.T, byteData []byte, height int, width int, outputPath string) {

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

func readFromFile(t *testing.T, path string) []byte {
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
	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			idx := (y*width + x) * 3
			r, g, b, _ := decode.At(x, y).RGBA()
			img[idx] = byte(r)
			img[idx+1] = byte(g)
			img[idx+2] = byte(b)
		}
	}
	return img
}

//func TestStableDiffusionTextToImage(t *testing.T) {
//	sd, err := NewCStableDiffusion(getLibrary())
//	if err != nil {
//		t.Log(err)
//	}
//	ctx := sd.NewStableDiffusionCtx(8, true, true, "", CUDA_RNG)
//	defer ctx.Close()
//	ctx.StableDiffusionLoadFromFile("./models/miniSD-ggml-model-q5_0.bin", DEFAULT)
//	data, _ := ctx.StableDiffusionTextToImage("A lovely cat, high quality", "", 7.0, 256, 256, EULER_A, 20, 42, 1)
//	writeToFile(t, data[1], 256, 256, "./data/love_cat2.png")
//}
//
//func TestStableDiffusionImgToImage(t *testing.T) {
//	sd, err := NewCStableDiffusion(getLibrary())
//	if err != nil {
//		t.Log(err)
//	}
//	ctx := sd.NewStableDiffusionCtx(8, false, true, "", CUDA_RNG)
//	defer ctx.Close()
//	ctx.StableDiffusionLoadFromFile("./models/miniSD-ggml-model-q5_0.bin", DEFAULT)
//	img := readFromFile(t, "./data/love_cat2.png")
//	data, _ := ctx.StableDiffusionImageToImage(img, "A lovely cat that theme pink", "", 7.0, 256, 256, EULER_A, 20, 0.4, 42)
//	writeToFile(t, data, 256, 256, "./data/output1.png")
//}
//
//func TestBase64(t *testing.T) {
//	img := readFromFile(t, "./assets/love_cat2.png")
//	imgBase64 := base64.StdEncoding.EncodeToString(img)
//	t.Log(imgBase64)
//}
