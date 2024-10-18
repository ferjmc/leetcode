package count_max_or_subsets

// 2044. Count Number of Maximum Bitwise-OR Subsets
// Given an integer array nums, find the maximum possible bitwise
// OR of a subset of nums and return the number of different non-empty subsets with the maximum bitwise OR.
//
// An array a is a subset of an array b if a can be obtained from b by deleting some (possibly zero) elements of b.
// Two subsets are considered different if the indices of the elements chosen are different.
//
// The bitwise OR of an array a is equal to a[0] OR a[1] OR ... OR a[a.length - 1] (0-indexed).
func countMaxOrSubsets(nums []int) int {
	maxOr := 0
	for _, num := range nums {
		maxOr |= num
	}
	return countSubsets(nums, 0, 0, maxOr)
}

func countSubsets(nums []int, index, currentOr, targetOr int) int {
	// base case: reached the end of the array
	if index == len(nums) {
		if currentOr == targetOr {
			return 1
		}
		return 0
	}

	// don't include the current number
	countWithout := countSubsets(nums, index+1, currentOr, targetOr)
	// include the current number
	countWith := countSubsets(nums, index+1, currentOr|nums[index], targetOr)

	// return the sum of both cases
	return countWithout + countWith
}
