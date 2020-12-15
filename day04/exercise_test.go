package day04_test

import (
	"advent-of-code/day04"
	"gotest.tools/v3/assert"
	"io/ioutil"
	"strings"
	"testing"
)

func TestValidate(t *testing.T) {
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
