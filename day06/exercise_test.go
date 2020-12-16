package day06_test

import (
	"advent-of-code/day06"
	"advent-of-code/file"
	"gotest.tools/v3/assert"
	"testing"
)

func TestUniqueAnswers(t *testing.T) {
	groups, err := file.LoadGroups("answers.sample.txt")
	assert.NilError(t, err)

	assert.Equal(t, 11, day06.UniqueAnswers(groups))
}

func TestAllSameInGroup(t *testing.T) {
	groups, err := file.LoadGroups("answers.sample.txt")
	assert.NilError(t, err)

	assert.Equal(t, 6, day06.AllSameInGroup(groups))
}

func TestExercise(t *testing.T) {
	groups, err := file.LoadGroups("answers.txt")
	assert.NilError(t, err)

	t.Run("Part 1", func(t *testing.T) {
		t.Logf("Total yes answers: %d", day06.UniqueAnswers(groups))
	})

	t.Run("Part b", func(t *testing.T) {
		t.Logf("Total yes answers: %d", day06.AllSameInGroup(groups))
	})
}
