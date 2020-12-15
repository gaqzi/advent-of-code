package day04

import (
	"strings"
)

var requiredFields map[string]func(string) bool

func init() {
	requiredFields = map[string]func(string) bool{
		"byr": IsValidByr,
		"iyr": IsValidIyr,
		"eyr": IsValidEyr,
		"hgt": IsValidHgt,
		"hcl": IsValidHcl,
		"ecl": IsValidEcl,
		"pid": IsValidPid,
	}
}

func IsValidPassportSimple(passport []string) bool {
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

func IsValidPassport(passport []string) bool {
	var hasRequired int
	for _, line := range passport {
		for _, segment := range strings.Split(line, " ") {
			divideAt := strings.Index(segment, ":")

			if valid, ok := requiredFields[segment[:divideAt]]; ok {
				if valid(segment[divideAt+1:]) {
					hasRequired++
				}
			}
		}
	}

	return hasRequired == len(requiredFields)
}
