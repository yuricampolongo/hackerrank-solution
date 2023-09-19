package birthdaycakecandle

import "sort"

func birthdayCakeCandles(candles []int32) int32 {
	sort.Slice(candles, func(i, j int) bool { return candles[i] > candles[j] })
	lastValue := int32(-1)
	counter := int32(0)
	for _, v := range candles {
		if v > lastValue {
			lastValue = v
		}

		if lastValue == v {
			counter++
		}
	}

	return counter
}
