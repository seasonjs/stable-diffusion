package main

import (
	"github.com/seasonjs/hf-hub/api"
	sd "github.com/seasonjs/stable-diffusion"
	"io"
	"os"
)

func main() {
	options := sd.DefaultOptions

	model, err := sd.NewAutoModel(options)
	if err != nil {
		print(err.Error())
		return
	}
	defer model.Close()

	hapi, err := api.NewApi()
	if err != nil {
		print(err.Error())
		return
	}

	modelPath, err := hapi.Model("justinpinkney/miniSD").Get("miniSD.ckpt")
	if err != nil {
		print(err.Error())
		return
	}

	err = model.LoadFromFile(modelPath)
	if err != nil {
		print(err.Error())
		return
	}
	var writers []io.Writer
	filenames := []string{
		"../assets/love_cat0.png",
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

	err = model.Predict("british short hair cat, high quality", sd.DefaultFullParams, writers)
	if err != nil {
		print(err.Error())
	}
}
