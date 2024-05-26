// Package trimmedmean provides functionality to calculate trimmed means of slices.
package trimmedmean

import (
	"errors"
	"sort"
)

// TrimmedMean computes the trimmed mean of a slice of floats.
// The function can handle both symmetric and asymmetric trimming based on the provided arguments.
// If only one degreeOfTrimming argument is provided, symmetric trimming is assumed.
// If two arguments are provided, the first is the lower trimming proportion and the second is the upper trimming proportion.
func TrimmedMean(data []float64, degreeOfTrimming ...float64) (float64, error) {
	// Check for empty data
	if len(data) == 0 {
		return 0, errors.New("the slice is empty")
	}

	// Sort the data to prepare for trimming
	sort.Float64s(data)
	n := len(data)

	var lowerTrim, upperTrim float64
	if len(degreeOfTrimming) == 1 {
		// Symmetric trimming
		lowerTrim = degreeOfTrimming[0]
		upperTrim = degreeOfTrimming[0]
	} else if len(degreeOfTrimming) == 2 {
		// Asymmetric trimming
		lowerTrim = degreeOfTrimming[0]
		upperTrim = degreeOfTrimming[1]
	} else {
		return 0, errors.New("invalid number of degreeOfTrimming arguments")
	}

	// Validate the trimming proportions
	if lowerTrim < 0 || upperTrim < 0 || lowerTrim >= 0.5 || upperTrim >= 0.5 {
		return 0, errors.New("invalid trimming proportions: must be in the range [0, 0.5)")
	}

	// Calculate the number of elements to trim
	lowerIndex := int(float64(n) * lowerTrim)
	upperIndex := n - int(float64(n) * upperTrim)

	// Calculate the trimmed mean by summing the untrimmed elements
	sum := 0.0
	count := 0
	for i := lowerIndex; i < upperIndex; i++ {
		sum += data[i]
		count++
	}

	// Check if any elements are left after trimming
	if count == 0 {
		return 0, errors.New("all elements have been trimmed, no data left to compute the mean")
	}

	// Return the trimmed mean
	return sum / float64(count), nil
}
