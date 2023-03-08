package leetcode

func canThreePartsEqualSum(arr []int) bool {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	if sum%3 != 0 {
		return false
	}
	avg := sum / 3
	sum = 0
	cnt := 0
	for _, v := range arr {
		sum += v
		if sum == avg {
			sum = 0
			cnt++
		}
	}
	return cnt >= 3
}
