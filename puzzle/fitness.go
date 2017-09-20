package puzzle

// Evaluate returns the fitness score of a puzzle.
func Evaluate(n int, vm [][]int) (int, []int) {
	depthBFS := BFS(n, vm)
	moves := depthBFS[len(depthBFS)-1]
	if moves == 0 {
		unreached := 0
		for _, dist := range depthBFS {
			if dist == 0 {
				unreached++
			}
		}
		return 1 - unreached, depthBFS
	}
	return moves, depthBFS
}

// BFS returns a matrix containing the minimum number of moves needed to reach that cell from the start cell.
func BFS(n int, vm [][]int) []int {
	depthBFS := make([]int, n*n)
	q := make([]int, 1, n*n)
	accum, depth, nodes := 0, 1, 1

	for len(q) > 0 {
		loc := q[0]
		q = q[1:]
		for _, move := range vm[loc] {
			if depthBFS[move] > depth || (depthBFS[move] == 0 && move != 0) {
				depthBFS[move] = depth
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

	return depthBFS
}
