package day04

import (
	"strings"
)

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
			kv := strings.Split(segment, ":")

			if simple {
				if isRequiredField(kv[0]) {
					hasRequired++
				}
				continue
			}

			if isValid(kv[0], kv[1]) {
				hasRequired++
			}
		}
	}

	return hasRequired == len(requiredFields)
}
