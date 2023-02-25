package probles

func maxDepth(s string) int {
	l, lMax := 0, 0
	ans := 0
	for _, v := range s {
		if v == '(' {
			l++
			if l > lMax {
				lMax = l
			}
		} else if v == ')' && l > 0 {
			l--
			if lMax-l > ans {
				ans = lMax - l
			}
		}
	}

	return ans
}
