package majority

func Majority(arr []int) bool {
	maj := majEleCandidate(arr)
	return majEleCheck(arr, maj)
}

func majEleCheck(arr []int, maj int) bool {
	cnt := 0
	for _, v := range arr {
		if v == maj {
			cnt++
		}
	}
	return cnt*2 > len(arr)
}

func majEleCandidate(arr []int) int {
	var maj int
	cnt := 0
	for _, v := range arr {
		if cnt == 0 {
			maj = v
		} else {
			if maj == v {
				cnt++
			} else {
				cnt--
			}
		}
	}
	return maj
}
