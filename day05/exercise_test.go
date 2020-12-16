package day05_test

import (
	"advent-of-code/day05"
	"advent-of-code/file"
	"gotest.tools/v3/assert"
	"sort"
	"testing"
)

func TestParseSeat(t *testing.T) {
	assert.Equal(t, 357, day05.ParseSeat("FBFBBFFRLR"))
}

func TestExercise(t *testing.T) {
	lines, err := file.Load("boarding-passes.txt")
	assert.NilError(t, err)

	t.Run("Find highest seat ID", func(t *testing.T) {
		var highestID int
		for _, l := range lines {
			if seatID := day05.ParseSeat(l); seatID > highestID {
				highestID = seatID
			}
		}

		t.Logf("Highest seat id: %d", highestID)
	})

	t.Run("Find your seat, the gap between two others", func(t *testing.T) {
		var seats []int
		for _, l := range lines {
			seatID := day05.ParseSeat(l)
			seats = append(seats, seatID)
		}

		sort.Ints(seats)

		for i := 0; i < len(seats)-1; i++ {
			if seats[i+1]-seats[i] != 1 {
				t.Logf("Gap seat: %d", seats[i]+1)
				break
			}
		}
	})
}
