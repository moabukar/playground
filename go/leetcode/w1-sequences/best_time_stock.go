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
	// if the length of the prices is 0, return 0
	if len(prices) == 0 {
		return 0
	}
	// make a variable to store the profit
	var profit int
	// make a variable to store the minimum price
	min := prices[0]
	// loop through the prices
	for i := 1; i < len(prices); i++ {
		// if the current price is less than the minimum price, set the minimum price to the current price
		if prices[i] < min {
			min = prices[i]
		} else {
			// else, set the profit to the difference between the current price and the minimum price
			tmp := prices[i] - min
			// if the profit is greater than the current profit, set the profit to the current profit
			if tmp > profit {
				profit = tmp
			}
		}
	}
	// return the profit
	return profit
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	mP := maxProfit(prices)
	fmt.Println(mP)
}
