package file

import (
	"fmt"
	"io/ioutil"
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
