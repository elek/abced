package ast

import (
	"fmt"
	"strconv"
)

type TreePrinter struct {
	prefix string
}

func (t TreePrinter) Visit(node Node) (w Visitor) {
	if node == nil {
		return
	}

	fmt.Printf("%s %d %T %s\n", t.prefix, node.Pos(), node, node)

	return TreePrinter{
		prefix: t.prefix + "   ",
	}
}

type Renderer struct {
	pos      Pos
	Rendered string
}

func (r *Renderer) Visit(node Node) (w Visitor) {
	if node == nil {
		return
	}
	for node.Pos() > r.pos {
		r.print(" ")
	}
	switch n := node.(type) {
	case *Line:
	case Note:
	case *Length:
	case Broken:
	case Symbol:
		r.print(n.Symbol)
	case Letter:
		r.print(n.Pitch)
	case *Number:
		r.print(strconv.Itoa(n.Value))
	case *Divider:
		for i := 0; i < n.Number; i++ {
			r.print("/")
		}
	case Bar:
		r.print(n.Symbol)
	default:
		panic(fmt.Sprintf("unsupported node %T", node))
	}
	return r
}

func (r *Renderer) print(s string) {
	r.Rendered += s
	r.pos += Pos(len(s))
}

func Render(node Node) string {
	r := &Renderer{}
	Walk(r, node)
	return r.Rendered
}
