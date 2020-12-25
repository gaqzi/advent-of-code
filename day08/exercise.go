package day08

import (
	"fmt"
	"strconv"
	"strings"
)

type instruction struct {
	operation string
	value     int
}

func Halting(instructions []string) int {
	var pos, acc int
	ins := parseInstructions(instructions)
	visited := map[int]struct{}{}

	for {
		if _, ok := visited[pos]; ok {
			break
		}
		visited[pos] = struct{}{}

		switch ins[pos].operation {
		case "jmp":
			pos += ins[pos].value
			continue
		case "acc":
			acc += ins[pos].value
		case "nop":
		}

		pos++
	}

	return acc
}

func parseInstructions(instructions []string) []instruction {
	var rets []instruction

	for _, instLine := range instructions {
		line := strings.SplitN(instLine, " ", 2)
		val, err := strconv.Atoi(line[1])
		if err != nil {
			panic(fmt.Sprintf("failed to convert '%s': %s", line[1], err))
		}
		rets = append(rets, instruction{operation: line[0], value: val})
	}

	return rets
}
