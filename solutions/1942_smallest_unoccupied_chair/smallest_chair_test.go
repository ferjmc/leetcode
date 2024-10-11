package smallest

import "testing"

func TestSmallestChair(t *testing.T) {
	t.Parallel()

	tests := []struct {
		times        [][]int
		targetFriend int
		expected     int
	}{
		{[][]int{{1, 4}, {2, 3}, {4, 6}}, 1, 1},
		{[][]int{{3, 10}, {1, 5}, {2, 6}}, 0, 2},
		{[][]int{{33889, 98676}, {80071, 89737}, {44118, 52565}, {52992, 84310}, {78492, 88209}, {21695, 67063}, {84622, 95452}, {98048, 98856}, {98411, 99433}, {55333, 56548}, {65375, 88566}, {55011, 62821}, {48548, 48656}, {87396, 94825}, {55273, 81868}, {75629, 91467}},
			6, 2,
		},
		{[][]int{{4, 5}, {12, 13}, {5, 6}, {1, 2}, {8, 9}, {9, 10}, {6, 7}, {3, 4}, {7, 8}, {13, 14}, {15, 16}, {14, 15}, {10, 11}, {11, 12}, {2, 3}, {16, 17}},
			15,
			0,
		},
	}
	for _, tt := range tests {
		if got := smallestChair(tt.times, tt.targetFriend); got != tt.expected {
			t.Errorf("with times %v and target %v , smallestChair() = %v, want %v", tt.times, tt.targetFriend, got, tt.expected)
		}
	}
}
