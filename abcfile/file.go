package abcfile

import "strings"

type AbcFile struct {
	Headers []*Header
	Tunes   []*Tune
}

type Headers []Header

func (f *AbcFile) AddScore(abc string) {
	f.Tunes = append(f.Tunes, NewScore(abc))
}

func (f *AbcFile) GetScore(ID string) *Tune {
	for _, s := range f.Tunes {
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

type Tune struct {
	Raw string
}

func NewScore(c string) *Tune {
	return &Tune{
		Raw: c,
	}
}

func (s *Tune) SetValue(content string) {
	s.Raw = content
}

func (s *Tune) HeadersAndLines() (Headers, []string) {
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

func (h Headers) Get(key string) string {
	for _, k := range h {
		if k.Key == key {
			return k.Value
		}
	}
	return ""
}
func (s *Tune) ID() string {
	return extract("X", s.Raw)
}

func (s *Tune) Title() string {
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
