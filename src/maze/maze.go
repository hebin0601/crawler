package main

import (
	"fmt"
	"os"
)

func readMazi(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	var row, col int
	fmt.Fscanf(file, "%d %d", &row, &col)

	maze := make([][]int, row)
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}

	return maze
}

type point struct {
	i, j int
}

var dirs = [4]point{
	{-1, 0}, {0, -1}, {1, 0}, {0, 1}}

func (p point) add(r point) point {
	return point{p.i + r.i, p.j + r.j}
}

func (p point) at(grid [][]int) (int, bool) {
	if p.i < 0 || p.j > len(grid) {
		return 0, false
	}
	if p.j < 0 || p.j > len(grid[p.i]) {
		return 0, flase
	}

	return grid(p.i, p.j), true

}

func walk(maze [][]int, start, end point) {
	steps := make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}

	Q := []point{start}

	for len(Q) > 0 {
		cur := Q[0]
		Q = Q[1:]

		for _, dir := range dirs {
			next := cur.add(dir)

			val, ok := next.at(maze)

			if ok != nil || val == 1 {
				continue
			}
			
			var ,ok = next.at(steps)
		}
	}
}

func main() {
	maze := readMazi("maze.in")

	for _, row := range maze {
		for _, col := range row {
			fmt.Printf("%d", col)
		}
		fmt.Println()
	}

	walk(maze, point{0, 0}, point{len(maze) - 1, maze[0] - 1})
}
