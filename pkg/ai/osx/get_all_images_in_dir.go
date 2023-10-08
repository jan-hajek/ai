package osx

import (
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

type ImageFile struct {
	Path   string
	Number int
}

func GetAllImagesInDir(dir string) (paths []ImageFile, _ error) {
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, errors.Errorf("error reading dir: %v", err)
	}

	for _, file := range files {
		if !file.IsDir() {
			fileBase := strings.TrimSuffix(file.Name(), filepath.Ext(file.Name()))
			number, err := strconv.Atoi(fileBase[len(fileBase)-1:])
			if err != nil {
				return nil, errors.Errorf("error parsing number from file name: %s", fileBase)
			}
			paths = append(paths, ImageFile{
				Path:   path.Join(dir, file.Name()),
				Number: number,
			})
		}
	}

	return paths, nil
}
