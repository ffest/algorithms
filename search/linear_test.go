package search

import (
	"testing"
)

func Test_Linear(t *testing.T) {
	cases := []struct {
		arr    []int
		find   int
		result int
	}{
		{
			[]int{1, 3, 5, 6, 7, 9, 11, 13, 14, 15, 16, 17, 22},
			7,
			4,
		},
		{
			[]int{-54, -23, -7, 0, 1, 3, 9, 11, 13, 14, 15, 16, 17, 22},
			-7,
			2,
		},
	}

	for i, c := range cases {
		res := Binary(c.arr, c.find)

		if res != c.result {
			t.Errorf("Unexpected result. Case %d. Expected %d, got %d", i, c.result, res)
		}
	}
}

func BenchmarkLinear(b *testing.B) {
	arr := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		arr[i] = i
	}
	for i := 0; i < b.N; i++ {
		_ = Linear(arr, i)
	}
}
