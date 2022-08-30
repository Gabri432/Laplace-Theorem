package laplace

// MatRow type is simply a set of floating point numbers.
//
// Each number is ideally a section of a column.
type MatRow struct {
	Columns []float64
}

// Matrix type is a set of Matrows.
type Matrix struct {
	Rows []MatRow
}

type Position struct {
	Row    int
	Column int
}
