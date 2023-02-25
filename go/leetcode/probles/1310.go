package probles

func xorQueries(arr []int, queries [][]int) []int {
	xors := make([]int, len(arr)+1)
	for i, v := range arr {
		xors[i+1] = xors[i] ^ v
	}

	ans := make([]int, 0, len(queries))
	for _, v := range queries {
		ans = append(ans, xors[v[0]]^xors[v[1]+1])
	}

	return ans
}
