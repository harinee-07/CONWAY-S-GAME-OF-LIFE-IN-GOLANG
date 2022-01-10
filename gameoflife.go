package main

import (
	"fmt"
	"time"
)

func main() {
	// Size of the grid in m x n
	var m, n int
	fmt.Scanf("%d", &m)
	fmt.Scanf("%d", &n)

	// No of generations to run
	var gen int
	fmt.Scanf("%d", &gen)

	grid := make([][]int, m)
	for i := range grid {
		grid[i] = make([]int, n)
	}

	// Get the initial live cells in the grid
	var getlivecell int
	fmt.Scanf("%d", &getlivecell)
	for i := 0; i < getlivecell; i++ {
		var r, c int
		fmt.Scanf("%d,%d", &r, &c)

		grid[r][c] = 1
	}
	fmt.Print("Initial generation")
	fmt.Printf("\n")

	for k := 0; k < m; k++ {
		for l := 0; l < n; l++ {
			fmt.Printf("%d ", grid[k][l])
		}
		fmt.Printf("\n")
	}

	s := time.Now()
	for i := 0; i < gen; i++ {
		grid = getnextgen(grid, m, n)

		extinct := true
		for k := 0; k < m; k++ {
			for l := 0; l < n; l++ {
				if grid[k][l] == 1 {
					extinct = false
					break
				}
			}
		}

		if extinct {
			gen = i + 1
			break
		}
	}

	fmt.Println("Grid for Generation ", gen)
	for k := 0; k < m; k++ {
		for l := 0; l < n; l++ {
			fmt.Printf("%d ", grid[k][l])
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")

	fmt.Println("time taken: ", time.Since(s))
}

func getnextgen(univ [][]int, m, n int) [][]int {
	nextGen := make([][]int, m)
	for i := range nextGen {
		nextGen[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			cnt := getliveneighbourcount(univ, i, j, m, n)

			if univ[i][j] == 1 {
				if cnt < 2 || cnt > 3 {
					nextGen[i][j] = 0
				}

				if cnt == 2 || cnt == 3 {
					nextGen[i][j] = 1
				}
			} else {
				if cnt == 3 {
					nextGen[i][j] = 1
				}
			}

		}
	}

	return nextGen
}

func getliveneighbourcount(univ [][]int, i, j, m, n int) int {
	count := 0
	for k := i - 1; k <= i+1; k++ {
		for l := j - 1; l <= j+1; l++ {
			x := k
			y := l
			if x < 0 {
				x = m - 1
			}

			if x > m-1 {
				x = 0
			}

			if y < 0 {
				y = n - 1
			}

			if y > n-1 {
				y = 0
			}

			if x == i && y == j {
				continue
			}

			if univ[x][y] == 1 {
				count++
			}
		}
	}

	return count
}
