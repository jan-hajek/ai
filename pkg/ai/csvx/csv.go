package csvx

import (
	"encoding/csv"
	"os"
	"strconv"

	"github.com/pkg/errors"
)

type Row = []string

func OpenFileForWriting(path string) (writeFn func(Row) error, closeFn func(), _ error) {
	file, err := os.Create(path)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}
	w := csv.NewWriter(file)

	return func(row Row) error {
			return errors.WithStack(w.Write(row))
		}, func() {
			w.Flush()
			file.Close()
		}, nil
}

func ReadFromFile(path string) ([]Row, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer file.Close()

	r := csv.NewReader(file)

	rows, err := r.ReadAll()

	return rows, errors.WithStack(err)
}

func ConvertFloatsToStrings(floats []float64) Row {
	result := make(Row, len(floats))
	for i, value := range floats {
		result[i] = strconv.FormatFloat(value, 'f', -1, 64)
	}
	return result
}

func ConvertIntsToStrings(ints []int) Row {
	result := make(Row, len(ints))
	for i, value := range ints {
		result[i] = strconv.Itoa(value)
	}
	return result
}
