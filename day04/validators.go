package day04

import (
	"fmt"
	"regexp"
	"strconv"
)

func rangeValidator(min, max int) func(string) bool {
	return func(s string) bool {
		v, err := strconv.Atoi(s)
		if err != nil {
			return false
		}

		return v >= min && v <= max
	}
}

func regexValidator(pattern string) func(string) bool {
	compiled, err := regexp.Compile(pattern)
	if err != nil {
		panic(fmt.Sprintf("failed to compile compiled '%s'", pattern))
	}

	return func(s string) bool {
		return compiled.MatchString(s)
	}
}

func setValidator(values ...string) func(string) bool {
	allValid := map[string]struct{}{}
	for _, v := range values {
		allValid[v] = struct{}{}
	}

	return func(s string) bool {
		_, ok := allValid[s]
		return ok
	}
}

var (
	IsValidByr   = rangeValidator(1920, 2002)
	IsValidIyr   = rangeValidator(2010, 2020)
	IsValidEyr   = rangeValidator(2020, 2030)
	isValidHgtCM = rangeValidator(150, 193)
	isValidHgtIN = rangeValidator(59, 76)
	IsValidHcl   = regexValidator(`^#[\da-fA-F]{6}$`)
	IsValidEcl   = setValidator("amb", "blu", "brn", "gry", "grn", "hzl", "oth")
	IsValidPid   = regexValidator(`^\d{9}$`)
)

var IsValidHgt = func(s string) bool {
	suffixPos := len(s) - 2

	switch s[suffixPos:] {
	case "cm":
		return isValidHgtCM(s[:suffixPos])
	case "in":
		return isValidHgtIN(s[:suffixPos])
	default:
		return false
	}
}
