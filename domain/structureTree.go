package domain

//import "fmt"

func CreateNodes(node *treeNode, structures map[string]Structure, depth int) {
	if depth == 0 {
		return
	}

	for _, st := range structures {
		child := node.addChild(st, node.Finance, node.Monies)
		child.calculate()
		st.increasePrice(child.Result.Quantity)
		CreateNodes(child, structures, depth - 1)
	}
}

func FindBestChild(node *treeNode, depth int, path string, hours int, cii Money) Result {
	results := make([]Result, node.Size)

	for i, child := range node.Children {
		results[i] = findBestPath(child, depth, path, hours, cii)
	}

	return findBest(results)
}

func findBestPath(node *treeNode, depth int, path string, hours int, cii Money) Result {
//	fmt.Printf("Hours: %v, CII: %v \n", node.Result.Hours(), node.Result.IncomeIncrease)
	hours += node.Result.Hours()
	cii += node.Result.IncomeIncrease
	path += node.String()
	ratio := calcRatio(hours, cii)

	//fmt.Printf("Hours: %v, CII: %v, Ratio: %v \n", hours, cii, ratio)

	if depth == 1 {
		return Result{path, ratio}
	}

	return FindBestChild(node, depth-1, path, hours, cii)
}

func calcRatio(hours int, cii Money) Money {
	ratio := Money(-1)
	if hours != 0 {
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
