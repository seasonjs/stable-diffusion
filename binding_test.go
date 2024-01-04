// Copyright (c) seasonjs. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package sd_test

import (
	"fmt"
	sd "github.com/seasonjs/stable-diffusion"
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
		return "./deps/darwin/libsd-abi.dylib"
	case "linux":
		return "./deps/linux/libstable-diffusion.so"
	case "windows":
		return "./deps/windows/sd-abi_avx2.dll"
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

func readFromFile(t *testing.T, path string) *sd.Image {
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
	return &sd.Image{
		Width:  uint32(width),
		Height: uint32(height),
		Data:   img,
	}
}

func TestNewCStableDiffusionText2Img(t *testing.T) {
	diffusion, err := sd.NewCStableDiffusion(getLibrary())
	if err != nil {
		t.Error(err)
		return
	}
	diffusion.SetLogCallBack(func(level sd.LogLevel, text string) {
		fmt.Printf("%s", text)
	})
	ctx := diffusion.NewCtx("./models/miniSD.ckpt", "", "", "", false, false, true, 4, sd.F16, sd.CUDA_RNG, sd.DEFAULT)
	defer diffusion.FreeCtx(ctx)

	images := diffusion.PredictImage(ctx, "british short hair cat, high quality", "", 0, 7.0, 256, 256, sd.EULER_A, 10, 42, 1)

	writeToFile(t, images[0].Data, 256, 256, "./assets/test.png")
}

func TestNewCStableDiffusionImg2Img(t *testing.T) {
	diffusion, err := sd.NewCStableDiffusion(getLibrary())
	if err != nil {
		t.Error(err)
		return
	}
	diffusion.SetLogCallBack(func(level sd.LogLevel, text string) {
		fmt.Printf("%s", text)
	})
	ctx := diffusion.NewCtx("./models/miniSD.ckpt", "", "", "", false, false, true, -1, sd.F16, sd.CUDA_RNG, sd.DEFAULT)
	defer diffusion.FreeCtx(ctx)

	img := readFromFile(t, "./assets/test.png")
	images := diffusion.ImagePredictImage(ctx, *img, "cat wears shoes, high quality", "", 0, 7.0, 256, 256, sd.EULER_A, 20, 0.4, 42, 1)

	writeToFile(t, images[0].Data, 256, 256, "./assets/test1.png")
}
