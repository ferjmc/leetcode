package min_groups

import (
	"container/heap"
	"leetcode/data_structure/heaps"
	"sort"
)

func minGroups(intervals [][]int) int {

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	groups := &heaps.MinHeap{}
	heap.Init(groups)

	for i := 0; i < len(intervals); i++ {
		// if the current interval starts after the earliest ending interval,
		// we can reuse the group, so pop the top of the heap
		if groups.Len() > 0 && (*groups)[0] < intervals[i][0] {
			heap.Pop(groups)
		}
		// push the current interval's end time onto the groups
		heap.Push(groups, intervals[i][1])
	}

	return groups.Len()
}
