package uis

import (
	"bufio"
	"fmt"
	"n-puzzle-problem/lists"
	mutils "n-puzzle-problem/matrix_utils"
	"n-puzzle-problem/nodes"
	"os"
	"strconv"
	"strings"
	"time"
)

func GetMatrixFromUser() [3][3]uint8 {
	var matrixNumbers [9]uint8
	for i := 0; i < 999; i++ {
		matrixNumbers = getUserInput()
		if !mutils.PuzzleHasSolution(matrixNumbers) {
			fmt.Println("This puzzle configuration has no solution. Please enter another one!")
		} else {
			break
		}
	}

	return mutils.ConvertArrayToMatrix(matrixNumbers)
}

func getUserInput() [9]uint8 {
	fmt.Println("\nEnter the matrix numbers separated by a whitespace:")
	fmt.Println("(ex: 1 8 2 0 4 3 7 6 5)")

	var input string
	var matrixNumbers [9]uint8

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	stringNumbers := strings.Fields(input)

	for i, strNumber := range stringNumbers {
		num, err := strconv.Atoi(strNumber)
		if err != nil {
			panic(err)
		}
		matrixNumbers[i] = uint8(num)
	}

	return matrixNumbers
}

func PrintResults(node *nodes.Node, closedList *lists.List[nodes.Node], elapsedTime *time.Duration) {
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
