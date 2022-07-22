package abcfile

import "strings"

type AbcFile struct {
	Headers []*Header
	Scores  []*Score
}

func (f *AbcFile) AddScore(abc string) {
	f.Scores = append(f.Scores, NewScore(abc))
}

func (f *AbcFile) GetScore(ID string) *Score {
	for _, s := range f.Scores {
		if s.ID() == ID {
			return s
		}
	}
	return nil
}

type Header struct {
	Key   string
	Value string
}

type Score struct {
	Raw string
}

func NewScore(c string) *Score {
	return &Score{
		Raw: c,
	}
}

func (s *Score) SetValue(content string) {
	s.Raw = content
}

func (s *Score) HeadersAndLines() ([]Header, []string) {
	var headers []Header
	var lines []string

	for _, l := range strings.Split(s.Raw, "\n") {
		l = strings.TrimSpace(l)
		if l == "" {
			continue
		}

		parts := strings.SplitN(l, ":", 2)
		if len(lines) > 0 || len(parts) == 1 {
			lines = append(lines, l)
			continue
		}
		headers = append(headers, Header{
			Key:   strings.TrimSpace(parts[0]),
			Value: strings.TrimSpace(parts[1]),
		})
	}
	return headers, lines
}

func (s *Score) ID() string {
	return extract("X", s.Raw)
}

func (s *Score) Title() string {
	return extract("T", s.Raw)
}

func extract(key string, lines string) string {
	for _, line := range strings.Split(lines, "\n") {
		if strings.HasPrefix(strings.ToLower(line), strings.ToLower(key)+":") {
			return strings.TrimSpace(strings.SplitN(line, ":", 2)[1])
		}
	}
	return ""
}
