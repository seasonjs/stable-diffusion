// Copyright (c) seasonjs. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package binding_test

import (
	"fmt"
	"reflect"

	"testing"
	"unsafe"

	"github.com/jupiterrider/ffi"
	"github.com/seasonjs/stable-diffusion/internal/binding"
	// "github.com/seasonjs/hf-hub/api"
	// "github.com/seasonjs/stable-diffusion/internal/binding"
)

// StructOffsetChecker 结构体偏移量检查工具
// 用于验证Go结构体与C结构体的内存布局是否匹配

// StructOffsetResult 存储结构体偏移量检查结果
type StructOffsetResult struct {
	StructSizeMatch    bool            // 结构体大小是否匹配
	StructSize         uint64          // 结构体大小
	StructTypeSize     uint64          // 结构体类型大小
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
	result.StructSize = structSize
	result.StructTypeSize = structType.Size
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
func PrintStructOffsetResult(t *testing.T,result StructOffsetResult, structName string) {
	t.Logf("=== %s 结构体偏移量检查结果 ===\n", structName)
	t.Logf("结构体大小匹配: %t, 结构体大小: %d, 结构体类型大小: %d\n", result.StructSizeMatch, result.StructSize, result.StructTypeSize)
	t.Logf("总字段数: %d, 匹配字段数: %d\n", result.TotalFields, result.MatchedFields)
	t.Log("字段偏移量匹配情况:")
	for fieldName, match := range result.FieldOffsetMatches {
		status := "✓"
		if !match {
			status = "✗"
		}
		t.Logf("  %s: %s\n", fieldName, status)
	}
	t.Log("==============================")
}

func CheckAnyStructOffsets(structType ffi.Type, structValue interface{}) StructOffsetResult {
	return CheckStructOffsets(structType, structValue)
}

// CheckCtxParamsOffsets 检查CtxParams结构体偏移量（示例：使用改进后的工具）
func TestCheckCtxParamsOffsets(t *testing.T) {
	var image binding.Image
	result := CheckAnyStructOffsets(binding.FFITypeImage, image)
	PrintStructOffsetResult(t,result, "Image")

	offsets := make([]uint64, 4)

    // 4. 调用 GetStructOffsets
    status := ffi.GetStructOffsets(ffi.DefaultAbi, &binding.FFITypeImage, &offsets[0])

    if status == ffi.OK {
        fmt.Printf("Total Size: %d, Alignment: %d\n", binding.FFITypeImage.Size, binding.FFITypeImage.Alignment)
        fmt.Printf("Offsets:\n")
        fmt.Printf("  width:   %d\n", offsets[0])
        fmt.Printf("  height:  %d\n", offsets[1])
        fmt.Printf("  channel: %d\n", offsets[2])
        fmt.Printf("  data:    %d\n", offsets[3])
    } else {
        fmt.Printf("Failed to get offsets: %v\n", status)
    }
}

// // CheckImageOffsets 检查Image结构体偏移量（示例：使用改进后的工具）
// func TestCheckImageOffsets(t *testing.T) {
// 	var img diffusion.Image
// 	result := CheckStructOffsets(diffusion.FFITypeImage, img)
// 	PrintStructOffsetResult(result, "Image")
// }

// // CheckSampleParamsOffsets 检查SampleParams结构体偏移量（示例：使用改进后的工具）
// func TestCheckSampleParamsOffsets(t *testing.T) {
// 	var sampleParams diffusion.SampleParams
// 	result := CheckStructOffsets(diffusion.FFITypeSampleParams, sampleParams)
// 	PrintStructOffsetResult(result, "SampleParams")
// }

// func TestNewCStableDiffusionText2Img(t *testing.T) {
// 	cd, err := diffusion.NewCStableDiffusion(getLibrary())
// 	if err != nil {
// 		t.Error(err)
// 		return
// 	}
// 	cd.SetLogCallback(func(level diffusion.LogLevel, text string) {
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
// 	ctxParams := cd.CtxParamsInit()
// 	ctxParams.ModelPath = diffusion.CString(modelPath)

// 	ctx := cd.NewCtx(&ctxParams)
// 	defer cd.FreeCtx(ctx)

// 	// 创建并初始化ImgGenParams
// 	imgGenParams := cd.ImgGenParamsInit()
// 	imgGenParams.Prompt = diffusion.CString("british short hair cat, high quality")
// 	imgGenParams.Width = 256
// 	imgGenParams.Height = 256
// 	imgGenParams.SampleParams.Scheduler = diffusion.DISCRETE
// 	imgGenParams.SampleParams.SampleMethod = diffusion.EULER_A
// 	imgGenParams.SampleParams.SampleSteps = 10
// 	imgGenParams.Seed = 43
// 	imgGenParams.BatchCount = 1

// 	images := cd.GenerateImage(ctx, &imgGenParams)

// 	if len(images) > 0 {
// 		writeToFile(t, goImageSlice(images, len(images))[0].Data, 256, 256, "./assets/aaaaaa.png")
// 	}
// }

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
