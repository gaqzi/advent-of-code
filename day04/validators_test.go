package day04

import (
	"gotest.tools/v3/assert"
	"testing"
)

func TestValidators(t *testing.T) {
	var isValidFor = func(typ string) func(string) bool {
		return func(s string) bool {
			return isValid(typ, s)
		}
	}

	testCases := []struct {
		name    string
		f       func(string) bool
		value   string
		isValid bool
	}{
		{name: "byr: <1920 invalid", f: isValidFor("byr"), value: "1919"},
		{name: "byr: >2002 invalid", f: isValidFor("byr"), value: "2003"},
		{name: "byr: >= 1920 valid", f: isValidFor("byr"), value: "1920", isValid: true},
		{name: "byr: <= 2002 valid", f: isValidFor("byr"), value: "2002", isValid: true},

		{name: "iyr: <2010 invalid", f: isValidFor("iyr"), value: "2009"},
		{name: "iyr: >2020 invalid", f: isValidFor("iyr"), value: "2021"},
		{name: "iyr: >= 2010 valid", f: isValidFor("iyr"), value: "2010", isValid: true},
		{name: "iyr: <= 2020 valid", f: isValidFor("iyr"), value: "2020", isValid: true},

		{name: "eyr: <2020 invalid", f: isValidFor("eyr"), value: "2019"},
		{name: "eyr: >2030 invalid", f: isValidFor("eyr"), value: "2031"},
		{name: "eyr: >= 2020 valid", f: isValidFor("eyr"), value: "2020", isValid: true},
		{name: "eyr: <= 2030 valid", f: isValidFor("eyr"), value: "2030", isValid: true},

		{name: "hgt: cm: <150 invalid", f: isValidFor("hgt"), value: "149cm"},
		{name: "hgt: cm: >193 invalid", f: isValidFor("hgt"), value: "194cm"},
		{name: "hgt: cm: >=150 valid", f: isValidFor("hgt"), value: "150cm", isValid: true},
		{name: "hgt: cm: <=193 valid", f: isValidFor("hgt"), value: "193cm", isValid: true},

		{name: "hgt: in: <59 invalid", f: isValidFor("hgt"), value: "58in"},
		{name: "hgt: in: >76 invalid", f: isValidFor("hgt"), value: "77in"},
		{name: "hgt: in: >=59 valid", f: isValidFor("hgt"), value: "59in", isValid: true},
		{name: "hgt: in: <=76 valid", f: isValidFor("hgt"), value: "76in", isValid: true},

		{name: "hcl: invalid length", f: isValidFor("hcl"), value: "#12345"},
		{name: "hcl: valid", f: isValidFor("hcl"), value: "#123456", isValid: true},
		{name: "hcl: valid", f: isValidFor("hcl"), value: "#abcdef", isValid: true},
		{name: "hcl: invalid char", f: isValidFor("hcl"), value: "#12345g"},

		{name: "ecl: unknown: invalid", f: isValidFor("ecl"), value: "unknown"},
		{name: "ecl: amb: valid", f: isValidFor("ecl"), value: "amb", isValid: true},
		{name: "ecl: blu: valid", f: isValidFor("ecl"), value: "blu", isValid: true},
		{name: "ecl: brn: valid", f: isValidFor("ecl"), value: "brn", isValid: true},
		{name: "ecl: gry: valid", f: isValidFor("ecl"), value: "gry", isValid: true},
		{name: "ecl: grn: valid", f: isValidFor("ecl"), value: "grn", isValid: true},
		{name: "ecl: hzl: valid", f: isValidFor("ecl"), value: "hzl", isValid: true},
		{name: "ecl: oth: valid", f: isValidFor("ecl"), value: "oth", isValid: true},

		{name: "pid: too short: invalid", f: isValidFor("pid"), value: "12345678"},
		{name: "pid: too long: invalid", f: isValidFor("pid"), value: "1234567890"},
		{name: "pid: correct length + letter: invalid", f: isValidFor("pid"), value: "123a56789"},
		{name: "pid: only numbers: valid", f: isValidFor("pid"), value: "123456789", isValid: true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.isValid, tc.f(tc.value))
		})
	}
}
