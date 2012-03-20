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

func findBestPath(node treeNode, depth int, path string, hours int, cii Money) Result {
	hours += node.Result.TotalHours;
	cii += node.Result.IncomeIncrease;
	path += node.String();
	ratio := calcRatio(hours, cii)

	if depth == 1 {
		return Result { path, ratio }
	}

	var results []Result
	for _, child := range node.Children {
		results = append(results, findBestPath(child, depth-1, path, hours, cii));
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
		if result.ratio > best.ratio {
			best = result
		}
	}

	return best
}
