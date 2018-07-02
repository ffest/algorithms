package search

// Binary search runs in at worst logarithmic time, making O(log n) comparisons
// Where n is element in the sorted array
func Binary(arr []int, n int) (result int) {
	left := 0
	right := len(arr)
	for {
		median := (left + right) / 2
		switch {
		case left == right:
			return -1
		case n < arr[median]:
			right = median
			continue
		case n > arr[median]:
			left = median + 1
			continue
		default:
			return median
		}
	}
}
