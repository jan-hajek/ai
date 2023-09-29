package imagemigrator

import (
	"context"
	"fmt"
	"image"
	"image/jpeg"
	"os"
	"path"

	"github.com/pkg/errors"
)

type extractImagesInput struct {
	imageName string
	coords    []coords
}

type coords struct {
	x, y          int
	width, height int
	number        int
}

type SubImager interface {
	SubImage(r image.Rectangle) image.Image
}

func (i *ImageMigrator) extractImages(ctx context.Context, input extractImagesInput) (paths []string, err error) {
	img, err := openImage(path.Join(i.sourceDir, input.imageName))
	if err != nil {
		return nil, err
	}

	for index, c := range input.coords {
		name := fmt.Sprintf("%s-%d-%d.jpg", input.imageName, index, c.number)
		destImagePath := path.Join(i.targetDir, name)

		cropSize := image.Rect(0, 0, c.width, c.height)
		cropSize = cropSize.Add(image.Point{X: c.x, Y: c.y})
		croppedImage := img.(SubImager).SubImage(cropSize)

		err = saveImage(croppedImage, destImagePath)
		paths = append(paths, destImagePath)
		if err != nil {
			return nil, err
		}
	}

	return paths, nil
}

func openImage(path string) (image.Image, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer f.Close()
	img, format, err := image.Decode(f)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	if format != "jpeg" {
		return nil, errors.New("it is not jpeg")
	}

	return img, nil
}

func saveImage(img image.Image, name string) error {
	f, err := os.Create(name)
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
