package main

import "fmt"

/*
You are given an integer array prices where prices[i] is the price of a given stock on the ith day.

On each day, you may decide to buy and/or sell the stock.
You can only hold at most one share of the stock at any time. However, you can buy it then immediately sell it on the same day.

Find and return the maximum profit you can achieve.


Example 1:
----------
Input: prices = [7,1,5,3,6,4]
Output: 7

Explanation: Buy on day 2 (price = 1) and sell on day 3 (price = 5), profit = 5-1 = 4.
Then buy on day 4 (price = 3) and sell on day 5 (price = 6), profit = 6-3 = 3.
Total profit is 4 + 3 = 7.


Example 2:
----------
Input: prices = [1,2,3,4,5]
Output: 4

Explanation: Buy on day 1 (price = 1) and sell on day 5 (price = 5), profit = 5-1 = 4.
Total profit is 4.


Example 3:
----------
Input: prices = [7,6,4,3,1]
Output: 0

Explanation: There is no way to make a positive profit, so we never buy the stock to achieve the maximum profit of 0.


Constraints:

    1 <= prices.length <= 3 * 10^4
    0 <= prices[i] <= 10^4

*/

func maxProfit(prices []int) int {
	profit := 0

	// i reflects the "current" day
	i := 0
	for i < len(prices) {
		// we use k and i to look ahead in "time" and find the day with the biggest sale difference (profit)

		k := i + 1
		// iterate prices[] until we find an upward price pattern or an inflection point
		// our entry point is the first instance where price today  < price tomorrow
		for k < len(prices) {
			if prices[i] < prices[k] {
				break
			} else {
				i = k
				k++
			}
		}
		// the case k==len(prices) (i.e. the first loop ran to completion and didn't break)
		// indicates that the price never went up since day 0
		// liquidation imminent, there is no profit to be made here. don't tell the wife.
		if k >= len(prices) {
			break
		}

		// otherwise, continue searching until we find the future price with the biggest difference
		for k+1 < len(prices) {
			if prices[k] > prices[k+1] {
				break
			}
			k++
		}

		fmt.Printf("evaluating day:%d vs day:%d:  %d > %d ? %t; profit: %d\n", k, i, prices[k], prices[i], prices[k] > prices[i], prices[k]-prices[i])
		if prices[k] > prices[i] {
			profit += prices[k] - prices[i]
			// we wouldn't buy a stock on the same day we sold it, otherwise why'd we sell it to start with?
			// we scan for the next good entry price starting on the day after our sale
			i = k + 1
		} else {
			i++
		}
	}
	return profit
}

func main() {
	prices := []int{6, 1, 3, 2, 4, 7}
	profit := maxProfit(prices)
	fmt.Printf("%v => %d\n", prices, profit)
}
