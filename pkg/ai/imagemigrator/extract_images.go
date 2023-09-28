package imagemigrator

import "context"

type extractImagesInput struct {
	imagePath string
	coords    []coords
}

type coords struct {
	x1, x2 int
	x3, x4 int
	number int
}

func (i *ImageMigrator) extractImages(ctx context.Context, input extractImagesInput) error {
	return nil
}
