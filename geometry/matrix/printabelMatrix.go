package matrix

import "fmt"

type PrintableComparable interface {
	fmt.Stringer
	comparable
}

func InitPrintable[C PrintableComparable](values []C, rows int, cols int) PrintableMatrix[C] {
	if len(values) != rows*cols {
		panic("unallowed matrix initialization")
	}
	return PrintableMatrix[C]{mat: Matrix[C]{values: values, rows: rows, cols: cols}}
}

type PrintableMatrix[E PrintableComparable] struct {
	mat Matrix[E]
}

func (m PrintableMatrix[E]) Get(row int, col int) E {
	return m.mat.Get(row, col)
}

func (m *PrintableMatrix[E]) Set(row int, col int, val E) {
	m.mat.Set(row, col, val)
}

func (m PrintableMatrix[E]) String() string {
	rep := ""
	for i := range m.mat.GetNrRows() - 1 {
		lineRep := ""
		for j := range m.mat.GetNrCols() {
			lineRep += m.mat.Get(i, j).String()
		}
		rep += lineRep + "\r\n"
	}
	lineRep := ""
	for j := range m.mat.GetNrCols() {
		lineRep += m.mat.Get(m.mat.rows-1, j).String()
	}
	rep += lineRep
	return rep
}

func (m PrintableMatrix[E]) FirstRowAndColOf(sought E) (int, int) {
	return m.mat.FirstRowAndColOf(sought)
}
