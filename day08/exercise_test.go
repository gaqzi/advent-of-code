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

func TestModifying(t *testing.T) {
	instructions, err := file.Load("instructions.sample.txt")
	assert.NilError(t, err)

	assert.Equal(t, 8, day08.Modifying(instructions))
}

func TestExercise(t *testing.T) {
	instructions, err := file.Load("instructions.txt")
	assert.NilError(t, err)

	t.Logf("Part 1 acc: %d", day08.Halting(instructions))
	t.Logf("Part 2 acc: %d", day08.Modifying(instructions))
}
