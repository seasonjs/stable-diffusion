package sd

import (
	"io"
	"os"
	"testing"
)

func TestNewStableDiffusionAutoModelPredict(t *testing.T) {
	options := DefaultStableDiffusionOptions
	options.Width = 256
	options.Height = 256
	options.BatchCount = 4
	options.SampleSteps = 8
	t.Log(options)
	model, err := NewStableDiffusionAutoModel(options)
	if err != nil {
		t.Error(err)
	}
	defer model.Close()
	err = model.LoadFromFile("./models/miniSD-f16.gguf")
	if err != nil {
		t.Error(err)
	}
	writers := make([]io.Writer, options.BatchCount)
	filenames := []string{
		"./assets/love_cat3.png",
		"./assets/love_cat4.png",
		"./assets/love_cat5.png",
		"./assets/love_cat6.png"}
	for idx, filename := range filenames {
		file, err := os.Create(filename)
		if err != nil {
			t.Error(err)
		}
		defer file.Close()
		writers[idx] = file
	}

	err = model.Predict("a lovely cat, high quality", writers)
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
	err = model.LoadFromFile("./models/v1-5-pruned-emaonly_ggml_f16.bin")
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
