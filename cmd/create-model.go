package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var createModelCmd = &cobra.Command{
	Use:   "create-model",
	Short: "Prepares the data and trains a model",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create-model called")
	},
}
