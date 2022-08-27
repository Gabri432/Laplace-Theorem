package main

import "math"

func main() {
	return
}

// det will calculate the determinant of a 2*2 matrix.
//
// The result will be rounded allowing 2 decimal places at maximum
func Det(a, b, c, d float64) float64 {
	return RoundFloat(a*c-b*d, 100)
}

// roundFloat will round f according to the precision.
//
// A precision of 100 will mean 2 decimal places, 1000 for 3 decimal places, and so on.
func RoundFloat(f float64, precision int) float64 {
	return math.Round(f*float64(precision)) / float64(precision)
}

// IsSquareMatrix returns true if the matrix has as many rows as columns.
// Otherwise it gives false.
func (m Matrix) IsSquareMatrix() bool {
	for _, c := range m.Rows {
		if len(c.Row) != len(m.Rows) {
			return false
		}
	}
	return true
}

// It maps the position of each element of the matrix.
//
// It returns a map, each key is an element of the Matrix, each value is the relative row and column.
func (m Matrix) MapElements() map[float64]Position {
	positions := make(map[float64]Position)
	for rowValue := 0; rowValue < len(m.Rows); rowValue++ {
		for columnValue := 0; columnValue < len((m.Rows)); columnValue++ {
			positions[m.Rows[rowValue].Row[columnValue]] = Position{Row: rowValue, Column: columnValue}
		}
	}
	return positions
}

// SubMatrixes will generate all the sub matrixes available.
//
// Each sub-matrix is just another matrix with rows and columns 1 unit shorter than the mother matrix.
// If there is a matrix 3x3, each sub-matrix will be 2x2.
func (m Matrix) SubMatrixes() {}

func (m Matrix) LaplaceDet() {}
