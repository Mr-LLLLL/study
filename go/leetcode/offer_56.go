package leetcode

func Code_Offer_56(nums []int) []int {
	// return Code_Offer_56_v1(nums)
	return Code_Offer_56_v2(nums)
}

// time:O(n)
// space:O(n)
func Code_Offer_56_v1(nums []int) []int {
	m := make(map[int]struct{})

	for _, v := range nums {
		if _, ok := m[v]; ok {
			delete(m, v)
		} else {
			m[v] = struct{}{}
		}
	}

	res := make([]int, 0, 2)
	for k := range m {
		res = append(res, k)
	}

	return res
}

// time:O(n)
// space:O(1)
func Code_Offer_56_v2(nums []int) []int {
	res := 0
	for _, v := range nums {
		res ^= v
	}
	div := 1
	for div&res == 0 {
		div <<= 1
	}

	a, b := 0, 0
	for _, v := range nums {
		if v&div == div {
			a ^= v
		} else {
			b ^= v
		}
	}

	return []int{a, b}
}
