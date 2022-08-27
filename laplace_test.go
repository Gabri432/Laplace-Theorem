package main

import (
	"testing"
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
	r1 := MatRow{Columns: []float64{1, 2, 3, 8}}
	r2 := MatRow{Columns: []float64{7, 5, 3, 4}}
	r3 := MatRow{Columns: []float64{9, 2, 6, 1}}
	m := Matrix{[]MatRow{r1, r2, r3}}
	mappedValues := m.MapElements()
	if mappedValues[1].Row != 1 {
		t.Fatalf("Expected 1, got %d.", mappedValues[1].Row)
	}
}
