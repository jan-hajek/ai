package imagemigrator

import (
	"context"
	"fmt"
)

type ImageMigrator struct {
	annotationFileName string
	sourceDir          string
	targetDir          string
}

func NewImageMigrator(
	annotationFileName string,
	sourceDir string,
	targetDir string,
) *ImageMigrator {
	return &ImageMigrator{
		sourceDir:          sourceDir,
		annotationFileName: annotationFileName,
		targetDir:          targetDir,
	}
}

func (i *ImageMigrator) Migrate(ctx context.Context) error {
	inputs, err := i.readJson(ctx)
	if err != nil {
		return err
	}

	createdImages := 0
	for _, input := range inputs {
		imagesNames, err := i.extractImages(ctx, input)
		if err != nil {
			return err
		}
		createdImages += len(imagesNames)
	}

	fmt.Println("Created images: ", createdImages)

	return nil
}
