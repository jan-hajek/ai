package knn

type KnnStrategy struct {
	sourceDataDir string
	imageDataPath string
}

func New(
	sourceDataDir string,
	imageDataPath string,
) *KnnStrategy {
	return &KnnStrategy{
		sourceDataDir: sourceDataDir,
		imageDataPath: imageDataPath,
	}
}
