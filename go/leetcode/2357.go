package main

func minimumOperations(nums []int) int {
	dup := make(map[int]struct{})
	for _, v := range nums {
		dup[v] = struct{}{}
	}

	if _, ok := dup[0]; ok {
		return len(dup) - 1
	}

	return len(dup)
}
