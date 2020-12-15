package day04

import "strings"

var requiredFields map[string]struct{}

func init() {
	requiredFields = map[string]struct{}{
		"byr": {},
		"iyr": {},
		"eyr": {},
		"hgt": {},
		"hcl": {},
		"ecl": {},
		"pid": {},
	}
}

func IsValidPassport(passport []string) bool {
	var hasRequired int
	for _, line := range passport {
		for _, segment := range strings.Split(line, " ") {
			divideAt := strings.Index(segment, ":")

			if _, ok := requiredFields[segment[:divideAt]]; ok {
				hasRequired++
			}
		}
	}

	return hasRequired == len(requiredFields)
}
