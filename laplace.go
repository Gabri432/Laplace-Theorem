package main

import "math"

func main() {
	return
}

// det will calculate the determinant of a 2*2 matrix.
// The result will be rounded allowing 2 decimal places at maximum
func Det(a, b, c, d float64) float64 {
	return RoundFloat(a*c-b*d, 100)
}

// roundFloat will round f according to the precision.
// A precision of 100 will mean 2 decimal places, 1000 for 3 decimal places, and so on.
func RoundFloat(f float64, precision int) float64 {
	return math.Round(f*float64(precision)) / float64(precision)
}

func LaplaceDet()
