package day02

import (
	"fmt"
	"strconv"
	"strings"
)

type CharPolicy struct {
	First  int64
	Second int64
	Char   string
}

func NewCharPolicy(s string) (CharPolicy, error) {
	split := strings.SplitN(s, " ", 2)
	range_ := strings.SplitN(split[0], "-", 2)

	minMax := make([]int64, 2)
	for i, val := range range_ {
		v, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return CharPolicy{}, fmt.Errorf("unable to parse '%s' to int: %w", val, err)
		}
		minMax[i] = v
	}

	return CharPolicy{First: minMax[0], Second: minMax[1], Char: split[1]}, nil
}

func (p CharPolicy) ValidCount(s string) bool {
	count := int64(strings.Count(s, p.Char))
	return count >= p.First && count <= p.Second
}

func (p CharPolicy) ValidPositions(s string) bool {
	r := []rune(s)
	chr := rune(p.Char[0])
	first := r[p.First-1] == chr
	second := r[p.Second-1] == chr

	switch {
	case first && second:
		return false
	case first || second:
		return true
	default:
		return false
	}
}
