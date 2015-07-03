package unionfind

import (
	"testing"

)

func TestNewUnionFind(t *testing.T) {
	var uf *UnionFind

	uf = New(9)
	if !slicesEqual(uf.arr, []int{0, 1, 2, 3, 4, 5, 6, 7, 8}) {
		t.Error("wrong initial array mapping with N=9")
	}
	if uf.count != 9 {
		t.Error("wrong count with N=9")
	}

}

func TestUnion(t *testing.T) {
	var uf *UnionFind

	uf = New(9)
	uf.Union(3, 4)
	uf.Union(1, 3)


}

func TestConnected(t *testing.T) {
	uf := New(9)
	uf.Union(3, 4)
	uf.Union(1, 3)
	if !uf.Connected(1, 3) {
		t.Error("connected gives false negative on input union")
	}
	if !uf.Connected(1, 4) {
		t.Error("connected gives false negative on transitive union")
	}
	if uf.Connected(6, 7) {
		t.Error("connected gives false positive on random nodes")
	}
	if uf.Connected(3, 7) {
		t.Error("connected gives false positive on reassigned nodes")
	}
	if uf.Connected(1, 7) {
		t.Error("connected gives false positive on reassigned nodes")
	}
}



func slicesEqual(xs, ys []int) bool {
	for i, x := range xs {
		y := ys[i]
		if x != y {
			return false
		}
	}
	return true
}
