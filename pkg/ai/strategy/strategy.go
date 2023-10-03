package strategy

import "context"

type TrainFile struct {
	Path   string
	Number int
}

type Strategy interface {
	TestFile(path string) (number int, confidence float64, _ error)
	TrainFiles(context.Context, []TrainFile) error
}
