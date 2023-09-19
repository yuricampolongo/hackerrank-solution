package main

import (
	"fmt"
	"math"
)

type magicSquare struct {
	matrix [][]int32
}

func main() {
	matrix := [][]int32{
		{4, 9, 2},
		{3, 5, 7},
		{8, 1, 5},
	}
	result := formingMagicSquare(matrix)
	fmt.Println(result)
}

/*
 * Complete the 'formingMagicSquare' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts 2D_INTEGER_ARRAY s as parameter.
 */

func formingMagicSquare(s [][]int32) int32 {
	magicSquares := []*magicSquare{}
	magicSquares = append(magicSquares, &magicSquare{
		matrix: [][]int32{
			{8, 1, 6},
			{3, 5, 7},
			{4, 9, 2},
		},
	})
	magicSquares = append(magicSquares, &magicSquare{
		matrix: [][]int32{
			{6, 1, 8},
			{7, 5, 3},
			{2, 9, 4},
		},
	})
	magicSquares = append(magicSquares, &magicSquare{
		matrix: [][]int32{
			{4, 9, 2},
			{3, 5, 7},
			{8, 1, 6},
		},
	})
	magicSquares = append(magicSquares, &magicSquare{
		matrix: [][]int32{
			{2, 9, 4},
			{7, 5, 3},
			{6, 1, 8},
		},
	})
	magicSquares = append(magicSquares, &magicSquare{
		matrix: [][]int32{
			{8, 3, 4},
			{1, 5, 9},
			{6, 7, 2},
		},
	})
	magicSquares = append(magicSquares, &magicSquare{
		matrix: [][]int32{
			{4, 3, 8},
			{9, 5, 1},
			{2, 7, 6},
		},
	})
	magicSquares = append(magicSquares, &magicSquare{
		matrix: [][]int32{
			{6, 7, 2},
			{1, 5, 9},
			{8, 3, 4},
		},
	})
	magicSquares = append(magicSquares, &magicSquare{
		matrix: [][]int32{
			{2, 7, 6},
			{9, 5, 1},
			{4, 3, 8},
		},
	})

	minCost := int32(1000)

	for _, v := range magicSquares {
		cost := int32(0)
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				cost += int32(math.Abs(float64(s[i][j]) - float64(v.matrix[i][j])))
			}
		}
		if cost < minCost {
			minCost = cost
		}
	}

	return minCost
}
