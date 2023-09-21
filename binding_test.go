package stable_diffusion

import "testing"

func TestSD(t *testing.T) {
	sd, err := NewCStableDiffusion("./deps/windows/stable-diffusion.dll")
	if err != nil {
		t.Log(err)
	}
	info := sd.cGetStableDiffusionSystemInfo()
	t.Log(info)
}
