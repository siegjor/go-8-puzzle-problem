package matrix_utils

func PuzzleHasSolution(numbers [9]uint8) bool {
	count := 0
	for i := 0; i < 8; i++ {
		for j := i + 1; j < 9; j++ {
			if numbers[i] != 0 && numbers[j] != 0 && numbers[i] > numbers[j] {
				count++
			}
		}
	}

	return count%2 == 0
}

func ConvertArrayToMatrix(numbers [9]uint8) [3][3]uint8 {
	matrix := [3][3]uint8{}
	k := 0
	for i := range matrix {
		for j := range matrix[i] {
			matrix[i][j] = numbers[k]
			k++
		}
	}

	return matrix
}
