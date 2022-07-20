package abc

import "fmt"

type Beat struct {
	Nominator   int
	Denominator int
}

func (b Beat) Multiply(o Beat) Beat {
	return NewBeat(b.Nominator*o.Nominator, b.Denominator*o.Denominator).Simplify()
}

func (b Beat) Divide(o Beat) Beat {
	return b.Multiply(NewBeat(o.Denominator, o.Nominator))
}

func (b Beat) String() string {
	return fmt.Sprintf("%d/%d", b.Nominator, b.Denominator)
}

func (b Beat) AbcString() string {
	switch b {
	case NewBeat(1, 1):
		return ""
	case NewBeat(1, 2):
		return "/"
	case NewBeat(1, 4):
		return "//"
	case NewBeat(2, 1):
		return "2"
	case NewBeat(4, 1):
		return "4"
	case NewBeat(3, 2):
		return "3/2"
	default:
		return b.String()
	}
}

func (b Beat) Simplify() Beat {
	if b.Denominator == 0 {
		return b
	}
	d := gcd(b.Nominator, b.Denominator)
	return NewBeat(b.Nominator/d, b.Denominator/d)
}

func (b Beat) Add(o Beat) Beat {
	if o.Denominator == 0 {
		return b
	}
	if b.Denominator == 0 {
		return o
	}
	return NewBeat(b.Nominator*o.Denominator+o.Nominator*b.Denominator, b.Denominator*o.Denominator).Simplify()
}

func NewBeat(nom int, denom int) Beat {
	return Beat{
		Nominator:   nom,
		Denominator: denom,
	}
}

func gcd(a int, b int) int {
	if b == 0 {
		return a
	}
	if a == 0 {
		return b
	}
	if a > b {
		return gcd(a%b, b)
	}
	return gcd(a, b%a)
}
