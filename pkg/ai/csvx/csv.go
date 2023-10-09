package csvx

import (
	"encoding/csv"
	"os"
	"strconv"

	"github.com/pkg/errors"
)

func OpenFileForWriting(path string) (writeFn func([]string) error, closeFn func(), _ error) {
	file, err := os.Create(path)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}
	w := csv.NewWriter(file)

	return func(row []string) error {
			return errors.WithStack(w.Write(row))
		}, func() {
			w.Flush()
			file.Close()
		}, nil
}

func ReadFromFile(path string) ([][]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer file.Close()

	r := csv.NewReader(file)

	rows, err := r.ReadAll()

	return rows, errors.WithStack(err)
}

func ConvertFloatsToStrings(floats []float64) []string {
	result := make([]string, len(floats))
	for i, value := range floats {
		result[i] = strconv.FormatFloat(value, 'f', -1, 64)
	}
	return result
}

func ConvertIntsToStrings(ints []int) []string {
	result := make([]string, len(ints))
	for i, value := range ints {
		result[i] = strconv.Itoa(value)
	}
	return result
}
