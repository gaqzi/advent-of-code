package day07

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Holding struct {
	bags     map[string][]otherBag
	reversed map[string][]string
}

type otherBag struct {
	name  string
	count int
}

func Bagger(rules []string) (*Holding, error) {
	holder, err := regexp.Compile("^(.*?) bags contain")
	if err != nil {
		return nil, fmt.Errorf("failed to compile regex: %w", err)
	}
	bagsHeld, err := regexp.Compile(`([^ ]+) ([a-z]+ [a-z]+) bags?`)
	if err != nil {
		return nil, fmt.Errorf("failed to compile regex: %w", err)
	}

	bags := map[string][]otherBag{}
	for _, rule := range rules {
		bag := holder.FindStringSubmatch(rule)[1]
		bags[bag] = []otherBag{}
		for _, match := range bagsHeld.FindAllStringSubmatch(rule[strings.Index(rule, "contain"):], -1) {
			num, _ := strconv.Atoi(match[1])
			bags[bag] = append(bags[bag], otherBag{name: match[2], count: num})
		}
	}

	return &Holding{bags: bags, reversed: reversed(bags)}, nil
}

func reversed(bags map[string][]otherBag) map[string][]string {
	lookup := map[string][]string{}

	for k, v := range bags {
		for _, b := range v {
			if _, ok := lookup[b.name]; !ok {
				lookup[b.name] = []string{}
			}
			lookup[b.name] = append(lookup[b.name], k)
		}
	}

	return lookup
}

func (h Holding) CanHold(needle string) map[string]struct{} {
	bags := map[string]struct{}{}

	for _, bag := range h.reversed[needle] {
		bags[bag] = struct{}{}
		for k, v := range h.CanHold(bag) {
			bags[k] = v
		}
	}

	return bags
}

func (h Holding) Holds(needle string) int {
	var total int
	for _, bag := range h.bags[needle] {
		total += bag.count + (bag.count * h.Holds(bag.name))
	}

	return total
}
