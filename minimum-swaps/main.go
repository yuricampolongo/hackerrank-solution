package main

import "fmt"

func main() {
	result := minimumSwaps([]int32{2, 3, 4, 1, 5})
	fmt.Println(result)
}

// Complete the minimumSwaps function below.
func minimumSwaps(arr []int32) int32 {
	positions := make(map[int32]int32)
	values := make(map[int32]int32)

	for k, v := range arr {
		positions[v] = int32(k) + 1
		values[int32(k)+1] = v
	}

	swaps := 0
	for i := 1; i <= len(arr); i++ {
		k, v := int32(i), positions[int32(i)]

		if k != v {
			oldPosition := positions[k]
			positions[k] = k
			toSwap := values[k]
			positions[toSwap] = oldPosition

			values[k] = k
			values[oldPosition] = toSwap

			swaps++
		}
	}

	return int32(swaps)
}
