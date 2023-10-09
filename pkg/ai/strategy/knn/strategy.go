package knn

type Settings struct {
	ExtraDataSourceDir      string
	ExtraDataXFieldsCount   int
	ExtraDataYFieldsCount   int
	ExtractDataDestFilePath string
	TestingSetPath          string
	TrainingSetPath         string
	TrainingSetRatio        float64
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
