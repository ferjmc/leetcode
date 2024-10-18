package maximum_swap

import "testing"

func TestMaximumSwap(t *testing.T) {
	t.Parallel()

	tests := []struct {
		num    int
		expect int
	}{
		{2736, 7236},
		{9973, 9973},
		{98368, 98863},
		{99901, 99910},
	}
	for _, test := range tests {
		if got := maximumSwap2(test.num); got != test.expect {
			t.Errorf("For num %d, expected %d but got %d", test.num, test.expect, got)
		}
	}
}
