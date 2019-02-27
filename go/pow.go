func myPow(x float64, n int) float64 {
	if n == 0 || x == 1 {
		return 1
	}
	if n < 0 {
		n = -n
		x = 1 / x
	}
    p := myPow(x, n / 2)
    p = p * p
    if (n & 1) != 0 {
        p = x * p
    }
    return p
}