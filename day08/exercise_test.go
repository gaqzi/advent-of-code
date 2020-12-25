package day08_test

import (
	"advent-of-code/day08"
	"advent-of-code/file"
	"gotest.tools/v3/assert"
	"testing"
)

func TestHalting(t *testing.T) {
	instructions, err := file.Load("instructions.sample.txt")
	assert.NilError(t, err)

	assert.Equal(t, 5, day08.Halting(instructions))
}

func TestExercise(t *testing.T) {
	t.Run("Part 1", func(t *testing.T) {
		instructions, err := file.Load("instructions.txt")
		assert.NilError(t, err)

		t.Logf("Part 1 acc: %d", day08.Halting(instructions))
	})
}
