package utils

import (
	"unsafe"

	"github.com/seasonjs/stable-diffusion/internal/binding"
	"github.com/seasonjs/stable-diffusion/pkg/types"

)

func ImagePtrToImageSlice(imgPtr uintptr, size int) []binding.Image{
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
		Data:    unsafe.Slice((*byte)(dataPtr), cImage.Channel * cImage.Width*cImage.Height),
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

// func CImage(goImage types.Image, size int) uintptr {
// 	cImage := binding.Image{
// 		Channel: goImage.Channel,
// 		Width:   goImage.Width,
// 		Height:  goImage.Height,
// 		Data:    (*byte)(unsafe.SliceData(goImage.Data)),
// 	}
// 	return uintptr(unsafe.Pointer(&cImage))
// }
