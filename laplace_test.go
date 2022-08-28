package main

import (
	"testing"
)

var (
	r1     = MatRow{Columns: []float64{1, 0, 0, 0}}
	r2     = MatRow{Columns: []float64{7, 5, 3, 4}}
	r3     = MatRow{Columns: []float64{9, 2, 6, 1}}
	r4     = MatRow{Columns: []float64{3, 4, 9, 5}}
	mat4x4 = Matrix{[]MatRow{r1, r2, r3, r4}}
	row1   = MatRow{Columns: []float64{1, 2, 3}}
	row2   = MatRow{Columns: []float64{4, 5, 6}}
	row3   = MatRow{Columns: []float64{7, 8, 9}}
	mat3x3 = Matrix{[]MatRow{row1, row2, row3}}
)

func TestDet(t *testing.T) {
	determinant := Det(1, 2, 3, 6)
	if determinant != 0 {
		t.Fatalf("Expected determinant should be zero, got %f.", determinant)
	}
}

func TestRoundFloat(t *testing.T) {
	result := RoundFloat(5.234, 10)
	if result-5.2 != 0 {
		t.Fatalf("Expected result should be 5.2, got %f.", result)
	}
}

func TestIsSquareMatrix(t *testing.T) {
	r1 := MatRow{Columns: []float64{1, 2, 3, 4}}
	r2 := MatRow{Columns: []float64{1, 2, 3}}
	r3 := MatRow{Columns: []float64{1, 2, 3, 4}}
	m := Matrix{[]MatRow{r1, r2, r3}}
	if m.IsSquareMatrix() {
		t.Fatal("Expected false, got true.")
	}
}

func TestMapElements(t *testing.T) {
	mappedValues := mat4x4.MapElements()
	if mappedValues[Position{Row: 1, Column: 1}] != 1 {
		t.Fatal("Expected 1, got ", mappedValues[Position{Row: 1, Column: 1}], ".")
	}
}

func TestCreateSubMatrix3x3(t *testing.T) {
	subMatrix := mat4x4.CreateSubMatrix(1, 1)
	if len(subMatrix.Rows) != 3 {
		t.Fatalf("Expected amount of rows to be 3, got %d", len(subMatrix.Rows))
	}
	if len(subMatrix.Rows[0].Columns) != 3 {
		t.Fatalf("Expected amount of columns to be 3, got %d.", len(subMatrix.Rows[0].Columns))
	}
	if subMatrix.Rows[0].Columns[0] != 5 {
		t.Fatalf("Expected element on first row and first column to be 1, got %f.", subMatrix.Rows[0].Columns[0])
	}
}

func TestCreateSubMatrix2x2(t *testing.T) {
	subMatrix := mat3x3.CreateSubMatrix(1, 1)
	if len(subMatrix.Rows) != 2 {
		t.Fatalf("Expected amount of rows to be 2, got %d", len(subMatrix.Rows))
	}
	if len(subMatrix.Rows[0].Columns) != 2 {
		t.Fatalf("Expected amount of columns to be 2, got %d.", len(subMatrix.Rows[0].Columns))
	}
	if subMatrix.Rows[0].Columns[0] != 5 {
		t.Fatalf("Expected element on first row and first column to be 5, got %f.", subMatrix.Rows[0].Columns[0])
	}
}

func TestSubMatrixes3x3(t *testing.T) {
	subMats := mat4x4.SubMatrixes()
	if len(subMats) != 4 {
		t.Fatalf("Expected to have 4 sub-matrixes, got %d.", len(subMats))
	}
	if len(subMats[0].Rows) != len(subMats[0].Rows[0].Columns) {
		t.Fatal("Expected number of rows to be equal to the amount of columns, got not equal.", subMats)
	}
	if len(subMats[0].Rows) != 3 {
		t.Fatalf("Expected rows to be 3, got %d.", len(subMats[0].Rows))
	}
}

func TestSubMatrixes2x2(t *testing.T) {
	subMats := mat3x3.SubMatrixes()
	if len(subMats) != 3 {
		t.Fatalf("Expected to have 3 sub-matrixes, got %d.", len(subMats))
	}
	if len(subMats[0].Rows) != len(subMats[0].Rows[0].Columns) {
		t.Fatal("Expected number of rows to be equal to the amount of columns, got not equal.", subMats)
	}
	if len(subMats[0].Rows) != 2 {
		t.Fatalf("Expected rows to be 2, got %d.", len(subMats[0].Rows))
	}
}

func TestLaplaceDet3x3(t *testing.T) {
	r1 := MatRow{Columns: []float64{0, 0, 3}}
	r2 := MatRow{Columns: []float64{2, 3, 1}}
	r3 := MatRow{Columns: []float64{2, 1, 0}}
	m := Matrix{[]MatRow{r1, r2, r3}}
	result := m.LaplaceDet3x3()
	expectedResult := 3 * Det(2, 3, 2, 1)
	if result != expectedResult {
		t.Fatalf("Expected result to %f, got %f.", expectedResult, result)
	}

}

func TestLaplace4x4(t *testing.T) {
	result := mat4x4.LaplaceDet()
	if result != 63 {
		t.Fatalf("Expected result to be 63, got %f.", result)
	}
}
