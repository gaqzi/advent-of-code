package file

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func Load(f string) ([]string, error) {
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

func LoadGroups(f string) ([][]string, error) {
	file, err := os.Open(f)
	if err != nil {
		return nil, fmt.Errorf("failed to open file '%s': %w", f, err)
	}

	var group int
	rets := [][]string{{}}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() == "" {
			group++
			rets = append(rets, []string{})
			continue
		}

		rets[group] = append(rets[group], scanner.Text())
	}

	return rets, nil
}
