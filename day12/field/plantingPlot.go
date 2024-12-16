package field

type plantingPlot struct {
	id      string
	boarder swirl
}

func Init(id string) plantingPlot {
	return plantingPlot{id: id, boarder: InitSwirl()}
}

func (p plantingPlot) String() string {
	return p.id
}
