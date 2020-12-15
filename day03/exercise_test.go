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

	var rets []string
	for _, l := range strings.Split(string(file), "\n") {
		if l != "" {
			rets = append(rets, l)
		}
	}

	return rets, nil
}

func TestTraverse(t *testing.T) {
	t.Run("sample file", func(t *testing.T) {
		geography, err := loadFile("geography.sample.txt")
		assert.NilError(t, err, "failed to load geography")

		assert.Equal(t, int64(7), day03.Traverse(geography, 3, 1), "expected to have encountered 7 trees on the way down")
	})
}

func TestTraverseFull(t *testing.T) {
	geography, err := loadFile("geography.txt")
	assert.NilError(t, err, "failed to load geography")

	t.Run("Part 1", func(t *testing.T) {
		t.Logf("Hit %d trees on the way down", day03.Traverse(geography, 3, 1))
	})

	t.Run("Part 2", func(t *testing.T) {
		slopes := []struct {
			right int64
			down  int64
		}{
			{1, 1},
			{3, 1},
			{5, 1},
			{7, 1},
			{1, 2},
		}
		var total int64
		for _, slope := range slopes {
			treesHit := day03.Traverse(geography, slope.right, slope.down)
			t.Logf("Right %d, down %d = %d", slope.right, slope.down, treesHit)
			if total == 0 {
				total = treesHit
				continue
			}

			total *= treesHit
		}

		t.Logf("All trees hit multiplied: %d", total)
	})
}
