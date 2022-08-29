package laplace

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

var (
	matRow5x1 = MatRow{Columns: []float64{1, 0, 0, 0, 0}}
	matRow5x2 = MatRow{Columns: []float64{0, 1, 0, 0, 0}}
	matRow5x3 = MatRow{Columns: []float64{0, 0, 1, 0, 0}}
	matRow5x4 = MatRow{Columns: []float64{0, 0, 0, 1, 0}}
	matRow5x5 = MatRow{Columns: []float64{0, 0, 0, 0, 1}}
	mat5x5    = Matrix{[]MatRow{matRow5x1, matRow5x2, matRow5x3, matRow5x4, matRow5x5}}
)

var (
	matRow6x1 = MatRow{Columns: []float64{1, 0, 0, 0, 0, 0}}
	matRow6x2 = MatRow{Columns: []float64{0, 1, 0, 0, 0, 0}}
	matRow6x3 = MatRow{Columns: []float64{0, 0, 1, 0, 0, 0}}
	matRow6x4 = MatRow{Columns: []float64{0, 0, 0, 1, 0, 0}}
	matRow6x5 = MatRow{Columns: []float64{0, 0, 0, 0, 1, 0}}
	matRow6x6 = MatRow{Columns: []float64{0, 0, 0, 0, 0, 1}}
	mat6x6    = Matrix{[]MatRow{matRow6x1, matRow6x2, matRow6x3, matRow6x4, matRow6x5, matRow6x6}}
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

func TestCreateSubMatrix4x4(t *testing.T) {
	subMatrix := mat5x5.CreateSubMatrix(1, 1)
	if len(subMatrix.Rows) != 4 {
		t.Fatalf("Expected amount of rows to be 4, got %d", len(subMatrix.Rows))
	}
	if len(subMatrix.Rows[0].Columns) != 4 {
		t.Fatalf("Expected amount of columns to be 4, got %d.", len(subMatrix.Rows[0].Columns))
	}
	if subMatrix.Rows[0].Columns[0] != 1 {
		t.Fatalf("Expected element on first row and first column to be 1, got %f.", subMatrix.Rows[0].Columns[0])
	}
}

func TestSubMatrices3x3(t *testing.T) {
	subMats := mat4x4.SubMatrices()
	if len(subMats) != 4 {
		t.Fatalf("Expected to have 4 sub-matrices, got %d.", len(subMats))
	}
	if len(subMats[0].Rows) != len(subMats[0].Rows[0].Columns) {
		t.Fatal("Expected number of rows to be equal to the amount of columns, got not equal.", subMats)
	}
	if len(subMats[0].Rows) != 3 {
		t.Fatalf("Expected rows to be 3, got %d.", len(subMats[0].Rows))
	}
}

func TestSubMatrices2x2(t *testing.T) {
	subMats := mat3x3.SubMatrices()
	if len(subMats) != 3 {
		t.Fatalf("Expected to have 3 sub-matrices, got %d.", len(subMats))
	}
	if len(subMats[0].Rows) != len(subMats[0].Rows[0].Columns) {
		t.Fatal("Expected number of rows to be equal to the amount of columns, got not equal.", subMats)
	}
	if len(subMats[0].Rows) != 2 {
		t.Fatalf("Expected rows to be 2, got %d.", len(subMats[0].Rows))
	}
}

func TestSubMatrices4x4(t *testing.T) {
	subMats := mat5x5.SubMatrices()
	if len(subMats) != 5 {
		t.Fatalf("Expected to have 5 sub-matrices, got %d.", len(subMats))
	}
	if len(subMats[0].Rows) != len(subMats[0].Rows[0].Columns) {
		t.Fatal("Expected number of rows to be equal to the amount of columns, got not equal.", subMats)
	}
	if len(subMats[0].Rows) != 4 {
		t.Fatalf("Expected rows to be 4, got %d.", len(subMats[0].Rows))
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

func TestReducer5x5(t *testing.T) {
	result := LaplaceBig(mat5x5)
	if result != 1 {
		t.Fatalf("Expected det(I-5x5) to be 1, got %f.\n", result)
	}
}

func TestReducer6x6(t *testing.T) {
	result := LaplaceBig(mat6x6)
	if result != 1 {
		t.Fatalf("Expected det(I-6x6) to be 1, got %f.\n", result)
	}
}
