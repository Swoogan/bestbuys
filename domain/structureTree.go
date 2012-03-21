package domain

//import "fmt"

func CreateNodes(node treeNode, structures []Structure, numberOfBuys int) {
	if numberOfBuys == 0 {
		return
	}

	for i := 0; i < node.Size; i++ {
		child := node.addChild(i, structures[i], node.Finance, node.Monies)
		child.calculate()
		//cloned := make([]Structure, node.Size)
		//copy(cloned, structures)
		//cloned[i].increasePrice(child.Result.Quantity)
		//CreateNodes(child, cloned, numberOfBuys - 1)
	}
}

func FindBestChild(node treeNode, depth int, path string, hours int, cii Money) Result {
	results := make([]Result, node.Size)

	for i, child := range node.Children {
		results[i] = findBestPath(child, depth, path, hours, cii)
	}

	return findBest(results)
}

func findBestPath(node treeNode, depth int, path string, hours int, cii Money) Result {
	hours += node.Result.TotalHours
	cii += node.Result.IncomeIncrease
	path += node.String()
	ratio := calcRatio(hours, cii)

	if depth == 1 {
		return Result{path, ratio}
	}

	return FindBestChild(node, depth-1, path, hours, cii)
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
