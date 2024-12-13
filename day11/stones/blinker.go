package stones

type Blinker struct {
	rules []rule
}

func InitBlinker() Blinker {
	rules := []rule{Zero{}, SplitEvenNrDigits{}, MultiplyWith2024{}}
	return Blinker{rules: rules}
}

type StoneAndRepetitions struct {
	stone      int
	repetition int
}

func (b Blinker) NrOfStonesAfterBlinks(stones []int, nrBlinks int) int {
	uniqueStones := []StoneAndRepetitions{}
	for _, stone := range stones {
		uniqueStones = append(uniqueStones, StoneAndRepetitions{stone, 1})
	}
	return b.NrOfStonesAfterBlinksUnique(uniqueStones, nrBlinks)
}

func (b Blinker) NrOfStonesAfterBlinksUnique(stonesReps []StoneAndRepetitions, nrBlinks int) int {
	for i := 0; i < nrBlinks; i++ {
		stonesReps = b.BlinkRep(stonesReps)
	}
	sum := 0
	for _, stoneRep := range stonesReps {
		sum += stoneRep.repetition
	}
	return sum
}

func (b Blinker) BlinkRep(stonesReps []StoneAndRepetitions) []StoneAndRepetitions {
	nextStonesRep := []StoneAndRepetitions{}
	for _, stoneRep := range stonesReps {
		rule := b.relevantRule(stoneRep.stone)
		nextStones := rule.apply(stoneRep.stone)
		for _, stone := range nextStones {
			index, exists := inList(stone, nextStonesRep)
			if exists {
				nextStonesRep[index].repetition += stoneRep.repetition
			} else {
				nextStonesRep = append(nextStonesRep, StoneAndRepetitions{stone, stoneRep.repetition})
			}
		}
	}
	return nextStonesRep
}

func inList(element int, list []StoneAndRepetitions) (int, bool) {
	inList := false
	index := -1
	for i, el := range list {
		if el.stone == element {
			inList = true
			index = i
			break
		}
	}
	return index, inList
}

func (b Blinker) relevantRule(stone int) rule {
	for _, rule := range b.rules {
		if rule.applicable(stone) {
			return rule
		}
	}
	panic("no rule found to be applied")
}
