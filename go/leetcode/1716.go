package main

func Code_1716(n int) int {
	d := n / 7
	m := n % 7
	return 28*d + 7*d*(d-1)/2 + d*m + m*(m+1)/2
}
