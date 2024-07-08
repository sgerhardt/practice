package greedy_algorithms

import "errors"

//
//func getMaxProfit(stockPrices []int) (int, error) {
//	if len(stockPrices) < 2 {
//		return 0, errors.New("need at least two indexes")
//	}
//	minPrice := stockPrices[0]
//	maxProfit := stockPrices[1] - stockPrices[0]
//
//	for i := 1; i < len(stockPrices); i++ {
//		profit := stockPrices[i] - minPrice
//		if profit > maxProfit {
//			maxProfit = profit
//		}
//
//		if stockPrices[i] < minPrice {
//			minPrice = stockPrices[i]
//		}
//	}
//	return maxProfit, nil
//}

func getMaxProfit(stockPrices []int) (int, error) {
	if len(stockPrices) < 2 {
		return 0, errors.New("not enough indices")
	}

	// need to buy before you can sell...

	// first possible buy is at index 0, so set that to min to start
	minPrice := stockPrices[0]

	// first possible sale is at second index, so set max profit at first minus second index to start
	maxProfit := stockPrices[1] - stockPrices[0]

	for i := 1; i < len(stockPrices); i++ {
		profit := stockPrices[i] - minPrice
		if maxProfit < profit {
			maxProfit = profit
		}
		if stockPrices[i] < minPrice {
			minPrice = stockPrices[i]
		}
	}

	return maxProfit, nil
}
