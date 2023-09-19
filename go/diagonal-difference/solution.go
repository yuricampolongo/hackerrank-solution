package diagonaldifference

import "math"

func diagonalDifference(arr [][]int32) int32 {
	sumFirstDiagonal := float64(0)
	sumSecondDiagonal := float64(0)
	for i, j := 0, len(arr)-1; i < len(arr); i, j = i+1, j-1 {
		sumFirstDiagonal = sumFirstDiagonal + float64(arr[i][i])
		sumSecondDiagonal = sumSecondDiagonal + float64(arr[i][j])
	}
	return int32(math.Abs(sumFirstDiagonal - sumSecondDiagonal))
}
