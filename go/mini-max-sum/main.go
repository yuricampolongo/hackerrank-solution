package main

import (
	"fmt"
	"sort"
)

func main() {
	miniMaxSum([]int32{256741038, 623958417, 467905213, 714532089, 938071625})
}

func miniMaxSum(arr []int32) {
	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })
	min, max := int64(0), int64(0)
	for i := 0; i < 4; i++ {
		min += int64(arr[i])
		max += int64(arr[len(arr)-i-1])
	}
	fmt.Printf("%d %d", min, max)
}
