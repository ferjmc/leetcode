package max_k_elements

import "testing"

func TestMaxKelements(t *testing.T) {
	t.Parallel()

	tests := []struct {
		nums   []int
		k      int
		expect int64
	}{
		{[]int{10, 10, 10, 10, 10}, 5, 50},
		{[]int{1, 10, 3, 3, 3}, 3, 17},
	}

	for _, test := range tests {
		if got := maxKelements(test.nums, test.k); got != test.expect {
			t.Errorf("For nums %v and k %d, expected %d but got %d", test.nums, test.k, test.expect, got)
		}
	}
}
