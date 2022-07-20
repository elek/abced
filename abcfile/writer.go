package abcfile

import (
	"fmt"
	"io/ioutil"
)

func WriteFile(filename string, file *AbcFile) error {
	s := ""
	for _, header := range file.Headers {
		s += fmt.Sprintf("%s:%s\n", header.Key, header.Value)
	}
	if s != "" {
		s += "\n"
	}
	for _, c := range file.Scores {
		s += c.Raw + "\n"
	}
	return ioutil.WriteFile(filename, []byte(s), 0644)
}
