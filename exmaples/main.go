package main

import (
	"fmt"
	"github.com/Cyberhan123/libfetch"
	"github.com/seasonjs/hf-hub/api"
	"github.com/seasonjs/stable-diffusion/pkg/diffusion"
	"github.com/seasonjs/stable-diffusion/pkg/types"
	"io"
	"os"
)

func main() {
	var libDir = "./.cache"
	libfetchApi := libfetch.NewApi()
	libfetchApi.SetInstallDir(libDir)

	err := libfetchApi.Repo("leejet/stable-diffusion.cpp").Latest().Install(func(version string) string {
		return fmt.Sprintf("sd-master-%s-bin-win-cuda12-x64.zip", version)
	})
	if err != nil {
		print(err.Error())
		return
	}

	err = libfetchApi.Repo("leejet/stable-diffusion.cpp").Latest().Install(func(version string) string {
		return "cudart-sd-bin-win-cu12-x64.zip"
	})
	if err != nil {
		print(err.Error())
		return
	}

	model, err := diffusion.New(libDir)
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

	ctx, err := model.NewContext(&types.ContextParams{
		ModelPath: modelPath,
	})
	if err != nil {
		print(err.Error())
		return
	}
	defer ctx.Close()

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

	err = ctx.GenerateImage(&types.ImgGenParams{
		Prompt:     "british short hair cat, high quality",
		BatchCount: 1,
		Width:      256,
		Height:     256,
	}, types.PNG, writers)
	if err != nil {
		print(err.Error())
	}
}
