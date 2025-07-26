package main

import (
	"adventofcode/day16/model"
	"adventofcode/day16/state"
	"fmt"
	"os"
)

func main() {
	println("hej")
	bytes, error := os.ReadFile("input_exe.txt")
	if error != nil {
		println("wtf filen är borta")
	}
	content := string(bytes)
	m := model.Init(content)
	sv := model.InitValueFunc(m.GetNrRows(), m.GetNrCols())
	s := sv.FirstState()
	for i := 0; i < 1000; i++ {
		for sv.HasNext(s) {
			valueFunc := bellman(m, s, sv)
			sv.SetValueOf(s, valueFunc)
			s = sv.GetNext(s)
		}
		s = sv.FirstState()
	}
	fmt.Println(sv)
}

func bellman(m model.Model, s state.State, sv model.StateValue) float64 {
	actionValCandidate := ActionStateValue(m, s, model.ActionSpace[0], &sv)
	for i := 1; i < len(model.ActionSpace); i++ {
		actionValLoop := ActionStateValue(m, s, model.ActionSpace[i], &sv)
		actionValCandidate = min(actionValLoop, actionValCandidate)
	}
	return actionValCandidate
}

func ActionStateValue(m model.Model, s state.State, a model.Action, sv *model.StateValue) float64 {
	nextState := m.StateTransition(s, a)
	costTransition := m.Cost(s, model.ActionSpace[0])
	firstCandidate := sv.GetValueOf(nextState) + costTransition
	return firstCandidate
}
