package probles

func Code_1424(nums [][]int) []int {
	return Code_1424_v2(nums)
}

func Code_1424_v2(nums [][]int) []int {
	arr := make([][]int, 0)
	res := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums[i]); j++ {
			if i+j == len(arr) {
				arr = append(arr, make([]int, 0))
			}
			arr[i+j] = append(arr[i+j], nums[i][j])
		}
	}

	for i := 0; i < len(arr); i++ {
		for j := len(arr[i]) - 1; j >= 0; j-- {
			res = append(res, arr[i][j])
		}
	}

	return res
}

func Code_1424_v1(nums [][]int) []int {
	res := make([]int, 0)
	i := 0
	for {
		j := len(nums) - 1
		sign := true
		for {
			if i >= j {
				if i < len(nums[j])+j {
					res = append(res, nums[j][i-j])
					sign = false
				}
			}

			if j == 0 {
				break
			}

			j--
		}
		i++

		if sign {
			break
		}
	}

	return res
}
