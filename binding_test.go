package stable_diffusion

import (
	"fmt"
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

func TestSD(t *testing.T) {
	sd, err := NewCStableDiffusion(getLibrary())
	if err != nil {
		t.Log(err)
	}
	info := sd.cGetStableDiffusionSystemInfo()
	t.Log(info)
}
