package main

import (
	"fmt"
	"sort"
)

func main() {
	whatFlavors([]int32{1, 2, 3, 5, 6}, 5)
}

type flavor struct {
	price int32
	index int
}

func whatFlavors(cost []int32, money int32) {
	originalCost := make([]*flavor, len(cost))
	for i, cost := range cost {
		originalCost[i] = &flavor{
			price: cost,
			index: i + 1,
		}
	}

	//Create an ordered cost array
	orderedCost := make([]*flavor, len(cost))
	copy(orderedCost, originalCost)
	sort.Slice(orderedCost, func(i, j int) bool { return orderedCost[i].price < orderedCost[j].price })

	var id, id2 *flavor

	//Search for flavors
	for idx, flavor := range orderedCost {
		toSpend := money - flavor.price
		arrayToSearch := orderedCost[idx+1:]
		id = flavor
		id2 = nil
		sort.Search(len(arrayToSearch), func(i int) bool {
			if arrayToSearch[i].price == toSpend {
				id2 = arrayToSearch[i]
			}
			return arrayToSearch[i].price >= toSpend
		})
		if id2 != nil {
			break
		}
	}

	if id.index < id2.index {
		fmt.Printf("%d %d\n", id.index, id2.index)
	} else {
		fmt.Printf("%d %d\n", id2.index, id.index)
	}
}
