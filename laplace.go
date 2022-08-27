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
		if len(c.Columns) != len(m.Rows) {
			return false
		}
	}
	return true
}

// It maps the position of each element of the matrix.
//
// It returns a map, each key is an element of the Matrix, each value is the relative row and column.
//
// Each row and column start from 1, NOT 0.
func (m Matrix) MapElements() map[float64]Position {
	positions := make(map[float64]Position)
	for rowValue := 0; rowValue < len(m.Rows); rowValue++ {
		for columnValue := 0; columnValue < len((m.Rows)); columnValue++ {
			positions[m.Rows[rowValue].Columns[columnValue]] = Position{Row: rowValue + 1, Column: columnValue + 1}
		}
	}
	return positions
}

// CreateSubMatrix will generate a sub matrix according to the row and column to exclude.
//
// If row=1 and column=1 it means that the subMatrix will exclude (not contain) the first row and first column of the mother matrix.
func (m Matrix) CreateSubMatrix(row, column int) Matrix {
	subMatrix := Matrix{}
	for currentRow := 0; currentRow < len(m.Rows); currentRow++ {
		if currentRow+1 == row {
			continue
		}
		line := MatRow{}
		for currentColumn := 0; currentColumn < len(m.Rows); currentColumn++ {
			if currentColumn+1 == row {
				continue
			}
			line.Columns = append(line.Columns, m.Rows[currentRow].Columns[currentColumn])
		}
		subMatrix.Rows = append(subMatrix.Rows, line)
	}
	return subMatrix
}

// SubMatrixes will generate all the sub matrixes available.
//
// Each sub-matrix is just another matrix with rows and columns 1 unit shorter than the mother matrix.
// If there is a matrix 3x3, each sub-matrix will be 2x2.
func (m Matrix) SubMatrixes() []Matrix {
	subMatrixes := []Matrix{}
	for rowValue := 0; rowValue < len(m.Rows); rowValue++ {
		for columnValue := 0; columnValue < len((m.Rows)); columnValue++ {
			subMatrixes = append(subMatrixes, m.CreateSubMatrix(rowValue, columnValue))
		}
	}
	return subMatrixes
}

func (m Matrix) LaplaceDet() {}
