package longest_diverse_string

import "testing"

func TestLongestDiverseString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		a      int
		b      int
		c      int
		expect string
	}{
		{1, 1, 7, "ccaccbcc"},
		{7, 1, 0, "aabaa"},
		{4, 4, 3, "aabbccaabbc"},
		{0, 8, 11, "ccbccbbccbbccbbccbc"},
	}

	for _, test := range tests {
		if got := longestDiverseString3(test.a, test.b, test.c); got != test.expect {
			t.Errorf("For a=%d, b=%d, c=%d - Expected: %s, but got: %s", test.a, test.b, test.c, test.expect, got)
		}
	}
}
