package probles

func Code_647(s string) int {
	t := "$#"
	for _, v := range s {
		t += string(v) + "#"
	}

	n := len(t)
	t += "!"

	iMax, rMax, ans := 0, 0, 0
	f := make([]int, n)
	for i := 1; i < n; i++ {
		if i <= rMax {
			f[i] = min(rMax-i+1, f[2*iMax-i])
		} else {
			f[i] = 1
		}

		for t[i+f[i]] == t[i-f[i]] {
			f[i]++
		}

		if i+f[i]-1 > rMax {
			iMax = i
			rMax = i + f[i] - 1
		}
		ans += f[i] / 2
	}

	return ans
}

// manacher
// time:O(n) space:O(1)
