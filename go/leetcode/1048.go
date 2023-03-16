package leetcode

import (
	"sort"
)

func longestStrChain(words []string) int {
	m := make(map[string]int)
	maximum := 0

	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})

	for _, v := range words {
		for i := 0; i < len(v); i++ {
			pre := v[0:i] + v[i+1:]
			m[v] = max(m[pre]+1, m[v])
		}
		maximum = max(m[v], maximum)
	}

	return maximum
}
