// Copyright (c) seasonjs. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package sd_test

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"runtime"
	"testing"
	"unsafe"
	"reflect"

	"github.com/seasonjs/hf-hub/api"
	sd "github.com/seasonjs/stable-diffusion"
	"github.com/jupiterrider/ffi"
)

func getLibrary() string {
	switch runtime.GOOS {
	case "darwin":
		return "./deps/darwin/libsd-abi.dylib"
	case "linux":
		return "./deps/linux/libsd-abi.so"
	case "windows":
		// 使用CUDA版本的DLL文件
		return "./deps/windows/sd-abi_cuda12.dll"
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

// func readFromFile(t *testing.T, path string) *sd.Image {
// 	file, err := os.Open(path)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	defer file.Close()
// 	decode, err := png.Decode(file)
// 	if err != nil {
// 		t.Error(err)
// 	}

// 	bounds := decode.Bounds()
// 	width := bounds.Max.X - bounds.Min.X
// 	height := bounds.Max.Y - bounds.Min.Y
// 	size := width * height * 3
// 	img := make([]byte, size)
// 	for x := bounds.Min.X; x < bounds.Max.X; x++ {
// 		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
// 			idx := (y*width + x) * 3
// 			r, g, b, _ := decode.At(x, y).RGBA()
// 			img[idx] = byte(r)
// 			img[idx+1] = byte(g)
// 			img[idx+2] = byte(b)
// 		}
// 	}
// 	return &sd.Image{
// 		Width:  uint32(width),
// 		Height: uint32(height),
// 		Data:   img,
// 	}
// }

type goImage struct {
	Channel uint32
	Width   uint32
	Height  uint32
	Data    []byte
}

func goImageSlice(imgSlice []sd.Image, size int) []goImage {
	// We take the address and then dereference it to trick go vet from creating a possible misuse of unsafe.Pointer
	// ptr := *(*unsafe.Pointer)(unsafe.Pointer(&c))
	// if ptr == nil {
	// 	return nil
	// }
	// img := (*sd.Image)(ptr)
	goImages := make([]goImage, 0, size)
	// imgSlice := unsafe.Slice(img, size)
	for _, image := range imgSlice {
		var gImg goImage
		gImg.Channel = image.Channel
		gImg.Width = image.Width
		gImg.Height = image.Height
		dataPtr := *(*unsafe.Pointer)(unsafe.Pointer(&image.Data))
		gImg.Data = unsafe.Slice((*byte)(dataPtr), image.Channel*image.Width*image.Height)
		goImages = append(goImages, gImg)
	}
	return goImages
}


// StructOffsetChecker 结构体偏移量检查工具
// 用于验证Go结构体与C结构体的内存布局是否匹配

// StructOffsetResult 存储结构体偏移量检查结果
type StructOffsetResult struct {
	StructSizeMatch    bool            // 结构体大小是否匹配
	FieldOffsetMatches map[string]bool // 每个字段的偏移量是否匹配
	TotalFields        int             // 总字段数
	MatchedFields      int             // 匹配的字段数
}

// CheckStructOffsets 使用反射检查任意结构体的偏移量是否匹配
// 参数:
//   - structType: ffi.Type定义的结构体类型
//   - structValue: 要检查的Go结构体实例（使用interface{}类型接收任意结构体）
//
// 返回:
//   - StructOffsetResult: 检查结果
func CheckStructOffsets(structType ffi.Type, structValue interface{}) StructOffsetResult {
	result := StructOffsetResult{
		FieldOffsetMatches: make(map[string]bool),
	}

	// 使用反射获取结构体信息
	structVal := reflect.ValueOf(structValue)
	if structVal.Kind() == reflect.Ptr {
		structVal = structVal.Elem()
	}

	if structVal.Kind() != reflect.Struct {
		panic("structValue必须是结构体或结构体指针类型")
	}

	structTypeReflect := structVal.Type()
	fieldCount := structTypeReflect.NumField()
	result.TotalFields = fieldCount

	// 获取结构体大小
	structSize := uint64(unsafe.Sizeof(structValue))
	result.StructSizeMatch = structSize == structType.Size

	// 获取结构体字段偏移量
	offsets := make([]uint64, fieldCount)
	if status := ffi.GetStructOffsets(ffi.DefaultAbi, &structType, &offsets[0]); status != ffi.OK {
		panic("获取结构体偏移量失败: " + status.String())
	}

	// 检查每个字段的偏移量
	for i := 0; i < fieldCount; i++ {
		field := structTypeReflect.Field(i)
		fieldName := field.Name
		fieldOffset := field.Offset

		match := uint64(fieldOffset) == offsets[i]
		result.FieldOffsetMatches[fieldName] = match
		if match {
			result.MatchedFields++
		}
	}

	return result
}

// PrintStructOffsetResult 打印结构体偏移量检查结果
func PrintStructOffsetResult(result StructOffsetResult, structName string) {
	fmt.Printf("=== %s 结构体偏移量检查结果 ===\n", structName)
	fmt.Printf("结构体大小匹配: %t\n", result.StructSizeMatch)
	fmt.Printf("总字段数: %d, 匹配字段数: %d\n", result.TotalFields, result.MatchedFields)
	fmt.Println("字段偏移量匹配情况:")
	for fieldName, match := range result.FieldOffsetMatches {
		status := "✓"
		if !match {
			status = "✗"
		}
		fmt.Printf("  %s: %s\n", fieldName, status)
	}
	fmt.Println("==============================")
}

// CheckAnyStructOffsets 检查任意结构体的偏移量是否匹配（最通用的接口）
// 参数:
//   - structType: ffi.Type定义的结构体类型
//   - structValue: 要检查的Go结构体实例或指针
//
// 返回:
//   - StructOffsetResult: 检查结果
func CheckAnyStructOffsets(structType ffi.Type, structValue interface{}) StructOffsetResult {
	return CheckStructOffsets(structType, structValue)
}

// CheckCtxParamsOffsets 检查CtxParams结构体偏移量（示例：使用改进后的工具）
func TestCheckCtxParamsOffsets(t *testing.T) {
	var ctxParams sd.CtxParams
	result := CheckStructOffsets(sd.FFITypeCtxParams, ctxParams)
	PrintStructOffsetResult(result, "CtxParams")
}

// CheckImageOffsets 检查Image结构体偏移量（示例：使用改进后的工具）
func TestCheckImageOffsets(t *testing.T) {
	var img sd.Image
	result := CheckStructOffsets(sd.FFITypeImage, img)
	PrintStructOffsetResult(result, "Image")
}

// CheckSampleParamsOffsets 检查SampleParams结构体偏移量（示例：使用改进后的工具）
func TestCheckSampleParamsOffsets(t *testing.T) {
	var sampleParams sd.SampleParams
	result := CheckStructOffsets(sd.FFITypeSampleParams, sampleParams)
	PrintStructOffsetResult(result, "SampleParams")
}

func TestNewCStableDiffusionText2Img(t *testing.T) {
	diffusion, err := sd.NewCStableDiffusion(getLibrary())
	if err != nil {
		t.Error(err)
		return
	}
	diffusion.SetLogCallback(func(level sd.LogLevel, text string) {
		fmt.Printf("%s", text)
	})
	hapi, err := api.NewApi()
	if err != nil {
		print(err.Error())
		return
	}

	modelPath, err := hapi.Model("justinpinkney/miniSD").Get("miniSD.ckpt")
	if err != nil {
		print(err.Error())
		return
	}

	// 创建并初始化CtxParams
	ctxParams := diffusion.CtxParamsInit()
	ctxParams.ModelPath = sd.CString(modelPath)
	// ctxParams.VaeDecodeOnly = false
	// ctxParams.FreeParamsImmediately = true
	// ctxParams.NThreads = 4
	// ctxParams.Wtype = sd.F16
	// ctxParams.RngType = sd.CUDA_RNG

	ctx := diffusion.NewCtx(&ctxParams)
	defer diffusion.FreeCtx(ctx)

	// 创建并初始化ImgGenParams
	var imgGenParams sd.ImgGenParams
	diffusion.ImgGenParamsInit(&imgGenParams)
	imgGenParams.Prompt = sd.CString("british short hair cat, high quality")
	imgGenParams.NegativePrompt = sd.CString("")
	imgGenParams.ClipSkip = 0
	imgGenParams.Width = 256
	imgGenParams.Height = 256
	imgGenParams.SampleParams.Scheduler = sd.DISCRETE
	imgGenParams.SampleParams.SampleMethod = sd.EULER_A
	imgGenParams.SampleParams.SampleSteps = 10
	imgGenParams.SampleParams.Guidance.TxtCfg = 7.0
	imgGenParams.Seed = 43
	imgGenParams.BatchCount = 1

	images := diffusion.GenerateImage(ctx, &imgGenParams)

	if len(images) > 0 {
		writeToFile(t, goImageSlice(images, len(images))[0].Data,  256, 256, "./assets/aaaaaa.png")
	}
}

// func TestNewCStableDiffusionImg2Img(t *testing.T) {
// 	diffusion, err := sd.NewCStableDiffusion(getLibrary())
// 	if err != nil {
// 		t.Error(err)
// 		return
// 	}
// 	diffusion.SetLogCallback(func(level sd.LogLevel, text string) {
// 		fmt.Printf("%s", text)
// 	})
// 	hapi, err := api.NewApi()
// 	if err != nil {
// 		print(err.Error())
// 		return
// 	}

// 	modelPath, err := hapi.Model("justinpinkney/miniSD").Get("miniSD.ckpt")
// 	if err != nil {
// 		print(err.Error())
// 		return
// 	}

// 	// 创建并初始化CtxParams
// 	var ctxParams sd.CtxParams
// 	diffusion.CtxParamsInit(&ctxParams)
// 	ctxParams.ModelPath = unsafe.StringData(modelPath)
// 	ctxParams.VaeDecodeOnly = false
// 	ctxParams.FreeParamsImmediately = true
// 	ctxParams.NThreads = -1
// 	ctxParams.Wtype = sd.F16
// 	ctxParams.RngType = sd.CUDA_RNG

// 	ctx := diffusion.NewCtx(&ctxParams)
// 	defer diffusion.FreeCtx(ctx)

// 	img := readFromFile(t, "./assets/test.png")

// 	// 创建并初始化ImgGenParams
// 	var imgGenParams sd.ImgGenParams
// 	diffusion.ImgGenParamsInit(&imgGenParams)
// 	imgGenParams.Prompt = unsafe.StringData("cat wears shoes, high quality")
// 	imgGenParams.NegativePrompt = unsafe.StringData("")
// 	imgGenParams.ClipSkip = 0
// 	imgGenParams.Width = 256
// 	imgGenParams.Height = 256
// 	imgGenParams.SampleParams.Scheduler = sd.DISCRETE
// 	imgGenParams.SampleParams.SampleMethod = sd.EULER_A
// 	imgGenParams.SampleParams.SampleSteps = 20
// 	imgGenParams.SampleParams.Guidance.TxtCfg = 7.0
// 	imgGenParams.Strength = 0.4
// 	imgGenParams.Seed = 42
// 	// imgGenParams.InitImage = *img

// 	images := diffusion.GenerateImage(ctx, &imgGenParams)

// 	if len(images) > 0 {
// 		writeToFile(t, images[0].Data, 256, 256, "./assets/test1.png")
// 	}
// }
