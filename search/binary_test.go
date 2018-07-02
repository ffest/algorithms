package search

import (
	"testing"
)

func Test_Binary(t *testing.T) {
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
	}

	for i, c := range cases {
		res := Binary(c.arr, c.find)

		if res != c.result {
			t.Errorf("Unexpected result. Case %d. Expected %d, got %d", i, c.result, res)
		}
	}
}

func BenchmarkBinary(b *testing.B) {
	arr := []int{1, 3, 5, 6, 7, 9, 11, 13, 14, 15, 16, 17, 22}
	for i := 0; i < b.N; i++ {
		_ = Binary(arr, 16)
	}
}
