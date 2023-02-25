package leetcode

func Code_832(image [][]int) [][]int {
	for i := 0; i < len(image); i++ {
		left, right := 0, len(image[i])-1
		for left < right {
			if image[i][left] == image[i][right] {
				image[i][left] ^= 1
				image[i][right] ^= 1
			}
			left++
			right--
		}
		if left == right {
			image[i][left] ^= 1
		}
	}
	return image
}
