package leetcode

func hasGroupsSizeX(deck []int) bool {
	m := make(map[int]int, len(deck))
	var g int
	for _, v := range deck {
		m[v]++
		g = m[v]
	}

	for _, v := range m {
		g = gcd(g, v)
	}

	return g >= 2
}

func gcd(i, j int) int {
	if i == 0 {
		return j
	}
	return gcd(j%i, i)
}
