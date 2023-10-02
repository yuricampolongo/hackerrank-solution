package specialstringagain

import "strings"

func substrCount(n int32, s string) int64 {
	count := int64(n)

	arr := strings.Split(s, "")

	for i := 0; i < int(n)-1; i++ {
		for k := i + 1; k < int(n); k++ {
			if arr[i] == arr[k] {
				count++
			} else {
				lastIndex := int32(k*2 - i)
				if lastIndex < n && s[i:k] == s[k+1:lastIndex+1] {
					count++
				}
				break
			}
		}
	}

	return count
}
