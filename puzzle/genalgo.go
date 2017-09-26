package puzzle

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

//"github.com/gonum/plot"
//"github.com/gonum/plot/plotter"
//"github.com/gonum/plot/plotutil"
//"github.com/gonum/plot/vg"

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

	for i := elitism; i < len(survived); i++ {
		random := rand.Intn(sum) + 1
		sumCount := 0
		for j := 0; j < len(population); j++ {
			sumCount += populationFitness[j]
			if random <= sumCount {
				survived[i] = append([]int{}, population[j]...)
				break
			}
		}
	}
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

func GetPlot(n int, gens int, survRate float32, mutRate float32, count int) {
	lines := make([][]int, count)
	for i := range lines {
		_, _, _, line := GeneticPuzzle(n, gens, survRate, mutRate)
		lines[i] = line
	}

	p, _ := plot.New()
	p.Title.Text = "Genetic Algorithm, N = " + strconv.Itoa(n)
	p.X.Label.Text = "generation"
	p.Y.Label.Text = "fitness"
	p.X.Min = 0
	p.X.Max = float64(gens)
	p.Y.Min = 0
	p.Y.Max = float64(n*n - 1)
	p.Add(plotter.NewGrid())
	pts := make(plotter.XYs, len(lines[0]))

	for i := range pts {
		pts[i].X = float64(i * 10)
		sum := 0
		for _, line := range lines {
			sum += line[i]
		}
		pts[i].Y = float64(sum) / float64(count)
	}

	l, _ := plotter.NewLine(pts)
	p.Add(l)
	p.Save(6*vg.Inch, 6*vg.Inch, "genalgo"+strconv.Itoa(n)+".png")
}

func GeneticPuzzle(n int, gens int, survRate float32, mutRate float32) ([]int, []int, int, []int) {
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
	genFitness := make([]int, 0, gens/10)

	for i := 0; i < gens; i++ {
		survivors := pickSurvivors(population, populationFitness, elitism, survRate)

		if populationFitness[0] > bestFitness {
			bestPuzzle = population[0]
			bestFitness = populationFitness[0]
		}

		if i%10 == 0 {
			genFitness = append(genFitness, bestFitness-n*n)
		}

		var wg sync.WaitGroup

		for j := elitism; j < len(population); j += 2 {
			index := j
			wg.Add(1)
			go func() {
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
				defer wg.Done()
			}()
		}
		wg.Wait()
	}

	fit, dbfs := Evaluate(n, bestPuzzle)
	PrintTable(n, bestPuzzle)
	fmt.Println()
	PrintTable(n, dbfs)
	fmt.Println(fit - n*n)

	return bestPuzzle, dbfs, fit, genFitness
}
