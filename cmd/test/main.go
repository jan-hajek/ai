package main

import (
	"context"
	"fmt"
	"os"
	"path"

	"github.com/jelito/ai/pkg/ai/strategy/random"
	"github.com/jelito/ai/pkg/ai/tester"
)

func main() {
	ctx := context.Background()

	rootDir, err := os.Getwd()
	if err != nil {
		fmt.Println("error:")
		fmt.Println(err.Error())
		return
	}

	t := tester.NewTester(
		path.Join(rootDir, "data", "migrateddata", "test"),
		random.New(),
	)
	results, err := t.Test(ctx)
	if err != nil {
		fmt.Println("error:")
		fmt.Println(err.Error())
		return
	}
	fmt.Println("done without error")
	fmt.Println("results: ")
	fmt.Println("success rate: " + fmt.Sprintf("%f", results.GetSuccessRate()))
	fmt.Println("-------")
	for i := 0; i < 10; i++ {
		successRate := 0.0
		if _, ok := results.Numbers[i]; ok {
			successRate = results.Numbers[i].GetSuccessRate()
		}
		fmt.Println(fmt.Sprintf("%d - %f", i, successRate))
	}
}
