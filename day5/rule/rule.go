package rule

import "slices"

type Rule struct {
	first  int
	second int
}

type RuleBook struct {
	rules []Rule
}

func InitBook(rules []Rule) RuleBook {
	return RuleBook{rules: rules}
}

func (b RuleBook) Pass(update []int) bool {
	for _, rule := range b.rules {
		if !rule.Pass(update) {
			return false
		}
	}
	return true
}

func (b RuleBook) Order(update []int) []int {
	//japp jag tror jag har en approtch
	ordered := []int{}
	for _, page := range update {
		insertIndex := len(ordered)
		for i, orderedEl := range ordered {
			if b.correctOrder(page, orderedEl) {
				insertIndex = i
				break
			}
		}
		ordered = slices.Insert(ordered, insertIndex, page)
	}
	return ordered
}

func (b RuleBook) correctOrder(first int, second int) bool {
	return b.Pass([]int{first, second})
}

func Init(first int, second int) Rule {
	return Rule{first: first, second: second}
}

func (r Rule) Pass(update []int) bool {
	firstIndex, found1 := findIndexOf(r.first, update)
	secondIndex, found2 := findIndexOf(r.second, update)
	if !(found1 && found2) {
		return true
	}
	return firstIndex < secondIndex
}

func findIndexOf(seeked int, values []int) (int, bool) {
	for i, val := range values {
		if val == seeked {
			return i, true
		}
	}
	return -1, false
}
