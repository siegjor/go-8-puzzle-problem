package main

import (
	"fmt"
	"n-puzzle-problem/lists"
	"n-puzzle-problem/nodes"
)

func main() {
	matrix := [3][3]uint8{{2, 3, 4}, {1, 7, 8}, {5, 0, 6}}
	fmt.Println("matrix:", len(matrix))
	originNode := nodes.Node{State: matrix}
	fmt.Println(originNode)

	// openList := []*nodes.Node{&originNode}
	openList := lists.List[nodes.Node]{}
	openList.Insert(originNode, originNode.TotalCost)
	closedList := []*nodes.Node{}

	// foundGoal := false
	for i := range [5]uint8{} {
		fmt.Println(i)
		currentNode := openList.GetFirst()

		generatedChildren := currentNode.GenChildren()

		closedList = append(closedList, &currentNode)
		openList.RemoveFirst()

		for _, child := range generatedChildren {
			if openList.Contains(child) {

			}
			// openList.ForEach(func(node nodes.Node) {
			// 	if child.Equals(&node) {

			// 	}
			// })
			// if !slices.Contains(openList, &node) {
			// 	openList = append(openList, &node)
			// }
		}
	}
}
