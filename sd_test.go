package sd_test

import (
	sd "github.com/seasonjs/stable-diffusion"
	"io"
	"os"
	"testing"
)

func TestNewStableDiffusionAutoModelPredict(t *testing.T) {
	options := sd.DefaultOptions
	t.Log(options)
	model, err := sd.NewAutoModel(options)
	if err != nil {
		t.Error(err)
		return
	}
	defer model.Close()
	err = model.LoadFromFile("./models/miniSD.ckpt")
	if err != nil {
		t.Error(err)
		return
	}
	var writers []io.Writer
	filenames := []string{
		"./assets/love_cat0.png",
		"./assets/love_cat1.png",
	}
	for _, filename := range filenames {
		file, err := os.Create(filename)
		if err != nil {
			t.Error(err)
			return
		}
		defer file.Close()
		writers = append(writers, file)
	}

	err = model.Predict("british short hair cat，high quality", sd.DefaultFullParams, writers)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestNewStableDiffusionAutoModelImagePredict(t *testing.T) {
	options := sd.DefaultOptions
	options.VaeDecodeOnly = false
	t.Log(options)
	model, err := sd.NewAutoModel(options)
	if err != nil {
		t.Error(err)
		return
	}
	defer model.Close()
	err = model.LoadFromFile("./models/mysd.safetensors")
	if err != nil {
		t.Error(err)
		return
	}
	inFile, err := os.Open("./assets/love_cat0.png")
	if err != nil {
		t.Error(err)
		return
	}
	defer inFile.Close()

	var writers []io.Writer
	filenames := []string{
		"./assets/love_cat0_m.png",
		"./assets/love_cat1_m.png",
		//"./assets/love_cat5.png",
		//"./assets/love_cat6.png"
	}
	for _, filename := range filenames {
		file, err := os.Create(filename)
		if err != nil {
			t.Error(err)
			return
		}
		defer file.Close()
		writers = append(writers, file)
	}

	err = model.ImagePredict(inFile, "the cat that wears shoe", sd.DefaultFullParams, writers)
	if err != nil {
		t.Error(err)
		return
	}
}
