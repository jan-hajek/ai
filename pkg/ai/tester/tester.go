package tester

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/jelito/ai/pkg/ai/strategy"
	"github.com/pkg/errors"
)

type Tester struct {
	dir      string
	strategy strategy.Strategy
}

func NewTester(dir string, strategy strategy.Strategy) *Tester {
	return &Tester{
		dir:      dir,
		strategy: strategy,
	}
}

type Results struct {
	Numbers map[int]*Result
}

func (r Results) GetSuccessRate() float64 {
	if len(r.Numbers) == 0 {
		return 0
	}
	var successRate float64
	for _, result := range r.Numbers {
		successRate += result.GetSuccessRate()
	}
	return successRate / float64(len(r.Numbers))
}

type Result struct {
	Successful int
	Count      int
}

func (r Result) GetSuccessRate() float64 {
	if r.Count == 0 {
		return 0
	}
	return float64(r.Successful) / float64(r.Count)
}

func (t *Tester) Test(ctx context.Context) (results Results, _ error) {
	testFiles, err := t.getAllFilesInDir()
	if err != nil {
		return results, err
	}

	results.Numbers = make(map[int]*Result)

	for _, testFile := range testFiles {
		number, confidence, err := t.strategy.TestFile(testFile.Path)
		if err != nil {
			return results, err
		}
		_ = confidence

		if _, ok := results.Numbers[testFile.Number]; !ok {
			results.Numbers[testFile.Number] = &Result{}
		}
		if number == testFile.Number {
			results.Numbers[testFile.Number].Successful++
		}
		results.Numbers[testFile.Number].Count++
	}

	return results, nil
}

type TestFile struct {
	Path   string
	Number int
}

func (t *Tester) getAllFilesInDir() (paths []TestFile, _ error) {
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
			paths = append(paths, TestFile{
				Path:   file.Name(),
				Number: number,
			})
		}
	}

	return paths, nil
}
