package puzzle

import "fmt"

// PrintPuzzle prints the puzzle in the command line
func PrintPuzzle(n int, puzzle []int) {
	for index, val := range puzzle {
		fmt.Printf("%3d", val)
		if (index+1)%n == 0 {
			fmt.Println()
		}
	}
	fmt.Println()
	fmt.Println()
}
