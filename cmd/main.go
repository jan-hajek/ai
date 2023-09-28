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

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("error:")
		fmt.Println(err.Error())
		return
	}

	origDataDir := path.Join(dir, "data", "origdata")
	migratedDataDir := path.Join(dir, "data", "origdata")
	annotationFileName := "_annotations.coco.json"

	migrator := imagemigrator.NewImageMigrator(annotationFileName, origDataDir, migratedDataDir)

	err = migrator.Migrate(ctx)
	if err != nil {
		fmt.Println("error:")
		fmt.Println(err.Error())
		return
	}

	fmt.Println("done without error")
}
