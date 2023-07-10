package main

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

func Dataset() *mat.Dense {
	// Create data for the data frame
	data := []float64{
		1.0, 2.0, 3.0,
		4.0, 5.0, 6.0,
		7.0, 8.0, 9.0,
	}

	// Define the dimensions of the data frame
	rows, cols := 3, 3

	// Create a dense data frame
	df := mat.NewDense(rows, cols, data)

	// Print the data frame
	fmt.Println(df)

	// Access individual elements of the data frame
	val := df.At(1, 2)
	fmt.Println("Value at row 1, column 2:", val)

	// Modify individual elements of the data frame
	df.Set(0, 1, 10.0)
	fmt.Println("Modified data frame:")
	fmt.Println(df)

	return df
}
