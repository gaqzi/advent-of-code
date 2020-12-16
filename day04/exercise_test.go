package day04_test

import (
	"advent-of-code/day04"
	"advent-of-code/file"
	"fmt"
	"gotest.tools/v3/assert"
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
	passports, err := file.LoadGroups("passports.txt")
	assert.NilError(t, err, "failed to load passport.txt")

	t.Run("Part 1", func(t *testing.T) {
		var validPassports int64
		for _, passport := range passports {
			if day04.IsValidPassportSimple(passport) {
				validPassports++
			}
		}

		t.Logf("Valid passports: %d", validPassports)
	})

	t.Run("Part 2", func(t *testing.T) {
		var validPassports int64
		for _, passport := range passports {
			if day04.IsValidPassport(passport) {
				validPassports++
			}
		}

		t.Logf("Valid passports: %d", validPassports)
	})
}
