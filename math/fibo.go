package math

func FastFibo(n int64) int64 {
	if n <= 1 {
		return n
	}
	a, _ := fibo(n)
	return a
}

func fibo(n int64) (int64, int64) {
	if n == 0 {
		return 0, 1
	}

	a, b := fibo(n / 2)
	c := a * (b * 2 - a)
	d := a * a + b * b
	if n % 2 == 0{
		return c, d
	} else {
		return d, c + d
	}
}

func Fibo(n int64) int64{
	if n <= 1 {
		return n
	}
	return Fibo(n -1) + Fibo(n - 2)
}

