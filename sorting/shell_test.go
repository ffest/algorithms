package sorting

import (
	"math/rand"
	"testing"
)

func Test_Shell(t *testing.T) {
	cases := []struct {
		arr []int
		res []int
	}{
		{
			[]int{3, 5, 1, 34, 5, 7, 2, 5, 7, 8, 9, 1, 3, 10, 43},
			[]int{1, 1, 2, 3, 3, 5, 5, 5, 7, 7, 8, 9, 10, 34, 43},
		},
		{
			[]int{5, -3, -54, 7, 419, 7, 124, 0, 0, 0, 9, -9, 3, 10, 1},
			[]int{-54, -9, -3, 0, 0, 0, 1, 3, 5, 7, 7, 9, 10, 124, 419},
		},
		{
			[]int{0, 0, 0, 0, 0, 0},
			[]int{0, 0, 0, 0, 0, 0},
		},
	}

	for i, c := range cases {
		res := Shell(c.arr)

		if len(res) != len(c.res) {
			t.Errorf("Unexpected length of slice. Case %d. Expected %d, got %d", i, len(c.res), len(res))
		}

		for j := range c.res {
			if res[j] != c.res[j] {
				t.Errorf("Unexpected element of slice. Case %d. Expected %d, got %d", i, c.res[j], res[j])
			}
		}
	}
}

func BenchmarkShell100(b *testing.B) {
	arr := make([]int, 100)
	for i := 0; i < 100; i++ {
		arr[i] = rand.Intn(1000)
	}
	for i := 0; i < b.N; i++ {
		_ = Shell(arr)
	}
}

func BenchmarkShell10000(b *testing.B) {
	arr := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		arr[i] = rand.Intn(10000)
	}
	for i := 0; i < b.N; i++ {
		_ = Shell(arr)
	}
}
