package smalles_range

import (
	"reflect"
	"testing"
)

func TestSmallestRange(t *testing.T) {
	tests := []struct {
		nums     [][]int
		expected []int
	}{
		{[][]int{{4, 10, 15, 24, 26}, {0, 9, 12, 20}, {5, 18, 22, 30}}, []int{20, 24}},
		{[][]int{{1, 2, 3}, {1, 2, 3}, {1, 2, 3}}, []int{1, 1}},
	}

	for _, test := range tests {
		if got := smallestRange(test.nums); !reflect.DeepEqual(got, test.expected) {
			t.Errorf("Expected %v, but got %v", test.expected, got)
		}
	}
}
