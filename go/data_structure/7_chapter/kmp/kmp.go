package algorithm

func Match(str, pattern []byte) int {
	next := buildNext(pattern)
	m, n := len(str), len(pattern)
	i, j := 0, 0
	for j < n && i < m {
		if j < 0 || str[i] == pattern[j] {
			i++
			j++
		} else {
			j = next[j]
		}
	}
	if j == n {
		return i - j
	} else {
		return -1
	}
}

func buildNext(pattern []byte) []int {
	next := make([]int, len(pattern))
	next[0] = -1
	i, j := 0, -1
	m := len(pattern)
	for i < m-1 {
		if j < 0 || pattern[j] == pattern[i] {
			i++
			j++
			if pattern[i] == pattern[j] {
				next[i] = next[j]
			} else {
				next[i] = j
			}
		} else {
			j = next[j]
		}
	}
	return next
}
