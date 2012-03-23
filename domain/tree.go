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
		purchase := FullPurchase{}
		purchase.Purchase(st, f)
		child := newNode(purchase)
		structures[key].IncreasePrice()
		child.addChildren(t.Size, structures, f, depth-1)
		t.Children[i] = child
		i++
	}
}

func (t *Tree) Print(logger *log.Logger) {
	for _, child := range t.Children {
		child.Print(logger)
	}
}

func (t *Tree) FindBestPath(depth int, buys []Buy, hours int, cii Money) Result {
	results := make([]Result, t.Size)

	for i, child := range t.Children {
		results[i] = child.findBestPath(depth, buys, hours, cii)
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
