package knn

import (
	"context"
	"path"
	"strconv"

	"github.com/jelito/ai/pkg/ai/csvx"
	"github.com/jelito/ai/pkg/ai/imagedataextractor"
	"github.com/jelito/ai/pkg/ai/imagex"
	"github.com/jelito/ai/pkg/ai/osx"
)

func (b *KnnStrategy) DataExtraction(ctx context.Context) error {
	ide := imagedataextractor.NewImageDataExtractor(b.settings.ExtraDataXFieldsCount, b.settings.ExtraDataYFieldsCount)

	files, err := osx.GetAllImagesInDir(b.settings.ExtraDataSourceDir)
	if err != nil {
		return err
	}

	writer, closeFile, err := csvx.OpenFileForWriting(b.settings.ExtractDataDestFilePath)
	if err != nil {
		return err
	}
	defer closeFile()

	expectedNumberOfFields := b.settings.ExtraDataXFieldsCount * b.settings.ExtraDataYFieldsCount

	for _, file := range files {
		img, err := imagex.OpenImage(path.Join(b.settings.ExtraDataSourceDir, file.Name))
		if err != nil {
			return err
		}

		fieldsAvgs := ide.ExtractFields(ctx, img)
		if err != nil {
			return err
		}

		if len(fieldsAvgs) != expectedNumberOfFields {
			continue
		}

		row := csvx.Row{
			strconv.Itoa(file.Number),
			file.Name,
		}
		row = append(row, csvx.ConvertFloatsToStrings(fieldsAvgs)...)

		if err := writer(row); err != nil {
			return err
		}
	}

	return nil
}
