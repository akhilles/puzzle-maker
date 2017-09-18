package puzzle

// Evaluate returns the fitness score of a puzzle.
func Evaluate(p []int) int {
	return 0
}

// BFS returns a matrix containing the minimum number of moves needed to reach that cell from the start cell.
func BFS(n int, p []int, vm [][]int) []int {
	minDist := make([]int, n*n)
	q := []int{0}
	accum, depth, nodes := 0, 1, 1

	for len(q) > 0 {
		loc := q[0]
		q = q[1:]
		for _, move := range vm[loc] {
			if minDist[move] > depth || (minDist[move] == 0 && move != 0) {
				minDist[move] = depth
				q = append(q, move)
				accum++
			}
		}

		nodes--
		if nodes == 0 {
			depth++
			nodes = accum
			accum = 0
		}
	}

	PrintPuzzle(n, minDist)
	return minDist
}
