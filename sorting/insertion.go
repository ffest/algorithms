package sorting

// Insertion sort is a simple sorting algorithm that builds the final sorted array (or list) one item at a time.
// Has a worst-case and average complexity of О(n2), but efficient for (quite) small data sets
func Insertion(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i
		for j > 0 && arr[j-1] > key {
			arr[j] = arr[j-1]
			j--
		}
		arr[j] = key
	}

	return arr
}
