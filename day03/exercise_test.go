package day03_test

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"

	"gotest.tools/v3/assert"

	"advent-of-code/day03"
)

func loadFile(f string) ([]string, error) {
	file, err := ioutil.ReadFile(f)
	if err != nil {
		return nil, fmt.Errorf("failed to read file '%s': %w", f, err)
	}

	return strings.Split(string(file), "\n"), nil
}

func TestTraverse(t *testing.T) {
	t.Run("sample file", func(t *testing.T) {
		geography, err := loadFile("geography.sample.txt")
		assert.NilError(t, err, "failed to load geography")

		day03.Traverse(geography, 3, 1)
	})
}
