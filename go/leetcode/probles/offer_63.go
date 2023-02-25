package probles

import "math"

func Code_Offer_63(prices []int) int {
	minBuy := math.MaxInt32
	maxProfit := 0
	for _, v := range prices {
		if v < minBuy {
			minBuy = v
		}
		if v-minBuy > maxProfit {
			maxProfit = v - minBuy
		}
	}

	return maxProfit
}
