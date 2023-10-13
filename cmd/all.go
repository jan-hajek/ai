package cmd

import (
	"github.com/spf13/cobra"
)

var allCmd = &cobra.Command{
	Use:   "all",
	Short: "Prepares the data and trains a model",
	Args: func(cmd *cobra.Command, args []string) error {
		if err := cobra.ExactArgs(1)(cmd, args); err != nil {
			return err
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := getStrategyByName(args[0])
		if err != nil {
			printErrorWithStack(err)
			return err
		}

		err = s.DataExtraction(cmd.Context())
		if err != nil {
			printErrorWithStack(err)
			return err
		}

		err = s.PrepareSets(cmd.Context())
		if err != nil {
			printErrorWithStack(err)
			return err
		}

		err = s.TrainModel(cmd.Context())
		if err != nil {
			printErrorWithStack(err)
			return err
		}

		err = s.TestModel(cmd.Context())
		if err != nil {
			printErrorWithStack(err)
			return err
		}

		return nil
	},
}
