package diskfrag

import "fmt"

func expand(numRep []int) []Block {
	expanded := []Block{}
	id := 0
	for i, num := range numRep {
		free := (i+1)%2 == 0
		for j := 0; j < num; j++ {
			expanded = append(expanded, Block{free: free, id: id})
		}
		if i%2 == 0 {
			id++
		}
	}
	return expanded
}

func contract(mem []Block) []int {
	concentrated := []int{}
	startIndex := 0
	endIndex := len(mem) - 1
	for startIndex <= endIndex {
		if mem[endIndex].free {
			endIndex--
			continue
		}
		if !mem[startIndex].free {
			concentrated = append(concentrated, mem[startIndex].id)
			startIndex++
			continue
		}
		if mem[startIndex].free && !mem[endIndex].free {
			concentrated = append(concentrated, mem[endIndex].id)
			endIndex--
			startIndex++
		}

	}
	return concentrated
}

func Fragment(mem []int) []int {
	expanded := expand(mem)
	contracted := contract(expanded)
	return contracted
}

func FragmentFiles(mem []int) []int {
	expanded := expand(mem)
	contractedBlocks := contractedBlocks(expanded)
	return replaceEmptyWithZero(contractedBlocks)
}

type segment struct {
	startIndex int
	len        int
}

func contractedBlocks(expanded []Block) []Block {
	file := findFileBefore(len(expanded), expanded)
	for file.len != 0 {
		//fmt.Println(visualize(expanded))
		nextFiel := findFileBefore(file.startIndex, expanded)
		emptyPlace, found := findFirstEmptySegment(file, expanded)
		if found {
			expanded = interchangeValues(file, emptyPlace, expanded)
		}
		file = nextFiel
	}
	return expanded
}

func visualize(blocks []Block) string {
	retString := ""
	for _, b := range blocks {
		if b.free {
			retString += "."
		} else {
			retString += fmt.Sprint(b.id)
		}
	}
	return retString
}

func interchangeValues(s1 segment, s2 segment, blocks []Block) []Block {
	if s1.len != s2.len {
		panic("illegal input")
	}
	for i := 0; i < s1.len; i++ {
		blocks = changeValuesOfIndex(s1.startIndex+i, s2.startIndex+i, blocks)
	}
	return blocks
}

func findFirstEmptySegment(segOrg segment, blocks []Block) (segment, bool) {
	var seg segment
	found := false
	for i := range segOrg.startIndex {
		if canFitt(segment{i, segOrg.len}, blocks) {
			seg = segment{i, segOrg.len}
			found = true
			break
		}
	}
	return seg, found
}

func canFitt(seg segment, blocks []Block) bool {
	fits := true
	for i := seg.startIndex; i < seg.startIndex+seg.len; i++ {
		fits = fits && blocks[i].free
	}
	return fits
}

func findFileBefore(j int, expanded []Block) segment {
	firstNonEmptyBlock := j
	for i := j - 1; i >= 0; i-- {
		if !expanded[i].free {
			firstNonEmptyBlock = i
			break
		}
	}
	seg := segment{firstNonEmptyBlock, 1}
	for i := firstNonEmptyBlock - 1; i >= 0; i-- {
		if expanded[i].Eq(expanded[firstNonEmptyBlock]) {
			seg.startIndex = seg.startIndex - 1
			seg.len = seg.len + 1
		} else {
			return seg
		}
	}
	seg.len = 0
	return seg
}

func replaceEmptyWithZero(blocks []Block) []int {
	nums := []int{}
	for _, b := range blocks {
		if b.free {
			nums = append(nums, 0)
		} else {
			nums = append(nums, b.id)
		}
	}
	return nums
}

func changeValuesOfIndex(index1 int, index2 int, blocks []Block) []Block {
	valOfIndex2 := blocks[index2]
	blocks[index2] = blocks[index1]
	blocks[index1] = valOfIndex2
	return blocks
}
