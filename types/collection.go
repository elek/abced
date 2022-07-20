package abc

type Abc struct {
	Tunes  []Tune
	Fields map[string]string
}

func NewAbc() *Abc {
	return &Abc{
		Fields: make(map[string]string),
		Tunes:  make([]Tune, 0),
	}
}
