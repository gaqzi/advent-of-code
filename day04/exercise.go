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
	return isValidPassport(passport, true)

}

func IsValidPassport(passport []string) bool {
	return isValidPassport(passport, false)
}

func isValidPassport(p []string, simple bool) bool {
	var hasRequired int
	for _, line := range p {
		for _, segment := range strings.Split(line, " ") {
			divideAt := strings.Index(segment, ":")

			if valid, ok := requiredFields[segment[:divideAt]]; ok {
				if simple || valid(segment[divideAt+1:]) {
					hasRequired++
				}
			}
		}
	}

	return hasRequired == len(requiredFields)
}
