package ast

import "fmt"

type Visitor interface {
	Visit(node Node) (w Visitor)
}

func Walk(v Visitor, node Node) {
	if v = v.Visit(node); v == nil {
		return
	}
	switch n := node.(type) {
	case *Line:
		for _, c := range n.Nodes {
			Walk(v, c)
		}
	case Note:
		Walk(v, n.Letter)
		Walk(v, n.Length)
	case *Length:
		if n.First != nil {
			Walk(v, n.First)
		}
		if n.Divider != nil {
			Walk(v, n.Divider)
		}
		if n.Last != nil {
			Walk(v, n.Last)
		}

	case Broken:
		Walk(v, n.First)
		Walk(v, n.Symbol)
		Walk(v, n.Second)
	case *Number:
	case *Divider:
	case Letter:
	case Symbol:
	case Bar:
	default:
		panic(fmt.Sprintf("Unknown node type %T", n))
	}
	v.Visit(nil)
}
