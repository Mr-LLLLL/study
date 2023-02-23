package main

func numberOfSubarrays(nums []int, k int) int {
	pos := []int{-1}
	res := 0
	for i, v := range nums {
		if v&1 == 1 {
			pos = append(pos, i)
		}
		l := len(pos)
		if l > k {
			res += pos[l-k] - pos[l-k-1]
		}
	}

	return res
}
