package main

type MatRow struct {
	Row []float64
}

type Matrix struct {
	Rows []MatRow
}

type Position struct {
	Row    int
	Column int
}
