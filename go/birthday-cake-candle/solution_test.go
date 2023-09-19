package birthdaycakecandle

import "testing"

func TestBirthdayCakeCandles(t *testing.T) {
	input := []int32{4, 4, 3, 1}
	result := birthdayCakeCandles(input)

	if result != 2 {
		t.Errorf("Expected 2 got %d", result)
	}

}
