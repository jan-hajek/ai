package main

import (
	"context"
	"fmt"
	"os"
)

func main() {
	ctx := context.Background()

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("error:")
		fmt.Println(err.Error())
		return
	}

	_ = dir
	_ = ctx

	fmt.Println("done without error")
}
