package puzzle

import "math/rand"
import "fmt"

type Puzzle struct {
	n               int
	Cells, DepthBFS []int
	Fitness         int
}

func ValidMoves(n int, cell int, val int) []int {
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
func RandomPuzzle(n int, cm []int) *Puzzle {
	rp := make([]int, len(cm))
	for index, max := range cm {
		rp[index] = rand.Intn(max) + 1
	}
	p := Puzzle{n: n, Cells: rp}
	p.Evaluate()
	return &p
}

func (p *Puzzle) String() string {
	out := ""

	for index, val := range p.Cells {
		out += fmt.Sprintf("%3d", val)
		if (index+1)%p.n == 0 {
			out += fmt.Sprintln()
		}
	}
	out += fmt.Sprintln()
	out += fmt.Sprintln()
	for index, val := range p.DepthBFS {
		out += fmt.Sprintf("%3d", val)
		if (index+1)%p.n == 0 {
			out += fmt.Sprintln()
		}
	}

	out += fmt.Sprint("fitness: ") + fmt.Sprintln(p.Fitness-p.n*p.n)
	return out
}
