package main

import "fmt"

func main() {
	result := minimumDistances([]int32{3, 2, 1, 2, 3})
	fmt.Println(result)
}

func minimumDistances(a []int32) int32 {
	mapValues := make(map[int32]int32)
	minimum := -1

	for i := 0; i < len(a); i++ {
		if _, ok := mapValues[a[i]]; !ok {
			mapValues[a[i]] = int32(i) + 1
		} else {
			mapValues[a[i]] = (int32(i) - mapValues[a[i]]) + 1

			if minimum == -1 {
				minimum = int(mapValues[a[i]])
			} else if minimum > int(mapValues[a[i]]) {
				minimum = int(mapValues[a[i]])
			}
		}
	}

	return int32(minimum)
}
