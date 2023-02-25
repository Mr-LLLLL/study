package probles

func Code_1732(gain []int) int {
	maxHeight := 0
	cur := 0
	for _, v := range gain {
		cur += v
		if cur > maxHeight {
			maxHeight = cur
		}
	}
	return maxHeight
}
