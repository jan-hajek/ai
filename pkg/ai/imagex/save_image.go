package imagex

import (
	"image"
	"image/jpeg"
	"os"

	"github.com/pkg/errors"
)

func SaveJpegImage(img image.Image, path string) error {
	f, err := os.Create(path)
	if err != nil {
		return errors.WithStack(err)
	}
	defer f.Close()
	err = jpeg.Encode(f, img, nil)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}
