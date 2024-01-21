package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"testing"
	"time"
)

// Enable logging
var logFile *os.File
var start time.Time

func init() {
	// Create or open a log file
	var err error
	logFile, err = os.Create("program_log.log")
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(logFile)
	start = time.Now()
}

// First Test
func TestLinearRegression(t *testing.T) {
	// Test with synthetic data
	x := []float64{1, 2, 3, 4, 5}
	y := []float64{2, 4, 5, 4, 5}

	beta1, beta0 := LinearRegression(x, y)

	// Verify the expected coefficients with a tolerance for floating-point comparisons
	epsilon := 1e-6
	expectedBeta1 := 0.7
	expectedBeta0 := 1.5

	if math.Abs(beta1-expectedBeta1) > epsilon || math.Abs(beta0-expectedBeta0) > epsilon {
		t.Errorf("LinearRegression test failed. Expected Beta1: %.4f, Beta0: %.4f. Got Beta1: %.4f, Beta0: %.4f",
			expectedBeta1, expectedBeta0, beta1, beta0)
	}
	// Log information
	log.Printf("LinearRegression test passed. Beta1: %.4f, Beta0: %.4f", beta1, beta0)
}

// Second Test
func TestMean(t *testing.T) {
	// Test with synthetic data
	values := []float64{2, 4, 5, 4, 5}

	meanValue := mean(values)

	// Verify the expected mean with a tolerance for floating-point comparisons
	epsilon := 1e-6
	expectedMean := 4

	if math.Abs(meanValue-float64(expectedMean)) > epsilon {
		t.Errorf("Mean test failed. Expected Mean: %.4f. Got Mean: %.4f", float64(expectedMean), meanValue)
	}
	// Log information
	log.Printf("Mean test passed. Mean: %.4f", meanValue)
}

// BenchmarkLinearRegression benchmarks the LinearRegression function
func BenchmarkLinearRegression(b *testing.B) {
	x := []float64{1, 2, 3, 4, 5}
	y := []float64{2, 4, 5, 4, 5}
	for i := 0; i < b.N; i++ {
		LinearRegression(x, y)
	}
}

// BenchmarkMean benchmarks the mean function
func BenchmarkMean(b *testing.B) {
	values := []float64{2, 4, 5, 4, 5}
	for i := 0; i < b.N; i++ {
		mean(values)
	}
}

// closemain is not a standard function for testing; it's usually used for cleanup in a real program
func closemain() {
	// Close the log file
	logFile.Close()

	fmt.Printf("Process finished --- %s seconds ---", time.Since(start))
}
