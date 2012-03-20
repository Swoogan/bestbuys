package domain

type structureTree struct {
	root treeNode
	numberOfStructs int
}

func newStructureTree(numberOfStructs int) structureTree {
	return structureTree{ newRootNode(numberOfStructs), numberOfStructs }
}

func (st structureTree) findBestPath(depth int) Result {
	results := make([]Result, st.numberOfStructs)

	for i := 0; i< st.numberOfStructs; i++ {
		results[i] = findBestPath(st.root.Children[i], depth, "", 0, 0)
	}

	return findBest(results)
}

func (st structureTree) createNodes(numberOfBuys int, structures []Structure, finance Finance, monies Monies) {
	for i := 0; i < st.numberOfStructs; i++ {
		child := st.root.addChild(i, structures[i])
		child.calculate(finance, monies)
		var cloned []Structure
		copy(cloned, structures)
		cloned[i].increasePrice(child.Result.Quantity)
		createNodes(child, cloned, numberOfBuys - 1);
	}
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

func createNodes(node treeNode, structures []Structure, numberOfBuys int) {
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
		if result.ratio > best.ratio {
			best = result
		}
	}

	return best
}
