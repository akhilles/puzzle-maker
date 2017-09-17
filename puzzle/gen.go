package puzzle

import "math/rand"

// ConstraintMatrix returns a matrix containing the maximum legal value for each cell in the puzzle.
func ConstraintMatrix(n int) []int {
	cm := make([]int, n*n-1)

	for index := range cm {
		left := index % n
		right := n - 1 - left
		up := index / n
		down := n - 1 - up

		max := left
		if right > max {
			max = right
		}
		if up > max {
			max = up
		}
		if down > max {
			max = down
		}
		cm[index] = max
	}
	return cm
}

// RandomPuzzle returns a random legal puzzle of size n along with its constraint matrix.
func RandomPuzzle(n int) ([]int, []int) {
	cm := ConstraintMatrix(n)
	rp := make([]int, len(cm))

	for index, max := range cm {
		rp[index] = rand.Intn(max) + 1
	}

	return rp, cm
}
