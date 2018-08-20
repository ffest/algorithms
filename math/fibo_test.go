package math


import (
	"math/rand"
	"testing"
	"time"
)

func Test_FastFibo(t *testing.T) {
	cases := []struct {
		n      int64
		result int64
	}{
		{
			0,
			0,
		},
		{
			1,
			1,
		},
		{
			7,
			13,
		},
		{
			8,
			21,
		},
		{
		 15,
		 610,
		},
	}

	for i, c := range cases {
		res := FastFibo(c.n)

		if res != c.result {
			t.Errorf("Unexpected result. Case %d. Expected %d, got %d", i, c.result, res)
		}
	}
}

func Test_Fibo(t *testing.T) {
	cases := []struct {
		n      int64
		result int64
	}{
		{
			0,
			0,
		},
		{
			1,
			1,
		},
		{
			7,
			13,
		},
		{
			8,
			21,
		},
		{
			15,
			610,
		},
	}

	for i, c := range cases {
		res := Fibo(c.n)

		if res != c.result {
			t.Errorf("Unexpected result. Case %d. Expected %d, got %d", i, c.result, res)
		}
	}
}

func BenchmarkFastFibo(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < b.N; i++ {
		a := rand.Int63n(10000)
		_ = FastFibo(a)
	}
}

func BenchmarkFibo(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < b.N; i++ {
		a := rand.Int63n(40)
		_ = Fibo(a)
	}
}
