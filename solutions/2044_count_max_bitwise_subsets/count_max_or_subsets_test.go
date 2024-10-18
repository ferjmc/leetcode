package count_max_or_subsets

import "testing"

func TestCountMaxOrSubsets(t *testing.T) {
	t.Parallel()

	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{3, 1}, 2},
		{[]int{2, 2, 2}, 7},
		{[]int{3, 2, 1, 5}, 6},
	}

	for _, test := range tests {
		if got := countMaxOrSubsets(test.nums); got != test.expected {
			t.Errorf("For nums %v, expected %d but got %d", test.nums, test.expected, got)
		}
	}
}
