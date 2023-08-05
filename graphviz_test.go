package fuzz

import (
	"testing"

	"github.com/awalterschulze/gographviz"
)

func FuzzGraphviz(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte) {
		graphAst, err := gographviz.Parse(b)
		if err != nil {
			t.Skip()
		}
		graph := gographviz.NewGraph()
		if err := gographviz.Analyse(graphAst, graph); err != nil {
			t.Errorf("error analyzing parsed graph: %v", err)
		}
	})
}
