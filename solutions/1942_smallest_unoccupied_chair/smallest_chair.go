package smallest

import (
	"container/heap"
	"fmt"
	"leetcode/data_structure"
	"sort"
)

// smallestChair
// https://leetcode.com/problems/the-number-of-the-smallest-unoccupied-chair/description/?envType=daily-question&envId=2024-10-11
func smallestChair(times [][]int, targetFriend int) int {
	n := len(times)
	if n < 2 {
		panic(fmt.Sprintf("times should be longer than %d", n))
	}

	info := make([]FriendInformation, 0)
	for i := 0; i < n; i++ {
		info = append(info, FriendInformation{
			id:      i,
			time:    times[i][0],
			leaving: false,
		})
		info = append(info, FriendInformation{
			id:      i,
			time:    times[i][1],
			leaving: true,
		})
	}

	// Sort the times array by arrival time
	sort.Slice(info, func(i, j int) bool {
		if info[i].time == info[j].time {
			return info[i].leaving
		}
		return info[i].time < info[j].time
	})

	chairsByFriend := make(map[int]int, n)

	freeChairHeap := &data_structure.MinHeap{}
	heap.Init(freeChairHeap)

	chairNumber := 0

	for i := 0; i < len(info); i++ {
		if chair, ok := chairsByFriend[info[i].id]; ok {
			// put the chair on free chairs heap
			heap.Push(freeChairHeap, chair)
		} else if freeChairHeap.Len() > 0 {
			// take from used free chairs
			chairsByFriend[info[i].id] = heap.Pop(freeChairHeap).(int)
		} else {
			// take new chair
			chairsByFriend[info[i].id] = chairNumber
			chairNumber++
		}

		if info[i].id == targetFriend {
			return chairsByFriend[info[i].id]
		}
	}

	panic(fmt.Sprintf("chair number %d is not found", targetFriend))
}

type FriendInformation struct {
	id      int
	time    int
	leaving bool
}
