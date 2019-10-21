package math

import (
	"math/rand"
	"testing"
	"time"
)

func Test_Rand(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	cases := []struct {
		a int
		b int
	}{
		{
			5,
			10,
		},
		{
			11,
			31,
		},
	}

	for i, c := range cases {
		res := Rand(c.a, c.b)
		if res < c.a || res > c.b {
			t.Errorf("Unexpected result. Case %d. Expected [%d,%d], got %d", i, c.a, c.b, res)
		}
	}
}

func BenchmarkRand(b *testing.B) {
	a := rand.Intn(10000)
	for i := 0; i < b.N; i++ {
		_ = Rand(a, a+rand.Intn(10000))
	}
}
