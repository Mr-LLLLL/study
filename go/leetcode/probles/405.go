package probles

func toHex(num int) string {
	if num == 0 {
		return "0"
	}
	n := uint32(num)

	ans := make([]byte, 0)
	for n != 0 {
		v := byte(n & 15)
		if v < 10 {
			ans = append(ans, v+'0')
		} else {
			ans = append(ans, v-10+'a')
		}
		n >>= 4
	}
	l := len(ans)
	for i := 0; i < len(ans)/2; i++ {
		ans[i], ans[l-1-i] = ans[l-1-i], ans[i]
	}

	return string(ans)
}
