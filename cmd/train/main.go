package main

import (
	"context"
	"fmt"
	"os"
	"path"

	"github.com/jelito/ai/pkg/ai/strategy/knn"
	"github.com/jelito/ai/pkg/ai/trainer"
)

func main() {
	ctx := context.Background()

	rootDir, err := os.Getwd()
	if err != nil {
		fmt.Println("error:")
		fmt.Println(err.Error())
		return
	}

	t := trainer.NewTrainer(
		//path.Join(rootDir, "data", "migrateddata", "test"),
		path.Join(rootDir, "data", "honza"),
		knn.New(),
	)
	err = t.Train(ctx)
	if err != nil {
		fmt.Println("error:")
		fmt.Println(err.Error())
		return
	}
}
