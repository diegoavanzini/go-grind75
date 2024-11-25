package besttimetobuyandsellstock

func maxProfit(prices []int) int {
	var currentBuy = prices[0]
	var maxProfit int
	for _, currentPrice := range prices {
		if currentPrice < currentBuy {
			// if the current price is lower than the current buy price
			// the new buy price becomne the current (cheaper) price
			currentBuy = currentPrice
		} else { // the new price is not the lowest
			currentProfit := currentPrice - currentBuy
			// if the profit calculated with the current price
			// and the current purchase price is greather than the
			// last maximum profit
			if currentProfit > maxProfit {
				// the new maximum profit is
				// the new calculated profit
				maxProfit = currentProfit
			}
		}
	}
	return maxProfit
}
