package fio

import (
	"io/ioutil"
	"strings"
)

// Reads a text file (without validation), and returns a []string of lines
func ReadLinesFromFile(fpath string) ([]string, error) {
	bytes, err := ioutil.ReadFile(fpath)
	if err != nil {
		return nil, err
	}

	sContent := string(bytes)

	return strings.Split(strings.TrimSpace(sContent), "\n"), nil
}
