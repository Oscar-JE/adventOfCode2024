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

func (m PrintableMatrix[E]) String() string {
	rep := ""
	for i := range m.mat.GetNrRows() {
		lineRep := ""
		for j := range m.mat.GetNrCols() {
			lineRep += m.mat.Get(i, j).String()
		}
		rep += lineRep + "\r\n"
	}
	return rep
}
