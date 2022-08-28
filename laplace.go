package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	mat := getUserInput()
	if !(mat).IsSquareMatrix() {
		fmt.Println("This isn't a square matrix.\n Retry.")
		return
	} else {
		fmt.Println(validateMatrixType(mat))
	}
}

func getUserInput() Matrix {
	scanner := bufio.NewScanner(os.Stdin)
	newMatrix := Matrix{}
	fmt.Println("Insert Input:\n 1) each line creates a row of the Matrix.")
	fmt.Println("2) each number of the line must be separated by a semicolon;\n 3) once you have finished write END to start running the program.")
	fmt.Println("Example: \n 3;2\n 1;5\n END")
	for scanner.Scan() {
		if scanner.Text() == "END" {
			break
		}
		newMatrix.Rows = append(newMatrix.Rows, generateMatRow(scanner.Text()))
	}
	fmt.Println(newMatrix)
	return newMatrix
}

func generateMatRow(s string) MatRow {
	newMatRow := MatRow{}
	values := strings.Split(s, ";")
	for _, v := range values {
		column, err := strconv.ParseFloat(v, 64)
		if err != nil {
			fmt.Println(err.Error())
			break
		}
		newMatRow.Columns = append(newMatRow.Columns, column)
	}
	return newMatRow
}

func validateMatrixType(m Matrix) float64 {
	switch len(m.Rows) {
	case 1:
		return m.Rows[0].Columns[0]
	case 2:
		return Det(m.Rows[0].Columns[0], m.Rows[0].Columns[1], m.Rows[1].Columns[0], m.Rows[1].Columns[1])
	case 3:
		return m.LaplaceDet3x3()
	case 4:
		return m.LaplaceDet()
	case 5:
		return m.LaplaceBig()
	default:
		fmt.Println("Warning: the program is not tested properly for matrices of such size.")
		return m.LaplaceBig()
	}
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
// It returns a map, each key is the relative row and column, each value is an element of the first row of the matrix.
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
// Example:
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
	if len(m.Rows[0].Columns) > 3 {
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

// Reducer generates sub-matrices4x4 from bigger matrices.
func Reducer(matrices []Matrix) []Matrix {
	if len(matrices[0].Rows) > 4 {
		newSubMatrices := []Matrix{}
		for _, subM := range matrices {
			newSubMatrices = append(newSubMatrices, subM.SubMatrices()...)
		}
		return Reducer(newSubMatrices)
	}
	return matrices
}

// LaplageBig will calculate determinants for matrix-5x5 by using the Laplace Theorem.
//
// ATTENTION: It currently does not work properly for bigger matrices (matrix-6x6 or more).
//
// Check 'TestLaplaceBig6x6' function to see the problem better.
func (m Matrix) LaplaceBig() float64 {
	matrices4x4 := Reducer(m.SubMatrices())
	result := 0.0
	for _, mat := range matrices4x4 {
		result += mat.LaplaceDet()
	}
	return result
}
