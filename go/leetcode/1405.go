package leetcode

import "sort"

func longestDiverseString(a int, b int, c int) string {
	ans := []byte{}
	cnt := []struct {
		c    int
		char byte
	}{
		{a, 'a'},
		{b, 'b'},
		{c, 'c'},
	}
	for {
		sort.Slice(cnt, func(i, j int) bool {
			return cnt[i].c > cnt[j].c
		})
		hasNext := false
		for i, p := range cnt {
			if p.c == 0 {
				break
			}
			m := len(ans)
			if m >= 2 && ans[m-2] == p.char && ans[m-1] == p.char {
				continue
			}
			hasNext = true
			ans = append(ans, p.char)
			cnt[i].c--
			break
		}
		if !hasNext {
			return string(ans)
		}
	}
}
