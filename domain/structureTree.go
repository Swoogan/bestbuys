package domain

type structureTree struct {
	root treeNode
	numberOfStructs int
}

func newStructureTree(numberOfStructs int) structureTree {
	return structureTree{ newRootNode(numberOfStructs), numberOfStructs }
}
