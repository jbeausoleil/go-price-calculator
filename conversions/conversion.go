package conversions

import (
	"errors"
	"strconv"
)

// StringsToFloats converts a slice of strings to a slice of float64 values.
// It iterates through each string value in the input slice and converts it to a float64.
// If any string cannot be converted, it returns an error and a nil slice.
// Otherwise, it returns a new slice containing the converted float64 values.
func StringsToFloats(strings []string) ([]float64, error) {
	var floats []float64

	for _, stringVal := range strings {
		floatVal, err := strconv.ParseFloat(stringVal, 64)

		if err != nil {
			return nil, errors.New("failed to convert string to float")
		}

		floats = append(floats, floatVal)
	}
	return floats, nil
}
