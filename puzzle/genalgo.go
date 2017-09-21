package puzzle

import "math/rand"

import "sort"
import "fmt"

type byFitness []*Puzzle

func (a byFitness) Len() int           { return len(a) }
func (a byFitness) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byFitness) Less(i, j int) bool { return a[i].Fitness < a[j].Fitness }

func initPopulation(n int, size int, cm []int) []*Puzzle {
	population := make([]*Puzzle, size)
	for i := range population {
		population[i] = RandomPuzzle(n, cm)
	}
	sort.Sort(sort.Reverse(byFitness(population)))
	return population
}

func pickSurvivors(population []*Puzzle, elitism int, survRate float32) []*Puzzle {
	survived := make([]*Puzzle, 0, int(float32(len(population))*survRate))
	survived = append(survived, population[0:elitism]...)

	sum := 0
	for i := elitism; i < len(population); i++ {
		sum += population[i].Fitness
	}

	for i := elitism; i < cap(survived); i++ {
		rand := rand.Intn(sum) + 1
		sumCount := 0
		for j := elitism; j < len(population); j++ {
			sumCount += population[j].Fitness
			if rand <= sumCount {
				survived = append(survived, population[j])
				break
			}
		}
	}
	return survived
}

func crossover(disabled bool, parentA *Puzzle, parentB *Puzzle) (childA *Puzzle, childB *Puzzle) {
	if !disabled {
		crossPoint := int(0.25 * float32(len(parentA.Cells)) * (1 + 2*rand.Float32()))
		childACells := make([]int, 0, len(parentA.Cells))
		childACells = append(childACells, parentA.Cells[:crossPoint]...)
		childACells = append(childACells, parentB.Cells[crossPoint:]...)
		childBCells := make([]int, 0, len(parentA.Cells))
		childBCells = append(childBCells, parentB.Cells[:crossPoint]...)
		childBCells = append(childBCells, parentA.Cells[crossPoint:]...)

		childA := Puzzle{
			n:     parentA.n,
			Cells: childACells,
		}
		childB := Puzzle{
			n:     parentA.n,
			Cells: childBCells,
		}
		return &childA, &childB
	}
	return parentA, parentB
}

func (child *Puzzle) doMutate(cm []int) int {
	mutIndex := rand.Intn(len(child.Cells))
	oldVal := child.Cells[mutIndex]
	for oldVal == child.Cells[mutIndex] {
		child.Cells[mutIndex] = rand.Intn(cm[mutIndex]) + 1
	}
	child.Evaluate()
	return child.Fitness
}

func (child *Puzzle) mutate(cm []int, mutRate float32) {
	for rand.Float32() < mutRate {
		mutIndex := rand.Intn(len(child.Cells))
		oldVal := child.Cells[mutIndex]
		for oldVal == child.Cells[mutIndex] {
			child.Cells[mutIndex] = rand.Intn(cm[mutIndex]) + 1
		}
		child.Evaluate()
		mutRate /= 3
	}
}

// goodMut, relevantMut
func GeneticPuzzle(n int, sizePop int, gens int, elitism int, survRate float32, mutRate float32) *Puzzle {
	cm := ConstraintMatrix(n)
	population := initPopulation(n, sizePop, cm)

	genFitness := make([]int, gens)
	bestPuzzle := population[0]

	for i := 0; i < gens; i++ {
		survivors := pickSurvivors(population, elitism, survRate)
		for j := elitism; j < len(population); j += 2 {
			parentAIndex := rand.Intn(len(survivors))
			parentBIndex := parentAIndex
			for parentAIndex == parentBIndex {
				parentBIndex = rand.Intn(len(survivors))
			}
			childA, childB := crossover(false, survivors[parentAIndex], survivors[parentBIndex])
			childA.mutate(cm, mutRate)
			childB.mutate(cm, mutRate)
			population[j] = childA
			population[j+1] = childB
		}

		sort.Sort(sort.Reverse(byFitness(population)))
		genFitness[i] = population[0].Fitness - n*n
		if population[0].Fitness > bestPuzzle.Fitness {
			bestPuzzle = population[0]
		}
	}

	fmt.Println(bestPuzzle)
	return bestPuzzle
}
