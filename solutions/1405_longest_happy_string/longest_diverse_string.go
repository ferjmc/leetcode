package longest_diverse_string

import (
	"container/heap"
	"sort"
	"strings"
)

type CurrentLimit struct {
	letter byte
	limit  int
}

type CurrentLimitHeap []*CurrentLimit

func (h CurrentLimitHeap) Len() int { return len(h) }
func (h CurrentLimitHeap) Less(i, j int) bool {
	if h[i].limit == h[j].limit {
		return h[i].letter < h[j].letter
	}
	return h[i].limit > h[j].limit
}
func (h CurrentLimitHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *CurrentLimitHeap) Push(x interface{}) { *h = append(*h, x.(*CurrentLimit)) }
func (h *CurrentLimitHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[:n-1]
	return item
}

// 1405. Longest Happy String
// A string s is called happy if it satisfies the following conditions:
//
// s only contains the letters 'a', 'b', and 'c'.
// s does not contain any of "aaa", "bbb", or "ccc" as a substring.
// s contains at most a occurrences of the letter 'a'.
// s contains at most b occurrences of the letter 'b'.
// s contains at most c occurrences of the letter 'c'.
// Given three integers a, b, and c, return the longest possible happy string.
// If there are multiple longest happy strings, return any of them.
// If there is no such string, return the empty string "".
//
// A substring is a contiguous sequence of characters within a string.
func longestDiverseString(a int, b int, c int) string {
	// initialize priority queue
	h := make(CurrentLimitHeap, 0)
	heap.Init(&h)
	// push only non 0 values
	if a > 0 {
		heap.Push(&h, &CurrentLimit{letter: 'a', limit: a})
	}
	if b > 0 {
		heap.Push(&h, &CurrentLimit{letter: 'b', limit: b})
	}
	if c > 0 {
		heap.Push(&h, &CurrentLimit{letter: 'c', limit: c})
	}

	// use string builder to construct response
	var response strings.Builder

	for h.Len() > 0 {
		// get highest limit letter
		top1 := heap.Pop(&h).(*CurrentLimit)

		if response.Len() > 1 &&
			response.String()[response.Len()-1] == top1.letter &&
			response.String()[response.Len()-2] == top1.letter {

			if h.Len() == 0 {
				break
			}

			top2 := heap.Pop(&h).(*CurrentLimit)
			response.WriteByte(top2.letter)
			top2.limit--
			if top2.limit > 0 {
				response.WriteByte(top2.letter)
				top2.limit--
				if top2.limit > 0 {
					heap.Push(&h, top2)
				}
			}
			if top1.limit > 0 {
				heap.Push(&h, top1)
			}
		} else {
			response.WriteByte(top1.letter)
			top1.limit--
			if top1.limit > 0 {
				response.WriteByte(top1.letter)
				top1.limit--
				if top1.limit > 0 {
					heap.Push(&h, top1)
				}
			}
		}

		//	// at init is safe to add current char
		//	if response.Len() <= 1 {
		//		response.WriteByte(top1.letter)
		//		top1.limit--
		//		if top1.limit > 0 {
		//			response.WriteByte(top1.letter)
		//			top1.limit--
		//			if top1.limit > 0 {
		//				heap.Push(&h, top1)
		//			}
		//		}
		//	} else {
		//		// check validation of two consecutive same chars
		//		if response.String()[response.Len()-1] == top1.letter && response.String()[response.Len()-2] == top1.letter {
		//			// use the second top limit
		//			top2 := heap.Pop(&h).(*CurrentLimit)
		//			response.WriteByte(top2.letter)
		//			top2.limit--
		//			if top2.limit > 0 {
		//				response.WriteByte(top2.letter)
		//				top2.limit--
		//			}
		//
		//			// push back the updated limits
		//			if top2.limit > 0 {
		//				heap.Push(&h, top2)
		//			}
		//			if top1.limit > 0 {
		//				heap.Push(&h, top1)
		//			}
		//		} else {
		//			response.WriteByte(top1.letter)
		//			top1.limit--
		//			if top1.limit > 0 {
		//				response.WriteByte(top1.letter)
		//				top1.limit--
		//				if top1.limit > 0 {
		//					heap.Push(&h, top1)
		//				}
		//			}
		//		}
		//	}
		//}

		// handle special case when len = 1
		//if h.Len() == 1 {
		//	top1 := heap.Pop(&h).(*CurrentLimit)
		//
		//	if top1.limit <= 1 {
		//		response.WriteByte(top1.letter)
		//	} else {
		//		response.WriteByte(top1.letter)
		//		response.WriteByte(top1.letter)
		//	}
	}

	return response.String()
}

func longestDiverseString2(a int, b int, c int) string {
	aLimit, bLimit, cLimit := 0, 0, 0
	totalIteractions := a + b + c
	var result strings.Builder

	for i := 0; i < totalIteractions; i++ {
		if (a >= b && a >= c && aLimit != 2) || (a > 0 && (bLimit == 2 || cLimit == 2)) {
			result.WriteByte('a')
			a--
			aLimit++
			bLimit, cLimit = 0, 0
		} else if (b >= a && b >= c && bLimit != 2) || b > 0 && (cLimit == 2 || aLimit == 2) {
			result.WriteByte('b')
			b--
			bLimit++
			aLimit, cLimit = 0, 0
		} else if (c >= a && c >= b && cLimit != 2) || c > 0 && (aLimit == 2 || bLimit == 2) {
			result.WriteByte('c')
			c--
			cLimit++
			aLimit, bLimit = 0, 0
		}
	}
	return result.String()
}

type Char struct {
	char  string
	count int
}

func longestDiverseString3(a int, b int, c int) string {
	array := []Char{
		{"a", a},
		{"b", b},
		{"c", c},
	}

	happyString := ""

	for a+b+c > 0 {
		sort.Slice(array, func(i, j int) bool {
			return array[i].count > array[j].count
		})

		left := false
		for i, item := range array {
			if item.count > 0 &&
				!(len(happyString) > 1 &&
					string(happyString[len(happyString)-2:]) == strings.Repeat(item.char, 2)) {
				happyString += item.char
				array[i].count--
				left = true
				break
			}
		}

		if !left {
			break
		}
	}

	return happyString
}
