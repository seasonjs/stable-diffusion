package main

import (
	sd "github.com/seasonjs/stable-diffusion"
	"io"
	"os"
	"path/filepath"
)

func main() {
	options := sd.DefaultStableDiffusionOptions
	options.Width = 512
	options.Height = 512

	model, err := sd.NewStableDiffusionAutoModel(options)
	if err != nil {
		print(err.Error())
		return
	}
	defer model.Close()

	err = model.LoadFromFile("./models/mysd.safetensors")
	if err != nil {
		print(err.Error())
		return
	}
	var writers []io.Writer

	girl, err := filepath.Abs("./assets/a_girl.png")
	if err != nil {
		print(err.Error())
		return
	}

	filenames := []string{
		girl,
	}
	for _, filename := range filenames {
		file, err := os.Create(filename)
		if err != nil {
			print(err.Error())
			return
		}
		defer file.Close()
		writers = append(writers, file)
	}

	err = model.Predict("a girl, high quality", writers)
	if err != nil {
		print(err.Error())
	}
}
