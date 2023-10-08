package cmd

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/jelito/ai/pkg/ai/strategy"
	"github.com/jelito/ai/pkg/ai/strategy/knn"
	"github.com/jelito/ai/pkg/ai/strategy/random"
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
	rootCmd.AddCommand(imageSplitCmd)
	rootCmd.AddCommand(createModelCmd)
	rootCmd.AddCommand(extractDataCmd)
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
			path.Join(rootDir, "data", "migrateddata", "test"),
			path.Join(rootDir, "data", "knn", "imageData.csv"),
		), nil
	}

	return nil, fmt.Errorf("unknown strategy: %s, allowed values are: [%s]", name, strings.Join([]string{"random", "knn"}, ","))
}
