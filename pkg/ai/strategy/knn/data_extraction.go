package knn

import (
	"context"
	"encoding/csv"
	"os"
	"strconv"

	"github.com/jelito/ai/pkg/ai/imagedataextractor"
	"github.com/jelito/ai/pkg/ai/imagex"
	"github.com/jelito/ai/pkg/ai/osx"
	"github.com/pkg/errors"
)

func (b *KnnStrategy) DataExtraction(ctx context.Context) error {
	ide := imagedataextractor.NewImageDataExtractor(3, 3)

	files, err := osx.GetAllImagesInDir(b.sourceDataDir)
	if err != nil {
		return errors.WithStack(err)
	}

	file, err := os.Create(b.imageDataPath)
	defer file.Close()
	if err != nil {
		return errors.WithStack(err)
	}
	w := csv.NewWriter(file)
	defer w.Flush()
	for _, file := range files {
		img, err := imagex.OpenImage(file.Path)
		if err != nil {
			return errors.WithStack(err)
		}

		fieldsAvgs := ide.ExtractFields(ctx, img)
		if err != nil {
			return errors.WithStack(err)
		}

		columns := []string{
			strconv.Itoa(file.Number),
		}
		columns = append(columns, fieldsAvgsAsStringArray(fieldsAvgs)...)

		if err := w.Write(columns); err != nil {
			return errors.WithStack(err)
		}
	}

	return nil
}

func fieldsAvgsAsStringArray(fieldsAvgs []float64) []string {
	result := make([]string, len(fieldsAvgs))
	for i, avg := range fieldsAvgs {
		result[i] = strconv.FormatFloat(avg, 'f', -1, 64)
	}
	return result
}
