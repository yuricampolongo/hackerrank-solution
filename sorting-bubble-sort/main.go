package main

import "fmt"

func main() {
	countSwaps([]int32{3, 2, 1})
}

func countSwaps(a []int32) {
	n := len(a)
	swaps := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
				swaps++
			}
		}
	}

	fmt.Printf("Array is sorted in %d swaps.\nFirst Element: %d\nLast Element: %d", swaps, a[0], a[len(a)-1])
}
