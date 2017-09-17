package puzzle

import "math/rand"

func ConstraintMatrix(n int) []int {
	var cm []int
	cm = make([]int, n*n-1)

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

func RandomPuzzle(n int) ([]int, []int) {
	cm := ConstraintMatrix(n)
	var rp []int
	rp = make([]int, len(cm))

	for index, max := range cm {
		rp[index] = rand.Intn(max) + 1
	}

	return rp, cm
}
