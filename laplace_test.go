package main

import (
	"testing"
)

var (
	r1  = MatRow{Columns: []float64{1, 2, 3, 8}}
	r2  = MatRow{Columns: []float64{7, 5, 3, 4}}
	r3  = MatRow{Columns: []float64{9, 2, 6, 1}}
	r4  = MatRow{Columns: []float64{3, 4, 9, 5}}
	mat = Matrix{[]MatRow{r1, r2, r3, r4}}
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
	mappedValues := mat.MapElements()
	if mappedValues[1].Row != 1 {
		t.Fatalf("Expected 1, got %d.", mappedValues[1].Row)
	}
}

func TestCreateSubMatrix(t *testing.T) {
	subMatrix := mat.CreateSubMatrix(1, 1)
	if len(subMatrix.Rows) != 3 {
		t.Fatalf("Expected amount of rows to be 2, got %d", len(subMatrix.Rows))
	}
	if len(subMatrix.Rows[0].Columns) != 3 {
		t.Fatalf("Expected amount of columns to be 3, got %d.", len(subMatrix.Rows[0].Columns))
	}
	if subMatrix.Rows[0].Columns[0] != 5 {
		t.Fatalf("Expected element on first row and first column to be 1, got %f.", subMatrix.Rows[0].Columns[0])
	}
}

func TestSubMatrixes(t *testing.T) {
	subMats := mat.SubMatrixes()
	if len(subMats) != 4 {
		t.Fatalf("Expected to have 4 sub-matrixes, got %d.", len(subMats))
	}
	if len(subMats[0].Rows) != len(subMats[0].Rows[0].Columns) {
		t.Fatal("Expected number of rows to be equal to the amount of columns, got not equal.")
	}
	if len(subMats[0].Rows) != 3 {
		t.Fatalf("Expected rows to be 3, got %d.", len(subMats[0].Rows))
	}
}
