package sorting

// Shellsort is a generalization of insertion sort that allows the exchange of items that are far apart.
// The idea is to arrange the list of elements so that, starting anywhere,
// considering every hth element gives a sorted list. Such a list is said to be h-sorted.
// In the worst case, it makes O(n2) comparisons
func Shell(arr []int) []int {
	for d := int(len(arr)/2); d > 0; d /= 2 {
		for i := d; i < len(arr); i++ {
			for j := i; j >= d && arr[j-d] > arr[j]; j -= d {
				arr[j], arr[j-d] = arr[j-d], arr[j]
			}
		}
	}

	return arr
}