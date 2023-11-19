package cmd

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/jan-hajek/ai/pkg/ai/strategy"
	"github.com/jan-hajek/ai/pkg/ai/strategy/knn"
	"github.com/jan-hajek/ai/pkg/ai/strategy/random"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ai",
	Short: "Description",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.AddCommand(allCmd)
	rootCmd.AddCommand(imageSplitCmd)
	rootCmd.AddCommand(extractDataCmd)
	rootCmd.AddCommand(extractShapesCmd)
	rootCmd.AddCommand(prepareSetsCmd)
	rootCmd.AddCommand(trainAlgorithmCmd)
	rootCmd.AddCommand(testAlgorithmCmd)
}

func getStrategyByName(name string) (strategy.Strategy, error) {
	rootDir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	switch name {
	case "random":
		return random.New(), nil
	case "knn":
		return knn.New(
			knn.Settings{
				ExtraDataSourceDir:      path.Join(rootDir, "data", "migrateddata", "all"),
				ExtraDataXFieldsCount:   3,
				ExtraDataYFieldsCount:   3,
				ExtractDataDestFilePath: path.Join(rootDir, "data", "knn", "imageData.csv"),
				TrainingSetPath:         path.Join(rootDir, "data", "knn", "trainingData.csv"),
				TrainingSetRatio:        0.8,
				TrainModelKList:         []int{3, 4, 5, 6, 7, 8},
				TrainModelBucketsCount:  5,
				TestingSetPath:          path.Join(rootDir, "data", "knn", "testingData.csv"),
			},
		), nil
	}

	return nil, errors.Errorf("unknown strategy: %s, allowed values are: [%s]", name, strings.Join([]string{"random", "knn"}, ","))
}

func printErrorWithStack(err error) {
	if err == nil {
		return
	}

	fmt.Printf("%+v\n", err)
}
