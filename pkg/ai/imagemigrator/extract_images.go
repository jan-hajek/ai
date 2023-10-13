package imagemigrator

import (
	"context"
	"fmt"
	"image"
	"path"

	"github.com/jan-hajek/ai/pkg/ai/imagex"
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
	img, err := imagex.OpenImage(path.Join(i.sourceDir, input.imageName))
	if err != nil {
		return nil, err
	}

	for index, c := range input.coords {
		name := fmt.Sprintf("%s-%d-%d.jpg", input.imageName, index, c.number)
		destImagePath := path.Join(i.targetDir, name)

		cropSize := image.Rect(0, 0, c.width, c.height)
		cropSize = cropSize.Add(image.Point{X: c.x, Y: c.y})
		croppedImage := img.(SubImager).SubImage(cropSize)

		err = imagex.SaveJpegImage(croppedImage, destImagePath)
		paths = append(paths, destImagePath)
		if err != nil {
			return nil, err
		}
	}

	return paths, nil
}
