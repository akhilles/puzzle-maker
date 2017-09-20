package puzzle

import "math/rand"
import "fmt"

type Puzzle struct {
	cells, constraints, depthBFS []int
	moves                        [][]int
	fitness                      int
}

func validMoves(n int, cell int, val int) []int {
	moves := make([]int, 0, 4)

	row := cell / n
	col := cell % n

	if row-val >= 0 {
		moves = append(moves, cell-n*val)
	}
	if row+val < n {
		moves = append(moves, cell+n*val)
	}
	if col-val >= 0 {
		moves = append(moves, cell-val)
	}
	if col+val < n {
		moves = append(moves, cell+val)
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
func RandomPuzzle(n int) *Puzzle {
	cm := ConstraintMatrix(n)
	rp := make([]int, len(cm))
	vm := make([][]int, len(cm)+1)

	for index, max := range cm {
		rp[index] = rand.Intn(max) + 1
		vm[index] = validMoves(n, index, rp[index])
	}

	fitness, depthBFS := Evaluate(n, rp, vm)

	return &Puzzle{rp, cm, depthBFS, vm, fitness}
}

// PrintPuzzle prints the puzzle in the command line.
func PrintPuzzle(p) {
	for index, val := range puzzle {
		fmt.Printf("%3d", val)
		if (index+1)%n == 0 {
			fmt.Println()
		}
	}
	fmt.Println()
	fmt.Println()
}
