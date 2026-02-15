package binding_test

import (
	"fmt"
	"testing"

	"github.com/jupiterrider/ffi"
	"github.com/seasonjs/stable-diffusion/v2/internal/binding"
	"github.com/seasonjs/stable-diffusion/v2/internal/utils"
)

func testSetupImageGen(lib ffi.Lib) error {
	return binding.LoadImgGenFuns(lib)
}

func TestLoadImgGen(t *testing.T) {
	lib := testSetup(t)
	defer testCleanup(lib)
	err := testSetupImageGen(lib)
	if err != nil {
		t.Fatal(err)
	}
}

func TestImgGenParamsInit(t *testing.T) {
	lib := testSetup(t)
	defer testCleanup(lib)

	err := testSetupImageGen(lib)
	if err != nil {
		t.Fatal(err)
	}

	params := binding.ImgGenParamsInit()
	t.Logf("ImgGenParamsInit returned params: %+v", params)
}

func TestImgGenParamsToStr(t *testing.T) {
	lib := testSetup(t)
	defer testCleanup(lib)
	err := testSetupImageGen(lib)
	if err != nil {
		t.Fatal(err)
	}

	params := binding.ImgGenParamsInit()

	// 测试将参数转换为字符串
	strPtr := binding.ImgGenParamsToStr(params)
	if strPtr == nil {
		t.Fatal("ImgGenParamsToStr returned nil")
	}

	// 将 *byte 转换为字符串并打印
	str := utils.GoString(strPtr)
	t.Logf("ImgGenParamsToStr returned: %s", str)
}

func TestGenerateImage(t *testing.T) {
	lib := testSetup(t)
	defer testCleanup(lib)

	libc := testSetupStd(t)
	defer testCleanup(libc)

	err := testSetupContext(lib)
	if err != nil {
		t.Fatal(err)
	}

	// 加载必要的函数
	err = testSetupImageGen(lib)
	if err != nil {
		t.Fatal(err)
	}

	// 初始化图像生成参数
	params := binding.ImgGenParamsInit()

	// 设置基本参数
	prompt := "a cat"
	promptPtr, err := utils.CString(prompt)
	if err != nil {
		t.Fatal(err)
	}
	params.Prompt = promptPtr
	params.Width = 256
	params.Height = 256
	params.BatchCount = 1
	params.SampleParams.SampleSteps = 10

	ctx := getDefaultContext(t)
	defer binding.FreeCtx(ctx)
	// 生成图像
	result := binding.GenerateImage(ctx, &params)

	images := utils.GoImageSlice(result, int(params.BatchCount))
	defer utils.FreeImageSlice(result, int(params.BatchCount))

	for i, image := range images {
		writeTestImageToFile(t, image.Data, int(image.Width), int(image.Height), fmt.Sprintf("../../testdata/images/image_gen%d.png", i))
	}

	t.Logf("GenerateImage returned %d images", result)
}
