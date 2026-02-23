package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func maxProfit(prices []int) int {
	minPrice := math.MaxInt
	maxProfit := 0

	for _, price := range prices {
		if price < minPrice {
			minPrice = price
		}
		if price-minPrice > maxProfit {
			maxProfit = price - minPrice
		}
	}
	return maxProfit
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	var prices []int

	for scanner.Scan() {
		price, err := strconv.Atoi(scanner.Text())
		if err != nil {
			break
		}
		prices = append(prices, price)
	}

	fmt.Println(maxProfit(prices))
}
