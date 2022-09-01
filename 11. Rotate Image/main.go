package main

func rotate(matrix [][]int) {
	matrixSize := len(matrix)

	startRow := 0
	endRow := matrixSize - 1
	startColumn := 0
	endColumn := matrixSize - 1
	for i := 0; i < matrixSize; i++ {
		totalCycles := endRow - startRow

		for j := 0; j < totalCycles; j++ {
			tmp := matrix[startRow][startColumn+j]
			matrix[startRow][startColumn+j] = matrix[endRow-j][startColumn]

			matrix[endRow-j][startColumn] = matrix[endRow][endColumn-j]

			matrix[endRow][endColumn-j] = matrix[startRow+j][endColumn]

			matrix[startRow+j][endColumn] = tmp
		}
		startRow = startRow + 1
		endRow = endRow - 1
		startColumn = startColumn + 1
		endColumn = endColumn - 1
	}
}
