package diagonaldifference

import "testing"

func TestDiagonalDifference(t *testing.T) {
	curValue := int32(1)
	input := make([][]int32, 3)
	for i := 0; i < len(input); i++ {
		input[i] = make([]int32, 3)
		for j := 0; j < len(input[i]); j++ {
			input[i][j] = curValue
			curValue++
		}
	}

	result := diagonalDifference(input)

	if result != 0 {
		t.Errorf("Expected 0, got %d", result)
	}
}

func TestDiagonalDifference2(t *testing.T) {
	curValue := int32(1)
	input := make([][]int32, 3)
	for i := 0; i < len(input); i++ {
		input[i] = make([]int32, 3)
		for j := 0; j < len(input[i]); j++ {
			input[i][j] = curValue
			curValue++
		}
	}

	input[2][0] = 9

	result := diagonalDifference(input)

	if result != 2 {
		t.Errorf("Expected 2, got %d", result)
	}
}
