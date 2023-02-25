package probles

import "math"

func Code_Lcp_33(bucket []int, vat []int) int {
	l := len(bucket)
	maxk := 0
	for _, v := range vat {
		if v > maxk {
			maxk = v
		}
	}
	if maxk == 0 {
		return 0
	}

	res := math.MaxInt32
	for k := 1; k <= maxk; k++ {
		cur := k
		for i := 0; i < l; i++ {
			least := int(math.Ceil(float64(vat[i] / k)))
			if least-bucket[i] > 0 {
				cur += least - bucket[i]
			}
		}
		if cur < res {
			res = cur
		}
	}

	return res
}
