# 8-puzzle problem in Go
This project presents a program capable of solving the 8-puzzle problem in three different ways:
1) Using uniform cost
2) Using a simple A* algorithm, where the manhattan distance is used as heuristic function
3) Using a more precise A* algorithm, where the manhattan distance + linear conflicts are used as heuristic function

### What is an 8-puzzle problem?
It's a puzzle consisting of a 3x3 grid, where there are 8 numbered tiles labeled from 1 to 8, and 1 empty square.
The puzzle starts with the squares shuffled randomnly, and the objective is to rearrange the grid by moving the tiles using the empty square and reach a goal state.
In this case, the empty tile is represented by the 0, and the goal state is:
```
1 2 3
4 5 6
7 8 0
```

### Motivations
In one of my AI classes in university, I had to build the same program as a project, but I did it in Python (the code is also here on my GitHub).
After finishing it, I thought reimplementing it in Go would be a great way to get a good grasp of the language, since I'm learning it.
That, and the fact that I was curious to see the differences in performance between the Python implementation and the Go implementation. Thus, this Go version of the program was born.

### Configuring the program
To choose which approach to solve the problem, head to `nodes/node.go` file and uncomment the `SELECTED_STRATEGY` enum that you wish to use, while commenting the other ones.
```
// const SELECTED_STRATEGY = heuristics.UNIFORM_COST
// const SELECTED_STRATEGY = heuristics.A_STAR_MANHATTAN
// const SELECTED_STRATEGY = heuristics.A_STAR_LINEAR_CONFLICT
```

If you wish to see the execution of the algorithm visually, head to `main.go` and change the `DEBUG` constant flag to `true`

### Running the program
To run the program:
1) Open it in a folder, and run `go run main.go` in the terminal
2) Enter a whitespace-separated sequence of unique numbers from 0 to 8 (some input examples are commented in `main.go`)
3) Enjoy :)
