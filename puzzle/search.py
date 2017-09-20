import numpy as np
from evaluate import *
import matplotlib.pyplot as plt

def hill_climbing(n = 5, iters = 500, report = 100, p=0, temperature = 0, decay = 0.9):
    """returns the resultant puzzle, its evaluation, and the iteration data.
    Set p to a nonzero value for random walk probability.
    Report chooses how often to store data.
    Set temp to a nonzero value for simulated annealing."""
    puzzle = gen_random(n)
    best = -n*n
    data = []
    walk_chance = np.random.uniform(0,1)

    for i in range(iters):
        index = np.random.randint(0,n*n-1)
        r = index/n
        c = index%n

        _, fitness = evaluate(puzzle)
        curr = puzzle[r,c]
        change = rand_value((r,c),n)
        while(change == curr):
            change = rand_value((r, c), n)
        puzzle[r,c] = change
        _, fitness2 = evaluate(puzzle)

        if(temperature != 0 and fitness > fitness2):
            p = np.exp((fitness2-fitness)/temperature)
        if(fitness2<fitness and walk_chance>p):
            puzzle[r,c] = curr
        else:
            best = fitness2

        if((i+1)%report == 0):
            print('Iteration {}: {}'.format(i+1, best))
            data.append(best)
        temperature *=decay

    return puzzle, best, np.array(data)

def hill_restarts(n = 5, restarts = 10, iters = 100, report = 100, p=0, temperature = 0, decay = 0.9):
    best_puzzle = []
    best_eval = -n*n
    data = []
    for i in range(restarts):
        puzzle, eval, temp = hill_climbing(n, iters = iters, report=report, p=p, temperature=temperature, decay=decay)
        if(len(data)>0):
            data = np.concatenate((data,temp[temp>np.max(data)]))
        else:
            data = temp
        if(eval>best_eval):
            best_puzzle = puzzle
            best_eval = eval

    return best_puzzle, best_eval, np.array(data)


def rand_value(loc, n):
    r,c = loc[0]+1,loc[1]+1
    return np.random.randint(1, max(n - (r), r - 1, n - (c), c - 1) + 1)



if __name__=='__main__':
    #puz = gen_random(501)
    m = np.mat('2 2 2 4 3; \
            2 2 3 3 3;\
            3 3 2 3 3;\
            4 3 2 2 2;\
            1 2 1 4 0')
    n = np.mat('3 3 2 4 3;\
                2 2 2 1 1;\
                4 3 1 3 4;\
                2 3 1 1 3;\
                1 1 3 2 0')
    puz = np.array(n)
    # print(puz)
    # print(evaluate(puz))
    # print(hill_climbing(puz, 1))

    puz,best, data = hill_restarts(5,restarts=2,iters=500,report=1, p= 0, temperature = 0, decay=0.9)
    print("Restarts: ", puz, best)
    plt.plot(data, 'ro')
    plt.show()

    # puz, best, data = hill_climbing(11, iters=100, report= 1,p = 0.03)
    # print("Basic: ", puz, best)
    # plt.plot(data, 'ro')
    # plt.show()
