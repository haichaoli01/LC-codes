package ltcode

import (
	"math"
)

// 解法一 DP
func maxProfit309(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	buy, sell := make([]int, len(prices)), make([]int, len(prices))
	for i := range buy {
		buy[i] = math.MinInt64
	}
	buy[0] = -prices[0]
	buy[1] = max(buy[0], -prices[1])
	sell[1] = max(sell[0], buy[0]+prices[1])
	for i := 2; i < len(prices); i++ {
		sell[i] = max(sell[i-1], buy[i-1]+prices[i])
		buy[i] = max(buy[i-1], sell[i-2]-prices[i])
	}
	return sell[len(sell)-1]
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

// 解法二 优化辅助空间的 DP
func maxProfit309_1(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	buy := []int{-prices[0], max(-prices[0], -prices[1]), math.MinInt64}
	sell := []int{0, max(0, -prices[0]+prices[1]), 0}

	for i := 2; i < len(prices); i++ {
		sell[i%3] = max(sell[(i-1)%3], buy[(i-1)%3]+prices[i])
		buy[i%3] = max(buy[(i-1)%3], sell[(i-2)%3]-prices[i])
	}
	return sell[(len(prices)-1)%3]
}

func maxProfit3(prices []int) int {
	n := len(prices)
	dp := make([][3]int, n)
	if n == 0 {
		return 0
	}
	//没有股票不处于冷冻期
	dp[0][0] = 0
	//没有股票处于冷冻期
	dp[0][1] = 0
	//有股票
	dp[0][2] = -prices[0]

	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1])
// err dp[i][1] = dp[i-1][2] + prices[i-1]
		dp[i][1] = dp[i-1][2] + prices[i]
		dp[i][2] = max(dp[i-1][0]-prices[i], dp[i-1][2])
	}
	return max(dp[n-1][0], dp[n-1][1])
}

