package imagemigrator

import (
	"context"
	"fmt"
)

type ImageMigrator struct {
	origDataDir        string
	annotationFileName string
	migratedDataDir    string
}

func NewImageMigrator(
	annotationFileName string,
	origDataDir string,
	migratedDataDir string,
) *ImageMigrator {
	return &ImageMigrator{
		origDataDir:        origDataDir,
		annotationFileName: annotationFileName,
		migratedDataDir:    migratedDataDir,
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
