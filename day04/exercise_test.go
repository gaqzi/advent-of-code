package day04_test

import (
	"advent-of-code/day04"
	"fmt"
	"gotest.tools/v3/assert"
	"io/ioutil"
	"strings"
	"testing"
)

func TestIsValidPassportSimple(t *testing.T) {
	testCases := []struct {
		name    string
		lines   []string
		isValid bool
	}{
		{
			name: "All fields and valid",
			lines: []string{
				"ecl:gry pid:860033327 eyr:2020 hcl:#fffffd",
				"byr:1937 iyr:2017 cid:147 hgt:183cm",
			},
			isValid: true,
		},
		{
			name: "Invalid missing field",
			lines: []string{
				"iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884",
				"hcl:#cfa07d byr:1929",
			},
			isValid: false,
		},
		{
			name: "Valid missing cid",
			lines: []string{
				"hcl:#ae17e1 iyr:2013",
				"eyr:2024",
				"ecl:brn pid:760753108 byr:1931",
				"hgt:179cm",
			},
			isValid: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.isValid, day04.IsValidPassportSimple(tc.lines))
		})
	}
}

func TestIsValidPassport(t *testing.T) {
	testCases := []struct {
		lines   []string
		isValid bool
	}{
		{
			lines: []string{
				"eyr:1972 cid:100",
				"hcl:#18171d ecl:amb hgt:170 pid:186cm iyr:2018 byr:1926",
			},
		},
		{
			lines: []string{
				"iyr:2019",
				"hcl:#602927 eyr:1967 hgt:170cm",
				"ecl:grn pid:012533040 byr:1946",
			},
		},
		{
			lines: []string{
				"hcl:dab227 iyr:2012",
				"ecl:brn hgt:182cm pid:021572410 eyr:2020 byr:1992 cid:277",
			},
		},
		{
			lines: []string{
				"hgt:59cm ecl:zzz",
				"eyr:2038 hcl:74454a iyr:2023",
				"pid:3556412378 byr:2007",
			},
		},
		{
			lines: []string{
				"pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980",
				"hcl:#623a2f",
			},
			isValid: true,
		},
		{
			lines: []string{
				"eyr:2029 ecl:blu cid:129 byr:1989",
				"iyr:2014 pid:896056539 hcl:#a97842 hgt:165cm",
			},
			isValid: true,
		},
		{
			lines: []string{
				"hcl:#888785",
				"hgt:164cm byr:2001 iyr:2015 cid:88",
				"pid:545766238 ecl:hzl",
				"eyr:2022",
			},
			isValid: true,
		},
		{
			lines: []string{
				"iyr:2010 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719",
			},
			isValid: true,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("isValid=%v: %d", tc.isValid, i), func(t *testing.T) {
			t.Logf("%v", tc.lines)
			assert.Equal(t, tc.isValid, day04.IsValidPassport(tc.lines))
		})
	}
}

func TestRun(t *testing.T) {
	passports, err := ioutil.ReadFile("passports.txt")
	assert.NilError(t, err, "failed to load passport.txt")

	t.Run("Part 1", func(t *testing.T) {
		var passport []string
		var validPassports int64
		for _, l := range strings.Split(string(passports), "\n") {
			if l == "" {
				if day04.IsValidPassportSimple(passport) {
					validPassports++
				}
				passport = []string{}
				continue
			}

			passport = append(passport, l)
		}

		t.Logf("Valid passports: %d", validPassports)
	})

	t.Run("Part 2", func(t *testing.T) {
		var passport []string
		var validPassports int64
		for _, l := range strings.Split(string(passports), "\n") {
			if l == "" {
				if day04.IsValidPassport(passport) {
					validPassports++
				}
				passport = []string{}
				continue
			}

			passport = append(passport, l)
		}

		t.Logf("Valid passports: %d", validPassports)
	})
}

func TestValidators(t *testing.T) {
	testCases := []struct {
		name    string
		f       func(string) bool
		value   string
		isValid bool
	}{
		{name: "byr: <1920 invalid", f: day04.IsValidByr, value: "1919"},
		{name: "byr: >2002 invalid", f: day04.IsValidByr, value: "2003"},
		{name: "byr: >= 1920 valid", f: day04.IsValidByr, value: "1920", isValid: true},
		{name: "byr: <= 2002 valid", f: day04.IsValidByr, value: "2002", isValid: true},

		{name: "iyr: <2010 invalid", f: day04.IsValidIyr, value: "2009"},
		{name: "iyr: >2020 invalid", f: day04.IsValidIyr, value: "2021"},
		{name: "iyr: >= 2010 valid", f: day04.IsValidIyr, value: "2010", isValid: true},
		{name: "iyr: <= 2020 valid", f: day04.IsValidIyr, value: "2020", isValid: true},

		{name: "eyr: <2020 invalid", f: day04.IsValidEyr, value: "2019"},
		{name: "eyr: >2030 invalid", f: day04.IsValidEyr, value: "2031"},
		{name: "eyr: >= 2020 valid", f: day04.IsValidEyr, value: "2020", isValid: true},
		{name: "eyr: <= 2030 valid", f: day04.IsValidEyr, value: "2030", isValid: true},

		{name: "hgt: cm: <150 invalid", f: day04.IsValidHgt, value: "149cm"},
		{name: "hgt: cm: >193 invalid", f: day04.IsValidHgt, value: "194cm"},
		{name: "hgt: cm: >=150 valid", f: day04.IsValidHgt, value: "150cm", isValid: true},
		{name: "hgt: cm: <=193 valid", f: day04.IsValidHgt, value: "193cm", isValid: true},

		{name: "hgt: in: <59 invalid", f: day04.IsValidHgt, value: "58in"},
		{name: "hgt: in: >76 invalid", f: day04.IsValidHgt, value: "77in"},
		{name: "hgt: in: >=59 valid", f: day04.IsValidHgt, value: "59in", isValid: true},
		{name: "hgt: in: <=76 valid", f: day04.IsValidHgt, value: "76in", isValid: true},

		{name: "hcl: invalid length", f: day04.IsValidHcl, value: "#12345"},
		{name: "hcl: valid", f: day04.IsValidHcl, value: "#123456", isValid: true},
		{name: "hcl: valid", f: day04.IsValidHcl, value: "#abcdef", isValid: true},
		{name: "hcl: invalid char", f: day04.IsValidHcl, value: "#12345g"},

		{name: "ecl: unknown: invalid", f: day04.IsValidEcl, value: "unknown"},
		{name: "ecl: amb: valid", f: day04.IsValidEcl, value: "amb", isValid: true},
		{name: "ecl: blu: valid", f: day04.IsValidEcl, value: "blu", isValid: true},
		{name: "ecl: brn: valid", f: day04.IsValidEcl, value: "brn", isValid: true},
		{name: "ecl: gry: valid", f: day04.IsValidEcl, value: "gry", isValid: true},
		{name: "ecl: grn: valid", f: day04.IsValidEcl, value: "grn", isValid: true},
		{name: "ecl: hzl: valid", f: day04.IsValidEcl, value: "hzl", isValid: true},
		{name: "ecl: oth: valid", f: day04.IsValidEcl, value: "oth", isValid: true},

		{name: "pid: too short: invalid", f: day04.IsValidPid, value: "12345678"},
		{name: "pid: too long: invalid", f: day04.IsValidPid, value: "1234567890"},
		{name: "pid: correct length + letter: invalid", f: day04.IsValidPid, value: "123a56789"},
		{name: "pid: only numbers: valid", f: day04.IsValidPid, value: "123456789", isValid: true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.isValid, tc.f(tc.value))
		})
	}
}
