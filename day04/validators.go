package day04

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var requiredFields map[string]func(string) bool

func init() {
	requiredFields = map[string]func(string) bool{
		"byr": isValidByr,
		"iyr": isValidIyr,
		"eyr": isValidEyr,
		"hgt": isValidHgt,
		"hcl": isValidHcl,
		"ecl": isValidEcl,
		"pid": isValidPid,
	}
}

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
	isValidByr   = rangeValidator(1920, 2002)
	isValidIyr   = rangeValidator(2010, 2020)
	isValidEyr   = rangeValidator(2020, 2030)
	isValidHgtCM = rangeValidator(150, 193)
	isValidHgtIN = rangeValidator(59, 76)
	isValidHcl   = regexValidator(`^#[\da-fA-F]{6}$`)
	isValidEcl   = setValidator("amb", "blu", "brn", "gry", "grn", "hzl", "oth")
	isValidPid   = regexValidator(`^\d{9}$`)
)

var isValidHgt = func(s string) bool {
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

func isValid(typ, val string) bool {
	valid, ok := requiredFields[strings.ToLower(typ)]
	return ok && valid(val)
}

func isRequiredField(name string) bool {
	_, ok := requiredFields[name]
	return ok
}
