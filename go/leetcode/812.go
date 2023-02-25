package leetcode

import (
	"math"
)

func Code_812(points [][]int) float64 {
	s := 0.0

	l := len(points)

	for i := 0; i < l-2; i++ {
		for j := i + 1; j < l-1; j++ {
			x1 := points[j][0] - points[i][0]
			y1 := points[j][1] - points[i][1]
			for k := j + 1; k < l; k++ {
				x2 := points[k][0] - points[i][0]
				y2 := points[k][1] - points[i][1]

				s = math.Max(s, math.Abs(float64(x1*y2-x2*y1)))
			}
		}
	}

	return s / 2
}
