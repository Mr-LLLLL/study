package leetcode

func lemonadeChange(bills []int) bool {
	five := 0
	ten := 0

	for _, v := range bills {
		switch v {
		case 5:
			five++
		case 10:
			if five > 0 {
				five--
				ten++
			} else {
				return false
			}
		default:
			if five > 0 && ten > 0 {
				five--
				ten--
			} else if five > 2 {
				five -= 3
			} else {
				return false
			}
		}
	}
	return true
}
