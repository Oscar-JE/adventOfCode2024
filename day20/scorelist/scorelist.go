package scorelist

type ScoreAndNumber struct {
	score  int
	number int
}

type ScoreList struct {
	scores []ScoreAndNumber
}

func Init(scores []int) ScoreList {
	sl := ScoreList{scores: []ScoreAndNumber{}}
	for _, score := range scores {
		index, found := sl.findIndexOfScore(score)
		if found {
			sl.scores[index].number++
		} else {
			res := []ScoreAndNumber{}
			res = append(res, sl.scores[:index]...)
			res = append(res, ScoreAndNumber{score: score, number: 1})
			res = append(res, sl.scores[index:]...)
			sl.scores = res
		}
	}
	return sl
}

func (sl ScoreList) Eq(other ScoreList) bool {
	if len(sl.scores) != len(other.scores) {
		return false
	}
	for i , el := range {
		
	}

}

func (sl ScoreList) findIndexOfScore(score int) (int, bool) {
	for index, el := range sl.scores { //kan skrivas om som binärsök om prestandan kräver det
		if score < el.score { // ger en lista i ökande storlek
			return index, false
		}
		if el.score == score {
			return index, true
		}
	}
	return 0, false
}
