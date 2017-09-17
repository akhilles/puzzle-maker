import numpy as np

def gen_random(n=5):
    puzzle = [[np.random.randint(1,max(n-(r), r-1, n-(c), c-1)+1) for r in range(1,n+1)] for c in range(1,n+1)]
    puzzle[n-1][n-1] = 0
    return np.array(puzzle)

def evaluate(puzzle):
    n = np.shape(puzzle)[1]
    visited = np.zeros(np.shape(puzzle), dtype=int)
    ans = np.zeros(np.shape(puzzle), dtype=int)
    #is_valid = lambda x: (x[0]>=0 and x[0]<n) and (x[1]>=0 and x[1]<n)

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
                q.append(i)
                visited[i] = 1
                accum+=1

    quality = ans[n-1,n-1]
    if(quality == 0):
        quality = -(np.sum(ans==0)-1)
    ans[0,0] = 0
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
    puz = gen_random(11)

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
    puz = np.array(m)
    print(puz)
    print(evaluate(puz))
