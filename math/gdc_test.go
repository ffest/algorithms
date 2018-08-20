package math

import (
	"math/rand"
	"testing"
	"time"
)

func Test_Gdc(t *testing.T) {
	cases := []struct {
		a      int
		b      int
		result int
	}{
		{
			10,
			5,
			5,
		},
		{
			31,
			11,
			1,
		},
		{
			160,
			50,
			10,
		},
		{
			256,
			128,
			128,
		},
		{
			67,
			32,
			1,
		},
		{
			91,
			49,
			7,
		},
	}

	for i, c := range cases {
		res := Gdc(c.a, c.b)

		if res != c.result {
			t.Errorf("Unexpected result. Case %d. Expected %d, got %d", i, c.result, res)
		}
	}
}

func BenchmarkGdc(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	a := rand.Intn(10000)
	c := rand.Intn(10000)
	for i := 0; i < b.N; i++ {
		_ = Gdc(a, c)
	}
}
