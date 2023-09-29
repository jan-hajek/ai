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

	dirs := []string{
		"test",
		"train",
		"valid",
	}

	for _, dir := range dirs {
		fmt.Println("processing dir: " + dir)

		origDataDir := path.Join(rootDir, "data", "origdata", dir)
		migratedDataDir := path.Join(rootDir, "data", "migrateddata", dir)

		migrator := imagemigrator.NewImageMigrator(annotationFileName, origDataDir, migratedDataDir)

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
