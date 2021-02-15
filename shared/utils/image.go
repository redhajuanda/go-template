package utils

import (
	"image"

	"github.com/nfnt/resize"
)

// ResizeImage resizes image based on given size
func ResizeImage(img image.Image, x, y uint) image.Image {

	var images []image.Image
	images = append(images, img) // append original image
	newImage := resize.Resize(x, y, img, resize.Lanczos2)

	return newImage
}
