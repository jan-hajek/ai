package knn

import (
	"context"
	"math/rand"
	"time"

	"github.com/jan-hajek/ai/pkg/ai/csvx"
)

func (b *KnnStrategy) PrepareSets(_ context.Context) error {
	rows, err := csvx.ReadFromFile(b.settings.ExtractDataDestFilePath)
	if err != nil {
		return err
	}

	trainingWriter, closeTrainingFile, err := csvx.OpenFileForWriting(b.settings.TrainingSetPath)
	if err != nil {
		return err
	}
	defer closeTrainingFile()

	testingWriter, closeTestingFile, err := csvx.OpenFileForWriting(b.settings.TestingSetPath)
	if err != nil {
		return err
	}
	defer closeTestingFile()

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for _, row := range rows {
		if r.Float64() < b.settings.TrainingSetRatio {
			err := trainingWriter(row)
			if err != nil {
				return err
			}
		} else {
			err := testingWriter(row)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
