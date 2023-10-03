package imagex

import (
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"

	"github.com/pkg/errors"
)

func OpenImage(path string) (image.Image, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer f.Close()
	img, _, err := image.Decode(f)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return img, nil
}
