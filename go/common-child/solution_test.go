package commonchild

import (
	"fmt"
	"testing"
)

func TestCommonChild(t *testing.T) {
	inputs := []struct {
		s1       string
		s2       string
		expected int32
	}{
		{
			s1:       "ABCD",
			s2:       "ABDC",
			expected: 3,
		},
		{
			s1:       "HARRY",
			s2:       "SALLY",
			expected: 2,
		},
		{
			s1:       "SHINCHAN",
			s2:       "NOHARAAA",
			expected: 3,
		},
		{
			s1:       "AA",
			s2:       "BB",
			expected: 0,
		},
		{
			s1:       "WEWOUCUIDGCGTRMEZEPXZFEJWISRSBBSYXAYDFEJJDLEBVHHKS",
			s2:       "FDAGCXGKCTKWNECHMRXZWMLRYUCOCZHJRRJBOAJOQJZZVUYXIC",
			expected: 15,
		},
		{
			s1:       "OUDFRMYMAW",
			s2:       "AWHYFCCMQX",
			expected: 2,
		},
	}

	for _, v := range inputs {
		t.Run(fmt.Sprintf("running case #%s", v.s1), func(t *testing.T) {
			debugging = true

			result := commonChild(v.s1, v.s2)

			if result != v.expected {
				t.Fatalf("expected %d got %d", v.expected, result)
			}
		})
	}
}
