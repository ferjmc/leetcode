package min_groups

import "testing"

func TestMinGroups(t *testing.T) {
	t.Parallel()

	tests := []struct {
		intervals [][]int
		expect    int
	}{
		{[][]int{{5, 10}, {6, 8}, {1, 5}, {2, 3}, {1, 10}}, 3},
		{[][]int{{1, 3}, {5, 6}, {8, 10}, {11, 13}}, 1},
		{[][]int{{159431, 428743}, {614908, 651142}, {431031, 806494}}, 2},
	}

	for _, test := range tests {
		if got := minGroups(test.intervals); got != test.expect {
			t.Errorf("minGroups(%v) = %d, want %d", test.intervals, got, test.expect)
		}
	}
}
