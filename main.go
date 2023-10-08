package main

import (
	"n-puzzle-problem/lists"
	"n-puzzle-problem/nodes"
	"n-puzzle-problem/uis"
	"time"
)

// Easy case:
// Input: 1 8 2 0 4 3 7 6 5

// Medium cases:
// Input: 4 0 6 7 2 3 1 8 5
// Input: 2 0 8 6 4 5 3 1 7

// Hard cases:
// Input: 8 6 7 2 5 4 3 0 1
// Input: 6 4 7 8 5 0 3 2 1

const DEBUG bool = false

func main() {
	matrix := uis.GetMatrixFromUser()
	startTime := time.Now()
	originNode := nodes.Node{State: matrix}

	openList := lists.List[nodes.Node]{}
	openList.Insert(originNode, originNode.TotalCost)
	closedList := lists.List[nodes.Node]{}

	foundGoal := false
	for !foundGoal {
		currentNode := openList.GetFirst()
		if DEBUG {
			uis.PrintNode(&currentNode)
		}

		generatedChildren := currentNode.GenChildren()
		if DEBUG {
			uis.PrintChildren(generatedChildren)
		}

		closedList.Insert(currentNode, currentNode.TotalCost)
		openList.RemoveFirst()

		for _, child := range generatedChildren {
			if !openList.Contains(child) && !closedList.Contains(child) {
				openList.Insert(*child, child.TotalCost)
			}
		}

		// fmt.Println("open list:")
		// costSlices := []uint8{}
		// openList.ForEach(func(node nodes.Node) {
		// 	costSlices = append(costSlices, node.TotalCost)
		// })
		// fmt.Println(costSlices)

		foundGoal = currentNode.IsGoal()
		if foundGoal {
			elapsedTime := time.Since(startTime)
			uis.PrintResults(&currentNode, &closedList, &elapsedTime)
		}
	}
}
