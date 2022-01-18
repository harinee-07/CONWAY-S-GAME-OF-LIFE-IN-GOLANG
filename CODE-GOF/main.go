/*
 *   Author - Harinee
 *   Project - Game of Life
 *   Language - Go Lang
 *
 *   This is an implementation of Conway's Game of Life,
 *   https://en.wikipedia.org/wiki/Conway's_Game_of_Life.
 */

// creating and defining the main package of the program

package main

//importing neccessary libararies
import (
	"fmt"
	"CODE-GOF/modules"
)


func main() {
	//Initializing the variables

	var m,n,gen,getlivecell int
	//Get the input from a function call
	m,n,gen,getlivecell=modules.Get_input()
	
	//Make the grid with the row and column
	grid := make([][]int, m)
	for i := range grid {
		grid[i] = make([]int, n)
	}

	
	//Get the intial values for the grid
	for i := 0; i < getlivecell; i++ {
		var r, c int
		fmt.Scanln(&r, &c)

		grid[r][c] = 1
	}
	//Print the initial generation with a function call
	fmt.Print("Initial generation \n")
	modules.PrinttheGrid(grid,m,n)

	//Loop To run for the generations
	for i := 0; i < gen; i++ {
		grid = modules.Getnextgen(grid, m, n)

		extinct := true
		extinct=modules.CheckExtinct(grid,m,n,extinct)
		
		if extinct {
			gen = i + 1
			break
		}
	}


	//Print  the Final Generation with the Print Function
	fmt.Println("Grid for Generation ", gen)
	modules.PrinttheGrid(grid,m,n)
	
}

