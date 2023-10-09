package cmd

import (
	"github.com/spf13/cobra"
)

var testAlgorithmCmd = &cobra.Command{
	Use:   "test-model",
	Short: "Tests a model",
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

		err = s.TestModel(cmd.Context())
		if err != nil {
			printErrorWithStack(err)
			return err
		}

		return nil
	},
}
