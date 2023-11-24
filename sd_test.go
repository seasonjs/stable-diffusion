package sd

import (
	"os"
	"testing"
)

func TestNewStableDiffusionAutoModelPredict(t *testing.T) {
	options := DefaultStableDiffusionOptions
	options.Width = 512
	options.Height = 512
	t.Log(options)
	model, err := NewStableDiffusionAutoModel(options)
	if err != nil {
		t.Error(err)
	}
	defer model.Close()
	err = model.LoadFromFile("./models/sd_v1-4_ggml_Q5.bin")
	if err != nil {
		t.Error(err)
	}
	file, err := os.Create("./assets/love_cat2.png")
	defer file.Close()
	if err != nil {
		t.Error(err)
	}
	err = model.Predict("A lovely cat, high quality", file)
	if err != nil {
		t.Error(err)
	}

}
func TestNewStableDiffusionAutoModelImagePredict(t *testing.T) {
	options := DefaultStableDiffusionOptions
	options.Width = 512
	options.Height = 512
	options.VaeDecodeOnly = false
	t.Log(options)
	model, err := NewStableDiffusionAutoModel(options)
	if err != nil {
		t.Error(err)
	}
	defer model.Close()
	err = model.LoadFromFile("./models/sd_v1-4_ggml_Q5.bin")
	if err != nil {
		t.Error(err)
	}
	inFile, err := os.Open("./assets/love_cat2.png")
	defer inFile.Close()
	if err != nil {
		t.Error(err)
	}
	outfile, err := os.Create("./assets/shoes_cat.png")
	if err != nil {
		t.Error(err)
	}
	defer outfile.Close()
	err = model.ImagePredict(inFile, "the cat that wears shoes", outfile)
	if err != nil {
		t.Error(err)
	}
}
