package puzzle

import "math/rand"

import "fmt"
import "sort"

type byFitness []*Puzzle

func (a byFitness) Len() int           { return len(a) }
func (a byFitness) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byFitness) Less(i, j int) bool { return a[i].Fitness < a[j].Fitness }

func initPopulation(n int, size int) []*Puzzle {
	population := make([]*Puzzle, size)
	for i := range population {
		population[i] = RandomPuzzle(n)
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

func doSex(parentA *Puzzle, parentB *Puzzle, mutRate float32) {

}

// goodMut, relevantMut
func GeneticPuzzle(n int, sizePop int, gens int, elitism int, survRate float32, mutRate float32) {
	population := initPopulation(n, sizePop)

	genFitness := make([]int, gens)
	bestPuzzle := population[0]

	for i := 0; i < gens; i++ {
		survivors := pickSurvivors(population, elitism, survRate)
		for j := elitism; j < len(population); j++ {

		}

		sort.Sort(sort.Reverse(byFitness(population)))
	}

	fmt.Println(population)
	fmt.Println(pickSurvivors(population, elitism, survRate))

}
