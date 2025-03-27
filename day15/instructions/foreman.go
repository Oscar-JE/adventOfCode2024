package instructions

import (
	"adventofcode/day15/directions"
	"strings"
)

type Foreman struct {
	directions   []directions.Direction
	currentIndex int
}

func (f Foreman) HasNext() bool {
	return f.currentIndex < len(f.directions)
}

func (f *Foreman) GetNext() directions.Direction {
	retVal := f.directions[f.currentIndex]
	f.currentIndex++
	return retVal
}

func (f Foreman) String() string {
	rep := ""
	for _, d := range f.directions {
		rep += d.String()
	}
	return rep
}

func InitFromRep(representation string) Foreman {
	rep := strings.Replace(representation, "\r\n", "", -1)
	dir := []directions.Direction{}
	for _, r := range rep {
		dir = append(dir, directions.Init(r))
	}
	return Foreman{currentIndex: 0, directions: dir}
}
