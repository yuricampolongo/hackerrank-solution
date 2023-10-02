package commonchild

import (
	"fmt"
	"slices"
	"strings"
)

var debugging = false

func commonChild(s1 string, s2 string) int32 {
	return int32(lcs(s1, s2))
}

func lcs(s1 string, s2 string) int {
	matrix := make([][]int, len(s1)+1)

	for i := 0; i < len(s1)+1; i++ {
		matrix[i] = make([]int, len(s2)+1)
	}

	printMatrix(matrix)

	s1Runes := strings.Split(s1, "")
	s2Runes := strings.Split(s2, "")

	total := 0

	for i := 0; i < len(s1Runes); i++ {
		for j := 0; j < len(s2Runes); j++ {
			if s1Runes[i] == s2Runes[j] {
				matrix[i+1][j+1] = matrix[i][j] + 1
				if matrix[i+1][j+1] > total {
					total = matrix[i+1][j+1]
				}
			} else if matrix[i+1][j] > matrix[i][j+1] {
				matrix[i+1][j+1] = matrix[i+1][j]
			} else {
				matrix[i+1][j+1] = matrix[i][j+1]
			}
		}
	}

	printMatrix(matrix)
	printCommonSub(matrix, s1, s2)

	return total

}

func printMatrix(matrix [][]int) {
	if debugging {
		for row := range matrix {
			for column := range matrix[row] {
				fmt.Print(matrix[row][column], " ")
			}
			fmt.Print("\n")
		}
	}
	fmt.Print("\n")
}

func printCommonSub(matrix [][]int, s1 string, s2 string) {
	if debugging {
		s1Runes := strings.Split(s1, "")
		s := make([]string, 0, matrix[len(s1)][len(s2)])
		for i, j := len(s1), len(s2); i != 0 && j != 0; {
			if matrix[i][j] == matrix[i-1][j] {
				i--
			} else if matrix[i][j] == matrix[i][j-1] {
				j--
			} else {
				s = append(s, s1Runes[i-1])
				i--
				j--
			}
		}

		slices.Reverse(s)
		fmt.Println(s)
		fmt.Println()
	}
}
