package main

import "math"

func main() {
	return
}

// det will calculate the determinant of a matrix-2x2.
//
// The result will be rounded allowing 2 decimal places at maximum
func Det(a, b, c, d float64) float64 {
	return RoundFloat(a*d-b*c, 100)
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

// It maps the position of each element of the first row of the matrix.
//
// It returns a map, each key is an element of the first row of the matrix, each value is the relative row and column.
//
// The first row or column will be 1, NOT 0.
func (m Matrix) MapElements() map[Position]float64 {
	positions := make(map[Position]float64)
	for columnValue := 0; columnValue < len(m.Rows); columnValue++ {
		positions[Position{Row: 1, Column: columnValue + 1}] = m.Rows[0].Columns[columnValue]
	}
	return positions
}

// CreateSubMatrix will generate a sub matrix according to the row and column to exclude.
//
// If row=1 and column=1 it means that the subMatrix will exclude (not contain) elements from the first row and first column of the mother matrix.
func (m Matrix) CreateSubMatrix(row, column int) Matrix {
	subMatrix := Matrix{}
	for currentRow := 0; currentRow < len(m.Rows); currentRow++ {
		if currentRow+1 == row {
			continue
		}
		line := MatRow{}
		line.Columns = append(line.Columns, m.Rows[currentRow].Columns[:column-1]...)
		if column < len(m.Rows[0].Columns) {
			line.Columns = append(line.Columns, m.Rows[currentRow].Columns[column:]...)
		}
		subMatrix.Rows = append(subMatrix.Rows, line)
	}
	return subMatrix
}

// SubMatrices will generate all the sub matrices that don't contains elements from the first row.
//
// Each sub-matrix is just another matrix with rows and columns 1 unit shorter than the mother matrix.
// If there is a matrix 3x3, each sub-matrix will be 2x2.
// If there is a matrix 4x4, each sub-matrix will be 3x3.
//
// Examples:
//
// a Matrix3x3: [1 2 3; 3 7 9; 1 2 1] => [3 7;1 2],[7 9;2 1]
func (m Matrix) SubMatrices() []Matrix {
	subMatrices := []Matrix{}
	for columnValue := 1; columnValue <= len(m.Rows); columnValue++ {
		subMatrices = append(subMatrices, m.CreateSubMatrix(1, columnValue))
	}
	return subMatrices
}

// LaplaceDet will calculate the matrix-3x3 determinant by using the Laplace Theorem.
func (m Matrix) LaplaceDet3x3() float64 {
	var result float64
	for pos, elem := range m.MapElements() {
		pow := float64(pos.Column + pos.Row)
		subMatrix := m.CreateSubMatrix(pos.Row, pos.Column)
		a, b := subMatrix.Rows[0].Columns[0], subMatrix.Rows[0].Columns[1]
		c, d := subMatrix.Rows[1].Columns[0], subMatrix.Rows[1].Columns[1]
		result += math.Pow(-1, pow) * float64(elem) * Det(a, b, c, d)
	}
	return result
}

// LaplaceDet will calculate the matrix-4x4 determinant by using the Laplace Theorem.
func (m Matrix) LaplaceDet() float64 {
	subMatrices := []Matrix{}
	if len(m.Rows[0].Columns) > 2 {
		for _, subM := range m.SubMatrices() {
			subMatrices = append(subMatrices, subM)
		}
	}
	result := 0.0
	for pos, elem := range m.MapElements() {
		power := float64(pos.Row + pos.Column)
		result += math.Pow(-1, power) * elem * subMatrices[pos.Column-1].LaplaceDet3x3()
	}
	return result
}
