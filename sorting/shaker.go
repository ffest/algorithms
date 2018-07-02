package sorting

// Shaker sort (cocktail sort, shake sort) is a stable sorting algorithm with quadratic asymptotic complexity О(n2).
// Shaker sort is a bidirectional version of bubble sort.
func Shaker(arr []int) []int {
	left := 0
	right := len(arr) - 1
	for left <= right {
		for i := left; i < right; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
			}
		}
		right--

		for i := right; i > left; i-- {
			if arr[i-1] > arr[i] {
				arr[i-1], arr[i] = arr[i], arr[i-1]
			}
		}
		left++
	}

	return arr
}
