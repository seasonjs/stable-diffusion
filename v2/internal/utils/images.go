package utils

import (
	"errors"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"io"
	"unsafe"

	"github.com/seasonjs/stable-diffusion/v2/internal/binding"
	"github.com/seasonjs/stable-diffusion/v2/pkg/types"
)

func ImagePtrToImageSlice(imgPtr uintptr, size int) []binding.Image {
	// We take the address and then dereference it to trick go vet from creating a possible misuse of unsafe.Pointer
	ptr := *(*unsafe.Pointer)(unsafe.Pointer(&imgPtr))
	if ptr == nil {
		return nil
	}
	img := (*binding.Image)(ptr)
	imgSlice := unsafe.Slice(img, size)
	return imgSlice
}

func GoImage(cImage binding.Image) types.Image {
	dataPtr := *(*unsafe.Pointer)(unsafe.Pointer(&cImage.Data))
	gImg := types.Image{
		Channel: cImage.Channel,
		Width:   cImage.Width,
		Height:  cImage.Height,
		Data:    unsafe.Slice((*byte)(dataPtr), cImage.Channel*cImage.Width*cImage.Height),
	}
	return gImg
}

func GoImageSlice(imgPtr uintptr, size int) []types.Image {
	cImgSlice := ImagePtrToImageSlice(imgPtr, size)
	goImages := make([]types.Image, 0, size)

	for _, image := range cImgSlice {
		goImages = append(goImages, GoImage(image))
	}
	return goImages
}

func CImage(goImage types.Image) binding.Image {
	data := unsafe.SliceData(goImage.Data)
	return binding.Image{
		Width:   goImage.Width,
		Height:  goImage.Height,
		Channel: goImage.Channel,
		Data:    uintptr(unsafe.Pointer(&data)),
	}
}

func FreeImage(image binding.Image) {
	if image.Data != 0 {
		Free(image.Data)
	}
}

func FreeImageSlice(imagePtr uintptr, size int) {
	cImgSlice := ImagePtrToImageSlice(imagePtr, size)
	for _, image := range cImgSlice {
		FreeImage(image)
	}
	if imagePtr != 0 {
		Free(imagePtr)
	}
}

func ImageToBytes(decode image.Image) types.Image {
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
	return types.Image{
		Width:   uint32(width),
		Height:  uint32(height),
		Data:    bytesImg,
		Channel: 3,
	}
}

func GoImageToImageInterface(goImage types.Image) image.Image {
	width := int(goImage.Width)
	height := int(goImage.Height)
	byteData := goImage.Data

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

func ImageToWriter(image image.Image, imageType types.ImageType, writer io.Writer) error {
	switch imageType {
	case types.PNG:
		err := png.Encode(writer, image)
		if err != nil {
			return err
		}
	case types.JPEG:
		err := jpeg.Encode(writer, image, &jpeg.Options{Quality: 100})
		if err != nil {
			return err
		}
	default:
		return errors.New("unknown image type")
	}
	return nil
}

func DecodeToCImage(reader io.Reader) (binding.Image, error) {
	decode, _, err := image.Decode(reader)
	if err != nil {
		return binding.Image{}, err
	}
	initImage := ImageToBytes(decode)
	initImagePtr := CImage(initImage)
	return initImagePtr, nil
}

func CImageEncode(image binding.Image, imageType types.ImageType, writer io.Writer) error {
	goImage := GoImage(image)
	return GoImageEncode(goImage, imageType, writer)
}

func GoImageEncode(image types.Image, imageType types.ImageType, writer io.Writer) error {
	imgInteface := GoImageToImageInterface(image)
	if err := ImageToWriter(imgInteface, imageType, writer); err != nil {
		return err
	}
	return nil
}
