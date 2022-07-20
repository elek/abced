package abcfile

import (
	"github.com/zeebo/errs/v2"
	"io/ioutil"
	"strings"
)

func ReadFile(filename string) (AbcFile, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return AbcFile{}, errs.Wrap(err)
	}
	return Read(string(content))
}

func Read(content string) (AbcFile, error) {
	a := AbcFile{}
	header := true
	current := ""
	for _, line := range strings.Split(string(content), "\n") {
		line = strings.TrimSpace(line)
		if header {
			if line == "" {
				header = false
				continue
			}
			if strings.HasPrefix(line, "%") {
				a.Headers = append(a.Headers, &Header{
					Key: line,
				})
				continue
			}
			parts := strings.SplitN(line, ":", 2)
			a.Headers = append(a.Headers, &Header{
				Key:   strings.TrimSpace(parts[0]),
				Value: strings.TrimSpace(parts[1]),
			})

		} else {
			if line == "" {
				if current != "" {
					a.AddScore(current)
				}
				current = ""
			} else {
				current += line + "\n"
			}
		}
	}
	if current != "" {
		a.AddScore(current)
	}
	return a, nil
}
