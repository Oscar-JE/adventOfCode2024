package main

import (
	"adventofcode/day5/rule"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	bytes, err := os.ReadFile("big.txt")
	if err != nil {
		panic("file was not found")
	}
	content := string(bytes)
	rulesAndUpdateSeq := strings.Split(content, "\r\n\r\n")
	ruleRep := rulesAndUpdateSeq[0]
	updatesRep := rulesAndUpdateSeq[1]
	rules := parseRules(ruleRep)
	updates := parseUpdates(updatesRep)
	ruleBook := rule.InitBook(rules)
	sum := 0
	for _, update := range updates {
		if !ruleBook.Pass(update.pages) {
			ordered := ruleBook.Order(update.pages)
			sum += ordered[len(ordered)/2]
		}
	}
	fmt.Println(sum)
}

type update struct {
	pages []int
}

func (u update) score() int {
	return u.pages[len(u.pages)/2]
}

func parseRules(strRep string) []rule.Rule {
	lines := strings.Split(strRep, "\r\n")
	rules := []rule.Rule{}
	for _, line := range lines {
		firstAndSecond := strings.Split(line, "|")
		first, err1 := strconv.Atoi(firstAndSecond[0])
		second, err2 := strconv.Atoi(firstAndSecond[1])
		if err1 != nil || err2 != nil {
			panic("error wile parsing rules")
		}
		rules = append(rules, rule.Init(first, second))
	}
	return rules
}

func parseUpdates(strRep string) []update {
	lines := strings.Split(strRep, "\r\n")
	updates := []update{}
	for _, line := range lines {
		valuesRep := strings.Split(line, ",")
		pages := []int{}
		for _, rep := range valuesRep {
			page, err := strconv.Atoi(rep)
			if err != nil {
				panic("error parsing updates")
			}
			pages = append(pages, page)
		}
		updates = append(updates, update{pages: pages})
	}
	return updates
}
