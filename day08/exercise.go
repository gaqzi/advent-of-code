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

func Modifying(instructions []string) int {
	var pos, acc int
	var prematureExit bool
	originalInstructions := parseInstructions(instructions)
	visited := map[int]instruction{}

	var previouslyModifiedPos int
	ins, previouslyModifiedPos := switchInstruction(originalInstructions, -1)
	for {
		for pos < len(ins) {
			if _, ok := visited[pos]; ok {
				prematureExit = true
				break
			}
			visited[pos] = ins[pos]

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

		if !prematureExit {
			break
		}
		prematureExit = false
		visited = map[int]instruction{}
		pos, acc = 0, 0
		ins, previouslyModifiedPos = switchInstruction(originalInstructions, previouslyModifiedPos)
	}

	return acc
}

func switchInstruction(instructions []instruction, pos int) ([]instruction, int) {
	ins := make([]instruction, len(instructions))
	copy(ins, instructions)

	for i := pos + 1; i < len(instructions); i++ {
		switch instructions[i].operation {
		case "nop":
			ins[i] = instruction{operation: "jmp", value: instructions[i].value}
			return ins, i
		case "jmp":
			ins[i] = instruction{operation: "nop", value: instructions[i].value}
			return ins, i
		}
	}

	return ins, pos
}
