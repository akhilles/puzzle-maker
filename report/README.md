# Local Search

CS 440 - Assignment 1  
Group: Akhil Velagapudi, Nithin Tammishetti

## Task 1

The GUI for vizualizing the puzzles was written in HTML and JavaScript. The algorithms that are responsible for actually generating the puzzles are written in Go and Python. The HTML front-end communicates with the back-end through REST calls.

\newpage
## Task 7

We implemented a genetic algorithm for our population-based approach. Each puzzle is represented as a chromosome by concatenating the values in each cell left to right and top to bottom. To make this simpler, the internal implementation used for the puzzle matrix was a flat array.


| Parameter | Value | Justification / Reasoning |
|----------------------------------------------------------------------------------------------|-----------------------------------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `n` size | given (5,7,9,11, ...) |  |
| number of generations | `4000` | Can be flexible but chosen to limit the computation time while also providing reasonably difficult puzzles. According to plots provided below, 4000 generations is well past the point of diminishing returns. |
| survival rate (subset of population that survives to then reproduce including elitists) | `0.3` | This factor is important in maintaining the genetic diversity of the population. A small survival rate can quickly arrive at local minimum by exploring only the very successful individuals. A higher survival rate can maintain a greater level of diversity in the population, leading to more exploration. |
| mutation rate (likelihood of each cell mutating after crossover) | `0.018` | Similar to survival rate, this parameter is important in determining the exploration vs exploitation of the genetic algorithm. High mutation rates are good for exploring new areas of the state space but come at cost of overlooking well-performing individuals. So, they slow down convergence to a local maximum but increase the likelihood of convergence to the global maximum. |
| population size | `n * n * 2` | This parameter has a positive correlation with diversity in the population. The cost of a large population is the increased computation necessary for simulating each generation. It is reasonable for this parameter to depend on the size of the state space, which is why it is represented as a function of n in our implementation. |
| elitism (number of individuals that survive each generation non-probabilistically) | `pop_size * surv_rate * 0.2` | Elitism is important for accelerating convergence because it allows the algorithm to not waste time re-discovering well-performing partial solutions that were discarded due to probabilistic selection. This introduces a form of greediness into the algorithm. Having too many elite individuals can hurt diversity. |

### Selection (w/ Elitism)

The selection model we employed is a fixed percent of the population (30%) survives and is responsible for repopulating to the original population size by creating offspring. These are the steps that take place each generation:

1. The fitness of each individual in the population is calculated
2. Elite individuals are automatically added to the survival pool
3. The fitness scores of all the individuals in the population are normalized to be positive by adding `n^2` to each score
4. A survivor is picked from the population using a random number generator with the normalized fitness of the individuals acting as weights
5. Step 4 is repeated until the survival pool is full (size of survival pool is 30% of population size)
6. Elite individuals are automatically added to the new population
7. 2 parents are chosen randomly from the survival pool and create 2 offspring (through crossover and mutation) which are added to the new population
8. Step 7 is repeated until the size of the new population is the same as the size of the initial population

### Crossover

After 2 parents are chosen for crossover, a crossover point is randomly picked between 25% and 75% of the length of the chromosomes. This method was implemented instead of picking a completely random crossover point in order to ensure atleast some of the features of each parent are passed on.

After a crossover point has been determined, all of the values to the right of the crossover point are swapped between the parent puzzles. The resulting offspring then go through a mutation process before being added to the new population.

### Mutation

Our original implementation of mutation simply tested each each cell in the puzzle to see if it should be mutated (using a random number generator and the cell mutation rate).

We found this to be inefficient and switched to determining whether a mutation should happen in the puzzle overall. If so, a random cell is picked and mutated to a new value.

In order to allow for multiple mutations, mutations will continue to occur as long as the random numbers being generated are less than the mutation rate. However, to avoid excessive multiple mutations, the mutation rate is reduced by a factor of 3 with each mutation in an individual.

### Puzzle

*See Figure 1*

![Puzzle with fitness of 68 generated by genetic algorithm](puzzle.png)

### Graphs

*See Figures 2 to 5*

![Fitness over number of generations for 5x5 puzzle](genalgo5.png)

![Fitness over number of generations for 7x7 puzzle](genalgo7.png)

![Fitness over number of generations for 9x9 puzzle](genalgo9.png)

![Fitness over number of generations for 11x11 puzzle](genalgo11.png)

\newpage
## Task 8

### n = 40 Puzzle

*See Figures 6 and 7*

method:			genetic algorithm  
fitness:		942  
generations:	160000  
pop size:		800  
survival rate:	0.3  
mutation rate:	0.018  
solution:		kjsdhfjksdhf

![40x40 puzzle generated by genetic algorithm with fitness 942](puzzle40.png)

![40x40 puzzle generated by genetic algorithm with fitness 942 (BFS process)](sol40.png)

### n = 20 Puzzle

*See Figures 8 and 9*

method:			genetic algorithm  
fitness:		323  
generations:	40000  
pop size:		400  
survival rate:	0.3  
mutation rate:	0.018  
solution:  
`D R D U D U D U D U R D U D U D U D L U R L R L R L R U D U L R L U D U R L D U U D U D R L U R L R L D R U R L D D U D U R D R D L R L R L R R U L D U U D U D U U L R L U D R U U U D L R L R L R L U D U U D D U U D D U R L R L L R L R L R L R D U D U D U L R R L R L R U D U D D L U D U D U D D U D D D U D L U R L L L R L R L R U D U L R L R R R U D U D D L U R U D U D L R U L U R D U D L L R L R L R L R L L R L R L U D U R L L U R R D U D D L R U L L R L D U U D U D U R L R D L R R R L R L D U U D U L R L D U L R L U D R L R L L R L R U D U L R U R L R D U U D D U D U D L U U D R L U R L R R L L R D L R L U R L D U D D U D D U R R D U D`

![20x20 puzzle generated by genetic algorithm with fitness 323](puzzle20.png)

![20x20 puzzle generated by genetic algorithm with fitness 323 (BFS process)](puzzle20.png)