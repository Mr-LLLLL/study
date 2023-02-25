package probles

func Code_447(points [][]int) int {
	res := 0
	l := len(points)
	for i := 0; i < l; i++ {
		m := make(map[int]int)
		for j := 0; j < l; j++ {
			xLen := points[j][0] - points[i][0]
			yLen := points[j][1] - points[i][1]
			length := xLen*xLen + yLen*yLen
			if length < 0 {
				length = -length
			}
			m[length]++
		}
		delete(m, 0)
		for _, v := range m {
			res += v * (v - 1)
		}
	}

	return res
}
