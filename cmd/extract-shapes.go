package cmd

import (
	"github.com/jan-hajek/ai/pkg/ai/imagex"
	"github.com/spf13/cobra"
	"os"
	"path"
	"github.com/jan-hajek/ai/pkg/ai/imagedataextractor"
)

var extractShapesCmd = &cobra.Command{
	Use:   "extract-shapes",
	Short: "Extracts shapes",
	Args: func(cmd *cobra.Command, args []string) error {
		if err := cobra.ExactArgs(1)(cmd, args); err != nil {
			return err
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		dir, err := os.Getwd()
		if err != nil {
			return err
		}

		img, err := imagex.OpenImage(path.Join(dir, "pkg/ai/imagedataextractor/test_data", "test_shape.jpg"))
		if err != nil {
			return err
		}

		ise := imagedataextractor.NewImageShapeExtractor()
		ise.ExtractShapes(nil, img)

		return nil
	},
}
