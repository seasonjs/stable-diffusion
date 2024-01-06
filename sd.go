// Copyright (c) seasonjs. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package sd

import (
	"errors"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"io"
	"log"
	"os"
	"path/filepath"
)

type OutputsImageType string

const (
	PNG  OutputsImageType = "PNG"
	JPEG                  = "JPEG"
)

type Options struct {
	VaePath               string
	TaesdPath             string
	LoraModelDir          string
	VaeDecodeOnly         bool
	VaeTiling             bool
	FreeParamsImmediately bool
	Threads               int
	Wtype                 WType
	RngType               RNGType
	Schedule              Schedule
	GpuEnable             bool
}

type FullParams struct {
	NegativePrompt   string
	ClipSkip         int
	CfgScale         float32
	Width            int
	Height           int
	SampleMethod     SampleMethod
	SampleSteps      int
	Strength         float32
	Seed             int64
	BatchCount       int
	OutputsImageType OutputsImageType
}

var DefaultOptions = Options{
	Threads:               -1, // auto
	VaeDecodeOnly:         true,
	FreeParamsImmediately: true,
	RngType:               CUDA_RNG,
	Wtype:                 F32,
	Schedule:              DEFAULT,
}

var DefaultFullParams = FullParams{
	NegativePrompt:   "out of frame, lowers, text, error, cropped, worst quality, low quality, jpeg artifacts, ugly, duplicate, morbid, mutilated, out of frame, extra fingers, mutated hands, poorly drawn hands, poorly drawn face, mutation, deformed, blurry, dehydrated, bad anatomy, bad proportions, extra limbs, cloned face, disfigured, gross proportions, malformed limbs, missing arms, missing legs, extra arms, extra legs, fused fingers, too many fingers, long neck, username, watermark, signature",
	CfgScale:         7.0,
	Width:            512,
	Height:           512,
	SampleMethod:     EULER_A,
	SampleSteps:      20,
	Strength:         0.4,
	Seed:             42,
	BatchCount:       1,
	OutputsImageType: PNG,
}

type Model struct {
	ctx                *CStableDiffusionCtx
	options            *Options
	csd                CStableDiffusion
	isAutoLoad         bool
	dylibPath          string
	diffusionModelPath string
	esrganPath         string
	upscalerCtx        *CUpScalerCtx
}

func NewAutoModel(options Options) (*Model, error) {
	file, err := dumpSDLibrary(options.GpuEnable)
	if err != nil {
		return nil, err
	}

	if options.GpuEnable {
		log.Printf("If you want to try offload your model to the GPU. " +
			"Please confirm the size of your GPU memory to prevent memory overflow.")
	}
	dylibPath := file.Name()
	model, err := NewModel(dylibPath, options)
	if err != nil {
		return nil, err
	}
	model.isAutoLoad = true
	return model, nil
}

func NewModel(dylibPath string, options Options) (*Model, error) {
	csd, err := NewCStableDiffusion(dylibPath)
	if err != nil {
		return nil, err
	}

	return &Model{
		dylibPath: dylibPath,
		options:   &options,
		csd:       csd,
	}, nil
}

func (sd *Model) LoadFromFile(path string) error {
	if sd.ctx != nil {
		sd.csd.FreeCtx(sd.ctx)
		sd.ctx = nil
		log.Printf("model already loaded, free old model")
	}

	_, err := os.Stat(path)
	if err != nil {
		return errors.New("the system cannot find the model file specified")
	}

	if !filepath.IsAbs(path) {
		sd.diffusionModelPath, err = filepath.Abs(path)
		if err != nil {
			return err
		}
	} else {
		sd.diffusionModelPath = path
	}

	ctx := sd.csd.NewCtx(path,
		sd.options.VaePath,
		sd.options.TaesdPath,
		sd.options.LoraModelDir,
		sd.options.VaeDecodeOnly,
		sd.options.VaeTiling,
		sd.options.FreeParamsImmediately,
		sd.options.Threads,
		sd.options.Wtype,
		sd.options.RngType,
		sd.options.Schedule)
	sd.ctx = ctx
	return nil
}

func (sd *Model) SetOptions(options Options) {
	if sd.ctx != nil {
		sd.csd.FreeCtx(sd.ctx)
		sd.ctx = nil
		log.Printf("model already loaded, free old model and set new options")
	}
	sd.options = &options
	ctx := sd.csd.NewCtx(
		sd.diffusionModelPath,
		sd.options.VaePath,
		sd.options.TaesdPath,
		sd.options.LoraModelDir,
		sd.options.VaeDecodeOnly,
		sd.options.VaeTiling,
		sd.options.FreeParamsImmediately,
		sd.options.Threads,
		sd.options.Wtype,
		sd.options.RngType,
		sd.options.Schedule)
	sd.ctx = ctx
}

func (sd *Model) Predict(prompt string, params FullParams, writer []io.Writer) error {
	if len(writer) != params.BatchCount {
		return errors.New("writer count not match batch count")
	}
	if sd.ctx == nil {
		return errors.New("model not loaded")
	}

	if params.Width%8 != 0 || params.Height%8 != 0 {
		return errors.New("width and height must be multiples of 8")
	}

	images := sd.csd.PredictImage(
		sd.ctx,
		prompt,
		params.NegativePrompt,
		params.ClipSkip,
		params.CfgScale,
		params.Width,
		params.Height,
		params.SampleMethod,
		params.SampleSteps,
		params.Seed,
		params.BatchCount,
	)

	if images == nil || len(images) != params.BatchCount {
		return errors.New("predict failed")
	}

	for i, img := range images {
		outputsImage := bytesToImage(img.Data, int(img.Width), int(img.Height))

		err := imageToWriter(outputsImage, params.OutputsImageType, writer[i])
		if err != nil {
			return err
		}
	}

	return nil
}

func (sd *Model) ImagePredict(reader io.Reader, prompt string, params FullParams, writer []io.Writer) error {

	if len(writer) != params.BatchCount {
		return errors.New("writer count not match batch count")
	}

	if sd.ctx == nil {
		return errors.New("model not loaded")
	}

	decode, _, err := image.Decode(reader)
	if err != nil {
		return err
	}
	initImage := imageToBytes(decode)
	images := sd.csd.ImagePredictImage(
		sd.ctx,
		initImage,
		prompt,
		params.NegativePrompt,
		params.ClipSkip,
		params.CfgScale,
		params.Width,
		params.Height,
		params.SampleMethod,
		params.SampleSteps,
		params.Strength,
		params.Seed,
		params.BatchCount,
	)
	for i, img := range images {
		outputsImage := bytesToImage(img.Data, int(img.Width), int(img.Height))
		err = imageToWriter(outputsImage, params.OutputsImageType, writer[i])
		if err != nil {
			return err
		}
	}
	return nil
}

func (sd *Model) UpscaleImage(reader io.Reader, esrganPath string, upscaleFactor uint32, writer io.Writer) error {
	if sd.upscalerCtx == nil {
		sd.upscalerCtx = sd.csd.NewUpscalerCtx(esrganPath, sd.options.Threads, sd.options.Wtype)
	}
	decode, _, err := image.Decode(reader)
	if err != nil {
		return err
	}
	initImage := imageToBytes(decode)
	img := sd.csd.UpscaleImage(sd.upscalerCtx, initImage, upscaleFactor)
	outputsImage := bytesToImage(img.Data, int(img.Width), int(img.Height))
	err = imageToWriter(outputsImage, PNG, writer)
	return err
}

func (sd *Model) SetLogCallback(cb CLogCallback) {
	sd.csd.SetLogCallBack(cb)
}

func (sd *Model) Close() error {
	if sd.ctx != nil {
		sd.csd.FreeCtx(sd.ctx)
		sd.ctx = nil
	}

	if sd.upscalerCtx != nil {
		sd.csd.FreeUpscalerCtx(sd.upscalerCtx)
		sd.upscalerCtx = nil
	}

	if sd.csd != nil {
		err := sd.csd.Close()
		if err != nil {
			return err
		}
	}

	if sd.isAutoLoad {
		err := os.Remove(sd.dylibPath)
		if err != nil {
			return err
		}
	}
	return nil
}

func imageToBytes(decode image.Image) Image {
	bounds := decode.Bounds()
	width := bounds.Max.X - bounds.Min.X
	height := bounds.Max.Y - bounds.Min.Y
	size := width * height * 3
	bytesImg := make([]byte, size)
	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			idx := (y*width + x) * 3
			r, g, b, _ := decode.At(x, y).RGBA()
			bytesImg[idx] = byte(r >> 8)
			bytesImg[idx+1] = byte(g >> 8)
			bytesImg[idx+2] = byte(b >> 8)
		}
	}
	return Image{
		Width:  uint32(width),
		Height: uint32(height),
		Data:   bytesImg,
	}
}

func bytesToImage(byteData []byte, width, height int) image.Image {
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			idx := (y*width + x) * 3
			img.Set(x, y, color.RGBA{
				R: byteData[idx],
				G: byteData[idx+1],
				B: byteData[idx+2],
				A: 255,
			})
		}
	}
	return img
}

func imageToWriter(image image.Image, imageType OutputsImageType, writer io.Writer) error {
	switch imageType {
	case PNG:
		err := png.Encode(writer, image)
		if err != nil {
			return err
		}
	case JPEG:
		err := jpeg.Encode(writer, image, &jpeg.Options{Quality: 100})
		if err != nil {
			return err
		}
	default:
		return errors.New("unknown image type")
	}
	return nil
}

//func chunkBytes(data []byte, chunks int) [][]byte {
//	length := len(data)
//	chunkSize := (length + chunks - 1) / chunks
//	result := make([][]byte, chunks)
//
//	for i := 0; i < chunks; i++ {
//		start := i * chunkSize
//		end := (i + 1) * chunkSize
//		if end > length {
//			end = length
//		}
//		result[i] = data[start:end:end]
//	}
//
//	return result
//}
