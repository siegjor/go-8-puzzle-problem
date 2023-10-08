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

func GetAStarManhattanHeuristicValue(state [3][3]uint8) uint8 {
	return GetSumOfManhattanDistance(state)
}

func GetAStarLinearConflictHeuristicValue(state [3][3]uint8) uint8 {
	manhattanDistanceSum := GetSumOfManhattanDistance(state)
	linearConflicts := calculateLinearConflicts(state)

	return manhattanDistanceSum + linearConflicts
}

func calculateManhattanDistance(posA mutils.Pos, posB mutils.Pos) uint8 {
	xDistance := math.Abs(float64(posA.X) - float64(posB.X))
	yDistance := math.Abs(float64(posA.Y) - float64(posB.Y))
	return uint8(xDistance + yDistance)
}

func GetSumOfManhattanDistance(state [3][3]uint8) uint8 {
	var sum uint8 = 0
	currentStatePosMap := getStatePositionsMap(state)
	for value, currentPos := range currentStatePosMap {
		solutionPos := solutionStateMap[value]
		sum += calculateManhattanDistance(currentPos, solutionPos)
	}

	return sum
}

func calculateLinearConflicts(state [3][3]uint8) uint8 {
	var linearConflicts uint8 = 0

	currentPosMap := getStatePositionsMap(state)
	for currentValue, currentPos := range currentPosMap {
		currentSolutionPos := solutionStateMap[currentValue]
		for tempX := range state {
			for tempY := range state[tempX] {
				tempValue := state[tempX][tempY]
				if tempValue == 0 {
					continue
				}

				tempSolutionPos := solutionStateMap[tempValue]
				if currentPos.X == int8(tempX) && currentSolutionPos.X == tempSolutionPos.X {
					if int8(tempY) < currentPos.Y && currentSolutionPos.Y < tempSolutionPos.Y {
						linearConflicts++
						continue
					}
				}

				if currentPos.Y == int8(tempY) && currentSolutionPos.Y == tempSolutionPos.Y && currentPos.Y == currentSolutionPos.Y {
					if int8(tempX) < currentPos.X && currentSolutionPos.X < tempSolutionPos.X {
						linearConflicts++
						continue
					}
				}
			}
		}
	}

	return linearConflicts * 2
}

func getStatePositionsMap(state [3][3]uint8) map[uint8]mutils.Pos {
	positionsMap := make(map[uint8]mutils.Pos)
	for i := range state {
		for j := range state[i] {
			value := state[i][j]
			if value == 0 {
				continue
			}
			positionsMap[value] = mutils.Pos{X: int8(i), Y: int8(j)}
		}
	}

	return positionsMap
}
