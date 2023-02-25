package leetcode

func Code_1869(s string) bool {
	oneNum, zeroNum := 0, 0
	pre := '#'
	cnt := 1

	for _, c := range s {
		if c == pre {
			cnt++
		} else {
			cnt = 1
		}
		pre = c

		if c == '1' {
			if cnt > oneNum {
				oneNum = cnt
			}
		} else if c == '0' {
			if cnt > zeroNum {
				zeroNum = cnt
			}
		}
	}

	return oneNum > zeroNum
}
