package main

import "fmt"

//Say you have an array for which the ith element is the price of a given stock on day i.

//If you were only permitted to complete at most one transaction (i.e., buy one and sell one share of the stock), design an algorithm to find the maximum profit.

//Note that you cannot sell a stock before you buy one.
//Example 1:
//
//Input: [7,1,5,3,6,4]
//Output: 5
//Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
//             Not 7-1 = 6, as selling price needs to be larger than buying price.
//Example 2:
//
//Input: [7,6,4,3,1]
//Output: 0
//Explanation: In this case, no transaction is done, i.e. max profit = 0.

//定义：min[i]表示第i天前最低股票价格，i>=1（i从0开始）
//初始化：min[0]=prices[0]，便于边界处理
//递推表达式：min[i]=MIN{min[i-1],prices[i]}，i>=1（i从0开始）
//求解：遍历求解每一天出售可获取的最大利润，最后返回最大值利润
func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	//prices := []int{7, 6, 4, 3, 1}
	max := bb(prices)
	fmt.Printf("max profit: %d\n", max)
}

func bb(p []int) int {
	if len(p) < 2 {
		return 0
	}
	fmt.Println(p)
	min := make([]int, len(p))
	min[0] = p[0]
	for i := 1; i < len(p); i++ {
		if p[i] < min[i-1] {
			min[i] = p[i]
		} else {
			min[i] = min[i-1]
		}
	}
	fmt.Printf("min:%d\n", min)

	//max := make([]int, len(p) - 1)
	max := 0

	for i := 1; i < len(p); i++ {
		if p[i]-min[i] > max {
			max = p[i] - min[i]
		}

	}
	return max
}
