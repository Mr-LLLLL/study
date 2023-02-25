package probles

func Code_Offer_62(n int, m int) int {
	return GetN(n, m)
}

func GetN(n, m int) int {
	if n == 1 {
		return 0
	}
	x := GetN(n-1, m)
	return (x + m) % n
}

func iterator(n, m int) int {
	f := 0
	for i := 2; i < n+1; i++ {
		f = (m + f) % i
	}
	return f
}
