package strategy

import "context"

type TrainFile struct {
	Path   string
	Number int
}

type Strategy interface {
	DataExtraction(ctx context.Context) error
	PrepareSets(ctx context.Context) error
	TrainModel(ctx context.Context) error
	TestModel(ctx context.Context) error
	RecognizeImage(path string) (number int, confidence float64, _ error)
}
