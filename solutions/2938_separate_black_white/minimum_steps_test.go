package minimum_steps

import "testing"

func TestMinimumSteps(t *testing.T) {
	t.Parallel()

	tests := []struct {
		s      string
		expect int64
	}{
		{"101", 1},
		{"100", 2},
		{"0111", 0},
		{"11000111", 6},
	}

	for _, test := range tests {
		if got := minimumSteps(test.s); got != test.expect {
			t.Errorf("For input %s, expected %d, but got %d", test.s, test.expect, got)
		}
	}
}
