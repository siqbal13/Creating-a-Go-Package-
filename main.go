// Example execution and usage of trimmedmean function 
// We have used here the random number generator to get the 100 numbers generated randomly 
// to calculate the trimmedmean and then the numbers are saved in a separate file to use them 
// in an R language program to verify the results.
// You may see the verification of the answers by running both programs, as for Go language program 
// and the R program, to get consistent results.
package main

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"github.com/dani67894/TrimmedMean/trimmedmean"
)

// saveDataToCSV saves a slice of float64 data to a CSV file with the specified filename.
func saveDataToCSV(data []float64, filename string) error {
	// Create the file
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Create a CSV writer
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write each data point to the CSV file
	for _, value := range data {
		err := writer.Write([]string{strconv.FormatFloat(value, 'f', -1, 64)})
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	// Generate a sample of 100 random integers between 0 and 999
	intData := make([]float64, 100)
	for i := 0; i < 100; i++ {
		intData[i] = float64(rand.Intn(1000))
	}

	// Generate a sample of 100 random floats between 0 and 1000
	floatData := make([]float64, 100)
	for i := 0; i < 100; i++ {
		floatData[i] = rand.Float64() * 1000
	}

	// Save the integer data to a CSV file
	if err := saveDataToCSV(intData, "int_data.csv"); err != nil {
		fmt.Println("Error saving integer data:", err)
		return
	}

	// Save the float data to a CSV file
	if err := saveDataToCSV(floatData, "float_data.csv"); err != nil {
		fmt.Println("Error saving float data:", err)
		return
	}

	// Compute the trimmed mean for the integer data with symmetric 0.05 trim
	intSymmetricTrimmedMean, err := trimmedmean.TrimmedMean(intData, 0.05)
	if err != nil {
		fmt.Println("Error computing symmetric trimmed mean for integers:", err)
		return
	}

	// Compute the trimmed mean for the float data with symmetric 0.05 trim
	floatSymmetricTrimmedMean, err := trimmedmean.TrimmedMean(floatData, 0.05)
	if err != nil {
		fmt.Println("Error computing symmetric trimmed mean for floats:", err)
		return
	}

	// Compute the trimmed mean for the integer data with asymmetric trimming
	// Using 0.05 for lower trim and 0.1 for upper trim
	intAsymmetricTrimmedMean, err := trimmedmean.TrimmedMean(intData, 0.05, 0.1)
	if err != nil {
		fmt.Println("Error computing asymmetric trimmed mean for integers:", err)
		return
	}

	// Compute the trimmed mean for the float data with asymmetric trimming
	// Using 0.05 for lower trim and 0.1 for upper trim
	floatAsymmetricTrimmedMean, err := trimmedmean.TrimmedMean(floatData, 0.05, 0.1)
	if err != nil {
		fmt.Println("Error computing asymmetric trimmed mean for floats:", err)
		return
	}

	// Print the trimmed means
	fmt.Println("Symmetric Trimmed Mean (0.05 trim) for integers:", intSymmetricTrimmedMean)
	fmt.Println("Symmetric Trimmed Mean (0.05 trim) for floats:", floatSymmetricTrimmedMean)
	fmt.Println("Asymmetric Trimmed Mean (0.05 lower, 0.1 upper trim) for integers:", intAsymmetricTrimmedMean)
	fmt.Println("Asymmetric Trimmed Mean (0.05 lower, 0.1 upper trim) for floats:", floatAsymmetricTrimmedMean)
}
