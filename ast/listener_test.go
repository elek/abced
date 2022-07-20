package ast

import (
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestASTSimple(t *testing.T) {
	tests := []struct {
		line string
	}{
		{
			line: "a a",
		},
		{
			line: "a3/2 a// a/",
		},
		{
			line: "a<B c>d",
		},
		{
			line: "A | B C | d",
		},
		{
			line: "aaaa | a b c d | e e e e",
		},
		{
			line: "f<f e/d/c/d/ | B<G B d",
		},
	}

	for _, ts := range tests {
		t.Run(strings.ReplaceAll(ts.line, "|", ""), func(t *testing.T) {
			ast, err := BuildAST(ts.line)
			Walk(TreePrinter{}, ast)
			require.NoError(t, err)
			require.Equal(t, ts.line, Render(ast))
		})
	}

}

func TestErrors(t *testing.T) {
	tests := []struct {
		line string
	}{
		{
			line: "dAGF | DAGF | GE A<",
		},
	}

	for _, ts := range tests {
		t.Run(strings.ReplaceAll(ts.line, "|", ""), func(t *testing.T) {
			ast, err := BuildAST(ts.line)
			Walk(TreePrinter{}, ast)
			require.Error(t, err)
		})
	}
}
