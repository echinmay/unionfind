// unionfind implements one path variant of the path compression union find algorithm.
// This is inspired from the Algorithms, Part I course @ http://www.coursera.org.
// This is given by Princeton university by Kevin Wayna and Robert Sedgewick
package unionfind

//
//	Unionfind has many applications. The two main operations of use are -
//	find - determine which subset a particular element belongs to.
//	union - join two subsets

type UnionFind struct {
	arr   []int
	count int
}

// New initializes a new UnionFind struct and returns a pointer.
func New(count int) *UnionFind {

	id := make([]int, count)
	for i := 0; i < count; i++ {
		id[i] = i
	}
	return &UnionFind{arr: id, count: count}
}

func (ufd UnionFind) validate(idx int) bool {
	if idx > 0 && idx >= len(ufd.arr) {
		return false
	}
	return true
}

func (ufd UnionFind) getroot(i int) int {
	if !ufd.validate(i) {
		return -1
	}

	for ufd.arr[i] != i {
		// Set to the grandparent here. This is the path compression part.
		// We are flattenning the tree.
		ufd.arr[i] = ufd.arr[ufd.arr[i]]
		i = ufd.arr[i]
	}
	return i
}

func (ufd UnionFind) Connected(p, q int) bool {
	if ufd.validate(p-1) && ufd.validate(q-1) {
		proot := ufd.getroot(p-1)
		qroot := ufd.getroot(q-1)
		return proot == qroot
	}
	return false

}

func (ufd UnionFind) GetNumClusters() int {
	return ufd.count
}

func (ufd *UnionFind) Union(p, q int) {
	if !ufd.validate(p-1) || !ufd.validate(q-1) {
		return
	}

	proot := ufd.getroot(p-1)
	qroot := ufd.getroot(q-1)

	if proot == qroot {
		// Already coonected. Nothing left to do.
		return
	}
	ufd.arr[proot] = ufd.arr[qroot]
	// Number of subsets decrease by one
	ufd.count -= 1
	return
}
