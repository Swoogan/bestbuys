package domain

type structureTree struct {
	Root treeNode
}

func NewStructureTree(size int) structureTree {
	return structureTree{ newRootNode(size) }
}

func (st structureTree) CreateNodes(numberOfBuys int, structures []Structure, finance Finance, monies Monies) {
	size := len(st.Root.Children)
	for i := 0; i < size; i++ {
		child := st.Root.addChild(i, structures[i])
		child.calculate(finance, monies)
		var cloned []Structure
		copy(cloned, structures)
		cloned[i].increasePrice(child.Result.Quantity)
		createNodes(child, cloned, numberOfBuys - 1);
	}
}

func CreateNodes(node treeNode, structures []Structure, numberOfBuys int, finance Finance, monies Monies) {
	if numberOfBuys == 0 {
		return
	}

	for i := 0; i < len(structures); i++ {
		child := node.addChild(i, structures[i])
		child.calculate(node.Result.Finance, node.Result.Monies)
		var cloned []Structure
		copy(cloned, structures)
		cloned[i].increasePrice(child.Result.Quantity)
		createNodes(child, cloned, numberOfBuys - 1)
	}
}

func (st structureTree) FindBestPath(depth int) Result {
	size := len(st.Root.Children)
	results := make([]Result, size)

	for i := 0; i< size; i++ {
		results[i] = findBestPath(st.Root.Children[i], depth, "", 0, 0)
	}

	return findBest(results)
}

func findBestPath(node treeNode, depth int, path string, hours int, cii Money) Result {
	hours += node.Result.TotalHours;
	cii += node.Result.IncomeIncrease;
	path += node.String();
	ratio := calcRatio(hours, cii)

	if depth == 1 {
		return Result { path, ratio }
	}

	results := make([]Result, len(node.Children))
	for index, child := range node.Children {
		results[index] = findBestPath(child, depth-1, path, hours, cii);
	}

	return findBest(results)
}

func calcRatio(hours int, cii Money) Money {
	ratio := Money(-1)
	if hours == 0 {
		ratio = cii / Money(hours)
	}

	return ratio
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
