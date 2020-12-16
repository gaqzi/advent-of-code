package day06

func UniqueAnswers(groups [][]string) int {
	var total int
	for _, group := range groups {
		answers := map[int32]struct{}{}
		for _, person := range group {
			for _, answer := range person {
				answers[answer] = struct{}{}
			}
		}

		total += len(answers)
	}

	return total
}

func AllSameInGroup(groups [][]string) int {
	var total int

	for _, group := range groups {
		answers := map[int32]int{}
		for _, person := range group {
			for _, answer := range person {
				answers[answer]++
			}
		}

		for _, v := range answers {
			if v == len(group) {
				total++
			}
		}
	}

	return total
}
