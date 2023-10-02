package specialstringagain

import (
	"fmt"
	"testing"
)

func TestSubstrCount(t *testing.T) {
	inputs := []struct {
		s        string
		expected int64
	}{
		{
			s:        "mnonopoo",
			expected: 12,
		},
	}

	for _, v := range inputs {
		t.Run(fmt.Sprintf("testing case %s", v.s), func(t *testing.T) {
			result := substrCount(int32(len(v.s)), v.s)

			if result != v.expected {
				t.Fatalf("expected %d got %d", v.expected, result)
			}
		})
	}
}
