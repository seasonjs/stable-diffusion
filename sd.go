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
)

type OutputsImageType string

const (
	PNG  OutputsImageType = "PNG"
	JPEG                  = "JPEG"
)

type StableDiffusionOptions struct {
	Threads               int
	VaeDecodeOnly         bool
	FreeParamsImmediately bool
	LoraModelDir          string
	RngType               RNGType

	Schedule Schedule

	NegativePrompt   string
	CfgScale         float32
	Width            int
	Height           int
	SampleMethod     SampleMethod
	SampleSteps      int
	Strength         float32
	Seed             int64
	BatchCount       int
	GpuEnable        bool
	OutputsImageType OutputsImageType
}

type StableDiffusionModel struct {
	ctx        *StableDiffusionCtx
	options    *StableDiffusionOptions
	params     *StableDiffusionFullParams
	csd        CStableDiffusion
	isAutoLoad bool
	dylibPath  string
}

var DefaultStableDiffusionOptions = StableDiffusionOptions{
	Threads:               -1, // auto
	VaeDecodeOnly:         true,
	FreeParamsImmediately: true,
	LoraModelDir:          "",
	RngType:               CUDA_RNG,

	Schedule: DEFAULT,

	NegativePrompt:   "out of frame, lowers, text, error, cropped, worst quality, low quality, jpeg artifacts, ugly, duplicate, morbid, mutilated, out of frame, extra fingers, mutated hands, poorly drawn hands, poorly drawn face, mutation, deformed, blurry, dehydrated, bad anatomy, bad proportions, extra limbs, cloned face, disfigured, gross proportions, malformed limbs, missing arms, missing legs, extra arms, extra legs, fused fingers, too many fingers, long neck, username, watermark, signature",
	CfgScale:         7.0,
	Width:            500,
	Height:           500,
	SampleMethod:     EULER_A,
	SampleSteps:      20,
	Strength:         0.4,
	Seed:             42,
	BatchCount:       1,
	OutputsImageType: PNG,
}

func (s *StableDiffusionOptions) toStableDiffusionFullParamsRef(c CStableDiffusion) *StableDiffusionFullParams {
	params := c.StableDiffusionFullDefaultParamsRef()
	if len(s.NegativePrompt) != 0 {
		c.StableDiffusionFullParamsSetNegativePrompt(params, s.NegativePrompt)
	}
	if s.CfgScale != 0 {
		c.StableDiffusionFullParamsSetCfgScale(params, s.CfgScale)
	}
	if s.Width != 0 {
		c.StableDiffusionFullParamsSetWidth(params, s.Width)
	}
	if s.Height != 0 {
		c.StableDiffusionFullParamsSetHeight(params, s.Height)
	}
	c.StableDiffusionFullParamsSetSampleMethod(params, s.SampleMethod)
	if s.SampleSteps != 0 {
		c.StableDiffusionFullParamsSetSampleSteps(params, s.SampleSteps)
	}
	if s.Strength != 0 {
		c.StableDiffusionFullParamsSetStrength(params, s.Strength)
	}
	if s.Seed != 0 {
		c.StableDiffusionFullParamsSetSeed(params, s.Seed)
	}
	//default batch count is 1 in c++
	if s.BatchCount != 0 && s.BatchCount != 1 {
		c.StableDiffusionFullParamsSetBatchCount(params, s.BatchCount)
	}

	return params
}

func NewStableDiffusionAutoModel(options StableDiffusionOptions) (*StableDiffusionModel, error) {
	file, err := dumpSDLibrary(options.GpuEnable)
	if err != nil {
		return nil, err
	}

	if options.GpuEnable {
		log.Printf("If you want to try offload your model to the GPU. " +
			"Please confirm the size of your GPU memory to prevent memory overflow." +
			"If the model is larger than GPU memory, please specify the layers to offload.")
	}
	dylibPath := file.Name()
	model, err := NewStableDiffusionModel(dylibPath, options)
	if err != nil {
		return nil, err
	}
	model.isAutoLoad = true
	return model, nil
}

func NewStableDiffusionModel(dylibPath string, options StableDiffusionOptions) (*StableDiffusionModel, error) {
	sd, err := NewCStableDiffusion(dylibPath)
	if err != nil {
		return nil, err
	}

	if options.BatchCount < 1 {
		options.BatchCount = 1
	}

	ctx := sd.StableDiffusionInit(options.Threads, options.VaeDecodeOnly, options.FreeParamsImmediately, options.LoraModelDir, options.RngType)
	params := options.toStableDiffusionFullParamsRef(sd)
	return &StableDiffusionModel{
		dylibPath: dylibPath,
		ctx:       ctx,
		options:   &options,
		params:    params,
		csd:       sd,
	}, nil
}

func (sd *StableDiffusionModel) LoadFromFile(path string) error {
	_, err := os.Stat(path)
	if err != nil {
		return errors.New("the system cannot find the model file specified")
	}
	sd.csd.StableDiffusionLoadFromFile(sd.ctx, path, sd.options.Schedule)
	return nil
}

func (sd *StableDiffusionModel) Predict(prompt string, writer []io.Writer) error {
	if len(writer) != sd.options.BatchCount {
		return errors.New("writer count not match batch count")
	}
	data := sd.csd.StableDiffusionPredictImage(
		sd.ctx,
		sd.params,
		prompt,
	)

	result := chunkBytes(data, sd.options.BatchCount)

	for i := 0; i < sd.options.BatchCount; i++ {
		outputsImage := bytesToImage(result[i], sd.options.Width, sd.options.Height)
		err := imageToWriter(outputsImage, sd.options.OutputsImageType, writer[i])
		if err != nil {
			return err
		}
	}

	return nil
}

func (sd *StableDiffusionModel) ImagePredict(reader io.Reader, prompt string, writer io.Writer) error {
	decode, _, err := image.Decode(reader)
	if err != nil {
		return err
	}
	bytesImg := imageToBytes(decode)
	outputsBytes := sd.csd.StableDiffusionImagePredictImage(
		sd.ctx,
		sd.params,
		bytesImg,
		prompt,
	)
	outputsImage := bytesToImage(outputsBytes, sd.options.Width, sd.options.Height)
	return imageToWriter(outputsImage, sd.options.OutputsImageType, writer)
}

func (sd *StableDiffusionModel) Close() error {
	if sd.ctx != nil {
		sd.csd.StableDiffusionFree(sd.ctx)
		sd.ctx = nil
	}

	if sd.params != nil {
		sd.csd.StableDiffusionFreeFullParams(sd.params)
		sd.params = nil

	}

	if sd.isAutoLoad {
		err := os.Remove(sd.dylibPath)
		if err != nil {
			return err
		}
	}
	return nil
}

func imageToBytes(decode image.Image) []byte {
	bounds := decode.Bounds()
	width := bounds.Max.X - bounds.Min.X
	height := bounds.Max.Y - bounds.Min.Y
	size := width * height * 3
	bytesImg := make([]byte, size)
	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			idx := (y*width + x) * 3
			r, g, b, _ := decode.At(x, y).RGBA()
			bytesImg[idx] = byte(r)
			bytesImg[idx+1] = byte(g)
			bytesImg[idx+2] = byte(b)
		}
	}
	return bytesImg
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

func chunkBytes(data []byte, chunks int) [][]byte {
	length := len(data)
	chunkSize := (length + chunks - 1) / chunks
	result := make([][]byte, chunks)

	for i := 0; i < chunks; i++ {
		start := i * chunkSize
		end := (i + 1) * chunkSize
		if end > length {
			end = length
		}
		result[i] = data[start:end:end]
	}

	return result
}
