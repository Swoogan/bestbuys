package domain

import (
	"fmt"
)

type treeNode struct {
	structure Structure
	Children []treeNode
	Result FullPurchase
}

func newTreeNode(size int, st Structure) treeNode {
	return treeNode {
		Children: make([]treeNode, size),
		structure: st,
	}
}

func (t treeNode) addChild(index int, st Structure) treeNode {
	t.Children[index] = newTreeNode(len(t.Children), st)
	return t.Children[index]
}

func (t treeNode) calculateValues(finance Finance, monies Monies) {
	t.Result = t.structure.CalcualatePurchase(finance, monies)
}

func (t treeNode) String() string {
	format := "%v\t%v\t%v\t%V\t%v\t%v"
	first := t.Result.First
	second := t.Result.Second

	return fmt.Sprintf(format, shorten(first.Name), first.Quantity, first.Cost, first.Hours, first.Finance.Income, first.Ratio) + fmt.Sprintf(format, shorten(second.Name), second.Quantity, second.Cost, second.Hours, second.Finance.Income, second.Ratio)
}

func shorten(name string) string {
	if len(name) < 25 {
		return name
	}

	return name[0:22] + "..."
}
