package probles

func myPow(x float64, n int) float64 {
	quickMul := func(x float64, n int) float64 {
		ans := 1.0
		for n != 0 {
			if n&1 == 1 {
				ans *= x
			}
			x *= x
			n >>= 1
		}
		return ans
	}
	if n >= 0 {
		return quickMul(x, n)
	}
	return 1 / quickMul(x, -n)
}
