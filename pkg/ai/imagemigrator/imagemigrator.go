package imagemigrator

import (
	"context"
)

type ImageMigrator struct {
	origDataDir        string
	annotationFileName string
	migratedDataDir    string
}

func NewImageMigrator(
	origDataDir string,
	annotationFileName string,
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

	for _, input := range inputs {
		err := i.extractImages(ctx, input)
		if err != nil {
			return err
		}
	}

	return nil
}