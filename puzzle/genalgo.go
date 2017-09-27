package puzzle

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func initPopulation(n int, size int, cm []int) ([][]int, []int) {
	population := make([][]int, size)
	fitness := make([]int, size)
	for i := range population {
		p := RandomPuzzle(n, cm)
		population[i] = p
		fitness[i], _ = Evaluate(n, p)
	}
	return population, fitness
}

func pickSurvivors(population [][]int, populationFitness []int, elitism int, survRate float32) [][]int {
	numSurvived := int(float32(len(population)) * survRate)
	survived := make([][]int, numSurvived)
	for i := 0; i < elitism; i++ {
		maxValue := populationFitness[i]
		for j := i + 1; j < len(population); j++ {
			if populationFitness[j] > maxValue {
				maxValue = populationFitness[j]
				population[i], population[j] = population[j], population[i]
				populationFitness[i], populationFitness[j] = populationFitness[j], populationFitness[i]
			}
		}
		survived[i] = append([]int{}, population[i]...)
	}

	sum := 0
	for _, val := range populationFitness {
		sum += val
	}

	var wg sync.WaitGroup

	for i := elitism; i < len(survived); i++ {
		indf := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			random := rand.Intn(sum) + 1
			sumCount := 0
			for j := 0; j < len(population); j++ {
				sumCount += populationFitness[j]
				if random <= sumCount {
					survived[indf] = append([]int{}, population[j]...)
					break
				}
			}
		}()
	}

	wg.Wait()

	return survived
}

func crossover(parentA []int, parentB []int) ([]int, []int) {
	crossPoint := int(float32(len(parentA)) * (0.25 + 0.5*rand.Float32()))
	childA := make([]int, len(parentA))
	childB := make([]int, len(parentA))
	for i := range parentA {
		if i < crossPoint {
			childA[i] = parentA[i]
			childB[i] = parentB[i]
		} else {
			childA[i] = parentB[i]
			childB[i] = parentA[i]
		}
	}

	return childA, childB
}

func mutate(n int, child []int, cm []int, mutRate float32) int {
	for rand.Float32() < mutRate {
		mutIndex := rand.Intn(len(child))
		newValue := child[mutIndex] + rand.Intn(cm[mutIndex]-1) + 1
		if newValue > cm[mutIndex] {
			newValue -= cm[mutIndex]
		}
		child[mutIndex] = newValue

		mutRate /= 3
	}

	fitness, _ := Evaluate(n, child)
	return fitness
}

func GeneticPuzzle(n int, gens int, survRate float32, mutRate float32) ([]int, []int, int, []string) {
	start := time.Now()

	sizePop := n * n * 2
	mutRate *= float32(n * n)
	elitism := int(float32(sizePop) * survRate * 0.2)
	if elitism%2 == 1 {
		elitism++
	}

	cm := ConstraintMatrix(n)
	population, populationFitness := initPopulation(n, sizePop, cm)

	bestPuzzle := population[0]
	bestFitness := populationFitness[0]

	for i := 0; i < gens; i++ {
		survivors := pickSurvivors(population, populationFitness, elitism, survRate)

		if populationFitness[0] > bestFitness {
			bestPuzzle = population[0]
			bestFitness = populationFitness[0]
		}

		var wg sync.WaitGroup

		for j := elitism; j < len(population); j += 2 {
			index := j
			wg.Add(1)
			go func() {
				defer wg.Done()
				numSurvived := len(survivors)
				parentAIndex := rand.Intn(numSurvived)
				parentBIndex := parentAIndex + rand.Intn(numSurvived-1) + 1
				if parentBIndex >= numSurvived {
					parentBIndex -= numSurvived
				}
				childA, childB := crossover(survivors[parentAIndex], survivors[parentBIndex])
				populationFitness[index] = mutate(n, childA, cm, mutRate)
				populationFitness[index+1] = mutate(n, childB, cm, mutRate)
				population[index] = childA
				population[index+1] = childB
			}()
		}
		wg.Wait()
	}

	fit, dbfs := Evaluate(n, bestPuzzle)

	elapsed := time.Since(start)
	fmt.Printf("Fitness:  %d, ", bestFitness-n*n)
	fmt.Printf("Duration: %s\n", elapsed)

	return bestPuzzle, dbfs, fit, Solution(n, dbfs)
}
