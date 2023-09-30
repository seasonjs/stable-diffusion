package sd

import (
	"os"
	"testing"
)

func TestNewStableDiffusionAutoModelPredict(t *testing.T) {
	options := DefaultStableDiffusionOptions
	options.Width = 256
	options.Height = 256
	t.Log(options)
	model, err := NewStableDiffusionAutoModel(options)
	if err != nil {
		t.Log(err)
	}
	defer model.Close()
	err = model.LoadFromFile("./models/miniSD-ggml-model-q5_0.bin")
	if err != nil {
		t.Log(err)
	}
	file, err := os.Create("./data/output2.png")
	defer file.Close()
	if err != nil {
		t.Log(err)
	}
	err = model.Predict("A lovely cat, high quality", file)
	if err != nil {
		t.Log(err)
	}

}
func TestNewStableDiffusionAutoModelImagePredict(t *testing.T) {
	options := DefaultStableDiffusionOptions
	options.Width = 256
	options.Height = 256
	options.VaeDecodeOnly = false
	t.Log(options)
	model, err := NewStableDiffusionAutoModel(options)
	if err != nil {
		t.Log(err)
	}
	defer model.Close()
	err = model.LoadFromFile("./models/miniSD-ggml-model-q5_0.bin")
	if err != nil {
		t.Log(err)
	}
	inFile, err := os.Open("./data/output2.png")
	defer inFile.Close()
	if err != nil {
		t.Log(err)
	}
	outfile, err := os.Create("./data/output3.png")
	if err != nil {
		t.Log(err)
	}
	defer outfile.Close()
	err = model.ImagePredict(inFile, "pink cat", outfile)
}
