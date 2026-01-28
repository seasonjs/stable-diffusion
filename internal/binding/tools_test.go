package binding_test

import (
	"testing"

	"github.com/jupiterrider/ffi"
	"github.com/seasonjs/stable-diffusion/internal/binding"
	"github.com/seasonjs/stable-diffusion/internal/utils"
	"github.com/seasonjs/stable-diffusion/pkg/types"
)

func testSetupTools(lib ffi.Lib) error {
	return binding.LoadToosFuncs(lib)
}

func TestLoadTools(t *testing.T) {
	lib := testSetup(t)
	defer testCleanup(lib)
	err := testSetupTools(lib)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSetProgressCallback(t *testing.T) {
	lib := testSetup(t)
	defer testCleanup(lib)
	err := testSetupTools(lib)
	if err != nil {
		t.Fatal(err)
	}

	// 测试设置进度回调
	// 注意：这里我们只是测试函数调用，而不是实际的回调功能
	var callback uintptr = 0
	binding.SetProgressCallback(callback)
	t.Log("SetProgressCallback called successfully")
}

func TestSetPreviewCallback(t *testing.T) {
	lib := testSetup(t)
	defer testCleanup(lib)
	err := testSetupTools(lib)
	if err != nil {
		t.Fatal(err)
	}

	// 测试设置预览回调
	// 注意：这里我们只是测试函数调用，而不是实际的回调功能
	var callback uintptr = 0
	mode := types.PREVIEW_NONE
	interval := 10
	denoised := true
	noisy := false
	binding.SetPreviewCallback(callback, mode, interval, denoised, noisy)
	t.Log("SetPreviewCallback called successfully")
}

func TestGetNumPhysicalCores(t *testing.T) {
	lib := testSetup(t)
	defer testCleanup(lib)
	err := testSetupTools(lib)
	if err != nil {
		t.Fatal(err)
	}

	// 测试获取物理核心数
	cores := binding.GetNumPhysicalCores()
	if cores <= 0 {
		t.Fatalf("GetNumPhysicalCores returned invalid value: %d", cores)
	}
	t.Logf("GetNumPhysicalCores returned: %d", cores)
}

func TestGetSystemInfo(t *testing.T) {
	lib := testSetup(t)
	defer testCleanup(lib)
	err := testSetupTools(lib)
	if err != nil {
		t.Fatal(err)
	}

	// 测试获取系统信息
	infoPtr := binding.GetSystemInfo()
	if infoPtr == nil {
		t.Fatal("GetSystemInfo returned nil")
	}

	// 将 *byte 转换为字符串并打印
	info := utils.GoString(infoPtr)
	t.Logf("GetSystemInfo returned: %s", info)
}

func TestTypeName(t *testing.T) {
	lib := testSetup(t)
	defer testCleanup(lib)
	err := testSetupTools(lib)
	if err != nil {
		t.Fatal(err)
	}

	// 测试获取类型名称
	testTypes := []types.SdType{
		types.F32,
		types.F16,
		types.Q4_0,
		types.Q4_1,
		types.Q5_0,
		types.Q5_1,
		types.Q8_0,
	}

	for _, tType := range testTypes {
		namePtr := binding.TypeName(tType)
		if namePtr == nil {
			t.Fatalf("TypeName returned nil for type: %d", tType)
		}

		// 将 *byte 转换为字符串并打印
		name := utils.GoString(namePtr)
		t.Logf("TypeName for %d returned: %s", tType, name)
	}
}

func TestStrToSdType(t *testing.T) {
	lib := testSetup(t)
	defer testCleanup(lib)
	err := testSetupTools(lib)
	if err != nil {
		t.Fatal(err)
	}

	// 测试将字符串转换为 SdType
	testStrings := []string{
		"f32",
		"f16",
		"q4_0",
		"q4_1",
		"q5_0",
		"q5_1",
		"q8_0",
	}

	for _, str := range testStrings {
		strPtr, err := utils.CString(str)
		if err != nil {
			t.Fatal(err)
		}

		tType := binding.StrToSdType(strPtr)
		t.Logf("StrToSdType for '%s' returned: %d", str, tType)
	}
}

func TestRngTypeName(t *testing.T) {
	lib := testSetup(t)
	defer testCleanup(lib)
	err := testSetupTools(lib)
	if err != nil {
		t.Fatal(err)
	}

	// 测试获取 RNG 类型名称
	testTypes := []types.RNGType{
		types.CPU_RNG,
		types.CUDA_RNG,
		types.STD_DEFAULT_RNG,
	}

	for _, tType := range testTypes {
		namePtr := binding.RngTypeName(tType)
		if namePtr == nil {
			t.Fatalf("RngTypeName returned nil for type: %d", tType)
		}

		// 将 *byte 转换为字符串并打印
		name := utils.GoString(namePtr)
		t.Logf("RngTypeName for %d returned: %s", tType, name)
	}
}

func TestStrToRngType(t *testing.T) {
	lib := testSetup(t)
	defer testCleanup(lib)
	err := testSetupTools(lib)
	if err != nil {
		t.Fatal(err)
	}

	// 测试将字符串转换为 RNGType
	testStrings := []string{
		"cpu",
		"cuda",
		"tpu",
	}

	for _, str := range testStrings {
		strPtr, err := utils.CString(str)
		if err != nil {
			t.Fatal(err)
		}

		tType := binding.StrToRngType(strPtr)
		t.Logf("StrToRngType for '%s' returned: %d", str, tType)
	}
}

func TestConvert(t *testing.T) {
	lib := testSetup(t)
	defer testCleanup(lib)
	err := testSetupTools(lib)
	if err != nil {
		t.Fatal(err)
	}

	// 跳过 Convert 函数测试，因为它需要有效的文件路径
	// 实际使用时需要提供有效的文件路径
	t.Skip("Skipping Convert test - requires valid file paths")
}

func TestCommit(t *testing.T) {
	lib := testSetup(t)
	defer testCleanup(lib)
	err := testSetupTools(lib)
	if err != nil {
		t.Fatal(err)
	}

	// 测试获取提交信息
	commitPtr := binding.Commit()
	if commitPtr == nil {
		t.Fatal("Commit returned nil")
	}

	// 将 *byte 转换为字符串并打印
	commit := utils.GoString(commitPtr)
	t.Logf("Commit returned: %s", commit)
}

func TestVersion(t *testing.T) {
	lib := testSetup(t)
	defer testCleanup(lib)
	err := testSetupTools(lib)
	if err != nil {
		t.Fatal(err)
	}

	// 测试获取版本信息
	versionPtr := binding.Version()
	if versionPtr == nil {
		t.Fatal("Version returned nil")
	}

	// 将 *byte 转换为字符串并打印
	version := utils.GoString(versionPtr)
	t.Logf("Version returned: %s", version)
}
