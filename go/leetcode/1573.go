package main

var MODULO = 1000000007

type test struct{}

func Code_1573(s string) int {
	ones := getOneIndex(s)

	l := len(ones)

	if l%3 != 0 {
		return 0
	}

	if l == 0 {
		res := (len(s) - 1) * (len(s) - 2) / 2
		return res % MODULO
	}

	index1, index2 := l/3, l/3*2
	count1 := ones[index1] - ones[index1-1]
	count2 := ones[index2] - ones[index2-1]

	res := count1 * count2 % MODULO

	return res
}

func getOneIndex(s string) []int {
	res := make([]int, 0)
	for i, v := range s {
		if v == '1' {
			res = append(res, i)
		}
	}

	return res
}
