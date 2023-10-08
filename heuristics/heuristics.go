package heuristics

import (
	"math"
	mutils "n-puzzle-problem/matrix_utils"
)

var solutionStateMap = map[uint8]mutils.Pos{
	1: {X: 0, Y: 0},
	2: {X: 0, Y: 1},
	3: {X: 0, Y: 2},
	4: {X: 1, Y: 0},
	5: {X: 1, Y: 1},
	6: {X: 1, Y: 2},
	7: {X: 2, Y: 0},
	8: {X: 2, Y: 1},
}

func calculateManhattanDistance(posA mutils.Pos, posB mutils.Pos) uint8 {
	xDistance := math.Abs(float64(posA.X) - float64(posB.X))
	yDistance := math.Abs(float64(posA.Y) - float64(posB.Y))
	return uint8(xDistance + yDistance)
}

func GetSumOfManhattanDistance(state [3][3]uint8) uint8 {
	var sum uint8 = 0
	for i := range state {
		for j := range state[i] {
			if state[i][j] == 0 {
				continue
			}
			currentPos := mutils.Pos{X: int8(i), Y: int8(j)}
			solutionPos := solutionStateMap[state[i][j]]
			// fmt.Println("current pos: (", state[i][j], "):", currentPos)
			// fmt.Println("solution pos: (", state[i][j], "):", solutionPos)
			// fmt.Println("manhattan for", currentPos, ":", calculateManhattanDistance(currentPos, solutionPos))
			sum += calculateManhattanDistance(currentPos, solutionPos)
		}
	}

	return sum
}
