package puzzle

// Evaluate returns the fitness score of a puzzle.
func Evaluate(n int, cells []int) (int, []int) {
	depthBFS := make([]int, n*n)
	goal := n*n - 1
	q := make([]int, 1, n*n)
	visited, accum, depth, nodes := 1, 0, 1, 1

	for len(q) > 0 {
		loc := q[0]
		q = q[1:]
		if loc != len(cells) {
			for _, move := range ValidMoves(n, loc, cells[loc]) {
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
	if depthBFS[goal] != 0 {
		return depthBFS[goal] + visited, depthBFS
	}
	return visited, depthBFS
}

func findParent(n int, depthBFS []int, index int) (int, string) {
	value := depthBFS[index]

	r := index / n
	c := index % n

	for j := c - 1; j >= 0; j-- {
		newIndex := r*n + j
		if depthBFS[newIndex] == value-1 {
			return newIndex, "R"
		}
	}

	for j := c + 1; j < n; j++ {
		newIndex := r*n + j
		if depthBFS[newIndex] == value-1 {
			return newIndex, "L"
		}
	}

	for j := r + 1; j < n; j++ {
		newIndex := j*n + c
		if depthBFS[newIndex] == value-1 {
			return newIndex, "U"
		}
	}

	for j := r - 1; j >= 0; j-- {
		newIndex := j*n + c
		if depthBFS[newIndex] == value-1 {
			return newIndex, "D"
		}
	}

	return 0, "WTF"
}

func Solution(n int, depthBFS []int) []string {
	index := len(depthBFS) - 1
	value := depthBFS[index]
	moves := make([]string, depthBFS[index])

	for i := len(moves) - 1; i >= 0; i-- {
		index, moves[i] = findParent(n, depthBFS, index)
		value--
	}

	return moves
}
