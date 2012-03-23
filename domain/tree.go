package domain

import "log"

type Tree struct {
	Children []*node
	Size     int
}

func NewTree(size int) *Tree {
	return &Tree{Size: size}
}

func (t *Tree) Build(structures Structures, f Finance, depth int) {
	i := 0
	t.Children = make([]*node, t.Size)
	for key, st := range structures {
		purchase := st.purchase(f)
		child := newNode(purchase)
		structures[key].increasePrice(purchase.Quantity)
		child.addChildren(t.Size, structures, f, depth-1)
		t.Children[i++] = child
		//log.Println("created", child, depth)
	}
}

func (t *Tree) Print(logger *log.Logger) {
	for _, child := range t.Children {
		child.Print(logger)
	}
}

func (t *Tree) FindBestPath(depth int, path string, hours int, cii Money) Result {
	results := make([]Result, t.Size)

	for i, child := range t.Children {
		results[i] = child.findBestPath(depth, path, hours, cii)
	}

	return findBest(results)
}

func findBest(results []Result) Result {
	var best Result

	for _, result := range results {
		if result.Ratio > best.Ratio {
			best = result
		}
	}

	return best
}

