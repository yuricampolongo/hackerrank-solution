package main

import (
	"fmt"
	"sort"
)

func main() {
	result := maximumToys([]int32{1, 2, 3, 4}, 7)
	fmt.Println(result)
}

func maximumToys(prices []int32, k int32) int32 {
	sort.Slice(prices, func(i, j int) bool { return prices[i] < prices[j] })
	gifts := 0
	for _, price := range prices {
		if price < k {
			gifts++
			k -= price
		} else {
			return int32(gifts)
		}
	}
	return int32(gifts)
}
