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
			break
		}
		if index2 == len(l2.scores) {
			sumScores = append(sumScores, l1.scores[index1:]...)
			break
		}
		sumScore1 := l1.scores[index1]
		sumScore2 := l2.scores[index2]
		if sumScore1.score < sumScore2.score {
			sumScores = append(sumScores, ScoreAndNumber{score: sumScore1.score, number: sumScore1.number})
			index1++
			continue
		} else if sumScore1.score > sumScore2.score {
			sumScores = append(sumScores, ScoreAndNumber{score: sumScore2.score, number: sumScore2.number})
			index2++
			continue
		} else {
			sumScores = append(sumScores, ScoreAndNumber{score: sumScore1.score, number: sumScore1.number + sumScore2.number})
			index1++
			index2++
		}
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
