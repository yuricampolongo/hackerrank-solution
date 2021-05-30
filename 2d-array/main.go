package main

import "fmt"

func main() {
	result := hourglassSum([][]int32{
		{-1, -1, 0, -9, -2, -2},
		{-2, -1, -6, -8, -2, -5},
		{-1, -1, -1, -2, -3, -4},
		{-1, -9, -2, -4, -4, -5},
		{-7, -3, -3, -2, -9, -9},
		{-1, -3, -1, -2, -4, -5},
	})

	fmt.Println(result)
}

func hourglassSum(arr [][]int32) int32 {
	maxSum := int32(-64)
	var currentSum int32

	var i, j int
	for {
		currentSum += arr[i][j]
		currentSum += arr[i][j+1]
		currentSum += arr[i][j+2]
		currentSum += arr[i+1][j+1]
		currentSum += arr[i+2][j]
		currentSum += arr[i+2][j+1]
		currentSum += arr[i+2][j+2]

		if currentSum > maxSum {
			maxSum = currentSum
		}

		currentSum = 0

		j++

		if j == 4 {
			i++
			if i == 4 && j == 4 {
				break
			}
			j = 0
		}
	}

	return maxSum
}
