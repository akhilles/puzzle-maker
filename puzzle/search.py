import numpy as np
from evaluate import *
import matplotlib.pyplot as plt
from flask import Flask
from flask import jsonify
from flask_cors import CORS, cross_origin
from flask import request
app = Flask(__name__)
cors = CORS(app)

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
        walk_chance = np.random.uniform(0, 1)

        _, fitness = evaluate(puzzle)
        curr = puzzle[r,c]
        change = rand_value((r,c),n)
        while(change == curr):
            change = rand_value((r, c), n)
        puzzle[r,c] = change
        _, fitness2 = evaluate(puzzle)

        prob = p
        if(temperature != 0 and fitness > fitness2):
            prob = np.exp((fitness2-fitness)/temperature)
        if(fitness2<fitness and walk_chance>prob):
            puzzle[r,c] = curr
        else:
            best = fitness2

        if((i+1)%report == 0):
            print('Iteration {}: {}'.format(i+1, best))
            data.append(best)
            print(temperature, prob)
        temperature *=decay

    return puzzle, best, np.array(data)

def hill_restarts(n = 5, restarts = 10, iters = 100, report = 100, p=0, temperature = 0, decay = 0.9):
    best_puzzle = []
    best_eval = -n*n
    data = []
    for i in range(restarts):
        puzzle, eval, temp = hill_climbing(n, iters = iters, report=report, p=p, temperature=temperature, decay=decay)
        if(len(data)>0):
            prevmax = np.max(data)
            temp[temp<prevmax] = prevmax
            data = np.concatenate((data,temp))
        else:
            data = temp
        if(eval>best_eval):
            best_puzzle = puzzle
            best_eval = eval

    return best_puzzle, best_eval, np.array(data)

@app.route('/climb')
@cross_origin()
def climbing(n=5):
    n = int(request.args.get('n'))
    iters = request.args.get('iters')
    if(iters==None):
        iters=5000
    else:
        iters=int(iters)

    puz, best, tempdata = hill_restarts(n, restarts=1, iters=iters, report=100, p=0, temperature=0, decay=0.999)
    depth,b = evaluate(puz)
    puz = puz.flatten()
    depth = depth.flatten()

    return jsonify(
        Cells=list(puz[:-1]),
        DepthBFS=list(depth),
        Fitness=best,
        Iterations=iters
    )

def rand_value(loc, n):
    r,c = loc[0]+1,loc[1]+1
    return np.random.randint(1, max(n - (r), r - 1, n - (c), c - 1) + 1)



if __name__=='__main__':
    #puz = gen_random(501)
    # m = np.mat('2 2 2 4 3; \
    #         2 2 3 3 3;\
    #         3 3 2 3 3;\
    #         4 3 2 2 2;\
    #         1 2 1 4 0')
    # n = np.mat('3 3 2 4 3;\
    #             2 2 2 1 1;\
    #             4 3 1 3 4;\
    #             2 3 1 1 3;\
    #             1 1 3 2 0')
    # puz = np.array(n)
    # print(puz)
    # print(evaluate(puz))

    sizes = [5,7,9,11]
    num_runs = 50

    iters = 1000
    report = 50

    for n in sizes:
        data = np.zeros(int(iters/report))

        for i in range(num_runs):
            puz, best, tempdata = hill_restarts(n,restarts=1,iters=iters,report=report, p= 0, temperature = 80, decay=0.95)
            print("Restarts: ", puz, best)
            data = data + tempdata

        plt.figure()
        plt.plot(np.arange(0,iters,report),data/num_runs)
        plt.ylabel('Fitness')
        plt.xlabel('Iterations')
        plt.title('{} by {} HC with Annealing'.format(n,n))
        plt.savefig('plots/{}by{}anneal.png'.format(n,n))


    # puz, best, data = hill_restarts(11, restarts=1, iters=10000, report=100, p=0.0025, temperature=0, decay=0)
    # print("Basic: ", puz, best)
    # plt.figure()
    # plt.plot(np.arange(0, 10000, 100), data)
    # plt.ylabel('Fitness')
    # plt.xlabel('Iterations')
    # plt.title('{} by {} HC with Annealing'.format(11, 11))
    # plt.savefig('plots/{}by{}annealdeep.png'.format(11, 11))



    #try different params

    # ps = np.array([.02,.01,.005,.002,.001,.0005])
    #
    # bestp = 0
    # bestval = -1
    # pgraph = []
    # for p in ps:
    #     data = np.zeros(30)
    #     for i in range(2):
    #         puz, best, datatemp = hill_restarts(11, restarts=1, iters=3000, report=100, p=p, temperature=0, decay=0.999)
    #         data += datatemp
    #     data = data/2
    #     pgraph.append(data[-1])
    #     if(data[-1]>bestval):
    #         bestp = p
    #         bestval = data[-1]
    #
    # print(bestp,bestval)
    #
    # plt.figure()
    # plt.plot(ps,np.array(pgraph))
    # plt.ylabel('Average Evaluation')
    # plt.xlabel('Value of p')
    # plt.title('Determining good value for p'.format(11, 11))
    # plt.savefig('plots/pgraph')

    # puz, best, data = hill_restarts(11, restarts=1, iters=10000, report=100, p=0, temperature=100, decay=0.995)
    # print("Basic: ", puz, best)
    # plt.figure()
    # plt.plot(np.arange(0, 10000, 100), data)
    # plt.ylabel('Fitness')
    # plt.xlabel('Iterations')
    # plt.title('{} by {} HC with Annealing'.format(11, 11))
    # plt.savefig('plots/{}by{}annealdeep.png'.format(11, 11))

