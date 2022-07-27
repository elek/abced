package music

func sliceToPitch(strings []string) []Pitch {
	res := make([]Pitch, len(strings))
	for ix, note := range strings {
		v := Pitch(0)
		for _, r := range note {
			if r == '♭' {
				v--
			} else if r == '♯' {
				v++
			}
			for pr, pv := range pitches {
				if pr == r {
					v += pv
				}
			}
		}
		res[ix] = v
	}
	return res
}
