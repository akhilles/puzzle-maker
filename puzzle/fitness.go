package puzzle

// Evaluate returns the fitness score of a puzzle.
func Evaluate(p []int) int {
	return 0
}

// BFS returns a matrix containing the minimum number of moves needed to reach that cell from the start cell.
func BFS(n int, p []int) []int {
	numCells := n * n
	minDist := make([]int, numCells)
	for index := range minDist {
		minDist[index] = numCells
	}

	return minDist
}
