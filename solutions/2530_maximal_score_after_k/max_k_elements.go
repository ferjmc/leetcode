package max_k_elements

import (
	"container/heap"
	"leetcode/data_structure/heaps"
	"math"
)

// 2530. Maximal Score After Applying K Operations
// https://leetcode.com/problems/maximal-score-after-applying-k-operations/description/?envType=daily-question&envId=2024-10-14
// You are given a 0-indexed integer array nums and an integer k. You have a starting score of 0.
//
// In one operation:
//
// choose an index i such that 0 <= i < nums.length,
// increase your score by nums[i], and
// replace nums[i] with ceil(nums[i] / 3).
// Return the maximum possible score you can attain after applying exactly k operations.
//
// The ceiling function ceil(val) is the least integer greater than or equal to val.
func maxKelements(nums []int, k int) int64 {
	score := int64(0)

	maxNums := &heaps.MaxHeap{}
	heap.Init(maxNums)
	for i := 0; i < len(nums); i++ {
		heap.Push(maxNums, nums[i])
	}

	for i := 0; i < k; i++ {
		maxNum := heap.Pop(maxNums).(int)
		score += int64(maxNum)
		floatMax := float64(maxNum) / 3
		ceilMax := math.Ceil(floatMax)
		intMax := int(ceilMax)
		heap.Push(maxNums, intMax)
	}

	return score
}
