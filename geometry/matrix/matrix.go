package matrix

// ska vara generic som kräver comparabel
type Matrix[C comparable] struct {
	values []C
	rows   int
	cols   int
}

func Init[C comparable](values []C, rows int, cols int) Matrix[C] {
	if len(values) != rows*cols {
		panic("unallowed matrix initialization")
	}
	return Matrix[C]{values: values, rows: rows, cols: cols}
}

func (m Matrix[C]) GetNrRows() int {
	return m.rows
}

func (m Matrix[C]) GetNrCols() int {
	return m.cols
}

func (m Matrix[C]) indexFramRowAndCol(row int, col int) int {
	return m.cols*row + col
}

func (m Matrix[C]) Inside(row int, col int) bool {
	return 0 <= row && row < m.rows && 0 <= col && col < m.cols
}

func (m Matrix[C]) Get(row int, col int) C {
	if !m.Inside(row, col) {
		panic("Attempt to access values outside matrix")
	}
	index := m.indexFramRowAndCol(row, col)
	return m.values[index]
}

func (m Matrix[C]) GetRow(row int) []C {
	res := []C{}
	for j := range m.cols {
		res = append(res, m.Get(row, j))
	}
	return res
}
