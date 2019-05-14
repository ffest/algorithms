package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	for m > 0 && n > 0 {
		if nums1[m-1] >= nums2[n-1] {
			nums1[m+n-1] = nums1[m-1]
			m--
		} else {
			nums1[m+n-1] = nums2[n-1]
			n--
		}
	}
	for n > 0 {
		nums1[n-1] = nums2[n-1]
		n--
	}
}

func main() {
	m := 3
	n := 3
	nums1 := []int{3, 3, 8, 0, 0, 0}
	nums2 := []int{1, 2, 6}
	merge(nums1, m, nums2, n)
}
