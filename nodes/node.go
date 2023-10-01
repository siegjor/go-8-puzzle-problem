package nodes

import (
	"fmt"
	"n-puzzle-problem/heuristics"
)

type Node struct {
	parent                           *Node
	State                            [3][3]uint8
	children                         []*Node
	depth, heuristicValue, TotalCost uint64
}

type Pos struct {
	X uint8
	Y uint8
}

const SELECTED_STRATEGY = heuristics.UNIFORM_COST

func getEmptyTileCoords(node *Node) Pos {
	emptyTilePos := Pos{}
	for i := range node.State {
		for j := range node.State[i] {
			if node.State[i][j] == 0 {
				emptyTilePos.X = uint8(i)
				emptyTilePos.Y = uint8(j)
				break
			}
		}
	}

	return emptyTilePos
}

func (node *Node) GenChildren() []*Node {
	emptyTilePos := getEmptyTileCoords(node)

	if emptyTilePos.X+1 <= 2 {
		fmt.Println("> moved down")
		moveEmptyTileDown(*node, emptyTilePos)
	}

	if emptyTilePos.X-1 >= 0 {
		fmt.Println("> moved up")
		moveEmptyTileUp(*node, emptyTilePos)
	}

	if emptyTilePos.Y+1 <= 2 {
		fmt.Println("> moved right")
		moveEmptyTileRight(*node, emptyTilePos)
	}

	if emptyTilePos.Y-1 >= 0 {
		fmt.Println("> moved left")
		moveEmptyTileLeft(*node, emptyTilePos)
	}

	for _, child := range node.children {
		child.depth = node.depth + 1
		calculateTotalCost(child)
	}

	return node.children
}

func calculateTotalCost(node *Node) {
	switch SELECTED_STRATEGY {
	case heuristics.UNIFORM_COST:
		node.TotalCost = node.depth
		break
	case heuristics.A_STAR_MANHATTAN:
		break
	case heuristics.A_STAR_LINEAR_CONFLICT:
		break
	default:
		break
	}
}

func moveEmptyTileUp(node Node, emptyTilePos Pos) {
	newState := node.State
	newState[emptyTilePos.X][emptyTilePos.Y] = node.State[emptyTilePos.X-1][emptyTilePos.Y]
	newState[emptyTilePos.X-1][emptyTilePos.Y] = 0

	node.children = append(node.children, &Node{State: newState, parent: &node})
}

func moveEmptyTileDown(node Node, emptyTilePos Pos) {
	newState := node.State
	newState[emptyTilePos.X][emptyTilePos.Y] = node.State[emptyTilePos.X+1][emptyTilePos.Y]
	newState[emptyTilePos.X+1][emptyTilePos.Y] = 0

	node.children = append(node.children, &Node{State: newState, parent: &node})
}

func moveEmptyTileLeft(node Node, emptyTilePos Pos) {
	newState := node.State
	newState[emptyTilePos.X][emptyTilePos.Y] = node.State[emptyTilePos.X][emptyTilePos.Y-1]
	newState[emptyTilePos.X][emptyTilePos.Y-1] = 0

	node.children = append(node.children, &Node{State: newState, parent: &node})
}

func moveEmptyTileRight(node Node, emptyTilePos Pos) {
	newState := node.State
	newState[emptyTilePos.X][emptyTilePos.Y] = node.State[emptyTilePos.X][emptyTilePos.Y+1]
	newState[emptyTilePos.X][emptyTilePos.Y+1] = 0

	node.children = append(node.children, &Node{State: newState, parent: &node})
}

func (node *Node) Equals(other *Node) bool {
	for i := range node.State {
		for j := range node.State[i] {
			if node.State[i][j] != other.State[i][j] {
				return false
			}
		}
	}

	return true
}
