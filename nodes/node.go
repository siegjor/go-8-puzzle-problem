package nodes

import (
	"n-puzzle-problem/heuristics"
	mutils "n-puzzle-problem/matrix_utils"
)

type Node struct {
	parent                           *Node
	State                            [3][3]uint8
	children                         []*Node
	Depth, HeuristicValue, TotalCost uint8
	lastMovement                     string
}

// const SELECTED_STRATEGY = heuristics.UNIFORM_COST

const SELECTED_STRATEGY = heuristics.A_STAR_MANHATTAN

// const SELECTED_STRATEGY = heuristics.UNIFORM_COST

func getEmptyTileCoords(node *Node) mutils.Pos {
	emptyTilePos := mutils.Pos{}
	for i := range node.State {
		for j := range node.State[i] {
			if node.State[i][j] == 0 {
				emptyTilePos.X = int8(i)
				emptyTilePos.Y = int8(j)
				break
			}
		}
	}

	return emptyTilePos
}

func (nodePtr *Node) GenChildren() []*Node {
	emptyTilePos := getEmptyTileCoords(nodePtr)

	if emptyTilePos.Y-1 >= 0 {
		moveEmptyTileLeft(nodePtr, emptyTilePos)
	}

	if emptyTilePos.Y+1 <= 2 {
		moveEmptyTileRight(nodePtr, emptyTilePos)
	}

	if emptyTilePos.X+1 <= 2 {
		moveEmptyTileDown(nodePtr, emptyTilePos)
	}

	if emptyTilePos.X-1 >= 0 {
		moveEmptyTileUp(nodePtr, emptyTilePos)
	}

	for _, child := range nodePtr.children {
		child.Depth = nodePtr.Depth + 1
		calculateTotalCost(child)
	}

	return nodePtr.children
}

func calculateTotalCost(node *Node) {
	switch SELECTED_STRATEGY {
	case heuristics.UNIFORM_COST:
		node.TotalCost = node.Depth
		break
	case heuristics.A_STAR_MANHATTAN:
		node.HeuristicValue = heuristics.GetSumOfManhattanDistance(node.State)
		node.TotalCost = node.HeuristicValue + node.Depth
		break
	case heuristics.A_STAR_LINEAR_CONFLICT:
		break
	default:
		break
	}
}

func moveEmptyTileUp(node *Node, emptyTilePos mutils.Pos) {
	newState := node.State
	newState[emptyTilePos.X][emptyTilePos.Y] = node.State[emptyTilePos.X-1][emptyTilePos.Y]
	newState[emptyTilePos.X-1][emptyTilePos.Y] = 0

	newNode := Node{State: newState, parent: node, lastMovement: "up"}
	node.children = append(node.children, &newNode)
}

func moveEmptyTileDown(node *Node, emptyTilePos mutils.Pos) {
	newState := node.State
	newState[emptyTilePos.X][emptyTilePos.Y] = node.State[emptyTilePos.X+1][emptyTilePos.Y]
	newState[emptyTilePos.X+1][emptyTilePos.Y] = 0

	newNode := Node{State: newState, parent: node, lastMovement: "down"}
	node.children = append(node.children, &newNode)
}

func moveEmptyTileLeft(node *Node, emptyTilePos mutils.Pos) {
	newState := node.State
	newState[emptyTilePos.X][emptyTilePos.Y] = node.State[emptyTilePos.X][emptyTilePos.Y-1]
	newState[emptyTilePos.X][emptyTilePos.Y-1] = 0

	newNode := Node{State: newState, parent: node, lastMovement: "left"}
	node.children = append(node.children, &newNode)
}

func moveEmptyTileRight(node *Node, emptyTilePos mutils.Pos) {
	newState := node.State
	newState[emptyTilePos.X][emptyTilePos.Y] = node.State[emptyTilePos.X][emptyTilePos.Y+1]
	newState[emptyTilePos.X][emptyTilePos.Y+1] = 0

	newNode := Node{State: newState, parent: node, lastMovement: "right"}
	node.children = append(node.children, &newNode)
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

func (node *Node) IsGoal() bool {
	goalState := [3][3]uint8{{1, 2, 3}, {4, 5, 6}, {7, 8, 0}}
	goalNode := Node{State: goalState}
	return node.Equals(&goalNode)
}

func (node *Node) GetMovesToSolution() []string {
	moves := []string{}
	tempNode := node
	for tempNode.parent != nil {
		moves = append(moves, tempNode.lastMovement)
		tempNode = tempNode.parent
	}

	for i, j := 0, len(moves)-1; i < j; i, j = i+1, j-1 {
		moves[i], moves[j] = moves[j], moves[i]
	}

	return moves
}
