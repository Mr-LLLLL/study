package main

func Code_Lcp_6(coins []int) int {
	res := 0
	for _, v := range coins {
		res += v/2 + v&1
	}
	return res
}
