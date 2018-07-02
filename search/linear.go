package search

// Lenear search is the simplest search algorithm, making O(log n) comparisons
// Where n is element in the array
func Linear(arr []int, n int) int {
	for i, element := range arr {
		if element == n {
			return i
		}
	}

	return -1
}
