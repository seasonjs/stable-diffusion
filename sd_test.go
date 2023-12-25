package sd

import (
	"image"
	"io"
	"os"
	"testing"
)

func TestNewStableDiffusionAutoModelPredict(t *testing.T) {
	options := DefaultStableDiffusionOptions
	options.Width = 256
	options.Height = 256
	options.BatchCount = 2
	//options.SampleSteps = 2
	t.Log(options)
	model, err := NewStableDiffusionAutoModel(options)
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

	err = model.Predict("british short hair cat，high quality", writers)
	if err != nil {
		t.Error(err)
		return
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

	outfile, err := os.Create("./assets/shoes_cat.png")
	if err != nil {
		t.Error(err)
		return
	}
	defer outfile.Close()

	err = model.ImagePredict(inFile, "the cat that wears shoe", outfile)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestImageToByte(t *testing.T) {
	inFile, err := os.Open("./assets/love_cat0.png")
	if err != nil {
		t.Error(err)
		return
	}
	defer inFile.Close()

	outfile, err := os.Create("./assets/test_cat0.png")
	if err != nil {
		t.Error(err)
		return
	}
	defer outfile.Close()

	decode, _, err := image.Decode(inFile)
	if err != nil {
		t.Error(err)
		return
	}
	outputsBytes := imageToBytes(decode)

	outputsImage := bytesToImage(outputsBytes, 256, 256)
	err = imageToWriter(outputsImage, "PNG", outfile)
	if err != nil {
		t.Error(err)
		return
	}
}
