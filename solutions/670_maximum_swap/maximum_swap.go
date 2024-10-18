package maximum_swap

import "strconv"

// 670. Maximum Swap
// https://leetcode.com/problems/maximum-swap/description/?envType=daily-question&envId=2024-10-17
// You are given an integer num. You can swap two digits at most once to get the maximum valued number.
//
// Return the maximum valued number you can get.
func maximumSwap(num int) int {
	numStr := []rune(strconv.Itoa(num))
	minIndex := 0
	maxIndex := len(numStr) - 1
	for i := 0; i < len(numStr); i++ {
		if numStr[i] < numStr[minIndex] {
			minIndex = i
		}
	}
	for i := len(numStr) - 1; i >= 0; i-- {
		if numStr[i] > numStr[minIndex] {
			maxIndex = i
		}
	}
	if numStr[minIndex] > numStr[maxIndex] {
		numStr[minIndex], numStr[maxIndex] = numStr[maxIndex], numStr[minIndex]
	}
	result, _ := strconv.Atoi(string(numStr))
	return result
}

func maximumSwap2(num int) int {
	nums := make([]int, 0)
	for num > 0 {
		nums = append([]int{num % 10}, nums...)
		num /= 10
	}

	n := len(nums)
	for i := 0; i < len(nums); i++ {
		index := i
		for j := i + 1; j < n; j++ {
			if nums[j] >= nums[index] {
				index = j
			}
		}
		if index != i && nums[index] != nums[i] {
			nums[i], nums[index] = nums[index], nums[i]
			break
		}
	}
	res := 0
	for _, v := range nums {
		res = res*10 + v
	}
	return res
}
