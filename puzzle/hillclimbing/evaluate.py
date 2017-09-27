import numpy as np
from flask import Flask


def gen_random(n=5):
    puzzle = [[np.random.randint(1,max(n-(r), r-1, n-(c), c-1)+1) for r in range(1,n+1)] for c in range(1,n+1)]
    puzzle[n-1][n-1] = 0
    return np.array(puzzle)

def readFromText(name = 'solve.txt'):
    f = open(name,'r+')
    n = int(f.readline())
    puzzle = []
    for line in f:
        puzzle.append(line.split())
    puzzle = np.array(puzzle,dtype=int)
    return puzzle


def solutionChain(sol):
    n = np.shape(sol)[0]
    curr = [n-1,n-1]
    chain = []

    while(curr != [0,0]):
        chain.insert(0,curr)
        curr = list(sol[curr[0],curr[1]])
    return chain

def solutionString(chain):
    x,y = 0,0
    ans = []

    if(len(chain)==1):
        return 'no solution'

    for i in chain:
        dx = x-i[1]
        dy = y-i[0]
        if(dx>0):
            ans.append('left')
        elif(dx<0):
            ans.append('right')
        elif(dy>0):
            ans.append('up')
        elif(dy<0):
            ans.append('down')
        y,x = i
    return ans




def evaluate(puzzle, getSol = False):
    n = 0

    if(len(np.shape(puzzle)) == 2):
        n = np.shape(puzzle)[1]
    else:
        n = int(np.sqrt(np.shape(puzzle)[0]))
        puzzle = puzzle.reshape(n,n)

    sol = np.zeros((n,n,2), dtype=int)

    visited = np.zeros(np.shape(puzzle), dtype=int)
    ans = np.zeros(np.shape(puzzle), dtype=int)

    q = []
    moves = valid_moves((0,0), puzzle[0,0], n)
    for i in moves:
        visited[i] = 1
    nodes_in_level = len(moves)
    level = 1
    q.extend(moves)
    accum = 0
    visited[0, 0] = 1

    while(len(q)>0):
        if(nodes_in_level == 0):
            level += 1
            nodes_in_level = accum
            accum = 0

        loc = q.pop(0)
        ans[loc] = level
        moves = valid_moves(loc, puzzle[loc], n)
        nodes_in_level -= 1

        for i in moves:
            if(visited[i] == 0):
                sol[i] = loc
                q.append(i)
                visited[i] = 1
                accum+=1

    quality = ans[n-1,n-1]
    if(quality == 0):
        quality = -(np.sum(ans==0)-1)
    ans[0,0] = 0
    if(getSol):
        return sol
    return ans, quality


def valid_moves(position, step, n):
    is_valid = lambda x: (x[0] >= 0 and x[0] < n) and (x[1] >= 0 and x[1] < n)

    all_moves = [(position[0],position[1]+step), (position[0],position[1]-step), (position[0]+step,position[1]),
                 (position[0]-step,position[1])]

    valid = []
    for i in all_moves:
        if(is_valid(i)):
            valid.append(i)

    return valid





if __name__=='__main__':
    puz = gen_random(501).flatten()

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
    print(puz)
    solution = evaluate(puz,getSol=True)
    print(solution)
    chain = solutionChain(solution)
    print(chain)
    direction = solutionString(chain)
    print(direction)

    sol = evaluate(readFromText('solve.txt'),getSol=True)
    chain = solutionChain(sol)
    dir = solutionString(chain)
    print(dir)

