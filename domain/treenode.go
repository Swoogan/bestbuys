package domain

import (
	"fmt"
)

type treeNode struct {
	Structure Structure
	Children []treeNode
	Finance Finance
	Monies Monies
	Result FullPurchase
	Size int
}

func newTreeNode(size int, st Structure, f Finance, m Monies) treeNode {
	return }

func NewRootNode(size int, f Finance, m Monies) treeNode {
	return treeNode { 
		Children: make([]treeNode, size), 
		Size: size,
		Finance: f,
		Monies: m,
	}
}

func (t treeNode) addChild(i int, st Structure, f Finance, m Monies) treeNode {
	child := treeNode {
		Children: make([]treeNode, t.Size),
		Structure: st,
		Finance: f,
		Monies: m,
		Size: t.Size,
	}

	t.Children[[i] = child

	return child
}

func (t treeNode) calculate() {
	t.Result = t.Structure.purchase(t.Finance, t.Monies)
}

func (t treeNode) String() string {
	format := "%v\t%v\t%v\t%V\t%v\t%v"
	f := t.Result.First
	s := t.Result.Second

	return fmt.Sprintf(format, shorten(f.Name), f.Quantity, f.Cost, f.Hours, f.Finance.Income, f.Ratio) + fmt.Sprintf(format, shorten(s.Name), s.Quantity, s.Cost, s.Hours, s.Finance.Income, s.Ratio)
}

func shorten(name string) string {
	if len(name) < 25 {
		return name
	}

	return name[0:22] + "..."
}
