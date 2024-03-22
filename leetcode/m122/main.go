package main

import "fmt"

// 122. Best Time to Buy and Sell Stock II
// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii

func maxProfit(prices []int) int {
	profit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			profit += prices[i] - prices[i-1]
			continue
		}
	}

	return profit
}

func maxProfit2(prices []int) int {
	buyPrice := prices[0]
	maximumProfit := 0
	for day := 1; day < len(prices); day++ {
		if buyPrice > prices[day] {
			buyPrice = prices[day]
		}

		profit := prices[day] - buyPrice
		if profit > 0 {
			maximumProfit += profit
			buyPrice = prices[day]
		}
	}

	return maximumProfit
}

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(maxProfit2([]int{7, 1, 5, 3, 6, 4}))
}
