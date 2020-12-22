package day07_test

import (
	"advent-of-code/day07"
	"advent-of-code/file"
	"gotest.tools/v3/assert"
	"testing"
)

func TestBagger(t *testing.T) {
	t.Run("CanHold", func(t *testing.T) {
		rules, err := file.Load("rules.sample.txt")
		assert.NilError(t, err, "failed to load sample rules")

		bags, err := day07.Bagger(rules)
		assert.NilError(t, err)
		canContainShinyGold := bags.CanHold("shiny gold")
		assert.Equal(t, 4, len(canContainShinyGold))
	})

	t.Run("Holds", func(t *testing.T) {
		rules, err := file.Load("rules.sample2.txt")
		assert.NilError(t, err)

		bags, err := day07.Bagger(rules)
		assert.NilError(t, err)
		assert.Equal(t, 2, bags.Holds("dark blue"), "dark blue")
		assert.Equal(t, 6, bags.Holds("dark green"), "dark green")
		assert.Equal(t, 14, bags.Holds("dark yellow"), "dark yellow")
		assert.Equal(t, 30, bags.Holds("dark orange"), "dark orange")
		assert.Equal(t, 62, bags.Holds("dark red"), "dark red")
		assert.Equal(t, 126, bags.Holds("shiny gold"))
	})
}

func TestExercise(t *testing.T) {
	rules, err := file.Load("rules.txt")
	assert.NilError(t, err, "failed to load rules")

	bags, err := day07.Bagger(rules)
	assert.NilError(t, err)

	t.Logf("Bags that can hold shiny gold bags: %d", len(bags.CanHold("shiny gold")))
	t.Logf("Bags that shiny gold bag holds: %d", bags.Holds("shiny gold"))
}
