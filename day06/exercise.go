package day06

func UniqueAnswers(groups [][]string) int {
	return reduce(groups, func(_ int, answers map[int32]int) int {
		return len(answers)
	})
}

func AllSameInGroup(groups [][]string) int {
	return reduce(groups, func(people int, answers map[int32]int) int {
		var total int

		for _, responders := range answers {
			if responders == people {
				total++
			}
		}

		return total
	})
}

func reduce(groups [][]string, counter func(people int, answers map[int32]int) int) int {
	var total int

	for _, group := range groups {
		answers := map[int32]int{}
		for _, person := range group {
			for _, answer := range person {
				answers[answer]++
			}
		}

		total += counter(len(group), answers)
	}

	return total
}
