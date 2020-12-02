package day01

import (
	"fmt"
	"sort"
)

func FindSum(sum int64, haystack []int64) ([]int64, error) {
	for i := 0; i < len(haystack); i++ {
		for j := i + 1; j < len(haystack); j++ {
			if haystack[i] + haystack[j] == sum {
				return []int64{haystack[i], haystack[j]}, nil
			}
		}
	}

	return nil, fmt.Errorf("couldn't find two numbers that summed to %d", sum)
}

func FindSumSorted(sum int64, haystack []int64) ([]int64, error) {
	sort.Slice(haystack, func(i, j int) bool { return haystack[i] < haystack[j] })

	for i, val := range haystack {
		toFind := sum - val
		ix := sort.Search(len(haystack[i:]), func(i int) bool { return haystack[i] == toFind})
		if ix < len(haystack[i:]) && haystack[ix] == toFind {
			return []int64{toFind, haystack[ix]}, nil
		}
	}

	return nil, fmt.Errorf("couldn't find two numbers that summed to %d", sum)
}


func FindSumMap(sum int64, haystack []int64) ([]int64, error) {
	mapStack := make(map[int64]struct{}, len(haystack))

	return findSumMap(sum, mapStack)
}


func findSumMap(sum int64, haystack map[int64]struct{}) ([]int64, error) {
	for val, _ := range haystack {
		toFind := sum - val
		if _, ok := haystack[toFind]; ok {
			return []int64{val, toFind}, nil
		}
	}

	return nil, fmt.Errorf("couldn't find two numbers that summed to %d", sum)
}

// FindSumSplits finds the sum across the number splits in the haystack.
// 2002 in 3 splits might be 346 + 1380 + 294
func FindSumSplits(sum int64, splits int64, haystack []int64) ([]int64, error) {
	switch {
	case splits < 2:
		return nil, fmt.Errorf("need minimum two splits, got: %d", splits)
	case splits == 2:
		return FindSum(sum, haystack)
	}

	for i, val := range haystack {
		toFind := sum - val
		values, err := FindSumSplits(toFind, splits - 1, haystack[i:])
		if err == nil {
			return append([]int64{val}, values...), nil
		}
	}

	return nil, fmt.Errorf("couldn't find %d numbers that summed to %d", sum, splits)
}

func FindSumSplitsSorted(sum int64, splits int64, haystack []int64) ([]int64, error) {
	switch {
	case splits < 2:
		return nil, fmt.Errorf("need minimum two splits, got: %d", splits)
	case splits == 2:
		return FindSumSorted(sum, haystack)
	}

	for i, val := range haystack {
		toFind := sum - val
		if values, err := FindSumSplitsSorted(toFind, splits - 1, haystack[i:]); err == nil {
			return append([]int64{val}, values...), nil
		}
	}

	return nil, fmt.Errorf("couldn't find %d numbers that summed to %d", sum, splits)
}


func FindSumSplitsMap(sum int64, splits int64, haystack []int64) ([]int64, error) {
	stackMap := make(map[int64]struct{}, len(haystack))
	for _, val := range haystack {
		stackMap[val] = struct{}{}
	}

	return findSumSplitsMap(sum, splits, stackMap)
}

func findSumSplitsMap(sum int64, splits int64, haystack map[int64]struct{}) ([]int64, error) {
	switch {
	case splits < 2:
		return nil, fmt.Errorf("need minimum two splits, got: %d", splits)
	case splits == 2:
		return findSumMap(sum, haystack)
	}

	for val, _ := range haystack {
		toFind := sum - val
		if values, err := findSumSplitsMap(toFind, splits - 1, haystack); err == nil {
			return append([]int64{val}, values...), nil
		}
	}

	return nil, fmt.Errorf("couldn't find %d numbers that summed to %d", sum, splits)
}
