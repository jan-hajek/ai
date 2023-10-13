package knn

type Settings struct {
	ExtraDataSourceDir      string
	ExtraDataXFieldsCount   int
	ExtraDataYFieldsCount   int
	ExtractDataDestFilePath string
	TrainingSetPath         string
	TrainingSetRatio        float64
	TrainModelKList         []int
	TrainModelBucketsCount  int
	TestingSetPath          string
}
type KnnStrategy struct {
	settings Settings
}

func New(
	settings Settings,
) *KnnStrategy {
	return &KnnStrategy{
		settings: settings,
	}
}
