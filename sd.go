package sd

import (
	"errors"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"io"
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
	OutputsImageType OutputsImageType
}

type StableDiffusionModel struct {
	ctx        *CSDCtx
	options    *StableDiffusionOptions
	isAutoLoad bool
	dylibPath  string
}

var DefaultStableDiffusionOptions = StableDiffusionOptions{
	Threads:               -1, // auto
	VaeDecodeOnly:         true,
	FreeParamsImmediately: true,
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
	OutputsImageType: PNG,
}

func NewStableDiffusionAutoModel(options StableDiffusionOptions) (*StableDiffusionModel, error) {
	file, err := dumpSDLibrary()
	if err != nil {
		return nil, err
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
	ctx := sd.NewStableDiffusionCtx(options.Threads, options.VaeDecodeOnly, options.FreeParamsImmediately, options.RngType)
	return &StableDiffusionModel{
		dylibPath: dylibPath,
		ctx:       ctx,
		options:   &options,
	}, nil
}

func (sd *StableDiffusionModel) LoadFromFile(path string) error {
	sd.ctx.StableDiffusionLoadFromFile(path, sd.options.Schedule)
	return nil
}

func (sd *StableDiffusionModel) Predict(prompt string, writer io.Writer) error {
	outputsBytes := sd.ctx.StableDiffusionTextToImage(prompt, sd.options.NegativePrompt, sd.options.CfgScale, sd.options.Width, sd.options.Height, sd.options.SampleMethod, sd.options.SampleSteps, sd.options.Seed)
	outputsImage := bytesToImage(outputsBytes, sd.options.Width, sd.options.Height)
	return imageToWriter(outputsImage, sd.options.OutputsImageType, writer)
}

func (sd *StableDiffusionModel) ImagePredict(reader io.Reader, prompt string, writer io.Writer) error {
	decode, _, err := image.Decode(reader)
	if err != nil {
		return err
	}
	bytesImg := imageToBytes(decode)
	outputsBytes := sd.ctx.StableDiffusionImageToImage(bytesImg, prompt, sd.options.NegativePrompt, sd.options.CfgScale, sd.options.Width, sd.options.Height, sd.options.SampleMethod, sd.options.SampleSteps, sd.options.Strength, sd.options.Seed)
	outputsImage := bytesToImage(outputsBytes, sd.options.Width, sd.options.Height)
	return imageToWriter(outputsImage, sd.options.OutputsImageType, writer)
}

func (sd *StableDiffusionModel) Close() error {
	sd.ctx.Close()
	if sd.ctx.csd.libSd != 0 {
		err := closeLibrary(sd.ctx.csd.libSd)
		if err != nil {
			return err
		}
	}
	sd.ctx.csd.libSd = 0
	if sd.isAutoLoad {
		err := os.Remove(sd.dylibPath)
		return err
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
