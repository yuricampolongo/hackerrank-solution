package main

import "fmt"

func main() {
	arr := []int32{1, 1, 0, -1, -1}

	plusMinus(arr)
}

func plusMinus(arr []int32) {
	positives, negatives, zeros := 0, 0, 0

	for _, v := range arr {
		if v < 0 {
			negatives++
		} else if v > 0 {
			positives++
		} else {
			zeros++
		}
	}

	fmt.Printf("%.6f\n", float64(positives)/float64(len(arr)))
	fmt.Printf("%.6f\n", float64(negatives)/float64(len(arr)))
	fmt.Printf("%.6f\n", float64(zeros)/float64(len(arr)))
}
