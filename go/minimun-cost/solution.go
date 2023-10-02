package minimuncost

import (
	"sort"
)

func getMinimumCost(k int32, c []int32) int32 {
	sort.Slice(c, func(i, j int) bool {
		return c[i] < c[j]
	})

	var flowerPrices []int32
	multiplier := 0
	currentFlower := len(c) - 1
	currentFriendBuying := 0
	for len(flowerPrices) < len(c) {
		price := c[currentFlower] * (int32(multiplier) + 1)
		flowerPrices = append(flowerPrices, price)
		currentFlower--
		currentFriendBuying++
		if currentFriendBuying%int(k) == 0 {
			multiplier++
		}
	}

	return sumPrices(flowerPrices)
}

func sumPrices(c []int32) int32 {
	price := int32(0)
	for _, v := range c {
		price += v
	}

	return price
}
