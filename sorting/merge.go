package sorting

// Mergesort is an efficient, general-purpose, comparison-based sorting algorithm.
// The algorithm takes O(n log n) comparisons to sort n items.
// TODO: make it faster
func Merge(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	middle := len(arr) / 2

	left := Merge(arr[:middle])
	right := Merge(arr[middle:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	res := make([]int, 0, len(left)+len(right))

	for len(left) > 0 || len(right) > 0 {
		if len(left) == 0 {
			return append(res, right...)
		}
		if len(right) == 0 {
			return append(res, left...)
		}
		if left[0] <= right[0] {
			res = append(res, left[0])
			left = left[1:]
		} else {
			res = append(res, right[0])
			right = right[1:]
		}
	}

	return res
}
