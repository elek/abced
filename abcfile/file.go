package abcfile

import (
	"regexp"
	"strings"
)

type AbcFile struct {
	Headers []*Header
	Tunes   []*Tune
}

type Headers []Header

type Score string

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

func (s *Tune) Items() []interface{} {
	var res []interface{}
	headerPattern := regexp.MustCompile("([a-zA-Z]):\\s*(.+)")
	for _, l := range strings.Split(s.Raw, "\n") {
		l = strings.TrimSpace(l)
		if l == "" {
			continue
		}

		parts := headerPattern.FindStringSubmatch(l)

		if len(parts) == 3 {
			res = append(res, Header{
				Key:   parts[1],
				Value: parts[2],
			})
			continue
		} else {
			res = append(res, Score(l))
		}
	}
	return res
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

func (s *Tune) HeadersAndScores() (h Headers, sc []Score) {
	for _, i := range s.Items() {
		switch t := i.(type) {
		case Score:
			sc = append(sc, t)
		case Header:
			h = append(h, t)
		}
	}
	return h, sc
}

func extract(key string, lines string) string {
	for _, line := range strings.Split(lines, "\n") {
		if strings.HasPrefix(strings.ToLower(line), strings.ToLower(key)+":") {
			return strings.TrimSpace(strings.SplitN(line, ":", 2)[1])
		}
	}
	return ""
}
