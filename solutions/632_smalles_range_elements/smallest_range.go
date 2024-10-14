package smalles_range

import (
	"math"
	"sort"
)

// https://leetcode.com/problems/smallest-range-covering-elements-from-k-lists/description/?envType=daily-question&envId=2024-10-13
// You have k lists of sorted integers in non-decreasing order.
// Find the smallest range that includes at least one number from each of the k lists.
//
// We define the range [a, b] is smaller than range [c, d] if b - a < d - c or a < c if b - a == d - c.
// smallestRange
func smallestRange(nums [][]int) []int {
	numIndex := make([][2]int, 0)
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums[i]); j++ {
			numIndex = append(numIndex, [2]int{nums[i][j], i})
		}
	}

	sort.Slice(numIndex, func(i, j int) bool {
		return numIndex[i][0] < numIndex[j][0]
	})

	response := []int{math.MinInt, math.MaxInt}

	left, right := 0, 0
	visited := make(map[int]int)
	for left < len(numIndex) {
		for left < len(numIndex) && len(visited) < len(nums) {
			visited[numIndex[left][1]]++
			left++
		}

		for right < left && len(visited) == len(nums) {
			if response[1] == math.MaxInt || response[1]-response[0] > numIndex[left-1][0]-numIndex[right][0] {
				response[0], response[1] = numIndex[right][0], numIndex[left-1][0]
			}
			visited[numIndex[right][1]]--
			if visited[numIndex[right][1]] == 0 {
				delete(visited, numIndex[right][1])
			}
			right++
		}
	}

	return response
}
