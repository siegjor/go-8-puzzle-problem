package main

import (
	"n-puzzle-problem/lists"
	"n-puzzle-problem/nodes"
	"n-puzzle-problem/uis"
	"time"
)

func main() {
	matrix := [3][3]uint8{{1, 8, 2}, {0, 4, 3}, {7, 6, 5}}

	startTime := time.Now()
	originNode := nodes.Node{State: matrix}

	openList := lists.List[nodes.Node]{}
	openList.Insert(originNode, originNode.TotalCost)
	closedList := lists.List[nodes.Node]{}

	foundGoal := false
	for !foundGoal {
		currentNode := openList.GetFirst()
		// uis.PrintNode(&currentNode)

		generatedChildren := currentNode.GenChildren()
		// uis.PrintChildren(generatedChildren)

		closedList.Insert(currentNode, currentNode.TotalCost)
		openList.RemoveFirst()

		for _, child := range generatedChildren {
			if !openList.Contains(child) && !closedList.Contains(child) {
				openList.Insert(*child, child.TotalCost)
			}
		}

		foundGoal = currentNode.IsGoal()
		if foundGoal {
			elapsedTime := time.Since(startTime)
			uis.PrintResults(&currentNode, &openList, &closedList, &elapsedTime)
		}
	}
}
