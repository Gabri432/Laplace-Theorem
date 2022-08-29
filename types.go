package laplacetheorem

type MatRow struct {
	Columns []float64
}

type Matrix struct {
	Rows []MatRow
}

type Position struct {
	Row    int
	Column int
}
