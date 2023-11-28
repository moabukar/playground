package main

import "fmt"

// You are given an array prices where prices[i] is the price of a given stock on the ith day.

// You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.

// Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.

// Example 1:

// Input: prices = [7,1,5,3,6,4]
// Output: 5
// Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
// Note that buying on day 2 and selling on day 1 is not allowed because you must buy before you sell

func maxProfit(prices []int) int {
	// make a variable to store the min price
	if len(prices) == 0 {
		return 0
	}

	// make a variable to store the max profit
	minPrice := prices[0]
	maxProfit := 0

	// loop through the prices
	for _, price := range prices {
		// if the current price is less than the min price, set the min price to the current price
		if price < minPrice {
			minPrice = price
		}
		// else if the current price minus the min price is greater than the max profit, set the max profit to the current price minus the min price
		if price-minPrice > maxProfit {
			maxProfit = price - minPrice
		}
	}
	// return the max profit
	return maxProfit
}

func main() {
	prices := []int{0, 1, 5, 3, -5, 15}
	fmt.Println(maxProfit(prices))
	// mP := maxProfit(prices)
	// fmt.Println(mP)
}
