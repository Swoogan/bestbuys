package domain

import (
	"fmt"
	"log"
)

type node struct {
	Children []*node
	Purchase FullPurchase
}

func newNode(purchase FullPurchase) *node {
	return &node{Purchase: purchase}
}

func (n *node) addChildren(size int, structures Structures, f Finance, depth int) {
	if depth == 0 {
		return
	}

	n.Children = make([]*node, size)

	i := 0
	for key, st := range structures {
		purchase := FullPurchase{}
		purchase.Purchase(st, f)
		child := newNode(purchase)
		structures[key].IncreasePrice()
		child.addChildren(size, structures, f, depth-1)
		n.Children[i] = child
		i++
	}
}

func (n *node) findBestPath(depth int, path string, hours int, cii Money) Result {
	hours += n.Result.Hours
	cii += n.Increase
	path += n.String()
	ratio := calcRatio(hours, cii)

	if depth == 1 {
		return Result{path, ratio}
	}

	results := make([]Result, len(n.Children))

	for i, child := range n.Children {
		results[i] = child.findBestPath(depth, path, hours, cii)
	}

	return findBest(results)
}

func calcRatio(hours int, cii Money) Money {
	ratio := Money(-1)
	if hours != 0 {
		ratio = cii / Money(hours)
	}

	return ratio
}

func (n *node) String() string {
	f := n.Purchase.First
	s := n.Purchase.Second

	format := "%v\t\t %v\t $%f\t %v\t $%f\t %f\n"
	return fmt.Sprintf(format, shorten(f.Name), f.Quantity, f.Cost, f.Hours, f.NewIncome, f.Ratio) +
		fmt.Sprintf(format, shorten(s.Name), s.Quantity, s.Cost, s.Hours, s.NewIncome, s.Ratio)
}

func shorten(name string) string {
	if len(name) < 25 {
		return name
	}

	return name[0:22] + "..."
}

func (n *node) Print(logger *log.Logger) {
	logger.Printf("%v %v %f %d %d", n.First, n.Second, n.Increase, n.Quantity, n.Hours)

	for _, child := range n.Children {
		child.Print(logger)
	}
}
