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

func Add(l1 ScoreList, l2 ScoreList) ScoreList {
	index1 := 0
	index2 := 0
	sumScores := []ScoreAndNumber{}
	for index1 != len(l1.scores) || index2 != len(l2.scores) {
		if index1 == len(l1.scores) {
			sumScores = append(sumScores, l2.scores[index2:]...)
		}
		if index2 == len(l2.scores) {
			sumScores = append(sumScores, l1.scores[index1:]...)
		}
		//här händer det lite mer intresanta saker
		sumScore1 := l1.scores[index1]
		sumScore2 := l2.scores[index2]
		if sumScore1.score // fortsätter hät senare
	}
	return ScoreList{sumScores}
}

func (sl ScoreList) Eq(other ScoreList) bool {
	if len(sl.scores) != len(other.scores) {
		return false
	}
	for i, el := range sl.scores {
		if el != other.scores[i] {
			return false
		}
	}
	return true
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
