package leetcode

import "math"

func Code_1189(text string) int {
	arr := [5]int{}
	for _, v := range text {
		switch v {
		case 'b':
			arr[0]++
		case 'a':
			arr[1]++
		case 'l':
			arr[2]++
		case 'o':
			arr[3]++
		case 'n':
			arr[4]++
		}
	}

	min := math.MaxUint32
	arr[2] >>= 1
	arr[3] >>= 1
	for _, v := range arr {
		if v < min {
			min = v
		}
	}
	return min
}
