package heaps

type MinRangeHeap [][]int

func (h MinRangeHeap) Len() int { return len(h) }
func (h MinRangeHeap) Less(i, j int) bool {
	if h[i][1]-h[i][0] == h[j][1]-h[j][0] {
		return h[i][0] < h[j][0]
	}
	return h[i][1]-h[i][0] < h[j][1]-h[j][0]
}
func (h MinRangeHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinRangeHeap) Push(x interface{}) { *h = append(*h, x.([]int)) }
func (h *MinRangeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[:n-1]
	return item
}
