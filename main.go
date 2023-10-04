package main

//Import FMT and Stats packaage from Montana Flynn

import (
	"fmt"

	"github.com/montanaflynn/stats"
)

// main function that inputs anscombe data
func main() {
	// use stats.coordinate to make the four data series seen in the python and r code
	// quartet 1
	dataxy1 := []stats.coordinate{
		{X: 10, Y: 8.04},
		{X: 8, Y: 6.95},
		{X: 13, Y: 7.58},
		{X: 9, Y: 8.81},
		{X: 11, Y: 8.33},
		{X: 14, Y: 9.96},
		{X: 6, Y: 7.24},
		{X: 4, Y: 4.26},
		{X: 12, Y: 10.84},
		{X: 7, Y: 4.82},
		{X: 5, Y: 5.68},
	}
	// quartet 2
	dataxy2 := []stats.coordinate{
		{X: 10, Y: 9.14},
		{X: 8, Y: 8.14},
		{X: 13, Y: 8.74},
		{X: 9, Y: 8.77},
		{X: 11, Y: 9.26},
		{X: 14, Y: 8.1},
		{X: 6, Y: 6.13},
		{X: 4, Y: 3.1},
		{X: 12, Y: 9.13},
		{X: 7, Y: 7.26},
		{X: 5, Y: 4.74},
	}
	// quartet 3
	dataxy3 := []stats.coordinate{
		{X: 10, Y: 7.46},
		{X: 8, Y: 6.77},
		{X: 13, Y: 12.74},
		{X: 9, Y: 7.11},
		{X: 11, Y: 7.81},
		{X: 14, Y: 8.84},
		{X: 6, Y: 6.08},
		{X: 4, Y: 5.39},
		{X: 12, Y: 8.15},
		{X: 7, Y: 6.42},
		{X: 5, Y: 5.73},
	}
	// quartet 4
	dataxy4 := []stats.coordinate{
		{X: 8, Y: 6.58},
		{X: 8, Y: 5.76},
		{X: 8, Y: 7.71},
		{X: 8, Y: 8.84},
		{X: 8, Y: 8.47},
		{X: 8, Y: 7.04},
		{X: 8, Y: 5.25},
		{X: 19, Y: 12.5},
		{X: 8, Y: 5.56},
		{X: 8, Y: 7.91},
		{X: 8, Y: 6.89},
	}

	//linear regression for each quartet
	q1, _ := stats.LinearRegression(dataxy1)
	fmt.Println(q1)

	q2, _ := stats.LinearRegression(dataxy2)
	fmt.Println(q2)

	q3, _ := stats.LinearRegression(dataxy3)
	fmt.Println(q3)

	q4, _ := stats.LinearRegression(dataxy4)
	fmt.Println(q4)
}

// create linear regression models by ordinary least squares (montana flynn stats package)
