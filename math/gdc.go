package math

func Gdc(a int, b int) int {
	if b != 0 {
		return Gdc(b, a%b)
	}

	return a
}
