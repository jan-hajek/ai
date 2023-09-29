package main

import (
	"context"
	"fmt"
	"os"
	"path"

	"github.com/jelito/ai/pkg/ai/imagemigrator"
)

func main() {
	ctx := context.Background()

	rootDir, err := os.Getwd()
	if err != nil {
		fmt.Println("error:")
		fmt.Println(err.Error())
		return
	}

	annotationFileName := "_annotations.coco.json"

	dirs := []struct {
		sourceDir string
		targetDir string
	}{
		{
			sourceDir: "test",
			targetDir: "test",
		},
		{
			sourceDir: "train",
			targetDir: "train",
		},
		{
			sourceDir: "valid",
			targetDir: "valid",
		},
	}

	for _, dir := range dirs {
		fmt.Println("processing dir: " + dir.sourceDir)

		sourceDir := path.Join(rootDir, "data", "origdata", dir.sourceDir)
		targetDir := path.Join(rootDir, "data", "migrateddata", dir.targetDir)

		migrator := imagemigrator.NewImageMigrator(annotationFileName, sourceDir, targetDir)

		err = migrator.Migrate(ctx)
		if err != nil {
			fmt.Println("error:")
			fmt.Println(err.Error())
			return
		}

		fmt.Println("-------------------")
	}

	fmt.Println("done without error")
}
