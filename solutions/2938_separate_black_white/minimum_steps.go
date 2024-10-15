package minimum_steps

// 2938. Separate Black and White Balls
// https://leetcode.com/problems/separate-black-and-white-balls/?envType=daily-question&envId=2024-10-15
// There are n balls on a table, each ball has a color black or white.
//
// You are given a 0-indexed binary string s of length n, where 1 and 0 represent black and white balls, respectively.
//
// In each step, you can choose two adjacent balls and swap them.
//
// Return the minimum number of steps to group all the black balls to the right and all the white balls to the left.
// minimumSteps
func minimumSteps(s string) int64 {
	count := 0

	left, right := 0, len(s)-1
	for left < right {
		for left < right && s[left] == '0' {
			left++
		}
		for left < right && s[right] == '1' {
			right--
		}
		if left < right {
			count += right - left
		}

		left++
		right--
	}

	return int64(count)
}
