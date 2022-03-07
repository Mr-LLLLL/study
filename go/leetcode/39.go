package main

func Code_39(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	backtrace := make([]int, 0)
	sum := 0
	var dfs func(int)

	dfs = func(index int) {
		if sum == target {
			tmp := make([]int, 0, len(backtrace))
			for _, v := range backtrace {
				tmp = append(tmp, v)
			}
			res = append(res, tmp)
		} else if sum > target {
			return
		} else {
			for i := index; i < len(candidates); i++ {
				backtrace = append(backtrace, candidates[i])
				sum += candidates[i]
				dfs(i)
				sum -= candidates[i]
				backtrace = backtrace[:len(backtrace)-1]
			}
		}
	}
	dfs(0)

	return res
}
