package uis

import (
	"fmt"
	"n-puzzle-problem/lists"
	"n-puzzle-problem/nodes"
	"time"
)

func PrintResults(node *nodes.Node, openList *lists.List[nodes.Node], closedList *lists.List[nodes.Node], elapsedTime *time.Duration) {
	fmt.Println("\nSuccess!")
	fmt.Println("> Elapsed time:", elapsedTime.Seconds())
	fmt.Println("> Moves:", node.GetMovesToSolution())
	fmt.Println("> Total visited nodes:", closedList.Length)
	fmt.Println("> Depth:", node.Depth)
	fmt.Println("> Final state:", node.State)
}

func PrintNode(node *nodes.Node) {
	fmt.Println("\n>>> Node:")
	fmt.Println("> Depth:", node.Depth)
	fmt.Println("> Heuristic value:", node.HeuristicValue)
	fmt.Println("> TotalCost:", node.TotalCost)
	printState(node.State)
}

func printState(state [3][3]uint8) {
	for i := range state {
		fmt.Println(state[i])
	}
}

func PrintChildren(children []*nodes.Node) {
	fmt.Println("\n>>> Children nodes:")

	firstRows := [][3]uint8{}
	secondRows := [][3]uint8{}
	thirdRows := [][3]uint8{}

	redirectRowsMap := make(map[int]func(row [3]uint8))
	redirectRowsMap[0] = func(row [3]uint8) { firstRows = append(firstRows, row) }
	redirectRowsMap[1] = func(row [3]uint8) { secondRows = append(secondRows, row) }
	redirectRowsMap[2] = func(row [3]uint8) { thirdRows = append(thirdRows, row) }

	for _, child := range children {
		for i, row := range child.State {
			redirectRowsMap[i](row)
		}
	}

	for _, row := range firstRows {
		fmt.Print(row, "   ")
	}
	fmt.Println()

	for _, row := range secondRows {
		fmt.Print(row, "   ")
	}
	fmt.Println()

	for _, row := range thirdRows {
		fmt.Print(row, "   ")
	}
	fmt.Println()
}
