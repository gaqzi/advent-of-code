package day04

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var requiredFields map[string]func(string) bool

func init() {
	var (
		isValidHgtCM = rangeValidator(150, 193)
		isValidHgtIN = rangeValidator(59, 76)
	)

	requiredFields = map[string]func(string) bool{
		"byr": rangeValidator(1920, 2002),
		"iyr": rangeValidator(2010, 2020),
		"eyr": rangeValidator(2020, 2030),
		"hcl": regexValidator(`^#[\da-fA-F]{6}$`),
		"ecl": setValidator("amb", "blu", "brn", "gry", "grn", "hzl", "oth"),
		"pid": regexValidator(`^\d{9}$`),
		"hgt": func(s string) bool {
			suffixPos := len(s) - 2

			switch s[suffixPos:] {
			case "cm":
				return isValidHgtCM(s[:suffixPos])
			case "in":
				return isValidHgtIN(s[:suffixPos])
			default:
				return false
			}
		},
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

func isValid(typ, val string) bool {
	valid, ok := requiredFields[strings.ToLower(typ)]
	return ok && valid(val)
}

func isRequiredField(name string) bool {
	_, ok := requiredFields[name]
	return ok
}
