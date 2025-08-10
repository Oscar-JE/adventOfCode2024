package containers

import "testing"

func TestSlizeZero(t *testing.T) {
	testSlize := []int{0, 1}
	subZeroSlize := testSlize[:0]
	if len(subZeroSlize) != 0 {
		t.Errorf("subslice of zero not as expected")
	}
	topSlize := testSlize[2:]
	if len(topSlize) != 0 {
		t.Errorf("this was unfortunate")
	}
}

func TestFindPriorityIndex(t *testing.T) {
	q := []elementAndPriority[int]{elementAndPriority[int]{0, -5}, elementAndPriority[int]{1, 0}, elementAndPriority[int]{2, 5}}
	queue := PriorityQueue[int]{queue: q}
	i1 := queue.findPriorityIndex(-10)
	i2 := queue.findPriorityIndex(-3)
	i3 := queue.findPriorityIndex(3)
	i4 := queue.findPriorityIndex(10)
	if i1 != 0 {
		t.Errorf("if lower than all priorities we should get index 0")
	}
	if i2 != 1 {
		t.Errorf("should be one")
	}
	if i3 != 2 {
		t.Errorf("should be 2")
	}
	if i4 != 3 {
		t.Errorf("largest priority should be placed last")
	}
}

func TestInsert(t *testing.T) {
	t.Skip("created for debugging purposes")
	q := []elementAndPriority[int]{elementAndPriority[int]{0, -5}, elementAndPriority[int]{1, 0}, elementAndPriority[int]{2, 5}}
	queue := PriorityQueue[int]{queue: q}
	queue.insert(elementAndPriority[int]{3, 3}, 0)
	queue.insert(elementAndPriority[int]{4, 4}, 1)
	queue.insert(elementAndPriority[int]{5, 5}, 20)
	queue.insert(elementAndPriority[int]{6, 6}, 5)
	queue.insert(elementAndPriority[int]{7, 7}, 7)
}

func TestRemove(t *testing.T) {
	t.Skip("debugging purposes")
	q := []elementAndPriority[int]{elementAndPriority[int]{0, -5}, elementAndPriority[int]{1, 0}, elementAndPriority[int]{2, 5}}
	queue := PriorityQueue[int]{queue: q}
	queue.remove(1)
	queue.remove(1)
}
