package leetcode

func Code_1769(arr string) []int {
	firstNum, leftOneCnt, rightOneCnt := getFirstCntAndLeftRightOneCnt(arr)
	sum := getSumNum(arr, firstNum, leftOneCnt, rightOneCnt)

	return sum
}

func getFirstCntAndLeftRightOneCnt(arr string) (int, []int, []int) {
	firstCnt := 0
	leftNum := make([]int, len(arr))
	rightNum := make([]int, len(arr))
	oneCnt := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] == '1' {
			firstCnt += i
		}

		leftNum[i] = oneCnt
		if arr[i] == '1' {
			oneCnt++
		}
	}

	oneCnt = 0
	for i := len(arr) - 1; i >= 0; i-- {
		rightNum[i] = oneCnt
		if arr[i] == '1' {
			oneCnt++
		}
	}

	return firstCnt, leftNum, rightNum
}

func getSumNum(arr string, firstNum int, leftOneCnt, rightOneCnt []int) []int {
	res := make([]int, len(arr))
	res[0] = firstNum
	for i := 1; i < len(arr); i++ {
		res[i] = res[i-1] + leftOneCnt[i] - rightOneCnt[i-1]
	}

	return res
}
