package leetcode

func longestStrChain(words []string) int {
	return findMax("", 0, 0, 0, initBinArr(words))
}

func findMax(word string, layer, maximum, currentMax int, binArr [17][]string) int {
	if layer >= 17 {
		return maximum
	}

	arr := binArr[layer]
	if len(arr) == 0 {
		return findMax("", layer+1, maximum, 0, binArr)
	}

	for _, v := range arr {
		if isPre(word, v) {
			maximum = max(maximum, currentMax+1)
			tmpMax := findMax(v, layer+1, maximum, currentMax+1, binArr)
			maximum = max(maximum, tmpMax)
		} else {
			if maximum >= 17-layer {
				continue
			}
			tmpMax := findMax(v, layer+1, maximum, 1, binArr)
			maximum = max(maximum, tmpMax)
		}
	}

	return maximum
}

func initBinArr(arr []string) [17][]string {
	res := [17][]string{}
	for _, v := range arr {
		res[len(v)] = append(res[len(v)], v)
	}

	return res
}

func isPre(w1, w2 string) bool {
	if w1 == "" {
		return true
	}

	diffCnt := false
	for i, j := 0, 0; i < len(w1); {
		if w1[i] != w2[j] {
			if diffCnt {
				return false
			}
			diffCnt = true
			j++
			continue
		}

		i++
		j++
	}

	return true
}
