package luckbalance

import "sort"

type Contest struct {
	Luck        int32
	IsImportant bool
	Win         bool
}

func luckBalance(k int32, contestsInput [][]int32) int32 {
	var contests []Contest
	totalImportantContests := 0
	for k := range contestsInput {
		important := false

		if contestsInput[k][1] == 1 {
			important = true
			totalImportantContests++
		}

		contests = append(contests, Contest{
			Luck:        contestsInput[k][0],
			IsImportant: important,
			Win:         false,
		})
	}

	sortContests(contests)

	// if she cannot lose all important contest, than we mark the less important ones
	difference := totalImportantContests - int(k)
	if totalImportantContests > int(k) {
		for i := 0; i < difference; i++ {
			contests[i].Win = true
		}
	}

	totalLuck := 0
	for _, v := range contests {
		if v.Win {
			totalLuck -= int(v.Luck)
		} else {
			totalLuck += int(v.Luck)
		}
	}

	return int32(totalLuck)
}

func sortContests(contests []Contest) {
	sort.Slice(contests, func(i, j int) bool {
		if contests[i].IsImportant && !contests[j].IsImportant {
			return true
		} else if !contests[i].IsImportant && contests[j].IsImportant {
			return false
		}

		// If both are important or non-important, then sort by Luck value
		return contests[i].Luck < contests[j].Luck
	})
}
