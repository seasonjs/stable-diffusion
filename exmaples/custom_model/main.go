package main

import (
	sd "github.com/seasonjs/stable-diffusion"
	"io"
	"log"
	"os"
	"path/filepath"
)

func main() {
	options := sd.DefaultOptions

	model, err := sd.NewAutoModel(options)
	if err != nil {
		print(err.Error())
		return
	}
	defer model.Close()

	model.SetLogCallback(func(level sd.LogLevel, msg string) {
		log.Println(msg)
	})

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

	params := sd.DefaultFullParams
	params.BatchCount = 1
	params.Width = 256
	params.Height = 256

	err = model.Predict("a girl, high quality", params, writers)
	if err != nil {
		print(err.Error())
		return
	}
}
