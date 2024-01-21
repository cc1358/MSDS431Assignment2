package main

import (
	"fmt"
	"math"
)

// Create Column to represent a column of float64s
type Column struct {
	Data   []float64
	Name   string
	Column int
}

// Create a simple dataframe
type DataFrame struct {
	Columns []Column
}

// Create LinearRegression function to fit a regression model and returns the coefficients
func LinearRegression(x, y []float64) (float64, float64) {
	meanX := mean(x)
	meanY := mean(y)

	var numerator, denominator float64
	for i := range x {
		numerator += (x[i] - meanX) * (y[i] - meanY)
		denominator += math.Pow(x[i]-meanX, 2)
	}

	beta1 := numerator / denominator
	beta0 := meanY - beta1*meanX

	return beta1, beta0
}

// Create mean function to calculate the mean of a slice of float64.
func mean(values []float64) float64 {
	var sum float64
	for _, v := range values {
		sum += v
	}
	return sum / float64(len(values))
}

// Use LinearRegression function to create another function fitAndPrintLinearRegression to fit a linear regression model and print the coefficients
func fitAndPrintLinearRegression(x, y []float64) {
	beta1, beta0 := LinearRegression(x, y)
	fmt.Printf("Beta1: %.4f, Beta0: %.4f\n", beta1, beta0)
}

// Obtain a valid column choice from user in app
func getUserColumnChoice() int {
	var columnChoice int
	fmt.Print("Choose a column (1, 2, 3, 4): ")
	_, err := fmt.Scan(&columnChoice)
	for err != nil || columnChoice < 1 || columnChoice > 4 {
		fmt.Print("Invalid choice. Please choose a valid column (1, 2, 3, 4): ")
		_, err = fmt.Scan(&columnChoice)
	}
	return columnChoice
}

// Create function to print message from input
func printMessage(message string) {
	fmt.Println("***************")
	fmt.Println(message)
	fmt.Println("***************")
}

func main() {

	// Define the Anscombe data using a data frame
	anscombe := DataFrame{
		Columns: []Column{
			{Data: []float64{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5}, Name: "x1"},
			{Data: []float64{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5}, Name: "x2"},
			{Data: []float64{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5}, Name: "x3"},
			{Data: []float64{8, 8, 8, 8, 8, 8, 8, 19, 8, 8, 8}, Name: "x4"},
			{Data: []float64{8.04, 6.95, 7.58, 8.81, 8.33, 9.96, 7.24, 4.26, 10.84, 4.82, 5.68}, Name: "y1"},
			{Data: []float64{9.14, 8.14, 8.74, 8.77, 9.26, 8.1, 6.13, 3.1, 9.13, 7.26, 4.74}, Name: "y2"},
			{Data: []float64{7.46, 6.77, 12.74, 7.11, 7.81, 8.84, 6.08, 5.39, 8.15, 6.42, 5.73}, Name: "y3"},
			{Data: []float64{6.58, 5.76, 7.71, 8.84, 8.47, 7.04, 5.25, 12.5, 5.56, 7.91, 6.89}, Name: "y4"},
		},
	}

	// Fit linear regression models by ordinary least squares
	columnChoice := getUserColumnChoice()

	switch columnChoice {
	case 1:
		fitAndPrintLinearRegression(anscombe.Columns[0].Data, anscombe.Columns[4].Data)
	case 2:
		fitAndPrintLinearRegression(anscombe.Columns[1].Data, anscombe.Columns[5].Data)
	case 3:
		fitAndPrintLinearRegression(anscombe.Columns[2].Data, anscombe.Columns[6].Data)
	case 4:
		fitAndPrintLinearRegression(anscombe.Columns[3].Data, anscombe.Columns[7].Data)
	default:
		printMessage("Invalid choice. Please choose a valid column.")
	}

}
