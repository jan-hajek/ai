package trainer

import (
	"context"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/jelito/ai/pkg/ai/strategy"
	"github.com/pkg/errors"
)

type Trainer struct {
	dir      string
	strategy strategy.Strategy
}

func NewTrainer(dir string, strategy strategy.Strategy) *Trainer {
	return &Trainer{
		dir:      dir,
		strategy: strategy,
	}
}

func (t *Trainer) Train(ctx context.Context) error {
	trainFiles, err := t.getAllFilesInDir()
	if err != nil {
		return err
	}

	err = t.strategy.TrainFiles(ctx, trainFiles)
	if err != nil {
		return err
	}

	return nil
}

func (t *Trainer) getAllFilesInDir() (paths []strategy.TrainFile, _ error) {
	files, err := os.ReadDir(t.dir)
	if err != nil {
		return nil, errors.Errorf("error reading dir: %v", err)
	}

	for _, file := range files {
		if !file.IsDir() {
			fileBase := strings.TrimSuffix(file.Name(), filepath.Ext(file.Name()))
			number, err := strconv.Atoi(fileBase[len(fileBase)-1:])
			if err != nil {
				fmt.Println(file.Name())
				fmt.Println(fileBase)
				return nil, errors.Errorf("error parsing number from file name: %s", fileBase)
			}
			paths = append(paths, strategy.TrainFile{
				Path:   path.Join(t.dir, file.Name()),
				Number: number,
			})
		}
	}

	return paths, nil
}
