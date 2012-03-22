package domain

import (
	"fmt"
	"log"
)

type node struct {
	Children []*node
	Result   FullPurchase
}

func newNode(st Structure, f Finance) *node {
	return &node{Result: st.purchase(f)}
}

func (n *node) Print(logger *log.Logger) {
	logger.Println(n.Result)

	for _, child := range n.Children {
		child.Print(logger)
	}
}

func (n *node) addChildren(size int, structures Structures, f Finance, depth int) {
	if depth == 0 {
		return
	}

	n.Children = make([]*node, size)

	i := 0
	for key, st := range structures {
		child := newNode(st, f)
		structures[key].increasePrice(child.Result.Quantity)
		child.addChildren(size, structures, f, depth-1)
		n.Children[i] = child
		i++
		//              log.Println("created", child, depth)
	}
}

func (n *node) findBestPath(depth int, path string, hours int, cii Money) Result {
	//	fmt.Printf("Hours: %v, CII: %v \n", node.Result.Hours(), node.Result.IncomeIncrease)
	hours += n.Result.Hours()
	cii += n.Result.IncomeIncrease
	path += n.String()
	ratio := calcRatio(hours, cii)

	//fmt.Printf("Hours: %v, CII: %v, Ratio: %v \n", hours, cii, ratio)

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
	f := n.Result.First
	s := n.Result.Second

	format := "%v\t\t %v\t $%f\t %v\t $%f\t %f\n"
	return fmt.Sprintf(format, shorten(f.Name), f.Quantity, f.Cost, f.Hours, f.Finance.Income, f.Ratio) +
		fmt.Sprintf(format, shorten(s.Name), s.Quantity, s.Cost, s.Hours, s.Finance.Income, s.Ratio)
}

func shorten(name string) string {
	if len(name) < 25 {
		return name
	}

	return name[0:22] + "..."
}
