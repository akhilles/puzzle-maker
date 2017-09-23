package puzzle

// Evaluate returns the fitness score of a puzzle.
func Evaluate(n int, cells []int, complete bool) (int, []int) {
	depthBFS := make([]int, n*n)
	goal := n*n - 1
	q := make([]int, 1, n*n)
	visited, accum, depth, nodes := 1, 0, 1, 1

	for len(q) > 0 {
		loc := q[0]
		q = q[1:]
		if loc != len(cells) {
			for _, move := range ValidMoves(n, loc, cells[loc]) {
				if !complete && move == goal {
					return depth + n*n, depthBFS
				}
				if move != 0 && (depthBFS[move] > depth || depthBFS[move] == 0) {
					depthBFS[move] = depth
					q = append(q, move)
					visited++
					accum++
				}
			}
		}

		nodes--
		if nodes == 0 {
			depth++
			nodes = accum
			accum = 0
		}
	}
	if complete && depthBFS[goal] != 0 {
		return depthBFS[goal] + n*n, depthBFS
	}
	return visited, depthBFS
}
