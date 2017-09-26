# Local Search

CS 440 - Assignment 1  
Group: Akhil Velagapudi, Nithin Tammishetti

## Task 1

The GUI for vizualizing the puzzles was written in HTML and JavaScript. The algorithms that are responsible for actually generating the puzzles are written in Go and Python. The HTML front-end communicates with the back-end through REST calls.

## Task 7

We implemented a genetic algorithm for our population-based approach. Each puzzle is represented as a chromosome by concatenating the values in each cell left to right and top to bottom. To make this simpler, the internal implementation used for the puzzle matrix was a flat array.


| Parameter | Value | Justification / Reasoning |
|-----------------------------------------------------------------------------------------|------------------------------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `n` size | given (5,7,9,11, ...) |  |
| number of generations | `4000` | Can be flexible but chosen to limit the computation time while also providing reasonably difficult puzzles. According to plots provided below, 4000 generations is well past the point of diminishing returns. |
| survival rate (subset of population that survives to then reproduce including elitists) | `0.3` (can be optimized through grid search) | This factor is important in maintaining the genetic diversity of the population. A small survival rate can quickly arrive at local minimum by exploring only the very successful individuals. A higher survival rate can maintain a greater level of diversity in the population, leading to more exploration. |
| mutation rate (likelihood of each cell mutating after crossover) | `0.018` (can be optimized through grid search) | Similar to survival rate, this parameter is important in determining the exploration vs exploitation of the genetic algorithm. High mutation rates are good for exploring new areas of the state space but come at cost of overlooking well-performing individuals. So, they slow down convergence to a local maximum but increase the likelihood of convergence to the global maximum. |
| population size | `n * n * 2` | This parameter has a positive correlation with diversity in the population. The cost of a large population is the increased computation necessary for simulating each generation. It is reasonable for this parameter to depend on the size of the state space, which is why it is represented as a function of n in our implementation. |
| elitism (number of individuals that survive each generation non-probabilistically) | `population size * survival rate * 0.2` | Elitism is important for accelerating convergence because it allows the algorithm to not waste time re-discovering well-performing partial solutions that were discarded due to probabilistic selection. This introduces a form of greediness into the algorithm. Having too many elite individuals can hurt diversity. |

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

![](puzzle.png)

### Graphs

![](genalgo5.png)

![](genalgo7.png)

![](genalgo9.png)

![](genalgo11.png)

## Task 8