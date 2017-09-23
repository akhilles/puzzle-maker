package puzzle

import "math/rand"
import "fmt"

func ValidMoves(n int, cell int, val int) []int {
	moves := make([]int, 4)

	row := cell / n
	col := cell % n

	if row-val >= 0 {
		moves[0] = cell - n*val
	}
	if row+val < n {
		moves[1] = cell + n*val
	}
	if col-val >= 0 {
		moves[2] = cell - val
	}
	if col+val < n {
		moves[3] = cell + val
	}
	return moves
}

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

// RandomPuzzle returns a random legal puzzle of size n along with its constraint matrix and valid moves table.
func RandomPuzzle(n int, cm []int) []int {
	p := make([]int, len(cm))
	for index, max := range cm {
		p[index] = rand.Intn(max) + 1
	}
	return p
}

func PrintTable(n int, p []int) {
	for index, val := range p {
		fmt.Printf("%3d", val)
		if (index+1)%n == 0 {
			fmt.Println()
		}
	}
	fmt.Println()
}
